package stats

import (
	"regexp"
	"strings"
)

var urlRegex *regexp.Regexp

func init() {
	urlRegex = regexp.MustCompile(`^(http|https):\/\/|[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,6}(:[0-9]{1,5})?(\/.*)?$`)
}

type urls map[string]uint

// NewURLs initializes the urls map.
func NewURLs() urls {
	return make(map[string]uint)
}

// addMessage looks for a url in the message and increments the appropriate
// entry in the urls map.
func (u urls) addMessage(m *Message) {
	words := strings.Split(m.Message, " ")
	for _, w := range words {
		if urlRegex.FindStringSubmatch(w) != nil {
			u[w]++
		}
	}
}
