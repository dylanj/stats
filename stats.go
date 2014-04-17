package stats

import "fmt"

type Stats struct {
	channels map[string]*Channel
	users    map[string]*User
	messages []*Message
}

func NewStats() *Stats {
	return &Stats{
		channels: make(map[string]*Channel),
		users:    make(map[string]*User),
	}
}

func (s *Stats) AddMessage(message *Message) {
	s.messages = append(s.messages, message)
}

func (s *Stats) MessageCount() int {
	return len(s.messages)
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
	name := string(user.Name)
	if s.users[name] == nil {
		fmt.Printf("Adding %s to users\n", user.Name)
		s.users[name] = user
	} else {
		fmt.Printf("Already have user\n")
	}
}

func (s *Stats) GetUser(name []byte) *User {
	u := s.users[string(name)]
	if u != nil {
		return u
	} else {
		return NewUser(name, nil)
	}
}

func (s *Stats) ListChannels() {
	fmt.Printf("\nListing Channels:\n")
	for key, channel := range s.channels {
		fmt.Printf("Channel (%s) Name: %s\n", key, channel.GetName())
	}
}

func (s *Stats) ListUsers() {
	fmt.Printf("\nListing Users:\n")
	for _, user := range s.users {
		fmt.Printf("User: %s - Mask: %s\n", user.Name, user.Hostmask)
	}
}

func (s *Stats) GetChannel(name string) *Channel {
	fmt.Printf("looking for %s in channels\n", name)
	s.ListChannels()
	fmt.Printf("end of channels")

	channel := s.channels["deviate2"]
	if channel == nil {
		fmt.Printf("channel is _REALLY_ nil\n")
	}
	fmt.Printf("name: %s\n", channel.Name)

	return channel
}

func (s *Stats) HasChannelByName(name []byte) bool {
	channel_name := string(name)
	channel := s.channels[channel_name]

	return channel != nil
}

func (s *Stats) HasChannelByChannel(channel *Channel) bool {
	return s.HasChannelByName(channel.Name)
}
