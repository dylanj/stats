package stats

import "fmt"
import "bytes"
import "encoding/gob"
import "log"
import "io/ioutil"

type Stats struct {
	channels map[string]*Channel
	users    map[string]*User
	Messages []*Message
}

func NewStats() *Stats {
	return &Stats{
		channels: make(map[string]*Channel),
		users:    make(map[string]*User),
	}
}

func (s *Stats) AddMessage(message *Message) {
	s.Messages = append(s.Messages, message)
}

func (s *Stats) MessageCount() int {
	return len(s.Messages)
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
	}
}

func (s *Stats) GetUser(name string) *User {
	u, ok := s.users[name]; if ok {
		return u
	} else {
		return NewUser(name, "")
	}
}

func (s *Stats) ListChannels() {
	fmt.Printf("\nListing Channels:\n")
	for _, c := range s.channels {
		fmt.Printf("%s\n", c)
	}
}

func (s *Stats) ListUsers() {
	fmt.Printf("\nListing Users:\n")
	for _, u := range s.users {
		fmt.Printf("%s\n", u)
	}
}

func (s *Stats) GetChannel(name string) *Channel {
	channel := s.channels[name]

	return channel
}

func (s *Stats) HasChannelByName(name string) bool {
	channel := s.channels[name]

	return channel != nil
}

func (s *Stats) HasChannelByChannel(channel *Channel) bool {
	return s.HasChannelByName(channel.Name)
}

func (s *Stats) Information() {
	s.ListChannels()
	s.ListUsers()

	fmt.Printf("Number of messages in stats: %d\n", s.MessageCount())
}

func (s *Stats) ExportData() {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(s.Messages)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	ioutil.WriteFile("data.db", buffer.Bytes(), 0x644)
}
