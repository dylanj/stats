package stats

import (
	"regexp"
	"strings"
)

var swearRegex = regexp.MustCompile(`[[:alpha:]]*(?:fuck|tits|whore|bitch|cunt|pussy|dick|fag|ass|shit|nigger|cock)[[:alpha:]]*`)

type SwearCounter struct {
	TokenCounter
}

func NewSwearCounter() SwearCounter {
	return SwearCounter{
		NewTokenCounter(),
	}
}

func (s *SwearCounter) addMessage(message *Message) {
	words := strings.Fields(message.Message)

	for _, word := range words {
		r := swearRegex.FindStringSubmatch(word)

		if len(r) > 0 {
			s.addToken(r[0])
		}
	}
}
