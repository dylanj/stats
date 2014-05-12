package stats

import "testing"

func TestEmoticonCounter(t *testing.T) {
	t.Parallel()

	tc := NewEmoticonCounter()

	if len(tc.Top) != 0 {
		t.Error("Top Emoticons should be empty.")
	}
	if len(tc.All) != 0 {
		t.Error("All Emoticons should be empty.")
	}

	m := &Message{Message: "he:Dllo :D world :D :("}
	tc.addMessage(m)

	if len(tc.Top) != 2 {
		t.Error("Top Emoticons should have two unique Emoticons.")
	}
	if len(tc.All) != 2 {
		t.Error("All Emoticons should have two unique Emoticons.")
	}

	if count, ok := tc.All[":("]; !ok {
		t.Error("Should have :( in All emoticons.")
	} else if count != 1 {
		t.Error("Should get correct count for emoticon.")
	}

	if count, ok := tc.All[":D"]; !ok {
		t.Error("Should have :D in All emoticons.")
	} else if count != 2 {
		t.Error("Should get correct count for emoticons.")
	}

	if tok := tc.Top[0]; tok.Token != ":D" || tok.Count != 2 {
		t.Error("Top emoticon is incorrect")
	}
}
