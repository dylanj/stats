package stats

import "time"
import "fmt"

type Message struct {
	Date    time.Time
	User    *User
	Channel *Channel
	Message string
}

func NewMessage(message []byte) *Message {
	m := Message{
		Message: string(message),
	}

	return &m
}

func (m *Message) Print() {
	fmt.Printf("(%s)\n", m.User.Name)
	fmt.Printf("(%s)\n", m.Message)
	fmt.Printf("(%s)\n", m.Channel.Name)
}
