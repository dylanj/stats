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

func (m *Message) String() string {
  return fmt.Sprintf("%s (%s): %s\n", m.User.Name, m.Channel.Name, m.Message)
}
