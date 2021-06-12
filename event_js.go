package webgl

import (
	"syscall/js"
)

type UIEvent struct {
	Event
}

type Event struct {
	event  js.Value
	Target js.Value
}

func (e Event) PreventDefault() {
	e.event.Call("preventDefault")
}

func (e Event) StopPropagation() {
	e.event.Call("stopPropagation")
}

func NewEvent(typ string) Event {
	return Event{
		event: js.Global().Get("Event").New(typ),
	}
}

func parseEvent(event js.Value) Event {
	return Event{
		event:  event,
		Target: event.Get("target"),
	}
}
