package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	cubes := ParseCubes("input_test.txt")
	area := GetSurfaceArea(cubes)

	actual := area
	expected := 64
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	cubes := ParseCubes("input_test.txt")
	area := GetExteriorSurfaceArea(cubes)

	actual := area
	expected := 58
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
