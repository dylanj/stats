package stats

import "fmt"

type User struct {
	Name     string
	Hostmask string
	Messages []*Message
}

func (u *User) String() string {
  return fmt.Sprintf("User: %s Messages:(%d)", u.Name, len(u.Messages))
}

func NewUser(name string, hostmask string) *User {
	user := User{
		Name:     name,
		Hostmask: hostmask,
		Messages: make([]*Message, 0, 10),
	}

	return &user
}

func (u *User) AddMessage(m []byte, c *Channel) *Message {
	message := NewMessage(m)
	message.User = u
	message.Channel = c

	c.AddMessage(message)

	u.Messages = append(u.Messages, message)

	return message
}

func (u *User) ListMessages() {
	for _, m := range u.Messages {
		fmt.Printf("%s", m)
	}
}
