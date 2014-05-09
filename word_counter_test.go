package stats

import (
	"fmt"
	"testing"
)

func TestTokenCounter_Word(t *testing.T) {
	t.Parallel()

	tc := NewWordCounter()

	if len(tc.Top) != 0 {
		t.Error("Top tokens should be empty.")
	}
	if len(tc.All) != 0 {
		t.Error("All tokens should be empty.")
	}

	m := &Message{Message: "foo bar bar baz"}
	tc.addMessage(m)

	if len(tc.Top) != 3 {
		t.Error("Top tokens should have three unique tokens.")
	}
	if len(tc.All) != 3 {
		t.Error("All tokens should have three unique tokens.")
	}

	if count, ok := tc.All["foo"]; !ok {
		t.Error("Should have foo in All tokens.")
	} else if count != 1 {
		t.Error("Should get correct count for token.")
	}

	if count, ok := tc.All["bar"]; !ok {
		t.Error("Should have bar in All tokens.")
	} else if count != 2 {
		t.Error("Should get correct count for token.")
	}

	if tok := tc.Top[0]; tok.Token != "bar" || tok.Count != 2 {
		t.Error("Top token is incorrect")
	}

	tc = NewWordCounter()

	for i := 'a'; i < 'z'; i++ {
		url := fmt.Sprintf("foo%c", i)
		m := &Message{Message: url}
		tc.addMessage(m)
		url = fmt.Sprintf("bar%c", i)
		m = &Message{Message: url}
		tc.addMessage(m)
		url = fmt.Sprintf("baz%c", i)
		m = &Message{Message: url}
		tc.addMessage(m)
	}

	for _, v := range tc.Top {
		if v.Count != uint(1) {
			t.Error("Count is incorrect. ")
		}
	}
}
