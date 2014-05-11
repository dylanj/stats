package stats

import "strings"

type BasicTextCounters struct {
	Words   uint64
	Letters uint64
	Lines   uint64
}

// addMessage
func (c *BasicTextCounters) addMessage(message *Message) {
	words := strings.Fields(message.Message)
	letters := strings.Replace(message.Message, " ", "", -1)

	// maybe use a regex to filter out ^a-z
	c.Letters += uint64(len(letters))
	c.Words += uint64(len(words))
	c.Lines++
}

// WordsPerLine returns the words per line.
func (c *BasicTextCounters) WordsPerLine() float64 {
	if c.Lines == 0 {
		return 0
	}

	return float64(c.Words) / float64(c.Lines)
}

// LettersPerLine returns the letters per line.
func (c *BasicTextCounters) LettersPerLine() float64 {
	if c.Lines == 0 {
		return 0
	}

	return float64(c.Letters) / float64(c.Lines)
}
