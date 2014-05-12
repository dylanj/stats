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

func TestUser_EmoticonCounter(t *testing.T) {
	t.Parallel()

	s := NewStats()

	s.AddMessage(Msg, network, channel, hostmask, time.Now(), "you wanna come over ;) ;)")

	u := s.Users[1]

	if tok := u.TopEmoticon(); tok.Token != ";)" || tok.Count != 2 {
		t.Error("Should have ;) as top emoticon with 2 uses")
	}
}

func TestUser_String(t *testing.T) {
	t.Parallel()

	u := &User{
		Nick:       "foo",
		MessageIDs: []uint{1, 2, 3},
	}

	if u.String() != "User: foo, Messages: 3" {
		t.Error("Didn't return correct string")
	}
}
