package utils

import (
	"testing"
)

func TestGameRectangle(t *testing.T) {
	rect1 := NewGameRect(&GameVector{X: 0, Y: 0}, &GameVector{X: 10, Y: 10})
	rect2 := NewGameRect(&GameVector{X: 5, Y: 5}, &GameVector{X: 10, Y: 10})

	if !rect1.Overlaps(rect2) {
		t.Error("Rectangles should overlap")
	}

	rect3 := NewGameRect(&GameVector{X: 0, Y: 0}, &GameVector{X: 10, Y: 10})
	rect4 := NewGameRect(&GameVector{X: 15, Y: 15}, &GameVector{X: 100, Y: 100})

	if rect3.Overlaps(rect4) {
		t.Error("Rectangles should not overlap")
	}
}
