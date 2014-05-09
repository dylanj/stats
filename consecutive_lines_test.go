package stats

import (
	"testing"
	"time"
)

func TestConsecutiveLines(t *testing.T) {
	t.Parallel()

	s := NewStats()
	s.AddMessage(Msg, network, channel, "aaron", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "aaron", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "zamn", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "aaron", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "aaron", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "aaron", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "zamn", time.Now(), "some foo")
	s.AddMessage(Msg, network, channel, "zamn", time.Now(), "some foo")

	cl := s.Channels[1].ConsecutiveLines

	if len(cl.TopUsers) != 2 {
		t.Error("Should only have two users in TopUsers")
	}

	if cl.TopUsers[0].Token != "aaron" {
		t.Error("Top user should be aaron.")
	}

	if cl.TopUsers[0].Count != 3 {
		t.Error("Top user should have 3 consecutive lines.")
	}
}
