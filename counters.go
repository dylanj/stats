package stats

import "strings"

type QuestionsCount uint64
type ExclamationsCount uint64
type AllCapsCount uint64
type BasicTextCounters struct {
	Words   uint64
	Letters uint64
	Lines   uint64
}
type SendRecvCounter struct {
	Sent     uint64
	Received uint64
}
type ModeCounter struct {
	Ops       uint64
	Deops     uint64
	Voices    uint64
	Devoices  uint64
	Halfops   uint64
	Dehalfops uint64
	Bans      uint64
	Unbans    uint64
}

// addMessage
func (m *ModeCounter) addMessage(message *Message) {
	var positive = true
	for _, c := range message.Message {
		switch c {
		case '+':
			positive = true
			continue
		case '-':
			positive = false
		case 'o': // op
			if positive {
				m.Ops++
			} else {
				m.Deops++
			}
		case 'v': // voice
			if positive {
				m.Voices++
			} else {
				m.Devoices++
			}
		case 'h': // halfop
			if positive {
				m.Halfops++
			} else {
				m.Dehalfops++
			}
		case 'b': // ban
			if positive {
				m.Bans++
			} else {
				m.Unbans++
			}
		}
	}
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

func countSuffixes(message string, suffix string) int {
	count := 0
	words := strings.Fields(message)

	for _, word := range words {
		if strings.HasSuffix(word, suffix) {
			count++
		}
	}

	return count
}

func (a *AllCapsCount) addMessage(message *Message) {
	hasCapitalChar := false

	for _, c := range message.Message {
		if c > 'A' && c < 'Z' {
			hasCapitalChar = true
		}

		if c > 'a' && c < 'z' {
			return
		}
	}

	if hasCapitalChar {
		*a++
	}
}

func (q *QuestionsCount) addMessage(message *Message) {
	*q += QuestionsCount(countSuffixes(message.Message, "?"))
}

func (e *ExclamationsCount) addMessage(message *Message) {
	*e += ExclamationsCount(countSuffixes(message.Message, "!"))
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
