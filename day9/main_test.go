package main

import (
	"advent-of-code-2022/lib"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	actions := ParseActions("input_test.txt")
	rope := makeRope(lib.Vector2{}, lib.Vector2{})
	rope.PerformActions(actions)
	visits := rope.GetNumTailUniquePositions()

	expected := 13
	actual := visits
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
