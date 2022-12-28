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

// A check returns its direction if it's free to move to
var checks = []func(m Map, p lib.Vector2) Direction{
	func(m Map, p lib.Vector2) Direction {
		if m.Get(p.Add(lib.Vector2{X: -1, Y: -1})) != nil ||
			m.Get(p.Add(lib.Vector2{X: 0, Y: -1})) != nil ||
			m.Get(p.Add(lib.Vector2{X: 1, Y: -1})) != nil {
			return None
		} else {
			return North
		}
	},
	func(m Map, p lib.Vector2) Direction {
		if m.Get(p.Add(lib.Vector2{X: -1, Y: 1})) != nil ||
			m.Get(p.Add(lib.Vector2{X: 0, Y: 1})) != nil ||
			m.Get(p.Add(lib.Vector2{X: 1, Y: 1})) != nil {
			return None
		} else {
			return South
		}
	},
	func(m Map, p lib.Vector2) Direction {
		if m.Get(p.Add(lib.Vector2{X: -1, Y: -1})) != nil ||
			m.Get(p.Add(lib.Vector2{X: -1, Y: 0})) != nil ||
			m.Get(p.Add(lib.Vector2{X: -1, Y: 1})) != nil {
			return None
		} else {
			return West
		}
	},
	func(m Map, p lib.Vector2) Direction {
		if m.Get(p.Add(lib.Vector2{X: 1, Y: -1})) != nil ||
			m.Get(p.Add(lib.Vector2{X: 1, Y: 0})) != nil ||
			m.Get(p.Add(lib.Vector2{X: 1, Y: 1})) != nil {
			return None
		} else {
			return East
		}
	},
}

type Checker struct {
	startIndex int
}

func (c *Checker) IncrementStart() {
	c.startIndex = (c.startIndex + 1) % 4
}

func (c *Checker) GetDirection(m Map, p lib.Vector2) (bestDir Direction, allDirs Direction) {
	idx := c.startIndex
	for i := 0; i < 4; i++ {
		dir := checks[idx](m, p)
		allDirs |= dir
		if dir != None && bestDir == 0 {
			bestDir = dir
		}
		idx = (idx + 1) % 4
	}
	return
}
