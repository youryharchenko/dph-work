package iface

type World interface {
	Item
	NewItem() Item
	NewSet() Set
}
