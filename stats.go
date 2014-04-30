package stats

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Stats struct {
	Channels map[uint]*Channel
	Messages map[uint]*Message
	Networks map[uint]*Network
	Users    map[uint]*User

	networkByName map[string]*Network
	channelByName map[string]*Channel
	userByName    map[string]*User

	NetworkIDCount uint
	MessageIDCount uint
	ChannelIDCount uint
	UserIDCount    uint
}

// NewStats initializes a Stats struct.
func NewStats() *Stats {
	// load from stats.db
	return &Stats{
		Channels: make(map[uint]*Channel),
		Messages: make(map[uint]*Message),
		Networks: make(map[uint]*Network),
		Users:    make(map[uint]*User),

		networkByName: make(map[string]*Network),
		channelByName: make(map[string]*Channel),
		userByName:    make(map[string]*User),

		NetworkIDCount: 0,
		MessageIDCount: 0,
		ChannelIDCount: 0,
		UserIDCount:    0,
	}
}

// AddMessage adds a message to the stats.
func (s *Stats) AddMessage(kind MsgKind, network string, channel string, hostmask string, date time.Time, message string) {

	n := s.getNetwork(network)

	c := s.getChannel(n, channel)
	u := s.getUser(n, hostmask)

	s.addMessage(kind, n, c, u, date, message)
}

// HourlyChart returns an array of integers with the number of messages said each hour
// in the given channel on the given network.
// The index of the array is the hour
func (s *Stats) HourlyChart(network string, channel string) ([24]int, bool) {
	var chart [24]int

	n, ok := s.networkByName[network]
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
	n.addMessageID(id)
	u.addMessageID(id)
}

func (s *Stats) addChannel(n *Network, name string) *Channel {
	id := n.stats.ChannelIDCount
	n.stats.ChannelIDCount++

	c := newChannel(id, n, name)

	s.channelByName[c.Name] = c
	s.Channels[c.ID] = c

	n.addChannel(c)

	return c
}

func (s *Stats) addUser(n *Network, nick string) *User {
	id := s.UserIDCount
	s.UserIDCount++

	u := NewUser(id, n, nick)

	s.userByName[u.Nick] = u
	s.Users[id] = NewUser(id, n, nick)

	n.addUser(u)

	return u
}

func (s *Stats) getUser(n *Network, name string) *User {
	if u, ok := s.userByName[name]; ok {
		return u
	} else {
		return s.addUser(n, name)
	}
}

func (s *Stats) getChannel(n *Network, name string) *Channel {
	if c, ok := s.channelByName[name]; ok {
		return c
	} else {
		return s.addChannel(n, name)
	}
}

func (s *Stats) getNetwork(name string) *Network {
	if n, ok := s.networkByName[name]; ok {
		return n
	} else {
		return s.addNetwork(name)
	}
}

func (s *Stats) addNetwork(name string) *Network {
	id := s.NetworkIDCount
	s.NetworkIDCount++

	n := &Network{
		Name:       name,
		ID:         id,
		stats:      s,
		ChannelIDs: make([]uint, 0),
		UserIDs:    make([]uint, 0),
		MessageIDs: make([]uint, 0),

		channels: make(map[string]*Channel),
		users:    make(map[string]*User),
	}

	s.Networks[id] = n
	s.networkByName[name] = n

	return n
}

// Save writes the statistics to data.db.
func (s *Stats) Save() {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(s)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	ioutil.WriteFile("data.db", buffer.Bytes(), 0644)
}

// loadDatabase reads data.db and populates a Stats struct.
func loadDatabase() *Stats {
	file, _ := os.Open("data.db")

	decoder := gob.NewDecoder(file)
	var stats Stats

	decoder.Decode(&stats)

	return &stats
}
