package stats

import "fmt"

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
		Nick:       string(nick),
		NetworkID:  network.ID,
		MessageIDs: make([]uint, 0),
	}

	return &user
}

func (u *User) AddMessageID(m_id uint) {
	u.MessageIDs = append(u.MessageIDs, m_id)
}

func (u *User) MessageCount() int {
	return len(u.MessageIDs)
}

func (u *User) ListMessages() {
	for id, m := range u.MessageIDs {
		fmt.Printf("[%d] %s", id, m)
	}
}
