package main

import "testing"

func TestExamplePart1(t *testing.T) {
	jetStream := ParseJetStream("input_test.txt")
	world := makeRockFallingWorld(*jetStream)
	world.Simulate(2022)
	height := world.GetHighestPoint()

	actual := height
	expected := 3068
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	jetStream := ParseJetStream("input_test.txt")
	world := makeRockFallingWorld(*jetStream)
	world.Simulate(1000000000000)
	height := world.GetHighestPoint()

	actual := height
	expected := 1514285714288
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
