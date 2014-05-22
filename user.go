package stats

import (
	"fmt"
	"time"
)

type User struct {
	HourlyChart
	WordCounter
	SwearCounter
	EmoticonCounter
	QuestionsCount
	ExclamationsCount
	AllCapsCount
	BasicTextCounters
	ModeCounters
	NickReferences

	KickCounters SendRecvCounters
	SlapCounters SendRecvCounters
	Quotes       quotes

	ID           uint
	Nick         string
	Hostmask     string
	NetworkID    uint
	MessageIDs   []uint
	ChannelUsers map[string]*User

	LastSeen       time.Time
	MaxConsecutive uint
}

func NewUser(id uint, networkID uint, nick string) *User {
	user := User{
		ID:           id,
		Nick:         nick,
		NetworkID:    networkID,
		MessageIDs:   make([]uint, 0),
		ChannelUsers: make(map[string]*User),

		WordCounter:     NewWordCounter(),
		SwearCounter:    NewSwearCounter(),
		EmoticonCounter: NewEmoticonCounter(),
		NickReferences:  make(NickReferences),
	}

	return &user
}

// newChannelUser
func (u *User) addChannelUser(channel string) *User {
	cu := NewUser(u.ID, u.NetworkID, u.Nick)
	u.ChannelUsers[channel] = cu
	return cu
}

func (u *User) addMessage(network *Network, channel *Channel, message *Message) {
	u.MessageIDs = append(u.MessageIDs, message.ID)

	if message.Kind == Msg {
		u.HourlyChart.addMessage(message)
		u.Quotes.addMessage(message)
		u.WordCounter.addMessage(message)
		u.SwearCounter.addMessage(message)
		u.EmoticonCounter.addMessage(message)
		u.BasicTextCounters.addMessage(message)
		u.QuestionsCount.addMessage(message)
		u.ExclamationsCount.addMessage(message)
		u.AllCapsCount.addMessage(message)
		u.NickReferences.addMessage(network, channel, message)
	}

	if message.Kind == Mode {
		u.ModeCounters.addMessage(message)
	}

	u.LastSeen = message.Date
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s, Messages: %d", u.Nick, len(u.MessageIDs))
}
