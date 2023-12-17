package eventworld

import (
	"testing"
	"time"

	"github.com/youryharchenko/dph-work/iface"
)

func TestNewEventInterval(t *testing.T) {
	var ts1 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts2 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts3 iface.Timestamp = NewEventTimestamp()

	var ei0 iface.Interval = NewEmptyEventInterval()

	var ei1 iface.Interval = NewEventInterval(ts1, ts2)
	var ei2 iface.Interval = NewEventInterval(ts1, ts2)

	var ei3 iface.Interval = NewEventInterval(ts1, ts3)

	if !ei0.IsEmpty() {
		t.Error("ei0.IsEmpty()")
	}

	if ei1.IsEmpty() {
		t.Error("ei1.IsEmpty()")
	}

	if !ei1.Equal(ei2) {
		t.Error("ei1 is not equal ei2")
	}

	if ei1.Equal(ei3) {
		t.Error("ei1 is equal ei3")
	}

}

func TestContainsInterval(t *testing.T) {
	var ts1 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts2 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts3 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts4 iface.Timestamp = NewEventTimestamp()

	var ei1 iface.Interval = NewEventInterval(ts1, ts2)
	var ei2 iface.Interval = NewEventInterval(ts1, ts2)

	var ei3 iface.Interval = NewEventInterval(ts1, ts3)
	var ei4 iface.Interval = NewEventInterval(ts2, ts4)

	var ei5 iface.Interval = NewEventInterval(ts1, ts4)
	var ei6 iface.Interval = NewEventInterval(ts2, ts3)

	if !ei4.Contains(ei4) {
		t.Error("ei4 does not contain ei4")
	}

	if !ei1.Contains(ei2) {
		t.Error("ei2 does not contain ei2")
	}

	if ei3.Contains(ei4) {
		t.Error("ei3 contains ei4")
	}

	if !ei5.Contains(ei3) {
		t.Error("ei5 does not contain ei3")
	}

	if !ei5.Contains(ei4) {
		t.Error("ei5 does not contain ei4")
	}

	if !ei5.Contains(ei6) {
		t.Error("ei5 does not contain ei6")
	}

}

func TestIntersectionInterval(t *testing.T) {
	var ts1 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts2 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts3 iface.Timestamp = NewEventTimestamp()
	time.Sleep(10 * time.Microsecond)
	var ts4 iface.Timestamp = NewEventTimestamp()

	var ei0 iface.Interval = NewEmptyEventInterval()

	var ei1 iface.Interval = NewEventInterval(ts1, ts2)
	var ei2 iface.Interval = NewEventInterval(ts1, ts2)

	var ei3 iface.Interval = NewEventInterval(ts1, ts3)
	var ei4 iface.Interval = NewEventInterval(ts2, ts4)

	var ei5 iface.Interval = NewEventInterval(ts1, ts4)
	var ei6 iface.Interval = NewEventInterval(ts2, ts3)

	var ei7 iface.Interval = NewEventInterval(ts3, ts4)

	if !ei1.Equal(ei1.Intersection(ei2)) {
		t.Error("ei1.Equal(ei1.Intersection(ei2))")
	}

	if !ei6.Equal(ei5.Intersection(ei6)) {
		t.Error("ei6.Equal(ei5.Intersection(ei6))")
	}

	if !ei6.Equal(ei3.Intersection(ei4)) {
		t.Error("ei6.Equal(ei3.Intersection(ei4))")
	}

	if !ei6.Equal(ei4.Intersection(ei3)) {
		t.Error("ei6.Equal(ei4.Intersection(ei3))")
	}

	if !ei0.Equal(ei1.Intersection(ei0)) {
		t.Error("ei0.Equal(ei1.Intersection(ei0))")
	}

	if !ei0.Equal(ei0.Intersection(ei1)) {
		t.Error("ei0.Equal(ei0.Intersection(ei1))")
	}

	if !ei0.Equal(ei1.Intersection(ei7)) {
		//log.Println(ei0)
		//log.Println("ei1", ei1)
		//log.Println("ei7", ei7)
		//log.Println(ei1.Intersection(ei7))
		t.Error("ei0.Equal(ei1.Intersection(ei7))")
	}

	if !ei0.Equal(ei7.Intersection(ei1)) {
		//log.Println(ei0)
		//log.Println("ei1", ei1)
		//log.Println("ei7", ei7)
		//log.Println(ei7.Intersection(ei1))
		t.Error("ei0.Equal(ei1.Intersection(ei7))")
	}

	if !ei6.Intersection(ei7).IsPoint() {
		t.Error("!ei6.Intersection(ei7).IsPoint()")
	}

	if ei1.Intersection(ei7).IsPoint() {
		t.Error("ei1.Intersection(ei7).IsPoint() ")
	}

}
