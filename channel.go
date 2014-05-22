package stats

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Channel struct {
	HourlyChart
	LastTopics
	URLCounter
	WordCounter
	SwearCounter
	EmoticonCounter
	ConsecutiveLines
	QuestionsCount
	ExclamationsCount
	AllCapsCount
	NickReferences

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
	Quotes              quotes
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
		LastTopics:       NewLastTopics(),
		NickReferences:   make(NickReferences),
	}
}

// String returns a the name of the channel and the number of messages inside.
func (c *Channel) String() string {
	return fmt.Sprintf("Channel: %s, Messages: %d", c.Name, len(c.MessageIDs))
}

// AddMessageID adds a message id to the list of message ids.
func (c *Channel) addMessage(network *Network, message *Message, user *User) {
	c.MessageIDs = append(c.MessageIDs, message.ID)

	c.addUserID(message.UserID)

	if message.Kind == Msg {
		c.HourlyChart.addMessage(message)
		c.Quotes.addMessage(message)
		c.URLCounter.addMessage(message)
		c.WordCounter.addMessage(message)
		c.SwearCounter.addMessage(message)
		c.EmoticonCounter.addMessage(message)
		c.ConsecutiveLines.addMessage(message, user)
		c.QuestionsCount.addMessage(message)
		c.ExclamationsCount.addMessage(message)
		c.AllCapsCount.addMessage(message)
		c.NickReferences.addMessage(network, c, message)
	}

	if message.Kind == Topic {
		c.LastTopics.addMessage(message)
	}

	c.LastActive = message.Date
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
