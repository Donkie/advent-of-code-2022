package lib

import (
	"fmt"
	"math"
)

// Abs return the absolute value of the input value
func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// Max returns the maximum value of the inputs
func Max(nums ...int) int {
	curMax := math.MinInt
	for _, num := range nums {
		if num > curMax {
			curMax = num
		}
	}
	return curMax
}

// Min returns the minimum value of the inputs
func Min(nums ...int) int {
	curMin := math.MaxInt
	for _, num := range nums {
		if num < curMin {
			curMin = num
		}
	}
	return curMin
}

// Vector2 represents a 2D vector with integer values
type Vector2 struct {
	X int
	Y int
}

// String returns a string representation of the vector
func (v Vector2) String() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

// Equal returns whether this vector is equal to another one
func (v Vector2) Equal(v2 Vector2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

// Add returns a new vector where this vector is added to another one
func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

// Sub returns a new vector where this vector is subtracted by another one
func (v Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

// GetNormalized returns a "normalized" version of the vector
// Basically like a mathematical regular vector with length of 1, but it's locked in to the cardinal and diagonal directions.
func (v Vector2) GetNormalized() (out Vector2) {
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
	return
}

// Distance returns the Chebyshev distance between the two vectors
// Which is like the Manhattan distance except diagonal moves are also allowed
func (v Vector2) Distance(v2 Vector2) int {
	sub := v.Sub(v2)
	return Max(Abs(sub.X), Abs(sub.Y))
}

// Distance returns the Manhattan distance between the two vectors
func (v Vector2) ManhattanDistance(v2 Vector2) int {
	sub := v.Sub(v2)
	return Abs(sub.X) + Abs(sub.Y)
}

func (v Vector2) IsTouching(v2 Vector2) bool {
	return v.Distance(v2) <= 1
}
