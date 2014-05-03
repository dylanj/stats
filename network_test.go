package stats

import (
	"testing"
	"time"
)

func TestNetwork_buildIndexes(t *testing.T) {
	t.Parallel()

	s := NewStats()
	s.AddMessage(Msg, network, channel, nick, time.Now(), "some foo")

	n := s.Networks[1]
	n.channels = nil
	n.users = nil
	n.stats = nil

	n.buildIndexes(s)

	if n.channels == nil {
		t.Error("channels index should have been created")
	}

	if n.users == nil {
		t.Error("users index should have been created")
	}

	if n.stats == nil {
		t.Error("should be pointer to stats")
	}

	if _, ok := n.channels[channel]; !ok {
		t.Error("should be able to retrieve channel from index")
	}

	if _, ok := n.users[nick]; !ok {
		t.Error("should be able to retrieve user from index")
	}
}
