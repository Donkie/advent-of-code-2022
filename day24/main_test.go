package main

import "testing"

func TestExamplePart1(t *testing.T) {
	world := ParseWorld("input_test.txt")
	bestTime := world.FindShortestPath(false)

	actual := bestTime
	expected := 18

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	world := ParseWorld("input_test.txt")
	bestTime := world.FindShortestPath(true)

	actual := bestTime
	expected := 54

	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
