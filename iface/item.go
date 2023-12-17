package iface

import "github.com/google/uuid"

type Item interface {
	ID() uuid.UUID
	Life() Interval
}
