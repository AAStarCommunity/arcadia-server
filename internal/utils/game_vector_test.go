package utils

import (
	"testing"
)

func TestGameVectorAll(t *testing.T) {
	vector := GameVector{X: 3.0, Y: 4.0}

	mapData, err := vector.ToMap()
	if err != nil {
		t.Error("Error serializing vector:", err)
	}

	newVector, err := GameVectorFromMap(mapData)
	if err != nil {
		t.Error("Error deserializing vector:", err)
	}
	if newVector.X != vector.X || newVector.Y != vector.Y {
		t.Error("Deserialized vector is incorrect")
	}

	// 其他方法示例
	scaledVector := vector.Scaled(2.0)
	if scaledVector.X != 6.0 || scaledVector.Y != 8.0 {
		t.Error("Scaled vector is incorrect")
	}

	vector.Negate()
	if vector.X != -3.0 || vector.Y != -4.0 {
		t.Error("Negated vector is incorrect")
	}

	vector.Absolute()
	if vector.X != 3.0 || vector.Y != 4.0 {
		t.Error("Absolute vector is incorrect")
	}

	vector.Add(GameVector{X: 1.0, Y: 1.0})
	if vector.X != 4.0 || vector.Y != 5.0 {
		t.Error("Added vector is incorrect")
	}

	vector.AddScaled(GameVector{X: 1.0, Y: 1.0}, 2.0)
	if vector.X != 6.0 || vector.Y != 7.0 {
		t.Error("Added scaled vector is incorrect")
	}

	vector.Sub(GameVector{X: 1.0, Y: 1.0})
	if vector.X != 5.0 || vector.Y != 6.0 {
		t.Error("Subtracted vector is incorrect")
	}

	clonedVector := vector.Clone()
	if clonedVector.X != 5.0 || clonedVector.Y != 6.0 {
		t.Error("Cloned vector is incorrect")
	}

	vector.Scale(2.0)
	if vector.X != 10.0 || vector.Y != 12.0 {
		t.Error("Scaled vector is incorrect")
	}
}
