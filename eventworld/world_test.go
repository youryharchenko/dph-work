package eventworld

import (
	"testing"

	"github.com/youryharchenko/dph-work/iface"
)

func TestNewEventWorld(t *testing.T) {
	var w iface.World = NewEventWorld()

	if w.ID().String() == "" {
		t.Error("w.ID().String() is empty")
	}

}
