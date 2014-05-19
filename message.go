package stats

import "time"

// MsgKind is the type of message
type MsgKind int

// These are the various message Kinds
const (
	// Msg is for PRIVMSG and NOTICE messages
	Msg MsgKind = iota
	Part
	Join
	Quit
	Kick
	Mode
	Topic
	Action
)

type Message struct {
	ID        uint
	Date      time.Time
	UserID    uint
	ChannelID uint
	Message   string
	Kind      MsgKind
}
