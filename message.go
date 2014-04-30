package stats

import "time"
import "fmt"

// MsgKind is the type of message
type MsgKind int

// These are the various message Kinds
const (
	// Msg is for PRIVMSG and NOTICE messages
	Msg MsgKind = iota
	// Part is for ...
	Part
	// Join is for ...
	Join
	// Quit is for ...
	Quit
	// Kick is for ...
	Kick
)

type Message struct {
	ID        uint
	Date      time.Time
	UserID    uint
	ChannelID uint
	Message   string
	Kind      MsgKind
}

func NewMessageString(message string) *Message {
	return NewMessage([]byte(message))
}

func NewMessage(message []byte) *Message {
	m := Message{
		Message: string(message),
	}

	return &m
}

func (m *Message) String() string {
	//return fmt.Sprintf("%s (%s): %s\n", m.User.Name, m.Channel.Name, m.Message)
	return fmt.Sprintf("hello")
}

// date hax
// 2013-08-07 16:49:40	-->	dylan (dylan@zqz.ca) has joined #deviate
// 2006-01-02 15:04:04
