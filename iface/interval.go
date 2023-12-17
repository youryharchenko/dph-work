package iface

type Interval interface {
	Start() Timestamp
	End() Timestamp
	Equal(Interval) bool
	Contains(Interval) bool
	SameStart(Interval) bool
	SameEnd(Interval) bool
	StartedEarlier(Interval) bool
	StartedLater(Interval) bool
	EndedEarlier(Interval) bool
	EndedLater(Interval) bool
	//Plus(Interval) Interval
	//Minus(Interval) Interval
	Intersection(Interval) Interval
	IsEmpty() bool
	IsPoint() bool
	String() string
}
