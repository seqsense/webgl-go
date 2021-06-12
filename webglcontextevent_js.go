package webgl

import (
	"syscall/js"
)

type WebGLContextEvent struct {
	Event
	StatusMessage string
}

func parseWebGLContextEvent(event js.Value) WebGLContextEvent {
	return WebGLContextEvent{
		Event:         parseEvent(event),
		StatusMessage: event.Get("statusMessage").String(),
	}
}
