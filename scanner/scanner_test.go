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
)

const weechatFile = weechatJoin + "\n" + weechatQuit + "\n" +
	weechatMessage + "\n" + weechatPart + "\n" +
	weechatPartMessage + "\n"

func TestScanner_ParseLine_Quit(t *testing.T) {
	t.Parallel()
	s := stats.NewStats()

	sc := NewDefaultScanner("file", "network", "#deviate", "weechat")

	sc.ParseLine(s, weechatQuit)

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

func TestScanner_ParseLine_Message(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc := NewDefaultScanner("file", "network", "#deviate", "weechat")

	sc.ParseLine(s, weechatMessage)

	var m *stats.Message
	if m = s.Messages[1]; m == nil {
		t.Error("Should be able to get first message from stats.")
	}

	if m.Kind != stats.Msg {
		t.Error("Kind should be Msg MsgKind.")
	}
}

func TestScanner_ParseLine_PartWithMessage(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc := NewDefaultScanner("file", "network", "#deviate", "weechat")

	sc.ParseLine(s, weechatPartMessage)

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

func TestScanner_ParseLine_Join(t *testing.T) {
	t.Parallel()

	s := stats.NewStats()

	sc := NewDefaultScanner("file", "network", "#deviate", "weechat")

	sc.ParseLine(s, weechatJoin)

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

	if m.Kind != stats.Join {
		t.Error("Kind should be Join MsgKind.")
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

func TestScanner_NewDefaultScanner(t *testing.T) {
	t.Parallel()

	var s *Scanner
	if s = NewDefaultScanner("file", "foo", "bar", "baz"); s != nil {
		t.Error("Should return nil when unknown scanner specified.")
	}

	if s = NewDefaultScanner("file", "foo", "bar", "weechat"); s == nil {
		t.Error("Should return weechat scanner.")
	}

	if s.filename != "file" {
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

	weechat.channel = "#deviate"
	weechat.network = "test_network"

	var s *stats.Stats
	var n *stats.Network

	if s = weechat.ParseReader(reader); s == nil {
		t.Error("Should not return nil.")
	}

	if n = s.GetNetwork("test_network"); n == nil {
		t.Error("Should be able to get network.")
	}

	if c := s.GetChannel("test_network", "#channel"); c == nil {
		t.Error("Should be able to get channel.")
	}

	if u := s.GetUser("test_network", "antoon"); u == nil {
		t.Error("Should be able to find user who quit.")
	}

	if len(s.Messages) != 1 {
		t.Error("Should have a message.")
	}
}

func TestScanner_ParseMessage(t *testing.T) {
	t.Parallel()
	t.SkipNow()
}

func TestScanner_ParseLine(t *testing.T) {
	t.Parallel()
	t.SkipNow()
}
