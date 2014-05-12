package stats

type ConsecutiveLines struct {
	UserID   uint
	Count    uint
	TopUsers TopTokenArray
}

// NewConsecutiveLines
func NewConsecutiveLines() ConsecutiveLines {
	return ConsecutiveLines{
		TopUsers: make(TopTokenArray, 0, topTokenMaxSize),
	}
}

// addMessage
func (cl *ConsecutiveLines) addMessage(message *Message, user *User) {
	if cl.UserID == user.ID {
		cl.Count++

		if user.MaxConsecutive < cl.Count {
			user.MaxConsecutive = cl.Count
		}
	} else {
		cl.UserID = user.ID
		cl.Count = 1
	}

	cl.TopUsers.insert(user.Nick, cl.Count)
}
