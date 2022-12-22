package main

import (
	"advent-of-code-2022/lib"
	"log"
)

type Square uint8

const (
	Void Square = iota
	Open
	Wall
)

type Direction uint8

const (
	Right Direction = iota
	Down
	Left
	Up
)

func (d Direction) Turn(by Direction) Direction {
	var add int
	if by == Left {
		add = -1
	} else if by == Right {
		add = 1
	} else {
		log.Panic("Invalid direction to turn by")
	}
	return Direction(lib.Mod(int(d)+add, 4))
}

func (d Direction) ToVec() lib.Vector2 {
	switch d {
	case Right:
		return lib.Vector2{X: 1, Y: 0}
	case Down:
		return lib.Vector2{X: 0, Y: 1}
	case Left:
		return lib.Vector2{X: -1, Y: 0}
	default:
		return lib.Vector2{X: 0, Y: -1}
	}
}

type Step struct {
	pos lib.Vector2
	dir Direction
}

type Operation struct {
	steps   int
	turnDir Direction
}

type Character struct {
	pos lib.Vector2
	dir Direction
}

type World struct {
	sqrs      [][]Square
	stepCache map[Step]lib.Vector2
	path      []Operation
	char      Character
}

func (w *World) Height() int {
	return len(w.sqrs)
}

func (w *World) Width() int {
	return len(w.sqrs[0])
}

func (w *World) Get(v lib.Vector2) Square {
	return w.sqrs[v.Y][v.X]
}

func (w *World) GetInitialPos() lib.Vector2 {
	for x := 0; x < w.Width(); x++ {
		test := lib.Vector2{X: x, Y: 0}
		if w.Get(test) == Open {
			return test
		}
	}
	log.Panic("Failed to find initial pos")
	return lib.Vector2{}
}

func (w *World) IsOOB(pos lib.Vector2) bool {
	return pos.X < 0 || pos.Y < 0 || pos.X >= w.Width() || pos.Y >= w.Height() || w.Get(pos) == Void
}

func (w *World) Step(pos lib.Vector2, dir Direction) lib.Vector2 {
	// First see if this step is trivial and not out of bounds
	targSqr := pos.Add(dir.ToVec())
	if !w.IsOOB(targSqr) {
		return targSqr
	}

	// Secondly see if we have done this step before
	step := Step{
		pos: pos,
		dir: dir,
	}
	target, ok := w.stepCache[step]
	if ok {
		return target
	}

	// If not, figure out the step destination
	origPos := pos
	dir = dir.Turn(Left).Turn(Left)
	for true {
		next := pos.Add(dir.ToVec())
		if w.IsOOB(next) {
			break
		} else {
			pos = next
		}
	}
	w.stepCache[Step{
		pos: origPos,
		dir: dir,
	}] = pos
	return pos
}

func (w *World) Walk(steps int) {
	for i := 0; i < steps; i++ {
		next := w.Step(w.char.pos, w.char.dir)
		if w.Get(next) == Wall {
			break
		}

		w.char.pos = next
	}
}

func (w *World) Traverse() {
	w.char.pos = w.GetInitialPos()
	w.char.dir = Right

	for _, op := range w.path {
		if op.steps == 0 {
			w.char.dir = w.char.dir.Turn(op.turnDir)
		} else {
			w.Walk(op.steps)
		}
	}
}

func (w *World) GetPassword() int {
	return (w.char.pos.Y+1)*1000 + (w.char.pos.X+1)*4 + int(w.char.dir)
}
