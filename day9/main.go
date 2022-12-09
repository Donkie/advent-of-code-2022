package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func charToDir(char byte) (Dir, error) {
	switch char {
	case 'U':
		return Up, nil
	case 'R':
		return Right, nil
	case 'D':
		return Down, nil
	case 'L':
		return Left, nil
	}
	return 0, fmt.Errorf("Unknown direction char %b", char)
}

// ParseActions parses the input file into a list of actions.
func ParseActions(fileName string) (actions []Action) {
	lib.ParseInputByLine(fileName, func(line string) error {
		args := strings.Split(line, " ")

		dir, err := charToDir(args[0][0])
		if err != nil {
			return err
		}

		times, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		for i := 0; i < times; i++ {
			actions = append(actions, Action{dir: dir})
		}
		return nil
	})
	return
}

func main() {
	actions := ParseActions("input.txt")
	rope := makeRope(2)
	rope.PerformActions(actions)
	visits := rope.GetNumTailUniquePositions()

	log.Printf("Part 1 - Number of visited positions: %d", visits)

	longRope := makeRope(10)
	longRope.PerformActions(actions)
	visits = longRope.GetNumTailUniquePositions()

	log.Printf("Part 2 - Number of visited positions: %d", visits)
}
