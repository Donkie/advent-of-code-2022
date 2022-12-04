package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	pairs := parseInput("input_test.txt")

	expected := 2
	actual := getNumPairsFullyContaining(pairs)
	if actual != expected {
		t.Errorf("Expected: %d, actual: %d\n", expected, actual)
	}
}
func TestExamplePart2(t *testing.T) {
	pairs := parseInput("input_test.txt")

	expected := 4
	actual := getNumPairsOverlapping(pairs)
	if actual != expected {
		t.Errorf("Expected: %d, actual: %d\n", expected, actual)
	}
}
