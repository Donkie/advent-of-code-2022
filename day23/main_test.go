package main

import "testing"

func TestExamplePart1(t *testing.T) {
	world := ParseWorld("input_test.txt")
	world.Simulate(10)
	emptyTiles := world.GetEmptySquaresInBoundingRegion()

	actual := emptyTiles
	expected := 110

	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart1Smaller(t *testing.T) {
	world := ParseWorld("input_test2.txt")
	world.Simulate(3)
	emptyTiles := world.GetEmptySquaresInBoundingRegion()

	actual := emptyTiles
	expected := 25

	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	world := ParseWorld("input_test.txt")
	round := world.SimulateUntilEnd()

	actual := round
	expected := 20

	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
