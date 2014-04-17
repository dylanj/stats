package stats

const ChannelLength = 32
const TopicLength = 256

type Channel struct {
	Name     []byte
	Topic    []byte
	Joins    int
	Parts    int
	Users    []*User
	Messages []*Message
}

func NewChannel(name []byte) *Channel {
	channel := Channel{
		Name:  make([]byte, ChannelLength),
		Topic: make([]byte, TopicLength),
		Joins: 0,
		Parts: 0,
	}

	copy(channel.Name, name)

	return &channel
}

func (c *Channel) GetName() string {
	return string(c.Name)
}

// func (c *Channel) AddUser(u *User) bool {
//   // s.users = append(u, 
//   // return
// }
