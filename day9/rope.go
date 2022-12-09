package main

import "advent-of-code-2022/lib"

// Dir represents a cardinal movement direction
type Dir int64

const (
	Up Dir = iota
	Right
	Down
	Left
)

// Action represents a movement action
type Action struct {
	dir Dir
}

// getMovementVector returns an equivalent 2D vector for the given action movement direction
func (act Action) getMovementVector() lib.Vector2 {
	switch act.dir {
	case Up:
		return lib.Vector2{X: 0, Y: 1}
	case Right:
		return lib.Vector2{X: 1, Y: 0}
	case Down:
		return lib.Vector2{X: 0, Y: -1}
	default:
		return lib.Vector2{X: -1, Y: 0}
	}
}

// Rope represents a rope with a number of rope pieces
// Also has a position recorder which records the position of the tail end of the rope
type Rope struct {
	pieces   []lib.Vector2
	recorder Recorder
}

// Head returns the position of the head piece of the rope
func (rope Rope) Head() lib.Vector2 {
	return rope.pieces[0]
}

// Tail returns the position of the tail piece of the rope
func (rope Rope) Tail() lib.Vector2 {
	return rope.pieces[len(rope.pieces)-1]
}

// makeRope creates a new rope with a number of pieces
func makeRope(numPieces int) (rope Rope) {
	rope.pieces = make([]lib.Vector2, numPieces)
	rope.recorder.RecordPosition(rope.Tail())
	return
}

// moveTail iteratively moves all the rope pieces towards the head if needed
func (rope *Rope) moveTail() {
	for i := 1; i < len(rope.pieces); i++ {
		p1 := rope.pieces[i-1]
		p2 := rope.pieces[i]
		if p1.IsTouching(p2) {
			return
		}

		diff := p1.Sub(p2).GetNormalized()
		rope.pieces[i] = rope.pieces[i].Add(diff)
	}

	rope.recorder.RecordPosition(rope.Tail())
}

// PerformAction performs a movement action on the head of the rope
// All subsequent rope pieces will be moved along with it
func (rope *Rope) PerformAction(action Action) {
	movement := action.getMovementVector()
	rope.pieces[0] = rope.pieces[0].Add(movement)
	rope.moveTail()
}

// PerformActions performs a list of actions sequentially
func (rope *Rope) PerformActions(actions []Action) {
	for _, act := range actions {
		rope.PerformAction(act)
	}
}

// GetNumTailUniquePositions returns the number of unique positions that the tail of the rope has had
func (rope Rope) GetNumTailUniquePositions() int {
	return rope.recorder.GetNumOfRecordedPositions()
}
