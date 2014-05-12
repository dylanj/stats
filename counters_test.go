package stats

import "testing"

func TestBasicTextCounters(t *testing.T) {
	t.Parallel()

	c := &BasicTextCounters{}
	m := &Message{Message: "I pity the fool"}

	c.addMessage(m)

	if c.Letters != 12 {
		t.Error("Should have 12 letters.")
	}

	if c.Words != 4 {
		t.Error("Should have 4 words.")
	}

	if c.Lines != 1 {
		t.Error("Should have 1 line.")
	}

	c.addMessage(m)

	if c.Letters != 24 {
		t.Error("Should have 24 letters.")
	}

	if c.Words != 8 {
		t.Error("Should have 8 words.")
	}

	if c.Lines != 2 {
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

func TestQuestionsCount(t *testing.T) {
	t.Parallel()

	var q QuestionsCount
	q.addMessage(&Message{Message: "hello?"})

	if q != 1 {
		t.Error("Should have added one question.")
	}

	q.addMessage(&Message{Message: "hello? is it me you're looking for?"})

	if q != 3 {
		t.Error("Should have added two more questions.")
	}
}

func TestExclamationsCount(t *testing.T) {
	t.Parallel()

	var e ExclamationsCount
	e.addMessage(&Message{Message: "No!"})

	if e != 1 {
		t.Error("Should have added on exclamation.")
	}

	e.addMessage(&Message{Message: "cant touch this! dun na na na"})

	if e != 2 {
		t.Error("Should have only picked up one more exclamation.")
	}
}

func TestAllCapsCount(t *testing.T) {
	t.Parallel()

	var a AllCapsCount
	a.addMessage(&Message{Message: "THIS WONT WoRK"})

	if a != 0 {
		t.Error("Should not have added an all caps sentence.")
	}

	a.addMessage(&Message{Message: "YOU CAN READ THIS BETTER IF I TYPE IN CAPS"})

	if a != 1 {
		t.Error("Should have added one all caps sentence.")
	}

	a.addMessage(&Message{Message: "!!#$^"})

	if a != 1 {
		t.Error("Should not have added another all caps sentence.")
	}
}
