package game

import (
	"arcadia/internal/events"
	"sync"
)

type gameClient struct {
	mu sync.Mutex
	c  map[string]*PoloClient
}

func (g *gameClient) Add() {

}

func JoinPlayer(event *events.JoinEventData, client *PoloClient) {

}
