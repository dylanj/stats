package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/DylanJ/stats"
)

type UserJSON struct {
	ID           uint
	Name         string
	MessageCount int
	Message      string
}

type ChannelStatsJSON struct {
	TopUsers    []*UserJSON
	HourlyChart stats.HourlyChart
	TopURLs     []*TopURL
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
		TopURLs:     TopURLs(channel.URLs, 5),
		TopUsers:    channelStats_TopUsers(s, channel),
	}

	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func channelStats_TopUsers(s *stats.Stats, c *stats.Channel) []*UserJSON {
	var users []*UserJSON
	users = make([]*UserJSON, 0)

	for id, _ := range c.UserIDs {
		if u, ok := s.Users[id]; ok {
			message := s.Messages[u.RandomMessageID()]

			user := &UserJSON{
				ID:           id,
				Name:         u.Nick,
				MessageCount: len(u.MessageIDs),
				Message:      message.Message,
			}

			users = append(users, user)
		}
	}

	sort.Sort(sort.Reverse(ByMessageCount(users)))

	return users
}
