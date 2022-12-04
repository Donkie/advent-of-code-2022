package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Gesture int

const (
	Rock Gesture = iota
	Paper
	Scissors
)

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

type Round struct {
	opponentMove Gesture
	myMove       Gesture
}

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

func (round Round) GetMyScore() int {
	return round.GetMyGestureScore() + round.GetMyOutcomeScore()
}

func parseInput(fileName string) []Round {
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
		round.myMove = parseGesture(line[2])
		rounds = append(rounds, round)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rounds
}

func getTotalScore(rounds []Round) int {
	sum := 0
	for _, round := range rounds {
		sum += round.GetMyScore()
	}
	return sum
}

func main() {
	rounds := parseInput("input.txt")

	totalScore := getTotalScore(rounds)
	fmt.Printf("Total score: %d\n", totalScore)
}
