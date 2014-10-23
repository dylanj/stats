package main

import (
	"bytes"
	"testing"

	"github.com/DylanJ/stats"
)

const (
	weechatJoin        = `2013-08-07 16:49:40	-->	dylan (dylan@zqz.ca) has joined #deviate`
	weechatQuit        = `2013-08-07 16:52:04	<--	knivey (knivey@zkpq-5EEAFC38.dhcp.embarqhsd.net) has quit (Ping timeout: 181 seconds)`
	weechatPartMessage = `2013-08-07 16:55:42	<--	dylan (dylan@zqz.ca) has left #deviate (peace out)`
	weechatPart        = `2013-08-07 16:55:42	<--	zamn (zamn@newjoizi.comcast.us) has left #deviate`
	weechatMessage     = `2013-08-07 16:50:02	@Aaron	dylan: Auth with my bot for +v`
	weechatTopic       = `2013-08-07 16:49:40	--	Topic set by Scott on Sun, 14 Jul 2013 07:02:11`
	weechatAction      = `2013-08-08 10:49:57	 *	Knio slaps knivey`
)

const weechatFile = weechatJoin + "\n" + weechatQuit + "\n" +
	weechatMessage + "\n" + weechatPart + "\n" +
	weechatPartMessage + "\n"

func Benchmark_parseLine(b *testing.B) {
	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		sc.parseLine(s, weechatMessage)
	}
}

func TestScanner_parseLine_Quit(t *testing.T) {
	t.Parallel()
	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatQuit)

	var m *stats.Message
	if m = s.Messages[1]; m == nil {
		t.Error("Should be able to get first message from stats.")
	}

	if m.Kind != stats.Quit {
		t.Error("Kind should be Quit MsgKind.")
	}

	if m.ChannelID != 0 {
		t.Error("ChannelID should be 0")
	}
}

func TestScanner_parseLine_Action(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatAction)

	if len(s.Messages) > 0 {
		t.Error("It should ignore action messages (for now)")
	}
}

func TestScanner_parseLine_Topic(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatTopic)

	if len(s.Messages) > 0 {
		t.Error("It should ignore topic messages (for now)")
	}
}

func TestScanner_parseLine_Message(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatMessage)

	var m *stats.Message
	if m = s.Messages[1]; m == nil {
		t.Error("Should be able to get first message from stats.")
	}

	if m.Kind != stats.Msg {
		t.Error("Kind should be Msg MsgKind.")
	}
}

func TestScanner_parseLine_PartWithMessage(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatPartMessage)

	var m *stats.Message
	if m = s.Messages[1]; m == nil {
		t.Error("Should be able to get first message from stats.")
	}

	if m.Message != "peace out" {
		t.Error("Should have part message inside Message.")
	}

	if m.Kind != stats.Part {
		t.Error("Kind should be Part MsgKind.")
	}
}

func TestScanner_parseLine_Join(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatJoin)

	var m *stats.Message
	if m = s.Messages[1]; m == nil {
		t.Error("Should be able to get first message from stats.")
	}

	if m.Kind != stats.Join {
		t.Error("Kind should be Join MsgKind.")
	}
}

func TestScanner_NewScanner(t *testing.T) {
	t.Parallel()

	var s *scanner
	var e error
	if s, e = newScanner("network", "#deviate", "weechat", "file"); e != nil {
		t.Fatal(e)
	} else if s == nil {
		t.Error("Should return weechat scanner.")
	}

	if s.filenames[0] != "file" {
		t.Error(`Should set filename to "file"`)
	}

	if s.network != "foo" {
		t.Error(`Should set network to "foo"`)
	}

	if s.channel != "bar" {
		t.Error(`Should set channel to "bar"`)
	}
}

func TestScanner_ParseReader(t *testing.T) {
	t.Parallel()

	reader := bytes.NewBufferString(weechatFile)

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	var s = stats.NewStats()
	var n *stats.Network

	if err = sc.parseReader(s, reader); err != nil {
		t.Fatal("Error parsing stats:", err)
	}

	if n = s.GetNetwork("test_network"); n == nil {
		t.Error("Should be able to get network.")
	}

	if c := s.GetChannel("test_network", "#deviate"); c == nil {
		t.Error("Should be able to get channel.")
	}

	if u := s.GetUser("test_network", "dylan"); u == nil {
		t.Error("Should be able to find user from log.")
	}

	if len(s.Messages) == 0 {
		t.Error("Should have a message.")
	}
}

func TestScanner_parseLine(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc, err := newScanner("network", "#deviate", "weechat", "file")
	if err != nil {
		t.Fatal(err)
	}

	sc.parseLine(s, weechatJoin)

	if n := s.GetNetwork("network"); n == nil {
		t.Error("Stats should have the network.")
	}

	if c := s.GetChannel("network", "#deviate"); c == nil {
		t.Error("Stats should have the channel.")
	}

	if u := s.GetUser("network", "dylan"); u == nil {
		t.Error("Stats should have user who joined.")
	}

	if len(s.Messages) != 1 {
		t.Error("There should only be one message.")
	}

	var m *stats.Message
	if m = s.Messages[1]; m == nil {
		t.Error("Should be able to get first message from stats.")
	}

	if m.Date.IsZero() {
		t.Error("Date should have been initialized.")
	}

	if m.ChannelID != s.Channels[1].ID {
		t.Error("ChannelID should not be nil.")
	}

	if m.UserID != s.Users[1].ID {
		t.Error("UserID should not be nil.")
	}
}
