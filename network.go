package stats

import (
	"fmt"
	"strings"
)

type Network struct {
	HourlyChart
	Quotes quotes

	ID         uint
	Name       string
	ChannelIDs []uint
	UserIDs    []uint
	MessageIDs []uint

	channels map[string]*Channel
	users    map[string]*User

	stats *Stats
}

func (n *Network) addChannel(c *Channel) {
	n.ChannelIDs = append(n.ChannelIDs, c.ID)
	n.channels[strings.ToLower(c.Name)] = c
}

func (n *Network) addUser(u *User) {
	n.UserIDs = append(n.UserIDs, u.ID)
	n.users[strings.ToLower(u.Nick)] = u
}

func (n *Network) getUser(name string) *User {
	return &User{}
}

func (n *Network) addMessage(m *Message) {
	n.MessageIDs = append(n.MessageIDs, m.ID)

	n.HourlyChart.addMessage(m)
	n.Quotes.addMessage(m)
}

// buildIndexes builds the internal maps that relate data
func (n *Network) buildIndexes(s *Stats) {
	n.channels = make(map[string]*Channel)
	n.users = make(map[string]*User)
	n.stats = s

	for _, cID := range n.ChannelIDs {
		c := n.stats.Channels[cID]

		n.channels[c.Name] = c
	}

	for _, uID := range n.UserIDs {
		u := n.stats.Users[uID]

		n.users[u.Nick] = u
	}
}

// String returns a the name of the channel and some basic stats.
func (n *Network) String() string {
	return fmt.Sprintf("Network: %s, Channels: %d, Messages: %d", n.Name, len(n.ChannelIDs), len(n.MessageIDs))
}
