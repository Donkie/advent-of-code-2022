package main

type Instruction int

const (
	NOOP Instruction = iota
	ADDX
)

type Op struct {
	instruction Instruction
	args        []int
}

var opFuncs = map[Instruction]func(machine *Machine, step int, args []int) bool{
	NOOP: func(machine *Machine, step int, args []int) bool {
		return true
	},
	ADDX: func(machine *Machine, step int, args []int) bool {
		if step == 0 {
			return false
		} else {
			machine.RegX += args[0]
			return true
		}
	},
}

type Program []Op

type Machine struct {
	program              Program
	crt                  CRT
	Op                   int
	OpStep               int
	Cycle                int
	RegX                 int
	SignalStrengthMetric int
}

func makeMachine(program Program) (machine Machine) {
	machine.program = program
	machine.crt = makeCRT()
	machine.RegX = 1
	machine.Cycle = 0
	return
}

func (machine *Machine) RecordSignalMetric() {
	if ((machine.Cycle - 20) % 40) == 0 {
		machine.SignalStrengthMetric += machine.Cycle * machine.RegX
	}
}

func (machine *Machine) RunUntilExit() {
	for machine.Op < len(machine.program) {
		machine.Cycle++
		machine.RecordSignalMetric()

		machine.crt.Draw(machine)

		op := machine.program[machine.Op]
		opFinished := opFuncs[op.instruction](machine, machine.OpStep, op.args)
		if opFinished {
			machine.Op++
			machine.OpStep = 0
		} else {
			machine.OpStep++
		}
	}
}
