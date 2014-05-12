package stats

import "testing"

func TestSwearCounter(t *testing.T) {
	t.Parallel()

	tc := NewSwearCounter()

	if len(tc.Top) != 0 {
		t.Error("Top swears should be empty.")
	}
	if len(tc.All) != 0 {
		t.Error("All swears should be empty.")
	}
	if tc.Count != 0 {
		t.Error("Swear count should be empty.")
	}

	m := &Message{Message: "fuck #fucking fuck!!!"}
	tc.addMessage(m)

	if len(tc.Top) != 2 {
		t.Error("Top swears should have two unique swears.")
	}
	if len(tc.All) != 2 {
		t.Error("All swears should have two unique swears.")
	}
	if tc.Count != 3 {
		t.Error("Should see three swears.")
	}

	if count, ok := tc.All["fucking"]; !ok {
		t.Error("Should have fucking in All swears.")
	} else if count != 1 {
		t.Error("Should get correct count for swear.")
	}

	if count, ok := tc.All["fuck"]; !ok {
		t.Error("Should have fuck in All swears.")
	} else if count != 2 {
		t.Error("Should get correct count for swear.")
	}

	if tok := tc.Top[0]; tok.Token != "fuck" || tok.Count != 2 {
		t.Error("Top swear is incorrect")
	}
}
