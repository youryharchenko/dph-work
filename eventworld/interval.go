package eventworld

import (
	"fmt"

	"github.com/youryharchenko/dph-work/iface"
)

type EventInterval struct {
	isEmpty bool
	b       iface.Timestamp
	e       iface.Timestamp
}

func NewEventInterval(b iface.Timestamp, e iface.Timestamp) EventInterval {
	if b.After(e) {
		return EventInterval{b: e, e: b}
	}
	return EventInterval{b: b, e: e}
}

func NewEmptyEventInterval() EventInterval {
	return EventInterval{isEmpty: true}
}

func (ei EventInterval) String() string {
	if ei.IsEmpty() {
		return "Interval{Empty}"
	}
	return fmt.Sprintf("Interval{Start: %d, End: %d}", ei.b.Num(), ei.e.Num())
}

func (ei EventInterval) Equal(i iface.Interval) bool {

	return ei.IsEmpty() && i.IsEmpty() || ei.SameStart(i) && ei.SameEnd(i)
}

func (ei EventInterval) Contains(i iface.Interval) bool {
	return i.IsEmpty() || ei.IsEmpty() && i.IsEmpty() ||
		(ei.SameStart(i) || ei.StartedEarlier(i)) && (ei.SameEnd(i) || ei.EndedLater(i))
}

func (ei EventInterval) SameStart(i iface.Interval) bool {
	return !ei.IsEmpty() && !i.IsEmpty() && ei.Start().Equal(i.Start())
}

func (ei EventInterval) SameEnd(i iface.Interval) bool {
	return !ei.IsEmpty() && !i.IsEmpty() && ei.End().Equal(i.End())
}

func (ei EventInterval) StartedEarlier(i iface.Interval) bool {
	return !ei.IsEmpty() && !i.IsEmpty() && ei.Start().Before(i.Start())
}

func (ei EventInterval) StartedLater(i iface.Interval) bool {
	return !ei.IsEmpty() && !i.IsEmpty() && ei.Start().After(i.Start())
}

func (ei EventInterval) EndedEarlier(i iface.Interval) bool {
	return !ei.IsEmpty() && !i.IsEmpty() && ei.End().Before(i.End())
}

func (ei EventInterval) EndedLater(i iface.Interval) bool {
	return !ei.IsEmpty() && !i.IsEmpty() && ei.End().After(i.End())
}

func (ei EventInterval) Start() iface.Timestamp {
	return ei.b
}

func (ei EventInterval) End() iface.Timestamp {
	return ei.e
}

func (ei EventInterval) IsEmpty() bool {
	return ei.isEmpty
}

func (ei EventInterval) IsPoint() bool {
	return !ei.IsEmpty() && ei.Start().Equal(ei.End())
}

func (ei EventInterval) Intersection(i iface.Interval) iface.Interval {
	if ei.IsEmpty() || i.IsEmpty() {
		return NewEmptyEventInterval()
	}
	if ei.Contains(i) {
		return NewEventInterval(i.Start(), i.End())
	}
	if i.Contains(ei) {
		return NewEventInterval(ei.Start(), ei.End())
	}

	if (ei.SameStart(i) || ei.StartedEarlier(i)) &&
		(ei.SameEnd(i) || ei.EndedEarlier(i)) &&
		(i.Start().Before(ei.End()) || i.Start().Equal(ei.End())) {
		//log.Println("(ei.SameStart(i) || ei.StartedEarlier(i)) && (ei.SameEnd(i) || ei.EndedEarlier(i))")
		return NewEventInterval(i.Start(), ei.End())
	}

	if (i.SameStart(ei) || i.StartedEarlier(ei)) &&
		(i.SameEnd(ei) || i.EndedEarlier(ei)) &&
		(ei.Start().Before(i.End()) || ei.Start().Equal(i.End())) {
		return NewEventInterval(ei.Start(), i.End())
	}

	return NewEmptyEventInterval()
}
