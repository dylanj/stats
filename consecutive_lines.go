package stats

type ConsecutiveLines struct {
	UserID uint
	Count  uint
}

// addMessage
func (c *ConsecutiveLines) addMessage(m *Message, u *User) {
	if m.Kind != Msg {
		return
	}

	if c.UserID == u.ID {
		c.Count++

		if u.MaxConsecutive < c.Count {
			u.MaxConsecutive = c.Count
		}
	}
}
