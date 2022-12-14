package main

import "testing"

func TestExamplePart1(t *testing.T) {
	world := ParseWorld("input_test.txt")
	restingSand := world.SimulateSand(false)

	expected := 24
	actual := restingSand
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	world := ParseWorld("input_test.txt")
	world.HasFloor = true
	restingSand := world.SimulateSand(false)
	world.Print()

	expected := 93
	actual := restingSand
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
