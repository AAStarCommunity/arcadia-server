package events

type EventData[T any] struct {
	Event string `json:"event"`
	Data  T      `json:"data"`
}

type MoveDirection int

const (
	Left MoveDirection = iota
	Right
	Up
	Down
	UpLeft
	UpRight
	DownLeft
	DownRight
)

type GameVector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ComponentStateModel struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Position      GameVector             `json:"position"`
	Size          GameVector             `json:"size"`
	Life          int                    `json:"life"`
	Speed         float64                `json:"speed"`
	Direction     *MoveDirection         `json:"direction,omitempty"`
	LastDirection *MoveDirection         `json:"lastDirection,omitempty"`
	Action        *string                `json:"action,omitempty"`
	Properties    map[string]interface{} `json:"properties"`
}
