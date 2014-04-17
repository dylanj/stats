package stats

// import "container/list"
//import "fmt"

const NameLength = 32
const HostmaskLength = 128

type User struct {
	Name     []byte
	Hostmask []byte
	Messages []*Message
}

func NewUser(name []byte, hostmask []byte) *User {
	user := User{
		Name:     make([]byte, NameLength),
		Hostmask: make([]byte, HostmaskLength),
		Messages: make([]*Message, 5),
	}

	copy(user.Name, name)
	copy(user.Hostmask, hostmask)

	return &user
}

func (u *User) AddMessage(m []byte, c *Channel) *Message {
	message := NewMessage(m)
	message.User = u
	message.Channel = c

	u.Messages = append(u.Messages, message)

	return message
}

func (u *User) ListMessages() {
	for _, val := range u.Messages {
		val.Print()
	}
}
