package stats

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID         uint
	Nick       string
	Hostmask   string
	NetworkID  uint
	MessageIDs []uint
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
	}

	return &user
}

func (u *User) addMessageID(m_id uint) {
	u.MessageIDs = append(u.MessageIDs, m_id)
}

func (u *User) RandomMessageID() uint {
	count := len(u.MessageIDs)
	id := rand.Intn(count)
	return u.MessageIDs[id]
}
