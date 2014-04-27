package stats

import "fmt"

type Channel struct {
	ID         uint
	Name       string
	Topic      string
	JoinCount  uint
	PartCount  uint
	UserIDs    []uint
	MessageIDs []uint
	NetworkID  uint
}

func (c *Channel) GetName() string {
	return string(c.Name)
}

func (c *Channel) String() string {
	return fmt.Sprintf("Channel: %s Messages:(%d)", c.Name, len(c.MessageIDs))
}

func (c *Channel) AddMessageID(m_id uint) {
	c.MessageIDs = append(c.MessageIDs, m_id)
}

//188097
