package stats

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Channel struct {
	HourlyChart
	URLCounter
	WordCounter
	SwearCounter
	EmoticonCounter
	Quotes quotes
	ConsecutiveLines
	QuestionsCount
	ExclamationsCount
	AllCapsCount

	ID         uint
	Name       string
	Topic      string
	JoinCount  uint
	PartCount  uint
	UserIDs    map[uint]struct{}
	MessageIDs []uint
	NetworkID  uint

	TopConsecutiveLines TopTokenArray
	LastActive          time.Time
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

		URLCounter:       NewURLCounter(),
		WordCounter:      NewWordCounter(),
		SwearCounter:     NewSwearCounter(),
		EmoticonCounter:  NewEmoticonCounter(),
		ConsecutiveLines: NewConsecutiveLines(),
	}
}

// String returns a the name of the channel and the number of messages inside.
func (c *Channel) String() string {
	return fmt.Sprintf("Channel: %s, Messages: %d", c.Name, len(c.MessageIDs))
}

// AddMessageID adds a message id to the list of message ids.
func (c *Channel) addMessage(m *Message, u *User) {
	c.MessageIDs = append(c.MessageIDs, m.ID)

	c.addUserID(m.UserID)

	if m.Kind == Msg {
		c.HourlyChart.addMessage(m)
		c.Quotes.addMessage(m)
		c.URLCounter.addMessage(m)
		c.WordCounter.addMessage(m)
		c.SwearCounter.addMessage(m)
		c.EmoticonCounter.addMessage(m)
		c.ConsecutiveLines.addMessage(m, u)
		c.QuestionsCount.addMessage(m)
		c.ExclamationsCount.addMessage(m)
		c.AllCapsCount.addMessage(m)
	}

	c.LastActive = m.Date
}

// AddUserID
func (c *Channel) addUserID(id uint) {
	c.UserIDs[id] = struct{}{}
}

// addKick
func (c *Channel) addKick(stats *Stats, message *Message) {
	network := stats.Networks[c.NetworkID]

	targetName := strings.ToLower(strings.Split(message.Message, " ")[0])
	kickerID := message.UserID

	kicker := stats.Users[kickerID]
	kicker.KickCounters.Sent++

	if target, ok := network.users[targetName]; ok {
		target.KickCounters.Received++
	}
}

var slapsRegex = regexp.MustCompile(`^slaps\s(\w+) around a bit with a large trout`)

// addAction
func (c *Channel) addAction(stats *Stats, message *Message) {
	network := stats.Networks[c.NetworkID]

	if m := slapsRegex.FindStringSubmatch(message.Message); m != nil {
		receiver := network.users[strings.ToLower(m[1])]
		sender := stats.Users[message.UserID]
		c.addSlap(sender, receiver)
	}
}

// addSlap
func (c *Channel) addSlap(sender *User, receiver *User) {
	sender.SlapCounters.Sent++

	if receiver != nil {
		receiver.SlapCounters.Received++
	}
}
