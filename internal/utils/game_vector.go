package utils

import (
	"encoding/json"
	"math"
)

type GameVector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v GameVector) Scaled(arg float64) GameVector {
	return GameVector{X: v.X * arg, Y: v.Y * arg}
}

func (v *GameVector) Negate() {
	v.X = -v.X
	v.Y = -v.Y
}

func (v *GameVector) Absolute() {
	v.X = math.Abs(v.X)
	v.Y = math.Abs(v.Y)
}

func (v *GameVector) Add(arg GameVector) {
	v.X += arg.X
	v.Y += arg.Y
}

func (v *GameVector) AddScaled(arg GameVector, factor float64) {
	v.X += arg.X * factor
	v.Y += arg.Y * factor
}

func (v *GameVector) Sub(arg GameVector) {
	v.X -= arg.X
	v.Y -= arg.Y
}

func (v GameVector) ToMap() (map[string]interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func GameVectorFromMap(data map[string]interface{}) (GameVector, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return GameVector{}, err
	}
	var vector GameVector
	err = json.Unmarshal(jsonData, &vector)
	return vector, err
}

func (v GameVector) Clone() GameVector {
	return GameVector{X: v.X, Y: v.Y}
}

func (v *GameVector) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}
