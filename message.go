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

func (m *Message) String() string {
	//return fmt.Sprintf("%s (%s): %s\n", m.User.Name, m.Channel.Name, m.Message)
	return fmt.Sprintf("hello")
}
