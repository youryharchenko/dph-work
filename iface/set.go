package iface

import "github.com/google/uuid"

type Timestamp interface {
	Num() int64
	Equal(Timestamp) bool
	Before(Timestamp) bool
	After(Timestamp) bool
}

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

type Item interface {
	ID() uuid.UUID
	Life() Interval
}

type Set interface {
	Item
	Contains(Item) []Interval
	Add(Item) Set
	Delete(Item) Set
	Map(func(Item) Item) Set
	Plus(Set) Set
	Minus(Set) Set
	Intersection(Set) Interval
}

type World interface {
	Item
	NewItem() Item
	NewSet() Set
}
