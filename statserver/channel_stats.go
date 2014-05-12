package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sort"

	"github.com/DylanJ/stats"
)

type UserJSON struct {
	ID             uint             `json:"-"`
	Name           string           `json:"name"`
	MessageCount   int              `json:"count"`
	Message        string           `json:"random"`
	HourlyChart    [24]int          `json:"hourly"`
	VocabularySize int              `json:"vocabulary"`
	TopSwears      []stats.TopToken `json:"swears"`
}

type ChannelStatsJSON struct {
	TopUsers    []*UserJSON       `json:"users"`
	HourlyChart stats.HourlyChart `json:"hourly"`
	TopURLs     []stats.TopToken  `json:"urls"`
	TopWords    []stats.TopToken  `json:"words"`
	TopSwears   []stats.TopToken  `json:"swears"`
}

type ByMessageCount []*UserJSON

func (a ByMessageCount) Len() int           { return len(a) }
func (a ByMessageCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMessageCount) Less(i, j int) bool { return a[i].MessageCount < a[j].MessageCount }

func ChannelStats(w http.ResponseWriter, s *stats.Stats, n, c string) {
	channel := s.GetChannel(n, c)

	if channel == nil {
		// should probably return a json error here and handle that
		// on the client
		fmt.Println("Couldnt load channel (%s,%s)", n, c)
		return
	}

	data := &ChannelStatsJSON{
		HourlyChart: channel.HourlyChart,
		TopURLs:     channel.URLCounter.Top[:15],
		TopWords:    channel.WordCounter.Top,
		TopSwears:   channel.SwearCounter.Top,
		TopUsers:    channelStats_TopUsers(s, channel),
	}

	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func userStats_RandomMessage(s *stats.Stats, u *stats.User) string {
	id := rand.Intn(len(u.MessageIDs))
	return s.MessageIDs[id]
}

func channelStats_TopUsers(s *stats.Stats, c *stats.Channel) []*UserJSON {
	var users []*UserJSON
	users = make([]*UserJSON, 0)

	for id, _ := range c.UserIDs {
		if u, ok := s.Users[id]; ok {
			message := userStats_RandomMessage(s, u)

			user := &UserJSON{
				ID:             id,
				Name:           u.Nick,
				MessageCount:   len(u.MessageIDs),
				Message:        message.Message,
				HourlyChart:    u.HourlyChart,
				VocabularySize: len(u.WordCounter.All),
				TopSwears:      u.SwearCounter.Top,
			}

			users = append(users, user)
		}
	}

	sort.Sort(sort.Reverse(ByMessageCount(users)))

	return users
}
