package models

import "encoding/json"

const UpdateStateEvent = "UPDATE_STATE"

type GameState struct {
	Players   []ComponentState `json:"players"`
	NPCs      []ComponentState `json:"npcs"`
	Timestamp int              `json:"timestamp"`
}

func (e *GameState) ToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func GameStateModelFromMap(data map[string]interface{}) (*GameState, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var event GameState
	err = json.Unmarshal(jsonData, &event)
	return &event, err
}
