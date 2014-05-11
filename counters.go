package stats

import "strings"

type BasicTextCounters struct {
	words   uint64
	letters uint64
	lines   uint64
}

// addMessage
func (c *BasicTextCounters) addMessage(message *Message) {
	words := strings.Fields(message.Message)
	letters := strings.Replace(message.Message, " ", "", -1)

	// maybe use a regex to filter out ^a-z
	c.letters += uint64(len(letters))
	c.words += uint64(len(words))
	c.lines++
}
