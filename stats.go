package stats

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

type Stats struct {
	Channels map[uint]*Channel
	Messages map[uint]*Message
	Networks map[uint]*Network
	Users    map[uint]*User

	network_id_by_name map[string]uint
	channel_id_by_name map[string]uint
	user_id_by_name    map[string]uint

	NetworkIDCount uint
	MessageIDCount uint
	ChannelIDCount uint
	UserIDCount    uint
}

func (s *Stats) addChannel(n *Network, name string) *Channel {
	id := n.stats.ChannelIDCount
	n.stats.ChannelIDCount++

	c := newChannel(id, n, name)

	s.channel_id_by_name[c.Name] = c.ID
	s.Channels[c.ID] = c

	n.AddChannel(c)

	return c
}

func (s *Stats) addUser(n *Network, nick string) *User {
	id := s.UserIDCount
	s.UserIDCount++

	u := NewUser(id, n, nick)

	s.user_id_by_name[u.Nick] = u.ID
	s.Users[id] = NewUser(id, n, nick)

	n.AddUser(u)

	return u
}

func (s *Stats) AddMessage(kind MsgKind, network string, channel string, hostmask string, date time.Time, message string) {

	n := s.getNetwork(network)

	c := s.getChannel(n, channel)
	u := s.getUser(n, hostmask)

	s.addMessage(kind, n, c, u, date, message)
}

func (s *Stats) addMessage(k MsgKind, n *Network, c *Channel, u *User, d time.Time, m string) {
	id := s.MessageIDCount
	s.MessageIDCount++

	message := &Message{
		ID:        id,
		Date:      d,
		UserID:    u.ID,
		ChannelID: c.ID,
		Message:   m,
		Kind:      k,
	}

	s.Messages[id] = message

	c.addMessageID(id)
	n.AddMessageID(id)
	u.AddMessageID(id)
}

func NewStats() *Stats {
	// load from stats.db
	return &Stats{
		Channels: make(map[uint]*Channel),
		Messages: make(map[uint]*Message),
		Networks: make(map[uint]*Network),
		Users:    make(map[uint]*User),

		network_id_by_name: make(map[string]uint),
		channel_id_by_name: make(map[string]uint),
		user_id_by_name:    make(map[string]uint),

		NetworkIDCount: 0,
		MessageIDCount: 0,
		ChannelIDCount: 0,
		UserIDCount:    0,
	}
}

func (s *Stats) GetNetworkByID(id uint) *Network {
	return s.Networks[id]
}

func (s *Stats) getUser(n *Network, name string) *User {
	if id, ok := s.user_id_by_name[name]; ok {
		return s.Users[id]
	} else {
		return s.addUser(n, name)
	}
}

func (s *Stats) getChannel(n *Network, name string) *Channel {
	if id, ok := s.channel_id_by_name[name]; ok {
		return s.Channels[id]
	} else {
		return s.addChannel(n, name)
	}
}

func (s *Stats) getNetwork(name string) *Network {
	if id, ok := s.network_id_by_name[name]; ok {
		return s.Networks[id]
	} else {
		return s.addNetwork(name)
	}
}

func (s *Stats) addNetwork(name string) *Network {
	id := s.NetworkIDCount
	s.NetworkIDCount++

	network := &Network{
		Name:       name,
		ID:         id,
		stats:      s,
		ChannelIDs: make([]uint, 0),
		UserIDs:    make([]uint, 0),
		MessageIDs: make([]uint, 0),

		channels: make(map[string]*Channel),
		users:    make(map[string]*User),
	}

	s.Networks[id] = network
	s.network_id_by_name[name] = id

	return network
}

func (s *Stats) MessageCount() int {
	return len(s.Messages)
}

func (s *Stats) ListChannels() {
	fmt.Printf("\nListing Channels: (%d)\n", len(s.Channels))
	for id, c := range s.Channels {
		fmt.Printf("[%d] %s\n", id, c)
	}
}

func (s *Stats) ListUsers() {
	fmt.Printf("\nListing Users:\n")
	for id, u := range s.Users {
		fmt.Printf("[%d] %s\n", id, u)
	}
}

func (s *Stats) Information() {
	s.ListChannels()
	s.ListUsers()

	fmt.Printf("Number of messages in stats: %d\n", s.MessageCount())
}

func (s *Stats) RandomMessageForUser(u *User) *Message {
	count := len(u.MessageIDs)
	var m *Message

	for i := 0; i < 3; i++ {
		id := u.MessageIDs[rand.Intn(count)]
		m = s.Messages[id]
		if len(m.Message) > 0 {
			return m
		}
	}

	return m
}

func (s *Stats) ExportData() {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(s)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	ioutil.WriteFile("data.db", buffer.Bytes(), 0644)
}

func ImportData() *Stats {
	file, _ := os.Open("data.db")

	decoder := gob.NewDecoder(file)
	var stats Stats

	decoder.Decode(&stats)

	return &stats
}

// HourlyChart returns an array of integers with the number of messages said each hour
// in the given channel on the given network.
// The index of the array is the hour
func (s *Stats) HourlyChart(network string, channel string) ([24]int, bool) {
	var chart [24]int

	nID, ok := s.network_id_by_name[network]
	if !ok {
		return chart, false
	}

	n, ok := s.Networks[nID]
	if !ok {
		return chart, false
	}

	c, ok := n.channels[channel]
	if !ok {
		return chart, false
	}

	for _, id := range c.MessageIDs {
		hour := s.Messages[id].Date.Hour()
		chart[hour]++
	}

	return chart, true
}
