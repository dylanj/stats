package stats

import "math/rand"

const randomQuoteProbability = 10

type quotes struct {
	Last   *Message
	Random *Message
}

func (q *quotes) addMessage(m *Message) {
	q.Last = m

	if rand.Intn(randomQuoteProbability) == 0 {
		q.Random = m
	}
}
