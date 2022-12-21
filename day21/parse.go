package main

import (
	"advent-of-code-2022/lib"
	"strconv"
	"strings"
)

func ParseMonkeyLine(line string) (*Monkey, error) {
	monk := new(Monkey)

	split1 := strings.Split(line, ": ")
	monk.name = split1[0]
	cmd := split1[1]

	if cmd[0] >= '0' && cmd[0] <= '9' {
		// It's a value monkey
		val, err := strconv.Atoi(cmd)
		if err != nil {
			return nil, err
		}
		monk.value = val
		monk.isSolved = true
		monk.isValueMonkey = true
		monk.op = None
	} else {
		monk.dependency1Name = cmd[0:4]
		monk.dependency2Name = cmd[7:11]
		switch cmd[5] {
		case '+':
			monk.op = Add
		case '-':
			monk.op = Sub
		case '*':
			monk.op = Mul
		case '/':
			monk.op = Div
		}
		monk.isSolved = false
		monk.isValueMonkey = false
	}

	return monk, nil
}

func ParseMonkeyRiddle(fileName string) (riddle Riddle) {
	monkies := make(map[string]*Monkey)
	lib.ParseInputByLine(fileName, func(line string) error {
		monk, err := ParseMonkeyLine(line)
		if err != nil {
			return err
		}
		monkies[monk.name] = monk
		return nil
	})

	// Resolve dependency pointers
	for _, monk := range monkies {
		if !monk.isSolved {
			monk.dependency1 = monkies[monk.dependency1Name]
			monk.dependency2 = monkies[monk.dependency2Name]
		}
	}

	riddle.monkieNames = monkies
	return
}
