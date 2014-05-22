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
	}

	return &user
}

// newChannelUser
func (u *User) addChannelUser(channel string) *User {
	cu := NewUser(u.ID, u.NetworkID, u.Nick)
	u.ChannelUsers[channel] = cu
	return cu
}

func (u *User) addMessage(m *Message) {
	u.MessageIDs = append(u.MessageIDs, m.ID)

	if m.Kind == Msg {
		u.HourlyChart.addMessage(m)
		u.Quotes.addMessage(m)
		u.WordCounter.addMessage(m)
		u.SwearCounter.addMessage(m)
		u.EmoticonCounter.addMessage(m)
		u.BasicTextCounters.addMessage(m)
		u.QuestionsCount.addMessage(m)
		u.ExclamationsCount.addMessage(m)
		u.AllCapsCount.addMessage(m)
	}

	if m.Kind == Mode {
		u.ModeCounters.addMessage(m)
	}

	u.LastSeen = m.Date
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s, Messages: %d", u.Nick, len(u.MessageIDs))
}
