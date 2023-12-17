package iface

type Timestamp interface {
	Num() int64
	Equal(Timestamp) bool
	Before(Timestamp) bool
	After(Timestamp) bool
}
