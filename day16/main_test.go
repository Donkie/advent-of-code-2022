package main

import "testing"

func TestExamplePart1(t *testing.T) {
	valveGraph := ParseValveGraph("input_test.txt")
	releasedPressure := valveGraph.GetOptimalPressureReleaseAmount(30)

	actual := releasedPressure
	expected := 1651
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	valveGraph := ParseValveGraph("input_test.txt")
	releasedPressure := valveGraph.GetOptimalPressureReleaseAmountTwoOperators(26)

	actual := releasedPressure
	expected := 1707
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
