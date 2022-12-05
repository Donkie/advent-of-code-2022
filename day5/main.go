package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Represents a crate moval in the yard
type Move struct {
	from int
	to   int
}

// Represents an action to take in the yard
// An action consists of a crate move that is done "times" number of times.
// The mode field defines how the move is performed
// mode = false - Crates are picked up and dropped off one-by-one, N times, where N is the "times" in the action
// mode = true  - The top N crates are picked up and dropped off as a bunch
type Action struct {
	move  Move
	mode  bool
	times int
}

// Represents a yard of crate stacks
type Yard struct {
	stacks []lib.Stack[byte]
}

// Prints the current state of the stacks of the yard to the log
func (yard Yard) PrintStacks() {
	log.Println("Problem Stacks:")
	for idx, stack := range yard.stacks {
		var sb strings.Builder
		for i := 0; i < stack.Len(); i++ {
			sb.WriteByte(*stack.Get(i))
			sb.WriteByte(' ')
		}
		log.Printf("%d: %s", idx, sb.String())
	}
}

// Performs a move in the yard, moving N amount of crates from one stack to another
func (yard Yard) ExecuteMove(move Move, amount int) {
	crates := yard.stacks[move.from].PopN(amount)
	if crates == nil {
		log.Fatal("Failed to execute move, from stack is not big enough for the specified amount")
	}
	yard.stacks[move.to].PushN(*crates)
}

// Executes the action, on the crates in the yard
func (yard Yard) ExecuteAction(act Action) {
	if act.mode == false {
		for i := 0; i < act.times; i++ {
			yard.ExecuteMove(act.move, 1)
		}
	} else {
		yard.ExecuteMove(act.move, act.times)
	}
}

// Returns a string representing the top labels of every crate in the yard
func (yard Yard) TopCrates() string {
	var sb strings.Builder
	for _, stack := range yard.stacks {
		sb.WriteByte(*stack.Peek())
	}
	return sb.String()
}

// Represents the problem
type Problem struct {
	yard    Yard
	actions []Action
}

// Executes all actions in the problem
func (prb Problem) ExecuteActions() {
	for _, act := range prb.actions {
		prb.yard.ExecuteAction(act)
	}
}

// Returns a string representing the top labels of every crate in the yard
func (prb Problem) TopCrates() string {
	return prb.yard.TopCrates()
}

// Parses an array of lines from the input file format containing the first part of the input text
// Returns a Yard object containing the parsed crate stack data
func parseStackYard(stackLines []string) Yard {
	numStacks := (len(stackLines[0]) + 1) / 4 // 4 characters per stack

	stacks := make([]lib.Stack[byte], numStacks)

	// Loop from the bottom and up
	for linei := len(stackLines) - 2; linei >= 0; linei-- {
		line := stackLines[linei]
		for stacki := 0; stacki < numStacks; stacki++ {
			crateLabel := line[stacki*4+1]
			if crateLabel != ' ' {
				stacks[stacki].Push(crateLabel)
			}
		}
	}

	var yard Yard
	yard.stacks = stacks
	return yard
}

var moveFormatRegex = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

// Parses an action string into the ActionTimes structure
// Returns an error if any parsing issue appears
func parseAction(actionLine string, moveMode bool) (*Action, error) {
	rs := moveFormatRegex.FindStringSubmatch(actionLine)
	if rs == nil {
		return nil, fmt.Errorf("Input string \"%s\" didn't fulfill the format.", actionLine)
	}

	times, err := strconv.Atoi(rs[1])
	if err != nil {
		return nil, err
	}
	from, err := strconv.Atoi(rs[2])
	if err != nil {
		return nil, err
	}
	to, err := strconv.Atoi(rs[3])
	if err != nil {
		return nil, err
	}

	act := new(Action)
	act.mode = moveMode
	act.times = times
	act.move.from = from - 1 // -1 to convert to 0-based indexing
	act.move.to = to - 1
	return act, nil
}

// Parses the input file into a problem structure
// See the Action type for information on the moveMode argument
// Exits the program if a parsing error occurs
func parseInput(fileName string, moveMode bool) Problem {
	var stackLines []string
	var actions []Action

	parsingStacks := true

	err := lib.ParseInputByLine(fileName, func(line string) error {
		if line == "" {
			parsingStacks = false
			return nil
		}

		if parsingStacks {
			stackLines = append(stackLines, line)
		} else {
			action, err := parseAction(line, moveMode)
			if err != nil {
				return err
			}
			actions = append(actions, *action)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	var prb Problem
	prb.yard = parseStackYard(stackLines)
	prb.actions = actions
	return prb
}

func main() {
	prb := parseInput("input.txt", false)
	prb.ExecuteActions()
	log.Printf("Part 1 - Top crates are: %s", prb.TopCrates())

	prb = parseInput("input.txt", true)
	prb.ExecuteActions()
	log.Printf("Part 2 - Top crates are: %s", prb.TopCrates())
}
