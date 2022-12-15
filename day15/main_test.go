package main

import (
	"advent-of-code-2022/lib"
	"testing"
)

func TestExamplePart1(t *testing.T) {
	world := ParseSensorReadings("input_test.txt")
	num := world.GetPositionsCoveredWithNoBeacon(10)

	expected := 26
	actual := num

	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	world := ParseSensorReadings("input_test.txt")
	pos := world.GetUncoveredPosition(0, 20, 0, 20)

	expected := lib.Vector2{X: 14, Y: 11}
	actual := pos

	if !actual.Equal(expected) {
		t.Errorf("expected %s, actual %s", expected.String(), actual.String())
	}
}
