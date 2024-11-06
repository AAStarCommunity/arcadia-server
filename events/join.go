package events

const JoinEventName = "JOIN"

type JoinEventData struct {
	Name string `json:"name"`
	Skin string `json:"skin"`
}
