package stats

import "math/rand"

const randomQuoteProbability = 10

type quotes struct {
	Last   uint
	Random uint
}

func (q *quotes) addMessage(m *Message) {
	q.Last = m.ID

	if rand.Intn(randomQuoteProbability) == 0 {
		q.Random = m.ID
	}
}
