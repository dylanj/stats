package main

import (
	"fmt"
	"sort"

	"github.com/DylanJ/stats"
)

type UserJSON struct {
	ID             uint                    `json:"-"`
	Name           string                  `json:"name"`
	MessageCount   uint                    `json:"count"`
	Message        string                  `json:"random"`
	HourlyChart    [24]int                 `json:"hourly"`
	VocabularySize int                     `json:"vocabulary"`
	TopSwears      []stats.TopToken        `json:"swears"`
	SwearCount     uint                    `json:"swearcount"`
	Vocabulary     []stats.TopToken        `json:"vocab"`
	Emoticons      []stats.TopToken        `json:"emoticons"`
	EmoticonCount  uint                    `json:"emoticoncount"`
	Questions      uint                    `json:"questions"`
	Exclamations   uint                    `json:"exclamations"`
	AllCaps        uint                    `json:"allcaps"`
	SKicks         uint                    `json:"skicks"`
	RKicks         uint                    `json:"rkicks"`
	SSlaps         uint                    `json:"sslaps"`
	RSlaps         uint                    `json:"rslaps"`
	NickReferences map[string]uint         `json:"nickreferences"`
	Modes          stats.ModeCounters      `json:"modes"`
	Basic          stats.BasicTextCounters `json:"basic"`
}

type ChannelStatsJSON struct {
	TopUsers    []*UserJSON       `json:"users"`
	HourlyChart stats.HourlyChart `json:"hourly"`
	TopURLs     []stats.TopToken  `json:"urls"`
	TopWords    []stats.TopToken  `json:"words"`
	TopSwears   []stats.TopToken  `json:"swears"`
	SwearCount  uint              `json:"swearcount"`
}

type ByMessageCount []*UserJSON

func (a ByMessageCount) Len() int           { return len(a) }
func (a ByMessageCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMessageCount) Less(i, j int) bool { return a[i].MessageCount < a[j].MessageCount }

func topUsers(s *stats.Stats, c *stats.Channel) []*UserJSON {
	var users []*UserJSON
	users = make([]*UserJSON, 0)

	for id, _ := range c.UserIDs {
		if u, ok := s.Users[id]; ok {

			fmt.Printf("%#v\n\n\n", u.Quotes)

			user := &UserJSON{
				ID:             id,
				Name:           u.Nick,
				MessageCount:   u.BasicTextCounters.Lines,
				HourlyChart:    u.HourlyChart,
				Vocabulary:     u.WordCounter.Top,
				VocabularySize: len(u.WordCounter.All),
				TopSwears:      u.SwearCounter.Top,
				SwearCount:     u.SwearCounter.Count,
				Emoticons:      u.EmoticonCounter.Top,
				EmoticonCount:  u.EmoticonCounter.Count,
				Questions:      uint(u.QuestionsCount),
				Exclamations:   uint(u.ExclamationsCount),
				AllCaps:        uint(u.AllCapsCount),
				SKicks:         u.KickCounters.Sent,
				RKicks:         u.KickCounters.Received,
				SSlaps:         u.SlapCounters.Sent,
				RSlaps:         u.SlapCounters.Received,
				NickReferences: u.NickReferences,
				Modes:          u.ModeCounters,
				Basic:          u.BasicTextCounters,
			}

			if m := u.Quotes.Random; m != nil {
				user.Message = u.Quotes.Random.Message
			}

			users = append(users, user)
		}
	}

	sort.Sort(sort.Reverse(ByMessageCount(users)))

	return users
}
