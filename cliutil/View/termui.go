package View

type (
	Attr uint16
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Text struct {
	posX   int
	posY   int
	text   []rune
	canvas []Cell
}

type Event struct {
	Type   EventType
	Key    Key
	Ch     rune
	KeyMod KeyMod
	Err    error
	MouseX int
	MouseY int
}

type (
	EventType uint8
	Key       uint16
	KeyMod    uint8
)
