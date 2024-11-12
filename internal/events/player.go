package events

import "arcadia/internal/models"

type PlayerEventData struct {
	Player *models.ComponentState `json:"player"`
}
