package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"runtime/pprof"
	"time"

	"github.com/DylanJ/stats"
)

func main() {
	profile, _ := os.Create("prof")
	pprof.StartCPUProfile(profile)
	defer pprof.StopCPUProfile()

	f, _ := os.Open("deviate.weechatlog")

	defer f.Close()

	sc := NewDefaultScanner("whogivesashit", "network", "#deviate", "weechat")
	stats := sc.ParseReader(f)
	stats.Save()
}

type Scanner struct {
	filename string
	network  string
	channel  string

	dateFormat string

	message *regexp.Regexp
	join    *regexp.Regexp
	part    *regexp.Regexp
	kick    *regexp.Regexp
	quit    *regexp.Regexp
	action  *regexp.Regexp
	mode    *regexp.Regexp
	topic   *regexp.Regexp
}

var weechat = &Scanner{
	dateFormat: "2006-01-02 15:04:05",

	message: regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t(?:[@&+])?(?P<nick>[^\s\-]+)\t(?P<message>.*)$`),
	join:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t-->\t(?P<nick>.*) \((?P<host>.*)\) has joined (?P<channel>(?:&|#)\w+)$`),
	quit:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t<--\t(?P<nick>.*) \((?P<host>.*)\) has quit \((?P<message>.*)\)$`),
	part:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t<--\t(?P<nick>.*) \((?P<host>.*)\) has left (?P<channel>(?:&|#)\w+)(?: \((?P<message>.*)\))?$`),
	kick:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t<--\t(?P<nick>.*) has kicked (?P<target>.*) \((?P<message>.*)\)$`),
	topic:   regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t--\t(?P<nick>.*) has changed topic for (?P<channel>(?:&|#)\w+) from "(?P<topic>.*)"$`),
	mode:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t--\tMode (?P<channel>(?:&|#)\w+) \[(?P<mode>[^\]]+)\] by (?P<nick>.*)$`),
	action:  regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t *\t(?P<nick>.*) (?P<action>.*)$`),
}

//2013-08-14 09:12:37	--	Aaron has changed topic for #deviate from "< Knio> Man your so good at bc2, wanna show me some pro tips later?  -  < Aaron> You should go with a 250cc bike, they are great for big pussies like me." to "<Scott> I bought a 250cc bike, because I have a soul devouring pussycancer that makes me into a huge pussy-bitch. It drives at the speed of smell."
// NewDefaultScanner
func NewDefaultScanner(filename, network, channel, scanner string) *Scanner {
	var sc *Scanner
	switch scanner {
	case "weechat":
		sc = weechat
	default:
		return nil
	}

	sc.network = network
	sc.channel = channel
	sc.filename = filename

	return sc
}

// ParseReader parses a reader into statistics
func (sc *Scanner) ParseReader(r io.Reader) *stats.Stats {
	stats := stats.NewStats() // import

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		sc.ParseLine(stats, scanner.Text())
	}

	return stats
}

func findIndex(haystack []string, needle string) int {
	for i, v := range haystack {
		if needle == v {
			return i
		}
	}

	return -1
}

func findData(regex *regexp.Regexp, line string) map[string]string {
	results := make(map[string]string)

	r := regex.FindStringSubmatch(line)

	if r == nil {
		return nil
	}

	names := regex.SubexpNames()

	for i, n := range names[1:] {
		results[n] = r[i+1]
	}

	if host := results["host"]; len(host) > 0 {
		results["nick"] += "!" + host
	}

	return results
}

func (sc *Scanner) ParseLine(s *stats.Stats, line string) {
	if r := findData(sc.join, line); r != nil {

		nick, dateString, channel := r["nick"], r["date"], r["channel"]

		if len(nick) == 0 || len(dateString) == 0 || len(channel) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Join, sc.network, channel, nick, date, "")

	} else if r := findData(sc.part, line); r != nil {

		nick, dateString, channel, message := r["nick"], r["date"], r["channel"], r["message"]

		if len(nick) == 0 || len(dateString) == 0 || len(channel) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Part, sc.network, channel, nick, date, message)

	} else if r = findData(sc.quit, line); r != nil {

		nick, dateString, message := r["nick"], r["date"], r["message"]

		if len(nick) == 0 || len(dateString) == 0 || len(message) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Quit, sc.network, "", nick, date, message)

	} else if r = findData(sc.message, line); r != nil {

		nick, dateString, message := r["nick"], r["date"], r["message"]

		if len(nick) == 0 || len(dateString) == 0 || len(message) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Msg, sc.network, sc.channel, nick, date, message)
	} else if r = findData(sc.kick, line); r != nil {
		nick, dateString, target := r["nick"], r["date"], r["target"]

		if len(nick) == 0 || len(dateString) == 0 || len(target) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}
		s.AddMessage(stats.Kick, sc.network, sc.channel, nick, date, target)
	} else if r = findData(sc.mode, line); r != nil {
		nick, dateString, mode := r["nick"], r["date"], r["mode"]

		if len(nick) == 0 || len(dateString) == 0 || len(mode) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Mode, sc.network, sc.channel, nick, date, mode)
	} else if r = findData(sc.topic, line); r != nil {
		nick, dateString, topic := r["nick"], r["date"], r["topic"]

		if len(nick) == 0 || len(dateString) == 0 || len(topic) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Topic, sc.network, sc.channel, nick, date, topic)
	} else if r = findData(sc.action, line); r != nil {
		nick, dateString, action := r["nick"], r["date"], r["action"]

		if len(nick) == 0 || len(dateString) == 0 || len(action) == 0 {
			return
		}

		date, err := time.Parse(sc.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Action, sc.network, sc.channel, nick, date, action)
	}
}
