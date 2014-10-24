package main

import (
	"errors"
	"net/http"

	"github.com/DylanJ/stats"
	"github.com/aarondl/jsonware"
)

const (
	assetURL       = "/assuts/"
	localAssetPath = "./html/assets"
)

var st *stats.Stats

func main() {
	st = stats.NewStats()

	StartServer(":8080", st)
}

// StartServer starts the webserver that will serve the stats pages.
func StartServer(bind string, s *stats.Stats) {
	http.Handle(assetURL, http.StripPrefix(assetURL, http.FileServer(http.Dir(localAssetPath))))
	http.Handle("/api.json", jsonware.JSON(testHandler))

	http.ListenAndServe(bind, nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) (*ChannelStatsJSON, error) {
	st.RLock()
	defer st.RUnlock()

	network := r.Form.Get("network")
	channel := r.Form.Get("channel")

	ch := st.GetChannel(network, channel)
	if ch == nil {
		return nil, jsonware.JSONErr{
			Status: 404,
			Err:    errors.New("Channel does not exist."),
		}
	}

	data := &ChannelStatsJSON{
		HourlyChart: ch.HourlyChart,
		TopURLs:     ch.URLCounter.Top[:15],
		TopWords:    ch.WordCounter.Top,
		TopSwears:   ch.SwearCounter.Top,
		TopUsers:    topUsers(st, ch),
		SwearCount:  ch.SwearCounter.Count,
	}

	return data, nil
}
