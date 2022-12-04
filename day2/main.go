package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Gesture int

// Represents a RPS gesture
const (
	Rock Gesture = iota
	Paper
	Scissors
)

// Returns the gesture as a string
func (gesture Gesture) String() string {
	switch gesture {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return fmt.Sprintf("%d", int(gesture))
	}
}

// Returns an integer indicating the win/draw/loss between this gesture and the opponents one
// 1: This gesture beats the supplied gesture
// 0: It's a draw
// -1: This gesture loses against the other gesture
func (gesture Gesture) Beats(otherGesture Gesture) int {
	if gesture == otherGesture {
		return 0
	}
	if gesture == Rock && otherGesture == Scissors ||
		gesture == Paper && otherGesture == Rock ||
		gesture == Scissors && otherGesture == Paper {
		return 1
	}
	return -1
}

// Returns what Gesture to use in order to win against this gesture
func (gesture Gesture) HowToBeat() Gesture {
	switch gesture {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		log.Panicf("Unknown beater for gesture: %s", gesture.String())
		return 0
	}
}

// Returns what Gesture to use in order to lose against this gesture
func (gesture Gesture) HowToLose() Gesture {
	switch gesture {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		log.Panicf("Unknown loser for gesture: %s", gesture.String())
		return 0
	}
}

// Parses the input file gesture codes to the Gesture enum
func parseGesture(gesture byte) Gesture {
	switch gesture {
	case 'A', 'X':
		return Rock
	case 'B', 'Y':
		return Paper
	case 'C', 'Z':
		return Scissors
	default:
		log.Panicf("Unknown gesture identifier: %b", gesture)
		return Rock
	}
}

// Represents a round of RPS
type Round struct {
	opponentMove Gesture
	myMove       Gesture
}

// Gets the score I get based on what gesture I picked
func (round Round) GetMyGestureScore() int {
	switch round.myMove {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		log.Panicf("Unknown score for gesture: %s", round.myMove.String())
		return 0
	}
}

// Gets the score I get based on the outcome of this round
func (round Round) GetMyOutcomeScore() int {
	switch round.myMove.Beats(round.opponentMove) {
	case -1:
		return 0 // Loss
	case 1:
		return 6 // Win
	default:
		return 3 // Draw
	}
}

// Gets my total score of this round
func (round Round) GetMyScore() int {
	return round.GetMyGestureScore() + round.GetMyOutcomeScore()
}

// Parses the input file
// secondColumnMode = false: The 2nd column indicates the move I should take in that round
// secondColmunMode = true: The 2nd column indicates what outcome I should make the round have
func parseInput(fileName string, secondColumnMode bool) []Round {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rounds []Round
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 3 {
			continue
		}

		var round Round
		round.opponentMove = parseGesture(line[0])
		if !secondColumnMode {
			round.myMove = parseGesture(line[2])
		} else {
			outcome := line[2]
			switch outcome {
			case 'X':
				round.myMove = round.opponentMove.HowToLose()
			case 'Y':
				round.myMove = round.opponentMove
			case 'Z':
				round.myMove = round.opponentMove.HowToBeat()
			default:
				log.Panicf("Unknown identifier: %b", outcome)
			}
		}
		rounds = append(rounds, round)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rounds
}

// Gets the total score for a slice of rounds
func getTotalScore(rounds []Round) int {
	sum := 0
	for _, round := range rounds {
		sum += round.GetMyScore()
	}
	return sum
}

func main() {
	rounds := parseInput("input.txt", false)
	totalScore := getTotalScore(rounds)
	fmt.Printf("Part 1: Total score: %d\n", totalScore)

	rounds = parseInput("input.txt", true)
	totalScore = getTotalScore(rounds)
	fmt.Printf("Part 2: Total score: %d\n", totalScore)
}
