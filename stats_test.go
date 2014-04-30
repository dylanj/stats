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
		t.Error("There should be a User.")
	}

	if c := s.getUser(n, nick); c == nil {
		t.Error("User should not be nil.")
	}

	if len(s.Users) != 1 {
		t.Error("There should be only one User.")
	}
}

// func TestStats_HourlyChart(t *testing.T) {
// 	t.Parallel()

// 	numMessagesPerHour := 4

// 	s := NewStats()
// 	n := s.NewNetwork(network)
// 	n.AddChannel(channel)
// 	n.AddUser(nick)

// 	// for i := 23; i >= 0; i++ {
// 	// 	for j := 0; j < i; i++ {
// 	// 		//new msg?
// 	// 	}
// 	// }

// 	for i := 0; i < numMessagesPerHour*24; i++ {
// 		m := NewMessageString("hello")
// 		date := time.Date(2014, time.April, 29, i%24, 30, 0, 1, time.UTC)
// 		m.Date = date

// 		s.AddMessage(m)
// 		//todo add message to channel
// 	}

// 	chart := s.HourlyChart("test_network", "#test")
// 	for i := 0; i < 24; i++ {
// 		if chart[i] != numMessagesPerHour {
// 			t.Error("Hour", i, "Does not have", numMessagesPerHour, "messages.")
// 		}
// 	}
// }
