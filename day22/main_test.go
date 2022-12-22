package main

import "testing"

func TestExamplePart1(t *testing.T) {
	world := ParseWorld("input_test.txt")
	world.Traverse()
	pw := world.GetPassword()

	actual := pw
	expected := 6032
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
