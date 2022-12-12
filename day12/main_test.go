package main

import "testing"

func TestExamplePart1(t *testing.T) {
	graph := ParseHeightMap("input_test.txt", StartToEnd)
	numSteps := graph.FindShortestPath()

	actual := numSteps
	expected := 31
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	graph := ParseHeightMap("input_test.txt", AnyToEnd)
	numSteps := graph.FindShortestPath()

	actual := numSteps
	expected := 29
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
