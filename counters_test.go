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

func TestBasicTextCounters_WordsPerLine(t *testing.T) {
	t.Parallel()

	c := &BasicTextCounters{}

	if c.WordsPerLine() != 0 {
		t.Error("Should have no words per line")
	}

	m := &Message{Message: "I pity the fool"}
	c.addMessage(m)

	if c.WordsPerLine() != 4 {
		t.Error("Should have 4 words per line")
	}

	m = &Message{Message: "Whaaa a"}
	c.addMessage(m)

	if c.WordsPerLine() != 3 {
		t.Error("Should have 3 words per line")
	}
}

func TestBasicTextCounters_LettersPerLine(t *testing.T) {
	t.Parallel()

	c := &BasicTextCounters{}

	if c.LettersPerLine() != 0 {
		t.Error("Should have no letters per line")
	}

	m := &Message{Message: "a b cd"}
	c.addMessage(m)

	if c.LettersPerLine() != 4 {
		t.Error("Should have 4 letters per line")
	}

	m = &Message{Message: "              a  s"}
	c.addMessage(m)

	if c.LettersPerLine() != 3 {
		t.Error("Should have 3 letters per line")
	}

}
