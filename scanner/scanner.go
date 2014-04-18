package main

import "stats"
import "strings"
import "os"
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

  stats.Information()
  stats.ExportData()

  return errors.New("gi?")
}

func ParseJoin(line []byte) (*stats.User, *stats.Channel) {
  joinRegex := regexp.MustCompile(`-->\t(?P<name>.*) \((?P<hostmask>.*)\) has joined (?P<channel>.*)`)

  n1 := joinRegex.SubexpNames()
  r2 := joinRegex.FindSubmatch(line)

  matches := make(map[string][]byte)
  for i := 1; i < len(r2); i++ {
    matches[n1[i]] = r2[i]
  }

  user := stats.NewUser(string(matches["name"]), string(matches["hostmask"]))
  channel := stats.NewChannel(matches["channel"])

  return user, channel
}

func ParseMessage(s *stats.Stats, matches map[string][]byte) (*stats.Message) {
  user_name := strings.TrimLeft(string(matches["cmd"]), "@+&")

  user := s.GetUser(user_name)
  channel := s.GetChannel("#deviate")
  s.AddUser(user)

  return user.AddMessage(matches["message"], channel)
}

func ParseLine(s *stats.Stats, line []byte) error {
  messageRegex := regexp.MustCompile(`(?P<date>.*)\t(?P<cmd>.*)\t(?P<message>.*)`)
  n1 := messageRegex.SubexpNames()
  r2 := messageRegex.FindSubmatch(line)

  matches := make(map[string][]byte)
  for i := 1; i < len(r2); i++ {
    matches[n1[i]] = r2[i]
  }

  switch(string(matches["cmd"])) {
    case "-->":
      user, channel := ParseJoin(line)
      s.AddUser(user)
      s.AddChannel(channel)
    case "<--":
      // someone has quit.
    case "--":
      // some kind of message.
    default:
      message := ParseMessage(s, matches)
      s.AddMessage(message)
  }

  return errors.New("da")
}

