package stats

import (
	"testing"
	"time"
)

func TestURLs(t *testing.T) {
	t.Parallel()

	u := NewURLs()
	m := &Message{Message: "i use http://google.com to search for things."}

	if len(u.urls) != 0 {
		t.Error("urls map should be empty")
	}

	u.addMessage(m)

	if count, ok := u.urls["http://google.com"]; !ok || count == 0 {
		t.Error("Should have found and incremented google.com url count")
	}
}

func TestURLs_TopURLs(t *testing.T) {
	t.Parallel()

	u := NewURLs()

	urls := []string{"http://google.com", "http://slashdot.org", "http://news.ycombinator.com", "github.com", "amazon.com", "freenode.net"}

	for i, url := range urls {
		m := &Message{Message: "i use " + url + " to do things"}

		for j := 0; j <= i; j++ {
			u.addMessage(m)
		}
	}

	nURLs := 5
	topURLs := u.TopURLs(nURLs)

	if len(topURLs) != nURLs {
		t.Errorf("Should return %d top urls", nURLs)
	}

	if topURLs[0].Count != uint(len(urls)) {
		t.Error("Incorrect count for top url")
	}

	if topURLs[0].URL != urls[len(urls)-1] {
		t.Error("Incorrect URL for top urls")
	}
}

func TestURLsUpdates(t *testing.T) {
	t.Parallel()

	s := NewStats()
	n := s.addNetwork(network)
	c := s.addChannel(n, channel)
	u := s.addUser(n, nick)

	if len(c.URLs.urls) != 0 {
		t.Error("Channel URLs should be empty")
	}

	s.addMessage(Msg, n, c, u, time.Now(), "hello world")

	if len(c.URLs.urls) != 0 {
		t.Error("Channel URLs should still be empty")
	}

	s.addMessage(Msg, n, c, u, time.Now(), "hello google.com world")

	if len(c.URLs.urls) == 0 {
		t.Error("Channel URLs should have had url added to it")
	}
}
