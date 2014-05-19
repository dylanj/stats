package stats

import (
	"bytes"
	"testing"
	"time"
)

func init() {
	// override loadDB to avoid fs calls
	fileOpener = &nilFileOpener{}
}

const (
	network  = "test_network"
	channel  = "#test"
	nick     = "phish"
	name     = "Dylan"
	host     = "foo.zqz.ca"
	hostmask = nick + "!" + name + "@" + host
)

func TestStats_GetterMethods(t *testing.T) {
	t.Parallel()

	s := NewStats()

	if n := s.GetNetwork(network); n != nil {
		t.Error("Network should be nil.")
	}

	if c := s.GetChannel(network, channel); c != nil {
		t.Error("Channel should be nil on network.")
	}

	if u := s.GetUser(network, nick); u != nil {
		t.Error("User should be nil on network.")
	}

	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "some foo")

	if n := s.GetNetwork(network); n == nil || n.Name != network {
		t.Error("Should be able to lookup the network.")
	}

	if c := s.GetChannel(network, channel); c == nil || c.Name != channel {
		t.Error("Should be able to lookup the channel on the network.")
	}

	if u := s.GetUser(network, nick); u == nil || u.Nick != nick {
		t.Error("Should be able to lookup the user on the network.")
	}
}

func TestStats_AddMessage(t *testing.T) {
	t.Parallel()

	s := NewStats()

	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "tree foo http://google.com")

	if len(s.Networks) != 1 {
		t.Error("It should add the network.")
	}

	if len(s.Channels) != 1 {
		t.Error("It should add the channel.")
	}

	if len(s.Users) != 1 {
		t.Error("It should add the user.")
	}

	if len(s.Messages) != 1 {
		t.Error("It should add the message.")
	}

	if len(s.Networks[1].WordCounter.All) != 2 {
		t.Error("It should add the two words in the message.")
	}

	if len(s.Channels[1].WordCounter.All) != 2 {
		t.Error("It should add the two words in the message.")
	}

	if len(s.Networks[1].URLCounter.All) != 1 {
		t.Error("It should add the URL in the message.")
	}

	if len(s.Channels[1].URLCounter.All) != 1 {
		t.Error("It should add the URL in the message.")
	}

	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "some foo")

	if len(s.Networks) > 1 {
		t.Error("It should not add another network.")
	}

	if len(s.Channels) > 1 {
		t.Error("It should not add another channel.")
	}

	if len(s.Users) > 1 {
		t.Error("It should not add another user.")
	}

	if len(s.Messages) != 2 {
		t.Error("It should add the message.")
	}
}

func TestStats_AddKickMessage(t *testing.T) {
	t.Parallel()

	s := NewStats()
	s.AddMessage(Kick, network, channel, "dylan", time.Now(), "fish")

	n := s.networkByName[network]
	u := n.users["dylan"]

	if u.KickCounters.Sent != 1 {
		t.Error("Should have incremented dylan's sent counter.")
	}
	if u.KickCounters.Received != 0 {
		t.Error("Should not have incremented dylan's received counter.")
	}

	// add a message by fish
	s.AddMessage(Msg, network, channel, "fish", time.Now(), "salut")
	s.AddMessage(Kick, network, channel, "dylan", time.Now(), "fish")

	u2 := n.users["fish"]

	if u.KickCounters.Sent != 2 {
		t.Error("Should have incremented dylan's sent counter.")
	}
	if u.KickCounters.Received != 0 {
		t.Error("Should not have incremented dylan's received counter.")
	}

	if u2.KickCounters.Sent != 0 {
		t.Error("Should not have incremented fish's sent counter.")
	}

	if u2.KickCounters.Received != 1 {
		t.Error("Should have incremented fish's received counter.")
	}
}

func TestStats_AddMessageBlankChannel(t *testing.T) {
	t.Parallel()

	s := NewStats()

	s.AddMessage(Msg, network, "", hostmask, time.Now(), "some foo")

	if len(s.Channels) != 0 {
		t.Error("It should not add a channel.")
	}
}

func TestStats_addNetwork(t *testing.T) {
	t.Parallel()

	s := NewStats()

	if len(s.Networks) != 0 {
		t.Error("Network should not exist at this point.")
	}

	n := s.addNetwork(network)

	if len(s.Networks) != 1 {
		t.Error("Network should exist.")
	}

	if n, ok := s.Networks[n.ID]; !ok {
		t.Error("Should be able to look up network by ID")
	} else if n.Name != network {
		t.Errorf("Expected: %v, Got: %v", network, n.Name)
	}
}

func TestStats_getNetwork(t *testing.T) {
	t.Parallel()

	s := NewStats()

	if len(s.Networks) != 0 {
		t.Error("Network should not exist at this point.")
	}

	s.getNetwork(network)

	if n := s.getNetwork(network); n == nil {
		t.Error("Network should not be nil")
	}

	if len(s.Networks) != 1 {
		t.Error("Network should exist.")
	}

	if n := s.getNetwork(network); n == nil {
		t.Error("Network should not be nil")
	}

	if len(s.Networks) != 1 {
		t.Error("It should not duplicate the network.")
	}
}

func TestStats_getChannel(t *testing.T) {
	t.Parallel()

	s := NewStats()

	n := s.addNetwork(network)

	if len(s.Channels) != 0 {
		t.Error("There should not be any channels.")
	}

	if c := s.getChannel(n, channel); c == nil {
		t.Error("Channel should not be nil.")
	}

	if len(s.Channels) != 1 {
		t.Error("There should be a channel.")
	}

	if c := s.getChannel(n, channel); c == nil {
		t.Error("Channel should not be nil.")
	}

	if len(s.Channels) != 1 {
		t.Error("There should be only one channel.")
	}
}

func TestStats_getUser(t *testing.T) {
	t.Parallel()

	s := NewStats()

	n := s.addNetwork(network)

	if len(s.Users) != 0 {
		t.Error("There should not be any users.")
	}

	if c := s.getUser(n, nick); c == nil {
		t.Error("User should not be nil.")
	}

	if len(s.Users) != 1 {
		t.Error("There should be a user.")
	}

	if c := s.getUser(n, nick); c == nil {
		t.Error("User should not be nil.")
	}

	if len(s.Users) != 1 {
		t.Error("There should be only one user. (", len(s.Users), ")")
	}
}

func TestStats_buildIndexes(t *testing.T) {
	t.Parallel()

	s := NewStats()
	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "some foo")
	s.networkByName = nil

	s.buildIndexes()

	if s.networkByName == nil {
		t.Error("networkByName should have been created")
	}

	if s.networkByName[network] == nil {
		t.Error("should be able to look up network")
	}
}

func TestStats_SaveLoadDB(t *testing.T) {
	t.Parallel()

	defer func() {
		fileOpener = &nilFileOpener{}
	}()

	s := NewStats()
	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "some foo")

	b := bytes.Buffer{}
	fileOpener = &fakeFileOpener{&b}

	if s.Save() != true {
		t.Error("Should be able to create data.db.")
	}

	s, e := loadDatabase()

	if e != nil {
		t.Error("Should not be nil.")
	}

	if len(s.Messages) != 1 {
		t.Error("Should have loaded 1 message.")
	}
}
