package stats

import (
	"testing"
	"time"
)

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

	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "some foo")

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

func TestStats_HourlyChart(t *testing.T) {
	t.Parallel()

	s := NewStats()
	n := s.addNetwork(network)
	c := s.addChannel(n, channel)
	u := s.addUser(n, nick)

	for i := 23; i >= 0; i-- {
		for j := 0; j < i; j++ {
			date := time.Date(2014, time.April, 29, i, 30, 0, 1, time.UTC)

			s.addMessage(Msg, n, c, u, date, "nihao")
		}
	}

	chart, success := s.HourlyChart("test_network", "#test")

	if !success {
		t.Errorf("success should be true")
	}

	for i := 0; i < 24; i++ {
		if chart[i] != i {
			t.Errorf("Hour[%d] has %d messages, expected: %d", i, chart[i], i)
		}
	}
}
