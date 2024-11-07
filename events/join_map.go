package events

const JoinMapEventName = "JOIN_MAP"


type MapModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type JoinMapEvent struct {
	State   ComponentStateModel   `json:"state"`
	Map     MapModel              `json:"map"`
	Players []ComponentStateModel `json:"players"`
	NPCs    []ComponentStateModel `json:"npcs"`
}