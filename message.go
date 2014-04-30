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

func NewMessageString(message string) *Message {
	return NewMessage([]byte(message))
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

// date hax
// 2013-08-07 16:49:40	-->	dylan (dylan@zqz.ca) has joined #deviate
// 2006-01-02 15:04:04
