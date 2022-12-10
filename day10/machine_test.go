package main

import "testing"

func TestMachine(t *testing.T) {
	machine := makeMachine(Program{
		Op{
			instruction: NOOP,
			args:        []int{},
		},
		Op{
			instruction: ADDX,
			args:        []int{5},
		},
		Op{
			instruction: ADDX,
			args:        []int{-7},
		},
	})
	machine.RunUntilExit()

	if machine.Cycle != 5 || machine.Op != 3 || machine.OpStep != 0 || machine.RegX != -1 {
		t.Errorf("Invalid machine state, cycle: %d, op: %d, opStep: %d, regX: %d", machine.Cycle, machine.Op, machine.OpStep, machine.RegX)
	}
}
