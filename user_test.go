package stats

import (
	"testing"
	"time"
)

func TestUser_BasicTextCounters(t *testing.T) {
	t.Parallel()

	s := NewStats()

	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "a b c d ef")

	u := s.Users[1]

	if u.WordsPerLine() != 5 {
		t.Error("Should have 5 words per line here")
	}

	if u.LettersPerLine() != 6 {
		t.Error("Should have 6 letters per line")
	}
}
