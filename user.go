package stats

import "bytes"
import "strconv"

type User struct {
	Name     string
	Hostmask string
	Messages []*Message
}

func (u *User) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("User: ")
	buffer.WriteString(u.Name)
	buffer.WriteString(" Messages:(")
	buffer.WriteString(strconv.Itoa(len(u.Messages)))
	buffer.WriteString(")")

	return buffer.String()
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
	for _, val := range u.Messages {
		val.Print()
	}
}
