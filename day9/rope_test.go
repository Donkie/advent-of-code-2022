package main

import (
	"advent-of-code-2022/lib"
	"testing"
)

func TestRopeMove(t *testing.T) {
	moves := []struct {
		dir        Dir
		newTailPos lib.Vector2
	}{
		{Down, lib.Vector2{X: 0, Y: 0}},
		{Down, lib.Vector2{X: 0, Y: -1}},
		{Right, lib.Vector2{X: 0, Y: -1}},
		{Up, lib.Vector2{X: 0, Y: -1}},
		{Right, lib.Vector2{X: 1, Y: -1}},
		{Right, lib.Vector2{X: 2, Y: -1}},
		{Up, lib.Vector2{X: 2, Y: -1}},
		{Up, lib.Vector2{X: 3, Y: 0}},
		{Up, lib.Vector2{X: 3, Y: 1}},
		{Left, lib.Vector2{X: 3, Y: 1}},
		{Left, lib.Vector2{X: 2, Y: 2}},
	}

	rope := makeRope(2)

	for idx, move := range moves {
		rope.PerformAction(Action{dir: move.dir})
		actual := rope.Tail()
		if !actual.Equal(move.newTailPos) {
			t.Errorf("Step %d. Expected: %s, Actual: %s", idx, move.newTailPos.String(), actual.String())
			return
		}
	}
}
