package main

import "testing"

func TestExamplePart1(t *testing.T) {
	monkeyRiddle := ParseMonkeyRiddle("input_test.txt")
	number := monkeyRiddle.GetValue(monkeyRiddle.GetMonkey("root"))

	expected := 152
	actual := number
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
