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

// GetUser returns a user by the name given if it exists, if it doesn't
// it creates a new User and returns it.
func (n *Network) GetUser(name string) *User {
	var user, ok = n.users[name]

	if ok {
		return user
	} else {
		return n.AddUser(name)
	}
}

// GetChannel returns a channel by the name given if it exists, if it doesn't
// it creates a new Channel and returns it.
func (n *Network) GetChannel(name string) *Channel {
	var channel, ok = n.channels[name]

	if ok {
		return channel
	} else {
		return n.AddChannel(name)
	}
}

// AddUser adds a user by the name to the Network.
func (n *Network) AddUser(name string) *User {
	user := NewUser(name, "")

	n.users[name] = user

	n.stats.AddUser(user)

	return user
}

// Adds a channel by the name to the Network.
func (n *Network) AddChannel(name string) *Channel {
	id := n.stats.ChannelIDCount
	n.stats.ChannelIDCount++

	channel := NewChannel(name, n)
	channel.ID = id

	n.ChannelIDs = append(n.ChannelIDs, channel.ID)
	n.stats.Channels[id] = channel
	n.channels[name] = channel

	return channel
}

func (n *Network) AddMessage(m *Message) {
	n.stats.AddMessage(m)
	n.AddMessageID(m.ID)
}

func (n *Network) AddMessageID(m_id uint) {
	n.MessageIDs = append(n.MessageIDs, m_id)
}

func (n *Network) ChannelCount() int {
	return len(n.ChannelIDs)
}

func (n *Network) UserCount() int {
	return len(n.UserIDs)
}
