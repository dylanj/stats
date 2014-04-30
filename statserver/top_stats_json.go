package main

import "github.com/DylanJ/stats"
import "encoding/json"
import "sort"

type UserJSON struct {
  ID uint
  Name string
  MessageCount int
  Message string
}

type ByMessageCount []*UserJSON
func (a ByMessageCount) Len() int           { return len(a) }
func (a ByMessageCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMessageCount) Less(i, j int) bool { return a[i].MessageCount < a[j].MessageCount }

func UserStatsJSON(s *stats.Stats) string {
  var users []*UserJSON
  users = make([]*UserJSON,0)
  user_map := s.Users

  for id, u := range(user_map) {
    if u != nil {
      user := &UserJSON{
        ID: id,
        Name: u.Name,
        MessageCount: u.MessageCount(),
        Message: s.RandomMessageForUser(u).Message,
      }
      users = append(users, user)
    }
  }

  sort.Sort(sort.Reverse(ByMessageCount(users)))

  b, _ := json.Marshal(&users)
  return string(b)
}

