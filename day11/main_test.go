package main

import "testing"

func TestExamplePart1(t *testing.T) {
	troop := ParseMonkeyTroop("input_test.txt")
	troop.PerformRounds(20, true)
	metric := troop.GetMonkeyBusinessLevel()

	actual := metric
	expected := 10605
	if actual != expected {
		t.Errorf("expected: %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	troop := ParseMonkeyTroop("input_test.txt")
	troop.PerformRounds(10000, false)
	metric := troop.GetMonkeyBusinessLevel()

	actual := metric
	expected := 2713310158
	if actual != expected {
		t.Errorf("expected: %d, actual %d", expected, actual)
	}
}
