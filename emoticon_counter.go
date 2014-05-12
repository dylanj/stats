package stats

import "strings"

var emoticons = map[string]struct{}{
	":D":  struct{}{},
	";D":  struct{}{},
	":P":  struct{}{},
	":p":  struct{}{},
	":c":  struct{}{},
	"XD":  struct{}{},
	":)":  struct{}{},
	":(":  struct{}{},
	":>":  struct{}{},
	":<":  struct{}{},
	":-)": struct{}{},
	":-(": struct{}{},
	";)":  struct{}{},
	":'(": struct{}{},
}

type EmoticonCounter struct {
	TokenCounter
}

func NewEmoticonCounter() EmoticonCounter {
	return EmoticonCounter{
		NewTokenCounter(),
	}
}

func (s *EmoticonCounter) addMessage(message *Message) {
	words := strings.Fields(message.Message)

	for _, word := range words {
		if _, ok := emoticons[word]; ok {
			s.addToken(word)
		}
	}
}

// TopEmoticon
func (s *EmoticonCounter) TopEmoticon() TopToken {
	return s.TokenCounter.Top[0]
}
