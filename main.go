package stats

import "github.com/aarondl/ultimateq/irc" // THESE ARE OPTIONAL - SEE LAST FUNCTIO
import "net/http"

var port = ":8080"
var local_asset_path = "./assets"

func web_api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{ \"users\": [{\"name\": \"dylan\", \"num_lines\": 5, \"last_message\": \"Hello World\"}, { \"name\":\"fish\", \"num_lines\":3, \"last_message\": \"no worries\"}, {\"name\": \"me\", \"num_lines\": 212, \"last_message\": \"wassap\"}, {\"name\": \"chanserv\", \"num_lines\": 0, \"last_message\": null}] }")
}
func web_root(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

//import "github.com/aarondl/ultimateq/parse"

// StartServer starts the webserver that will serve the stats pages.
func StartServer() error {
	var asset_path = "/assets/"
	var asset_dir = http.Dir(local_asset_path)

	http.Handle(asset_path, http.StripPrefix(asset_path, http.FileServer(asset_dir)))
	http.HandleFunc("/api.json", web_api)
	http.HandleFunc("/", web_root)

	http.ListenAndServe(":8080", nil)

	return errors.New("gi?")
}

// StopServer stops the server.
func StopServer() error {
	return errors.New("gi?")
}

// GetURL returns a link to the stats server.
func GetURL() string {
	return "doo"
}

// ParseMessage parses an ultimateq irc.Message event.
func ParseMessage(msg *irc.Message) error {
	return errors.New("asda")
}

// THIS ONE IS OPTIONAL - BUT IS THE EASIEST FORMAT TO DEAL WITH USING THE PARSE LIBRARY
