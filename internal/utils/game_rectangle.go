package utils

type GameRect struct {
	Position *GameVector `json:"position"`
	Size     *GameVector `json:"size"`
}

func NewGameRect(position *GameVector, size *GameVector) *GameRect {
	if position == nil {
		position = &GameVector{X: 0, Y: 0}
	}
	return &GameRect{
		Position: position,
		Size:     size,
	}
}

func NewGameRectZero() *GameRect {
	return &GameRect{
		Position: &GameVector{X: 0, Y: 0},
		Size:     &GameVector{X: 0, Y: 0},
	}
}

func (r *GameRect) TopLeft() *GameVector {
	return r.Position
}

func (r *GameRect) TopRight() *GameVector {
	return &GameVector{X: r.Right(), Y: r.Top()}
}

func (r *GameRect) BottomRight() *GameVector {
	return &GameVector{X: r.Right(), Y: r.Bottom()}
}

func (r *GameRect) BottomLeft() *GameVector {
	return &GameVector{X: r.Left(), Y: r.Bottom()}
}

func (r *GameRect) Overlaps(other *GameRect) bool {
	if r.Right() <= other.Left() || other.Right() <= r.Left() {
		return false
	}
	if r.Bottom() <= other.Top() || other.Bottom() <= r.Top() {
		return false
	}
	return true
}

func (r *GameRect) Left() float64 {
	return r.Position.X
}

func (r *GameRect) Right() float64 {
	return r.Position.X + r.Size.X
}

func (r *GameRect) Top() float64 {
	return r.Position.Y
}

func (r GameRect) Bottom() float64 {
	return r.Position.Y + r.Size.Y
}
