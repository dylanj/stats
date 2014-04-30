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

func (n *Network) GetUser(name string) *User {
	var user, ok = n.users[name]

	if ok {
		return user
	} else {
		return n.NewUser(name)
	}
}

func (n *Network) GetChannel(name string) *Channel {
	var channel, ok = n.channels[name]

	if ok {
		return channel
	} else {
		return n.NewChannel(name)
	}
}

func (n *Network) NewUser(name string) *User {
	user := NewUser(name, "")

	n.users[name] = user

	n.stats.AddUser(user)

	return user
}

func (n *Network) NewChannel(name string) *Channel {
	id := n.stats.ChannelIDCount
	n.stats.ChannelIDCount++

	channel := &Channel{
		ID:         id,
		Name:       name,
		JoinCount:  0,
		PartCount:  0,
		UserIDs:    make([]uint, 0),
		MessageIDs: make([]uint, 0),
		NetworkID:  n.ID,
	}

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
