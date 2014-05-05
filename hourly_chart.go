package stats

type HourlyChart [24]int

// addMessage adds a message to the chart
func (h *HourlyChart) addMessage(m *Message) {
	hour := m.Date.Hour()
	h[hour]++
}
