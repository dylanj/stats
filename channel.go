package stats

import "fmt"

type Channel struct {
	Name     string
	Topic    string
	Joins    int
	Parts    int
	Users    []*User
	Messages []*Message
}

func NewChannel(name []byte) *Channel {
	channel := Channel{
		Name:  string(name),
		Topic: "",
		Joins: 0,
		Parts: 0,
	}

	return &channel
}

func (c *Channel) GetName() string {
	return string(c.Name)
}

func (c *Channel) String() string {
  return fmt.Sprintf("Channel: %s Messages:(%d)", c.Name, len(c.Messages))
}

func (c *Channel) AddMessage(m *Message) {
	c.Messages = append(c.Messages, m)
}
