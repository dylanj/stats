package main

import (
	"testing"

	"github.com/DylanJ/stats"
)

func TestURLs_TopURLs(t *testing.T) {
	t.Parallel()

	u := stats.NewURLs()

	urls := []string{"http://google.com", "http://slashdot.org", "http://news.ycombinator.com", "github.com", "amazon.com", "freenode.net"}

	for i, url := range urls {
		u[url] = uint(i) + 1
	}

	nURLs := 5
	topURLs := TopURLs(u, nURLs)

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
