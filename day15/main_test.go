package main

import (
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
