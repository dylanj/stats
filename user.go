package stats

import "fmt"

type User struct {
	ID         uint
	Name       string
	Hostmask   string
	NetworkID  uint
	MessageIDs []uint
	JoinCount  uint
	PartCount  uint
}

func (u *User) String() string {
	return fmt.Sprintf("User: %s Messages:(%d)", u.Name, len(u.MessageIDs))
}

func NewUser(name string, hostmask string) *User {
	user := User{
		Name:       string(name),
		Hostmask:   hostmask,
		MessageIDs: make([]uint, 0),
		JoinCount:  0,
		PartCount:  0,
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
