package lib

import "fmt"

// Vector3 represents a 3D vector with integer values
type Vector3 struct {
	X int
	Y int
	Z int
}

// String returns a string representation of the vector
func (v Vector3) String() string {
	return fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z)
}

// Equal returns whether this vector is equal to another one
func (v Vector3) Equal(v2 Vector3) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v.Z
}

// Add returns a new vector where this vector is added to another one
func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
		Z: v.Z + v2.Z,
	}
}

// Sub returns a new vector where this vector is subtracted by another one
func (v Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
		Z: v.Z - v2.Z,
	}
}

// GetNormalized returns a "normalized" version of the vector
// Basically like a mathematical regular vector with length of 1, but it's locked in to the cardinal and diagonal directions.
func (v Vector3) GetNormalized() (out Vector3) {
	if Abs(v.X) > 0 {
		out.X = v.X / Abs(v.X)
	} else {
		out.X = 0
	}
	if Abs(v.Y) > 0 {
		out.Y = v.Y / Abs(v.Y)
	} else {
		out.Y = 0
	}
	if Abs(v.Z) > 0 {
		out.Z = v.Z / Abs(v.Z)
	} else {
		out.Z = 0
	}
	return
}

// Distance returns the Chebyshev distance between the two vectors
// Which is like the Manhattan distance except diagonal moves are also allowed
func (v Vector3) Distance(v2 Vector3) int {
	sub := v.Sub(v2)
	return Max(Abs(sub.X), Abs(sub.Y), Abs(sub.Z))
}

// Distance returns the Manhattan distance between the two vectors
func (v Vector3) ManhattanDistance(v2 Vector3) int {
	sub := v.Sub(v2)
	return Abs(sub.X) + Abs(sub.Y) + Abs(sub.Z)
}

func (v Vector3) IsTouching(v2 Vector3) bool {
	return v.Distance(v2) <= 1
}
