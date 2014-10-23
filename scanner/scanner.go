package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/DylanJ/stats"
)

var (
	parserFlag = flag.String("parser", "weechat", "A named internal log parser or a parser file to load. Internal: [weechat]")
	netFlag    = flag.String("network", "", "The network where the log file came from.")
	chanFlag   = flag.String("channel", "", "The channel where the log file came from.")
)

var usage = `
Scanner should be invoked with one or more filenames. Use * as to use standard in
as an input file.

To invoke scanner with a custom parser simply define a file that starts with a date
format and supplies a regex for the following in order, ensuring all named regex args
are supplied:
  message [date, nick, message] 
  join    [date, nick, host]
  part    [date, nick, host]
  kick    [date, nick, target, message]
  quit    [date, nick, host, message]
  action  [date, nick, action]
  mode    [date, mode, nick]
  topic   [date, nick, action]

eg.
2006-01-02 15:04:05
^(?P<date>[0-9:\- ]*)\t(?:[@&+])?(?P<nick>[^\s\-]+)\t(?P<message>.*)$
^(?P<date>[0-9:\- ]*)\t-->\t(?P<nick>.*) \((?P<host>.*)\) has joined (?P<channel>(?:&|#)\w+)$
...

Sample invocation:
scanner -network zkpq -channel #deviate -parser myCustomParser.parser #deviateLog.log

scanner [options] <filenames...>
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	remaining := flag.Args()
	if len(remaining) == 0 {
		fmt.Fprintln(os.Stderr, "Must pass in at least one file name.")
		os.Exit(1)
	}

	if len(*netFlag) == 0 {
	}
	if len(*chanFlag) == 0 {
	}

	sc, err := newScanner(*netFlag, *chanFlag, *parserFlag, remaining...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Problem creating scanner:", err)
		os.Exit(1)
	}

	stats, err := sc.parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed parsing sources:", err)
		os.Exit(1)
	}
	stats.Save()
}

type scanner struct {
	filenames []string
	network   string
	channel   string

	parser parser
}

type parser struct {
	dateFormat string
	message    *regexp.Regexp
	join       *regexp.Regexp
	part       *regexp.Regexp
	kick       *regexp.Regexp
	quit       *regexp.Regexp
	action     *regexp.Regexp
	mode       *regexp.Regexp
	topic      *regexp.Regexp
}

var weechat = parser{
	dateFormat: "2006-01-02 15:04:05",

	message: regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t(?:[@&+])?(?P<nick>[^\s\-]+)\t(?P<message>.*)$`),
	join:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t-->\t(?P<nick>.*) \((?P<host>.*)\) has joined (?P<channel>(?:&|#)\w+)$`),
	quit:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t<--\t(?P<nick>.*) \((?P<host>.*)\) has quit \((?P<message>.*)\)$`),
	part:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t<--\t(?P<nick>.*) \((?P<host>.*)\) has left (?P<channel>(?:&|#)\w+)(?: \((?P<message>.*)\))?$`),
	kick:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t<--\t(?P<nick>.*) has kicked (?P<target>.*) \((?P<message>.*)\)$`),
	topic:   regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t--\t(?P<nick>.*) has changed topic for (?P<channel>(?:&|#)\w+) from "(?P<topic>.*)"$`),
	mode:    regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t--\tMode (?P<channel>(?:&|#)\w+) \[(?P<mode>\S+)[^\]]*\] by (?P<nick>.*)$`),
	action:  regexp.MustCompile(`^(?P<date>[0-9:\- ]*)\t *\t(?P<nick>.*) (?P<action>.*)$`),
}

func newScanner(network, channel, parser string, files ...string) (*scanner, error) {
	sc := &scanner{
		network:   network,
		channel:   channel,
		filenames: files,
	}

	switch parser {
	case "weechat":
		sc.parser = weechat
	default:
		var err error
		if sc.parser, err = loadParser(parser); err != nil {
			return nil, err
		}
	}

	return sc, nil
}

func loadParser(filename string) (parser, error) {
	var p parser

	f, err := os.Open(filename)
	if err != nil {
		return p, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for line := 0; scanner.Scan(); line++ {
		switch line {
		case 0:
			p.dateFormat = scanner.Text()
		case 1:
			p.message, err = regexp.Compile(scanner.Text())
		case 2:
			p.join, err = regexp.Compile(scanner.Text())
		case 3:
			p.part, err = regexp.Compile(scanner.Text())
		case 4:
			p.kick, err = regexp.Compile(scanner.Text())
		case 5:
			p.quit, err = regexp.Compile(scanner.Text())
		case 6:
			p.action, err = regexp.Compile(scanner.Text())
		case 7:
			p.mode, err = regexp.Compile(scanner.Text())
		case 8:
			p.topic, err = regexp.Compile(scanner.Text())
		}

		if err != nil {
			return p, fmt.Errorf("Failed to parse line: %d of %s: %v", line, filename, err)
		}
	}

	if err = scanner.Err(); err != nil {
		return p, fmt.Errorf("Failed to parse file: %v", err)
	}
	return p, nil
}

func (sc *scanner) parse() (*stats.Stats, error) {
	stats := stats.NewStats()
	for _, file := range sc.filenames {
		if file == "*" {
			if err := sc.parseReader(stats, os.Stdin); err != nil {
				return nil, err
			}
		} else {
			f, err := os.Open(file)
			if err != nil {
				return nil, err
			}

			err = sc.parseReader(stats, f)
			f.Close()
			if err != nil {
				return nil, err
			}
		}
	}

	return stats, nil
}

func (sc *scanner) parseReader(s *stats.Stats, r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		sc.parseLine(s, scanner.Text())
	}

	return scanner.Err()
}

func (sc *scanner) parseLine(s *stats.Stats, line string) {
	if r := findData(sc.parser.join, line); r != nil {

		nick, dateString, _ := r["nick"], r["date"], r["channel"]

		if len(nick) == 0 || len(dateString) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Join, sc.network, sc.channel, nick, date, "")

	} else if r := findData(sc.parser.part, line); r != nil {

		nick, dateString, _, message := r["nick"], r["date"], r["channel"], r["message"]

		if len(nick) == 0 || len(dateString) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Part, sc.network, sc.channel, nick, date, message)

	} else if r = findData(sc.parser.quit, line); r != nil {

		nick, dateString, message := r["nick"], r["date"], r["message"]

		if len(nick) == 0 || len(dateString) == 0 || len(message) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Quit, sc.network, "", nick, date, message)

	} else if r = findData(sc.parser.message, line); r != nil {

		nick, dateString, message := r["nick"], r["date"], r["message"]

		if len(nick) == 0 || len(dateString) == 0 || len(message) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Msg, sc.network, sc.channel, nick, date, message)
	} else if r = findData(sc.parser.kick, line); r != nil {
		nick, dateString, target := r["nick"], r["date"], r["target"]

		if len(nick) == 0 || len(dateString) == 0 || len(target) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}
		s.AddMessage(stats.Kick, sc.network, sc.channel, nick, date, target)
	} else if r = findData(sc.parser.mode, line); r != nil {
		nick, dateString, mode := r["nick"], r["date"], r["mode"]

		if len(nick) == 0 || len(dateString) == 0 || len(mode) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Mode, sc.network, sc.channel, nick, date, mode)
	} else if r = findData(sc.parser.topic, line); r != nil {
		nick, dateString, topic := r["nick"], r["date"], r["topic"]

		if len(nick) == 0 || len(dateString) == 0 || len(topic) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Topic, sc.network, sc.channel, nick, date, topic)
	} else if r = findData(sc.parser.action, line); r != nil {
		nick, dateString, action := r["nick"], r["date"], r["action"]

		if len(nick) == 0 || len(dateString) == 0 || len(action) == 0 {
			return
		}

		date, err := time.Parse(sc.parser.dateFormat, dateString)
		if err != nil {
			return
		}

		s.AddMessage(stats.Action, sc.network, sc.channel, nick, date, action)
	}
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
