package main

import "stats"
import "os"
import "fmt"
import "bufio"
import "regexp"
import "errors"

func main() {
  ParseFile("deviate.weechatlog")
}

// ParseFile will for each line through the file and feed it into ParseLine.
func ParseFile(filename string) error {
  file, _ := os.Open(filename)
  scanner := bufio.NewScanner(file)

  var stats = stats.NewStats()

  for scanner.Scan() {
    ParseLine(stats, scanner.Bytes())
  }

  stats.ListChannels()
  stats.ListUsers()
 fmt.Printf("number of messages in stats: %i\n", stats.MessageCount())

  return errors.New("gi?")
}

func ParseJoin(irc_line []byte) (*stats.User, *stats.Channel) {
  joinRegex := regexp.MustCompile(`-->\t(?P<name>.*) \((?P<hostmask>.*)\) has joined (?P<channel>.*)`)

  n1 := joinRegex.SubexpNames()
  r2 := joinRegex.FindAllSubmatch(irc_line, -1)[0]

  matches := make(map[string][]byte)

  for i, n := range r2 {
    matches[n1[i]] = n
  }

  user := stats.NewUser(matches["name"], matches["hostmask"])
  channel := stats.NewChannel(matches["channel"])

  return user, channel
}

func ParseMessage(s *stats.Stats, matches map[string][]byte) (*stats.Message) {
  user_name := matches["cmd"]

  user := s.GetUser(user_name)
  channel := s.GetChannel("#deviate2")

  s.ListChannels()

  if channel == nil {
    fmt.Printf("channl is nil\n")
  }
  if user == nil {
    fmt.Printf("channl is nil\n")
  }

  return user.AddMessage(matches["message"], channel)
}

// ParseLine parses a single line of IRC directly from a socket.
// Will parse into irc.Message events using ultimateq's parse package (or write custom code)
func ParseLine(s *stats.Stats, line []byte) error {
  messageRegex := regexp.MustCompile(`(?P<date>.*)\t(?P<cmd>.*)\t(?P<message>.*)`)
  n1 := messageRegex.SubexpNames()
  r2 := messageRegex.FindAllSubmatch(line, -1)[0]

  matches := make(map[string][]byte)

  for i, n := range r2 {
    matches[n1[i]] = n
  }

  switch(string(matches["cmd"])) {
    case "-->":
      user, channel := ParseJoin(line)
      fmt.Printf("%s\n", line)
      fmt.Printf("username: %s\nhostmask: %s\nchannel: %s\n", user.Name, user.Hostmask, channel.Name)
      s.AddUser(user)
      s.AddChannel(channel)
    case "<--":
      // someone has quit.
    case "--":
      // some kind of message"
    default:
      message := ParseMessage(s, matches)
      if message != nil { 
        message.Print()
      }
      //fmt.Printf("%s: %s\n", matches["cmd"], matches["message"])
  }

  return errors.New("da")
}

