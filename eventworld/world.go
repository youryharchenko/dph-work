package eventworld

import (
	"github.com/google/uuid"
	"github.com/youryharchenko/dph-work/iface"
)

type EventWorld struct {
	iface.World
	id   uuid.UUID
	life iface.Interval
}

func NewEventWorld() (ew *EventWorld) {
	ew = new(EventWorld)
	ew.id = uuid.New()
	ew.life = NewEventInterval(NewEventTimestamp(), NewEventMaxTimestamp())
	return
}

func (ew *EventWorld) ID() uuid.UUID {
	return ew.id
}
