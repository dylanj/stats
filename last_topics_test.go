package stats

import (
	"fmt"
	"testing"
)

func TestLastTopics(t *testing.T) {
	t.Parallel()

	l := NewLastTopics()

	var msg string

	for i := 0; i < maxTopics+1; i++ {
		msg = fmt.Sprintf("message %d", i)
		m := &Message{Message: msg}
		l.addMessage(m)
	}

	if len(l.Topics) != maxTopics {
		t.Error("Should have 5 topics.")
	}

	if l.Topics[maxTopics-1].Message != msg {
		t.Error("Should have last topic as last element in topics array.")
	}
}
