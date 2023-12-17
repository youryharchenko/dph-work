package iface

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
