package main

// import "github.com/aarondl/ultimateq/irc"
import "github.com/DylanJ/stats"
import "net/http"
import "fmt"
import "errors"

var port = ":8080"
var local_asset_path = "./html/assets"
var s *stats.Stats

func main() {
	s = stats.NewStats()

	StartServer()
}

func web_api(w http.ResponseWriter, r *http.Request) {
	fmt.Println("someone hit api")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, UserStatsJSON(s))
}
func web_root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("someone hit /")
	http.ServeFile(w, r, "./html/index.html")
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
// func ParseMessage(msg *irc.Message) error {
// 	return errors.New("asda")
// }
