package main

import (
	"advent-of-code-2022/lib"
	"testing"
)

func assertVector2(t *testing.T, expected lib.Vector2, actual lib.Vector2) {
	if !expected.Equal(actual) {
		t.Errorf("expected %s, actual %s", expected.String(), actual.String())
	}
}

func TestGetBorderPositions(t *testing.T) {
	reading := makeReading(10, 10, 11, 10)

	c := reading.GetBorderPositions()
	assertVector2(t, lib.Vector2{X: 10, Y: 8}, <-c)
	assertVector2(t, lib.Vector2{X: 11, Y: 9}, <-c)
	assertVector2(t, lib.Vector2{X: 12, Y: 10}, <-c)
	assertVector2(t, lib.Vector2{X: 11, Y: 11}, <-c)
	assertVector2(t, lib.Vector2{X: 10, Y: 12}, <-c)
	assertVector2(t, lib.Vector2{X: 9, Y: 11}, <-c)
	assertVector2(t, lib.Vector2{X: 8, Y: 10}, <-c)
	assertVector2(t, lib.Vector2{X: 9, Y: 9}, <-c)
	assertVector2(t, lib.Vector2{X: 0, Y: 0}, <-c) // Channel closed
}
