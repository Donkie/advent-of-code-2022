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
