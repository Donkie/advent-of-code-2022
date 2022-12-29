package main

import "advent-of-code-2022/lib"

type Direction uint8

const (
	None  Direction = 0b0000
	North Direction = 0b0001
	South Direction = 0b0010
	West  Direction = 0b0100
	East  Direction = 0b1000
)

func (d Direction) ToVec() lib.Vector2 {
	switch d {
	case North:
		return lib.Vector2{X: 0, Y: -1}
	case South:
		return lib.Vector2{X: 0, Y: 1}
	case West:
		return lib.Vector2{X: -1, Y: 0}
	case East:
		return lib.Vector2{X: 1, Y: 0}
	default:
		return lib.Vector2{X: 0, Y: 0}
	}
}
