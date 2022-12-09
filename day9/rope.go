package main

import "advent-of-code-2022/lib"

type Dir int64

const (
	Up Dir = iota
	Right
	Down
	Left
)

type Action struct {
	dir Dir
}

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

type Rope struct {
	head     lib.Vector2
	tail     lib.Vector2
	recorder Recorder
}

func makeRope(head lib.Vector2, tail lib.Vector2) (rope Rope) {
	rope.head = head
	rope.tail = tail
	rope.recorder.RecordPosition(tail)
	return
}

func (rope *Rope) moveTail() {
	if rope.tail.IsTouching(rope.head) {
		return
	}

	diff := rope.head.Sub(rope.tail).GetNormalized()
	rope.tail = rope.tail.Add(diff)
	rope.recorder.RecordPosition(rope.tail)
}

func (rope *Rope) PerformAction(action Action) {
	movement := action.getMovementVector()
	rope.head = rope.head.Add(movement)
	rope.moveTail()
}

func (rope *Rope) PerformActions(actions []Action) {
	for _, act := range actions {
		rope.PerformAction(act)
	}
}

func (rope Rope) GetNumTailUniquePositions() int {
	return rope.recorder.GetNumOfRecordedPositions()
}
