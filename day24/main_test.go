package main

import "testing"

func TestExamplePart1(t *testing.T) {
	world := ParseWorld("input_test.txt")
	bestTime := world.FindShortestPath()

	actual := bestTime
	expected := 18

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
