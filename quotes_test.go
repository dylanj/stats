package stats

import (
	"math/rand"
	"testing"
	"time"
)

func TestQuotes(t *testing.T) {
	rand.Seed(81) // returns (0, 9)

	var q quotes
	m := &Message{ID: 4}

	if q.Random != 0 {
		t.Error("Random message should not be set.")
	}

	if q.Last != 0 {
		t.Error("Last message should not be set.")
	}

	q.addMessage(m)

	if q.Random != m.ID {
		t.Error("Random message should be set")
	}

	if q.Last != m.ID {
		t.Error("Last message should be set")
	}

	m2 := &Message{ID: 5}

	q.addMessage(m2)

	if q.Random != m.ID {
		t.Error("Random message should not change")
	}

	if q.Last != m2.ID {
		t.Error("Last message be updated")
	}
}

func TestQuotesUpdates(t *testing.T) {
	rand.Seed(111) // returns (0,0,0)
	s := NewStats()
	n := s.addNetwork(network)
	c := s.addChannel(n, channel)
	u := s.addUser(n, nick)

	if n.Quotes.Last != 0 && n.Quotes.Random != 0 {
		t.Error("Last message and random message should not be set.")
	}
	if c.Quotes.Last != 0 && c.Quotes.Random != 0 {
		t.Error("Last message and random message should not be set.")
	}
	if u.Quotes.Last != 0 && u.Quotes.Random != 0 {
		t.Error("Last message and random message should not be set.")
	}

	s.addMessage(Msg, n, c, u, time.Now(), "nihao")

	if n.Quotes.Random != 1 {
		t.Error("Random message should be set")
	}
	if c.Quotes.Random != 1 {
		t.Error("Random message should be set")
	}
	if u.Quotes.Random != 1 {
		t.Error("Random message should be set")
	}

	if n.Quotes.Last != 1 {
		t.Error("Last message should be set")
	}
	if c.Quotes.Last != 1 {
		t.Error("Last message should be set")
	}
	if u.Quotes.Last != 1 {
		t.Error("Last message should be set")
	}
}
