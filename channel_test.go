package stats

import "testing"

func TestChannel_GetName (t *testing.T) {
  t.Parallel()

  channel := Channel{Name: "foo"}

  if channel.GetName() != "foo" {
    t.Error("GetName() doesn't return correct name")
  }
}
