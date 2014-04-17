package stats

import "time"
import "fmt"

const MessageLength = 512

type Message struct {
	Date    time.Time
	User    *User
	Channel *Channel
	Message []byte
}

func NewMessage(message []byte) *Message {
  m := Message{
    Message: make([]byte, MessageLength),
  }

  copy(m.Message, message)

  return &m
}

func (m *Message) Print() {
  fmt.Printf("(%s)\n", m.User.Name)
  fmt.Printf("(%s)\n", m.Message)
  fmt.Printf("(%s)\n", m.Channel.Name)
}
