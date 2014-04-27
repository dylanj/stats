package stats

import "time"
import "fmt"

type Message struct {
	ID        uint
	Date      time.Time
	UserID    uint
	ChannelID uint
	Message   string
}

func NewMessage(message []byte) *Message {
	m := Message{
		Message: string(message),
	}

	return &m
}

func (m *Message) String() string {
	//return fmt.Sprintf("%s (%s): %s\n", m.User.Name, m.Channel.Name, m.Message)
	return fmt.Sprintf("hello")
}
