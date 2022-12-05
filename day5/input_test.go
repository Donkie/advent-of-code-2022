package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	prb := parseInput("input_test.txt", false)
	prb.ExecuteActions()
	expected := "CMZ"
	actual := prb.TopCrates()
	if actual != expected {
		t.Errorf("Expected: %s, actual: %s\n", expected, actual)
	}
}
func TestExamplePart2(t *testing.T) {
	prb := parseInput("input_test.txt", true)
	prb.ExecuteActions()
	expected := "MCD"
	actual := prb.TopCrates()
	if actual != expected {
		t.Errorf("Expected: %s, actual: %s\n", expected, actual)
	}
}
