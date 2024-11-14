package components

import (
	"arcadia/internal/game/polo"
	"arcadia/internal/models"
)

type Player struct {
	id     string
	state  *models.ComponentState
	client *polo.Client
}

func (p *Player) ID() string {
	return p.id
}

func (p *Player) State() *models.ComponentState {
	return p.state
}

func NewPlayer(id string, state *models.ComponentState, client *polo.Client) *Player {
	return &Player{
		id:     id,
		state:  state,
		client: client,
	}
}
