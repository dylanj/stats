package stats

import (
	"testing"
	"time"
)

func TestNickReferencs(t *testing.T) {
	t.Parallel()
	s := NewStats()
	s.AddMessage(Msg, network, channel, "Scott", time.Now(), "Hey fish")
	s.AddMessage(Msg, network, channel, "fish", time.Now(), "Scott: Don't even talk to me...")

	n := s.GetNetwork(network)

	scott := n.users["scott"]
	fish := n.users["fish"]

	if len(scott.NickReferences) > 0 {
		t.Error("Scott should have no referenced nicks")
	}

	if len(fish.NickReferences) != 1 && fish.NickReferences["scott"] == 1 {
		t.Error("Fish should have references scotts nick")
	}

}
