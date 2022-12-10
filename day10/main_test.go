package main

import "testing"

func TestExamplePart1(t *testing.T) {
	program := ParseProgram("input_test.txt")
	machine := makeMachine(program)
	machine.RunUntilExit()
	metric := machine.SignalStrengthMetric

	actual := metric
	expected := 13140
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
