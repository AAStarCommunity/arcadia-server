package events

type EventData[T any] struct {
	Event string `json:"event"`
	Data  T      `json:"data"`
}
