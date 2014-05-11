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

// WordsPerLine returns the words per line.
func (c *BasicTextCounters) WordsPerLine() float64 {
	if c.lines == 0 {
		return 0
	}

	return float64(c.words) / float64(c.lines)
}

// LettersPerLine returns the letters per line.
func (c *BasicTextCounters) LettersPerLine() float64 {
	if c.lines == 0 {
		return 0
	}

	return float64(c.letters) / float64(c.lines)
}
