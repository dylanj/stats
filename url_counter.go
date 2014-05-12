package stats

import "regexp"

var tokenRegexURL = regexp.MustCompile(`(?:(?:[^\s:/?#]+)://|www\.)(?:[^\s/?#]+\.)*(?:[A-Za-z0-9][^\s/?#]*\.[A-Za-z]{2,6})(?:/[^\s#\?]+)?/?(?:\?[^\s#]*)?(?:#[^\s]*)?`)

type URLCounter struct {
	TokenCounter
}

func NewURLCounter() URLCounter {
	return URLCounter{
		NewTokenCounter(),
	}
}

func (u *URLCounter) addMessage(m *Message) {
	results := tokenRegexURL.FindAllStringSubmatch(m.Message, -1)
	for _, v := range results {
		u.TokenCounter.addToken(v[0])
	}
}
