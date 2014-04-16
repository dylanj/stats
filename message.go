package stats

import "time"

type Message struct {
	Date    time.Time
	User    User
	Channel Channel
	Message []byte
}
