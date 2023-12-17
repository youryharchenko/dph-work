package eventworld

import (
	"testing"
	"time"

	"github.com/youryharchenko/dph-work/iface"
)

func TestNewEventTimestamp(t *testing.T) {
	var ts1 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts2 iface.Timestamp = NewEventTimestamp()

	if ts1.Equal(ts2) {
		t.Error("ts1 is equal ts2")
	}

	if ts1.After(ts2) {
		t.Error("ts1 after ts2")
	}

	if ts2.Before(ts1) {
		t.Error("ts2 before ts1")
	}

}

func TestNewEventMaxTimestamp(t *testing.T) {
	maxDate := time.UnixMicro(1<<63 - 1)
	if NewEventMaxTimestamp().Num() != maxDate.UnixMicro() {
		t.Error("NewEventMaxTimestamp().Num() != maxDate.UnixMicro()")
	}

}
