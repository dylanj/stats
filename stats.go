package stats

import "fmt"

type Stats struct {
	users    []User
	channels map[string]*Channel
}

func NewStats() *Stats {
	return &Stats{
		//users: make([]User),
		channels: make(map[string]*Channel),
	}
}

func (s *Stats) AddChannel(channel *Channel) {
	if !s.HasChannelByChannel(channel) {
		s.channels[string(channel.Name)] = channel
		fmt.Printf("Adding %s to Channels\n", channel.Name)
	} else {
		fmt.Printf("Already have %s in list of channels\n", channel.Name)
	}
}

func (s *Stats) ListChannels() {
	fmt.Printf("\nListing Channels:\n")
	for key, channel := range s.channels {
		fmt.Printf("Channel (%s) Name: %s\n", key, channel.GetName())
	}
}

func (s *Stats) HasChannelByName(name []byte) bool {
	channel_name := string(name)
	channel := s.channels[channel_name]

	return channel != nil
}

func (s *Stats) HasChannelByChannel(channel *Channel) bool {
	return s.HasChannelByName(channel.Name)
}
