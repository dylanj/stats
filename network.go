package stats

type Network struct {
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
	n.channels[c.Name] = c
}

func (n *Network) addUser(u *User) {
	n.UserIDs = append(n.UserIDs, u.ID)
	n.users[u.Nick] = u
}

func (n *Network) getUser(name string) *User {
	return &User{}
}

func (n *Network) addMessageID(m_id uint) {
	n.MessageIDs = append(n.MessageIDs, m_id)
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
