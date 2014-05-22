package stats

import (
	"testing"
	"time"
)

func TestHourlyChart(t *testing.T) {
	t.Parallel()

	var chart HourlyChart

	for i := 23; i >= 0; i-- {
		for j := 0; j < i; j++ {
			date := time.Date(2014, time.April, 29, i, 30, 0, 1, time.UTC)

			m := &Message{
				Date: date,
			}

			chart.addMessage(m)
		}
	}

	for i := 0; i < 24; i++ {
		if chart[i] != i {
			t.Errorf("Hour[%d] has %d messages, expected: %d", i, chart[i], i)
		}
	}
}

func TestHourlyChartUpdates(t *testing.T) {
	t.Parallel()

	s := NewStats()
	n := s.addNetwork(network)
	c := s.addChannel(n, channel)
	u := s.addUser(n, nick)
	cu := u.addChannelUser(channel)

	date := time.Now()
	hour := date.Hour()

	if n.HourlyChart[hour] != 0 {
		t.Error("Network's chart should not have data in it")
	}

	if c.HourlyChart[hour] != 0 {
		t.Error("Channel's chart should not have data in it")
	}

	if u.HourlyChart[hour] != 0 {
		t.Error("User's chart should not have data in it")
	}

	s.addMessage(Msg, n, c, u, cu, date, "nihao")

	if n.HourlyChart[hour] != 1 {
		t.Error("Networks's chart should have data in it")
	}

	if c.HourlyChart[hour] != 1 {
		t.Error("Channel's chart should have data in it")
	}

	if u.HourlyChart[hour] != 1 {
		t.Error("User's chart should have data in it")
	}
}
