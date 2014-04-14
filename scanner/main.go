package main

//import "github.com/aarondl/ultimateq/data"
import "stats"
import "os"
import "fmt"
import "bufio"
import "regexp"
import "errors"
import "time"

type Message struct {
  date time.Time
  user User
  channel Channel
  message []byte
}

type User struct {
  name []byte
  hostmask []byte
  messages []Message
}

func main() {
  ParseFile("deviate.weechatlog")
}

// ParseFile will for each line through the file and feed it into ParseLine.
func ParseFile(filename string) error {
  file, _ := os.Open(filename)
  scanner := bufio.NewScanner(file)

  stats = stats.NewStats()

  for scanner.Scan() {
    ParseLine(&stats, scanner.Bytes())
  }

  return errors.New("gi?")
}

func ParseJoin(irc_line []byte) (User, Channel) {
  joinRegex := regexp.MustCompile(`(?P<name>.*) ((?P<hostmask>.*)) has joined (?P<channel>.*)`)

  n1 := joinRegex.SubexpNames()
  r2 := joinRegex.FindAllSubmatch(irc_line, -1)[0]

  matches := make(map[string][]byte)

  for i, n := range r2 {
    matches[n1[i]] = n
  }

  user := User{}
  user.name = matches["name"]
  user.hostmask = matches["hostmask"]

  channel := Channel{}
  channel.name = matches["channel"]

  return user, channel
}

// ParseLine parses a single line of IRC directly from a socket.
// Will parse into irc.Message events using ultimateq's parse package (or write custom code)
func ParseLine(stats *Stats, irclogline []byte) error {
  messageRegex := regexp.MustCompile(`(?P<date>.*)\t(?P<cmd>.*)\t(?P<message>.*)`)
  n1 := messageRegex.SubexpNames()
  r2 := messageRegex.FindAllSubmatch(irclogline, -1)[0]

  matches := make(map[string][]byte)

  for i, n := range r2 {
    matches[n1[i]] = n
  }

  // fmt.Printf("date: %s\n", matches["date"])
  // fmt.Printf("name: %s\n", matches["cmd"])
  // fmt.Printf("message: %s\n", matches["message"])

  switch(string(matches["cmd"])) {
    case "-->":
      user, channel := ParseJoin(irclogline)
      fmt.Printf("(%s): [%s] - has joined the channel", user.name, user.hostmask)
      stats.AddChannel(channel)
   
      // user := User{}
      // user.name = matches["cmd"]
      // user.hostmask
      // someone joined the channela.
    case "<--":
      // someone has quit.
    case "--":
      // some kind of message"
    default:

      // users[string(matches["cmd"])] = 
      fmt.Printf("%s: %s\n", matches["cmd"], matches["message"])
  }

  return errors.New("da")
}

