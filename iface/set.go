package iface

import "github.com/google/uuid"

type Item interface {
	ID() uuid.UUID
}

type Set interface {
	Item
	Contains(Item) bool
	Add(Item) Set
	Delete(Item) Set
	Map(func(Item) Item) Set
	Plus(Set) Set
	Minus(Set) Set
}

type World interface {
	NewItem() Item
	NewSet() Set
}
