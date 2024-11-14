package game

import (
	"arcadia/internal/events"
	"arcadia/internal/game/components"
	"arcadia/internal/game/polo"
	"arcadia/internal/models"
	"arcadia/internal/utils"
	"math/rand"
	"sync"
)

type gameClient struct {
	mu      sync.Mutex
	players map[string]*components.Player
}

func (g *gameClient) GetPlayers() []models.ComponentState {
	g.mu.Lock()
	defer g.mu.Unlock()

	players := make([]models.ComponentState, 0, len(g.players))
	for _, player := range g.players {
		players = append(players, *player.State())
	}

	return players
}

const tileSize = 16

var game *gameClient = &gameClient{}

func (g *gameClient) Add(player *components.Player) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.players[player.ID()] = player
}

func JoinPlayer(evt *events.JoinEventData, client *polo.Client) {

	// initial position
	position := &utils.GameVector{
		X: float64((3 + rand.Intn(3)) * tileSize),
		Y: float64(11 * tileSize),
	}

	state := &models.ComponentState{
		ID:       client.ID(),
		Name:     evt.Name,
		Position: position,
		Size:     &utils.GameVector{X: 16, Y: 16},
		Life:     100,
		Properties: map[string]interface{}{
			"skin": evt.Skin,
		},
	}

	player := components.NewPlayer(client.ID(), state, client)

	game.Add(player)

	client.Send(
		events.JoinMapEventName,
		events.JoinMapEventData{
			State:   state,
			Players: game.GetPlayers(),
			NPCs:    []models.ComponentState{},
		},
	)
}
