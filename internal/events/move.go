package events

import (
	"arcadia/internal/models"
	"arcadia/internal/utils"
)

const MoveEvent = "MOVE"

type MoveEventData struct {
	Position  *utils.GameVector    `json:"position"`
	Time      string               `json:"time"`
	Direction models.MoveDirection `json:"direction,omitempty"`
	MapID     string               `json:"map"`
}
