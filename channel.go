package stats

import (
	"fmt"
	"time"
)

type Channel struct {
	HourlyChart
	Quotes quotes
	URLs   urls

	ID         uint
	Name       string
	Topic      string
	JoinCount  uint
	PartCount  uint
	UserIDs    map[uint]struct{}
	MessageIDs []uint
	NetworkID  uint

	LastActive time.Time
}

func newChannel(id uint, network *Network, name string) *Channel {
	return &Channel{
		ID:         id,
		Name:       name,
		JoinCount:  0,
		PartCount:  0,
		UserIDs:    make(map[uint]struct{}, 0),
		MessageIDs: make([]uint, 0),
		NetworkID:  network.ID,

		URLs: NewURLs(),
	}
}

// String returns a the name of the channel and the number of messages inside.
func (c *Channel) String() string {
	return fmt.Sprintf("Channel: %s Messages:(%d)", c.Name, len(c.MessageIDs))
}

// AddMessageID adds a message id to the list of message ids.
func (c *Channel) addMessage(m *Message) {
	c.MessageIDs = append(c.MessageIDs, m.ID)

	c.addUserID(m.UserID)

	// stats stuff
	c.HourlyChart.addMessage(m)
	c.Quotes.addMessage(m)
	c.URLs.addMessage(m)

	c.LastActive = m.Date
}

// AddUserID
func (c *Channel) addUserID(id uint) {
	c.UserIDs[id] = struct{}{}
}
