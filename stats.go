package stats

import "fmt"

type Stats struct {
	channels map[string]*Channel
	users    map[string]*User
}

func NewStats() *Stats {
	return &Stats{
		channels: make(map[string]*Channel),
		users:    make(map[string]*User),
	}
}

func (s *Stats) AddChannel(channel *Channel) {
	if !s.HasChannelByChannel(channel) {
		fmt.Printf("Adding %s to Channels\n", channel.Name)
		s.channels[string(channel.Name)] = channel
	} else {
		fmt.Printf("Already have %s in list of channels\n", channel.Name)
	}
}

func (s *Stats) AddUser(user *User) {
	hostmask := string(user.Hostmask)
	if s.users[hostmask] == nil {
		fmt.Printf("Adding %s to users\n", user.Name)
		s.users[hostmask] = user
	} else {
		fmt.Printf("Already have user\n")
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
