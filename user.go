package stats

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	HourlyChart
	Quotes quotes
	WordCounter

	ID         uint
	Nick       string
	Hostmask   string
	NetworkID  uint
	MessageIDs []uint

	LastSeen time.Time
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s Messages:(%d)", u.Nick, len(u.MessageIDs))
}

func NewUser(id uint, network *Network, nick string) *User {
	user := User{
		ID:         id,
		Nick:       nick,
		NetworkID:  network.ID,
		MessageIDs: make([]uint, 0),

		WordCounter: NewWordCounter(),
	}

	return &user
}

func (u *User) addMessage(m *Message) {
	u.MessageIDs = append(u.MessageIDs, m.ID)

	u.HourlyChart.addMessage(m)
	u.Quotes.addMessage(m)

	u.LastSeen = m.Date
}

func (u *User) RandomMessageID() uint {
	count := len(u.MessageIDs)
	id := rand.Intn(count)
	return u.MessageIDs[id]
}
