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

func newChannel(id uint, network *Network, name string) *Channel {
	return &Channel{
		ID:         id,
		Name:       name,
		JoinCount:  0,
		PartCount:  0,
		UserIDs:    make([]uint, 0),
		MessageIDs: make([]uint, 0),
		NetworkID:  network.ID,
	}
}

// String returns a the name of the channel and the number of messages inside.
func (c *Channel) String() string {
	return fmt.Sprintf("Channel: %s Messages:(%d)", c.Name, len(c.MessageIDs))
}

// AddMessageID adds a message id to the list of message ids.
func (c *Channel) addMessageID(mid uint) {
	c.MessageIDs = append(c.MessageIDs, mid)
}

//188097
