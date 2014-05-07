package stats

import (
	"regexp"
	"strings"
)

const topTokenMaxSize = 50

var tokenRegexURL = regexp.MustCompile(`(?:(?:[^\s:/?#]+)://)?(?:[^\s/?#]+\.)*(?:[A-Za-z0-9][^\s/?#]*\.[A-Za-z]{2,6})(?:/[^\s#\?]+)?/?(?:\?[^\s#]*)?(?:#[^\s]*)?`)
var tokenRegexWord = regexp.MustCompile(`^([a-zA-Z]+)[\?!;,\.]?$`)

type URLCounter struct {
	TokenCounter
}

func (u *URLCounter) addMessage(m *Message) {
	results := tokenRegexURL.FindAllStringSubmatch(m.Message, -1)
	for _, v := range results {
		u.TokenCounter.addToken(v[0])
	}
}

type WordCounter struct {
	TokenCounter
}

func (w *WordCounter) addMessage(m *Message) {
	words := strings.Fields(m.Message)
	for _, v := range words {
		if r := tokenRegexWord.FindStringSubmatch(v); r != nil {
			w.TokenCounter.addToken(r[1])
		}
	}
}

type TopToken struct {
	Token string
	Count uint
}

type TokenCounter struct {
	All map[string]uint
	Top []TopToken
}

// NewTokens initializes the Tokens map.
func NewTokenCounter() TokenCounter {
	return TokenCounter{
		All: make(map[string]uint),
		Top: make([]TopToken, 0, topTokenMaxSize),
	}
}

func NewWordCounter() WordCounter {
	return WordCounter{
		NewTokenCounter(),
	}
}
func NewURLCounter() URLCounter {
	return URLCounter{
		NewTokenCounter(),
	}
}

func (tc *TokenCounter) addToken(token string) {
	tc.All[token]++

	count := tc.All[token]

	insertAt := -1
	currentIndex := -1

	// check if token is in top

	if len(tc.Top) == 0 {
		tc.Top = append(tc.Top, TopToken{token, count})
		return
	}

	for i, t := range tc.Top {
		if insertAt == -1 && count > t.Count {
			insertAt = i
		}

		if currentIndex == -1 && token == t.Token {
			currentIndex = i
		}

		if currentIndex != -1 && insertAt != -1 {
			break
		}
	}

	if currentIndex >= 0 {
		tc.Top[currentIndex].Token, tc.Top[insertAt].Token =
			tc.Top[insertAt].Token, tc.Top[currentIndex].Token
		tc.Top[insertAt].Count = count
	} else if len(tc.Top) < topTokenMaxSize {
		tc.Top = append(tc.Top, TopToken{token, count})
	} else if insertAt >= 0 {
		tc.Top[insertAt].Token = token
		tc.Top[insertAt].Count = count
	}
}
