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
