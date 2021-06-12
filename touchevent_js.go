package webgl

import (
	"syscall/js"
)

type TouchEvent struct {
	UIEvent

	ChangedTouches []Touch
	TargetTouches  []Touch
	Touches        []Touch

	AltKey   bool
	CtrlKey  bool
	ShiftKey bool
}

func parseTouchEvent(event js.Value) TouchEvent {
	return TouchEvent{
		UIEvent: UIEvent{
			Event: Event{
				event: event,
			},
		},
		ChangedTouches: parseTouches(event.Get("changedTouches")),
		TargetTouches:  parseTouches(event.Get("targetTouches")),
		Touches:        parseTouches(event.Get("touches")),
		AltKey:         event.Get("altKey").Bool(),
		CtrlKey:        event.Get("ctrlKey").Bool(),
		ShiftKey:       event.Get("shiftKey").Bool(),
	}
}

type Touch struct {
	Identifier       int
	ScreenX, ScreenY int
	ClientX, ClientY int
	PageX, PageY     int
}

func parseTouches(touches js.Value) []Touch {
	n := touches.Length()
	ts := make([]Touch, 0, n)
	for i := 0; i < n; i++ {
		t := touches.Index(i)
		ts = append(ts, Touch{
			Identifier: t.Get("identifier").Int(),
			ScreenX:    t.Get("screenX").Int(),
			ScreenY:    t.Get("screenY").Int(),
			ClientX:    t.Get("clientX").Int(),
			ClientY:    t.Get("clientY").Int(),
			PageX:      t.Get("pageX").Int(),
			PageY:      t.Get("pageY").Int(),
		})
	}
	return ts
}
