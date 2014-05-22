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

	if q.Random != nil {
		t.Error("Random message should not be set.")
	}

	if q.Last != nil {
		t.Error("Last message should not be set.")
	}

	q.addMessage(m)

	if q.Random != m {
		t.Error("Random message should be set")
	}

	if q.Last != m {
		t.Error("Last message should be set")
	}

	m2 := &Message{ID: 5}

	q.addMessage(m2)

	if q.Random != m {
		t.Error("Random message should not change")
	}

	if q.Last != m2 {
		t.Error("Last message be updated")
	}
}

func TestQuotesUpdates(t *testing.T) {
	rand.Seed(7075) // returns (0,0,0,0) - dont ask
	s := NewStats()
	n := s.addNetwork(network)
	c := s.addChannel(n, channel)
	u := s.addUser(n, nick)
	cu := u.addChannelUser(channel)

	if n.Quotes.Last != nil && n.Quotes.Random != nil {
		t.Error("Last message and random message should not be set.")
	}
	if c.Quotes.Last != nil && c.Quotes.Random != nil {
		t.Error("Last message and random message should not be set.")
	}
	if u.Quotes.Last != nil && u.Quotes.Random != nil {
		t.Error("Last message and random message should not be set.")
	}

	m := s.addMessage(Msg, n, c, u, cu, time.Now(), "nihao")

	if n.Quotes.Random != m {
		t.Error("Random message should be set")
	}
	if c.Quotes.Random != m {
		t.Error("Random message should be set")
	}

	if u.Quotes.Random != m {
		t.Error("Random message should be set")
	}

	if n.Quotes.Last != m {
		t.Error("Last message should be set")
	}
	if c.Quotes.Last != m {
		t.Error("Last message should be set")
	}
	if u.Quotes.Last != m {
		t.Error("Last message should be set")
	}
}
