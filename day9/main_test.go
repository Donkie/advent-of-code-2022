package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	actions := ParseActions("input_test.txt")
	rope := makeRope(2)
	rope.PerformActions(actions)
	visits := rope.GetNumTailUniquePositions()

	expected := 13
	actual := visits
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	actions := ParseActions("input_test.txt")
	rope := makeRope(10)
	rope.PerformActions(actions)
	visits := rope.GetNumTailUniquePositions()

	expected := 1
	actual := visits
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
