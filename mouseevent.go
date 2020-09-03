package webgl

type MouseButton int

const (
	MouseButtonNull MouseButton = -1
)

type MouseEvent struct {
	UIEvent
	ClientX, ClientY int
	Button           MouseButton
}
