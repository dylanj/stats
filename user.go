package stats

import (
	"fmt"
	"time"
)

type User struct {
	HourlyChart
	Quotes quotes
	WordCounter
	SwearCounter
	EmoticonCounter
	BasicTextCounters
	QuestionsCount
	ExclamationsCount
	AllCapsCount

	ID         uint
	Nick       string
	Hostmask   string
	NetworkID  uint
	MessageIDs []uint

	LastSeen       time.Time
	MaxConsecutive uint
}

func NewUser(id uint, network *Network, nick string) *User {
	user := User{
		ID:         id,
		Nick:       nick,
		NetworkID:  network.ID,
		MessageIDs: make([]uint, 0),

		WordCounter:     NewWordCounter(),
		SwearCounter:    NewSwearCounter(),
		EmoticonCounter: NewEmoticonCounter(),
	}

	return &user
}

func (u *User) addMessage(m *Message) {
	u.MessageIDs = append(u.MessageIDs, m.ID)

	u.HourlyChart.addMessage(m)
	u.Quotes.addMessage(m)
	u.WordCounter.addMessage(m)
	u.SwearCounter.addMessage(m)
	u.EmoticonCounter.addMessage(m)
	u.BasicTextCounters.addMessage(m)

	u.LastSeen = m.Date
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s, Messages: %d", u.Nick, len(u.MessageIDs))
}
