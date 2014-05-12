package stats

import "testing"

func TestChannel_Stringer(t *testing.T) {
	t.Parallel()

	c := &Channel{
		Name:       "foo",
		MessageIDs: []uint{1, 2, 3},
	}

	if c.String() != "Channel: foo, Messages: 3" {
		t.Error("Did not return correct string.")
	}
}
