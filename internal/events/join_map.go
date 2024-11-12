package events

import "arcadia/internal/models"

const JoinMapEventName = "JOIN_MAP"

type MapModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type JoinMapEventData struct {
	State   *models.ComponentState  `json:"state"`
	Map     *MapModel               `json:"map"`
	Players []models.ComponentState `json:"players"`
	NPCs    []models.ComponentState `json:"npcs"`
}
