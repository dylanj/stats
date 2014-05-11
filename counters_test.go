package stats

import "testing"

func TestBasicTextCounters(t *testing.T) {
	t.Parallel()

	c := &BasicTextCounters{}
	m := &Message{Message: "I pity the fool"}

	c.addMessage(m)

	if c.letters != 12 {
		t.Error("Should have 12 letters.")
	}

	if c.words != 4 {
		t.Error("Should have 4 words.")
	}

	if c.lines != 1 {
		t.Error("Should have 1 line.")
	}

	c.addMessage(m)

	if c.letters != 24 {
		t.Error("Should have 24 letters.")
	}

	if c.words != 8 {
		t.Error("Should have 8 words.")
	}

	if c.lines != 2 {
		t.Error("Should have 2 lines.")
	}
}
