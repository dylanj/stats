package stats

import "math/rand"

const randomQuoteProbability = 10

type quotes struct {
	Last   uint
	Random uint
}

func (q *quotes) addMessage(m *Message) {
	if m.Kind != Msg {
		return
	}

	q.Last = m.ID

	if rand.Intn(randomQuoteProbability) == 0 {
		q.Random = m.ID
	}
}
