package stats

import "strings"

type NickReferences map[string]uint

var punctuationReplacer = strings.NewReplacer(
	".", "",
	",", "",
	":", "",
	";", "",
	"!", "",
	"?", "",
	"@", "",
)

func (r NickReferences) addMessage(network *Network, channel *Channel, message *Message) {
	if channel == nil {
		return
	}

	msg := punctuationReplacer.Replace(message.Message)
	msg = strings.ToLower(msg)
	words := strings.Fields(msg)

	for _, word := range words {
		var u *User
		var ok bool
		if u, ok = network.users[word]; !ok {
			continue
		}

		if _, ok = channel.UserIDs[u.ID]; ok {
			r[word]++
		}
	}
}
