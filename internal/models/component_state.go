package models

import "arcadia/internal/utils"

type MoveDirection int

const (
	Left MoveDirection = iota
	Right
	Up
	Down
	UpLeft
	UpRight
	DownLeft
	DownRight
)

type ComponentState struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Position      *utils.GameVector      `json:"position"`
	Size          *utils.GameVector      `json:"size"`
	Life          int                    `json:"life"`
	Speed         float64                `json:"speed"`
	Direction     *MoveDirection         `json:"direction,omitempty"`
	LastDirection *MoveDirection         `json:"lastDirection,omitempty"`
	Action        *string                `json:"action,omitempty"`
	Properties    map[string]interface{} `json:"properties"`
}
