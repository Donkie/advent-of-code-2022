package lib

import "fmt"

func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Vector2 struct {
	X int
	Y int
}

func (v Vector2) String() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

func (v Vector2) Equal(v2 Vector2) bool {
	return v.X == v2.X && v.Y == v2.Y
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector2) Sub(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

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

func (v Vector2) IsTouching(v2 Vector2) bool {
	return v.Distance(v2) <= 1
}
