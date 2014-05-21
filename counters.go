package stats

import "strings"

type QuestionsCount uint
type ExclamationsCount uint
type AllCapsCount uint
type BasicTextCounters struct {
	Words   uint
	Letters uint
	Lines   uint
}
type SendRecvCounters struct {
	Sent     uint
	Received uint
}
type ModeCounters struct {
	Ops       uint
	Deops     uint
	Voices    uint
	Devoices  uint
	Halfops   uint
	Dehalfops uint
	Bans      uint
	Unbans    uint
}

// addMessage
func (m *ModeCounters) addMessage(message *Message) {
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
	c.Letters += uint(len(letters))
	c.Words += uint(len(words))
	c.Lines++
}
