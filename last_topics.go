package stats

const maxTopics = 5

type LastTopics struct {
	Topics []*Message
}

// NewLastTopics
func NewLastTopics() LastTopics {
	return LastTopics{
		Topics: make([]*Message, 0, maxTopics),
	}
}

// addMessage
func (l *LastTopics) addMessage(message *Message) {
	if len(l.Topics) >= maxTopics {
		l.Topics = l.Topics[1:maxTopics]
	}
	l.Topics = append(l.Topics, message)
}
