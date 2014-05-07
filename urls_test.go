package stats

import (
	"testing"
	"time"
)

func TestURLs(t *testing.T) {
	t.Parallel()

	urls := NewURLs()
	m := &Message{Message: "i use http://google.com to search for things."}

	if len(urls) != 0 {
		t.Error("urls map should be empty")
	}

	urls.addMessage(m)

	if count, ok := urls["http://google.com"]; !ok || count == 0 {
		t.Error("Should have found and incremented google.com url count")
	}
}

func TestURLsUpdates(t *testing.T) {
	t.Parallel()

	s := NewStats()
	n := s.addNetwork(network)
	c := s.addChannel(n, channel)
	u := s.addUser(n, nick)

	if len(c.URLs) != 0 {
		t.Error("Channel URLs should be empty")
	}

	s.addMessage(Msg, n, c, u, time.Now(), "hello world")

	if len(c.URLs) != 0 {
		t.Error("Channel URLs should still be empty")
	}

	if len(n.URLs) != 0 {
		t.Error("Network URLs should still be empty")
	}

	s.addMessage(Msg, n, c, u, time.Now(), "hello google.com world")

	if len(c.URLs) == 0 {
		t.Error("Channel URLs should have had url added to it")
	}

	if len(n.URLs) == 0 {
		t.Error("Network URLs should have had url added to it")
	}
}
