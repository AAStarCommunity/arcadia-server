package events

type MoveEvent struct {
	Position  GameVector     `json:"position"`
	Time      string         `json:"time"`
	Direction *MoveDirection `json:"direction,omitempty"`
	MapID     string         `json:"map"`
}
