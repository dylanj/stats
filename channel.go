package stats

import "bytes"
import "strconv"

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
	var buffer bytes.Buffer

	buffer.WriteString("Channel: ")
	buffer.WriteString(c.Name)
	buffer.WriteString(" Messages: (")
	buffer.WriteString(strconv.Itoa(len(c.Messages)))
	buffer.WriteString(")")

	return buffer.String()
}

func (c *Channel) AddMessage(m *Message) {
	c.Messages = append(c.Messages, m)
}
