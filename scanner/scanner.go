package main

import "stats"
import "strings"
import "os"
import "bufio"
import "regexp"
import "errors"
//import "fmt"

func main() {
	// load existing stats
	// cmd options
	// default: scanner deviate.weechatlog network channel
	// help: ^
	// list: lists networks and channels in db.
	network := "zkpq.ca"
	channel := "#deviate"
	logfile := "deviate.weechatlog"

	ParseFile(network, channel, logfile)
}

// ParseFile will for each line through the file and feed it into ParseLine.
func ParseFile(network_name string, channel_name string, filename string) error {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

  var stats2 = stats.ImportData()

	var stats = stats.NewStats()
	var network = stats.GetNetwork(network_name)
	var channel = network.GetChannel(channel_name)

  var i = 0;
	for scanner.Scan() {
    i++
		ParseLine(network, channel, scanner.Bytes())
    //fmt.Printf("line %d\n", i)
	}

	stats2.Information()
	stats.ExportData()

	return errors.New("gi?")
}

func ParseJoin(n *stats.Network, line []byte) {
	// joinRegex := regexp.MustCompile(`-->\t(?P<name>.*) \((?P<hostmask>.*)\) has joined (?P<channel>.*)`)

	// n1 := joinRegex.SubexpNames()
	// r2 := joinRegex.FindSubmatch(line)

	// matches := make(map[string][]byte)
	// for i := 1; i < len(r2); i++ {
	// 	matches[n1[i]] = r2[i]
	// }

	// name := string(matches["name"])
	// channel_name := string(matches["channel"])

	// user := n.GetUser(name)
	// channel := n.GetChannel(channel_name)

	// channel.JoinCount++
	// user.JoinCount++
}

func ParseMessage(n *stats.Network, c *stats.Channel, matches map[string][]byte) {
	nick := strings.TrimLeft(string(matches["cmd"]), "@+&")
	user := n.GetUser(nick)

	message := &stats.Message{
		Message:   string(matches["message"]),
		UserID:    user.ID,
		ChannelID: c.ID,
		// todo date: ???
	}

  n.AddMessage(message)
}

func ParseLine(n *stats.Network, c *stats.Channel, line []byte) error {
  //messageRegex := regexp.MustCompile(`(?P<date>.*)\t(?P<cmd>.*)\t(?P<message>.*)`)
	messageRegex := regexp.MustCompile(`(?P<date>[0-9:\- ]*)\t(?P<cmd>[\w@+&]*)\t{1}(?P<message>.*)`)
	n1 := messageRegex.SubexpNames()
	r2 := messageRegex.FindSubmatch(line)

	matches := make(map[string][]byte)
	for i := 1; i < len(r2); i++ {
		matches[n1[i]] = r2[i]
	}

//  fmt.Printf("match: [%s]\n", matches["cmd"])
	switch string(matches["cmd"]) {
	case "-->":
		ParseJoin(n, line)
    break
	case "<--":
		// someone has quit.
    break
	case "--":
		// some kind of message.
    break
  case " *":
    break
	default:
		ParseMessage(n, c, matches)
    break
	}

	return errors.New("da")
}
