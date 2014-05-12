package stats

import "time"

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
