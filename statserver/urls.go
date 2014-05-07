package main

import "sort"

type TopURL struct {
	URL   string `json:"url"`
	Count uint   `json:"count"`
}

// TopURLs returns the top n most popular urls.
func TopURLs(u map[string]uint, n int) []*TopURL {
	list := make([]*TopURL, 0)

	if len(u) == 0 {
		return list
	}

	for url, count := range u {
		u := &TopURL{URL: url, Count: count}
		list = append(list, u)
	}

	sort.Sort(byCount(list))

	return list[0:n]
}

type byCount []*TopURL

func (a byCount) Len() int           { return len(a) }
func (a byCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byCount) Less(i, j int) bool { return a[i].Count > a[j].Count }
