package stats

import "fmt"
import "bytes"
import "encoding/gob"
import "log"
import "io/ioutil"

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

func (s *Stats) AddUser(u *User) {
	id := s.UserIDCount
	s.UserIDCount++

	u.ID = id
	s.Users[id] = u
}

func (s *Stats) AddMessage(m *Message) {
	id := s.MessageIDCount
	s.MessageIDCount++

	channel := s.Channels[m.ChannelID]
	user := s.Users[m.UserID]
	network := s.Networks[channel.ID]

	m.ID = id

	s.Messages[id] = m

	channel.AddMessageID(id)
	network.AddMessageID(id)
	user.AddMessageID(id)
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

func (s *Stats) GetNetwork(name string) *Network {
	id, ok := s.network_id_by_name[name]
	if ok {
		return s.Networks[id]
	} else {
		return s.NewNetwork(name)
	}
}

func (s *Stats) NewNetwork(name string) *Network {
	id := s.NetworkIDCount
	s.NetworkIDCount++

	network := &Network{
		Name:       name,
		ID:         id,
		stats:      s,
		ChannelIDs: make([]uint, 10),
		UserIDs:    make([]uint, 10),
		MessageIDs: make([]uint, 10),

		channels: make(map[string]*Channel),
		users:    make(map[string]*User),
	}

	s.Networks[id] = network
	s.network_id_by_name[name] = id

	return network
}

func (s *Stats) ImportData(filename string) *Stats {
	return nil
}

func (s *Stats) MessageCount() int {
	return len(s.Messages)
}

func (s *Stats) ListChannels() {
	fmt.Printf("\nListing Channels:\n")
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

func (s *Stats) ExportData() {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(s)

	if err != nil {
		log.Fatal("encode error:", err)
	}

	ioutil.WriteFile("data.db", buffer.Bytes(), 0x644)
}
