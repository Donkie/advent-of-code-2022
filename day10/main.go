package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func strToInstruction(name string) (Instruction, error) {
	switch name {
	case "noop":
		return NOOP, nil
	case "addx":
		return ADDX, nil
	default:
		return 0, fmt.Errorf("Unknown instruction: %s", name)
	}
}

func ParseProgram(fileName string) (program Program) {
	lib.ParseInputByLine(fileName, func(line string) error {
		parts := strings.Split(line, " ")

		instruction, err := strToInstruction(parts[0])
		if err != nil {
			return err
		}

		var op Op
		op.instruction = instruction

		if instruction == ADDX {
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				return err
			}
			op.args = []int{val}
		} else {
			op.args = []int{}
		}
		program = append(program, op)
		return nil
	})
	return
}

func main() {
	program := ParseProgram("input.txt")
	machine := makeMachine(program)
	machine.RunUntilExit()
	metric := machine.SignalStrengthMetric

	log.Printf("Part 1 - Signal strength metric: %d", metric)
}
