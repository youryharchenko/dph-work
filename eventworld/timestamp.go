package eventworld

import (
	"time"

	"github.com/youryharchenko/dph-work/iface"
)

type EventTimestamp struct {
	ts int64
}

func NewEventTimestamp() EventTimestamp {
	return EventTimestamp{ts: time.Now().UnixMicro()}
}

func NewEventMaxTimestamp() EventTimestamp {
	return EventTimestamp{ts: time.UnixMicro(1<<63 - 1).UnixMicro()}
}

func (ets EventTimestamp) Num() int64 {
	return ets.ts
}

func (ets EventTimestamp) Equal(ts iface.Timestamp) bool {
	return ets.ts == ts.Num()
}

func (ets EventTimestamp) Before(ts iface.Timestamp) bool {
	return ets.ts < ts.Num()
}

func (ets EventTimestamp) After(ts iface.Timestamp) bool {
	return ets.ts > ts.Num()
}
