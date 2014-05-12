package stats

import (
	"regexp"
	"strings"
)

var tokenRegexWord = regexp.MustCompile(`^([a-zA-Z]+)[\?!;,\.]?$`)

type WordCounter struct {
	TokenCounter
}

func NewWordCounter() WordCounter {
	return WordCounter{
		NewTokenCounter(),
	}
}

func (w *WordCounter) addMessage(m *Message) {
	words := strings.Fields(m.Message)
	for _, v := range words {
		if r := tokenRegexWord.FindStringSubmatch(v); r != nil {
			w.TokenCounter.addToken(strings.ToLower(r[1]))
		}
	}
}
