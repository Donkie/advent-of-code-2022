package main

import "testing"

func TestExamplePart1(t *testing.T) {
	graph := ParseHeightMap("input_test.txt")
	numSteps := graph.FindShortestPath()

	actual := numSteps
	expected := 31
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
