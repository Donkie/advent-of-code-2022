package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Represents a range of numbers
type Range struct {
	min int
	max int
}

// Represents a pair of Ranges from the input data
type RangePair struct {
	r1 Range
	r2 Range
}

// Returns true if this range contains the supplied number
func (r Range) Contains(num int) bool {
	return num >= r.min && num <= r.max
}

// Returns true if this range fully contains the other range
func (r Range) FullyContains(r2 Range) bool {
	return r.Contains(r2.min) && r.Contains(r2.max)
}

// Returns true if this range overlaps any part of the supplied range
func (r Range) Overlaps(r2 Range) bool {
	return r.Contains(r2.min) || r.Contains(r2.max)
}

// Returns a new Range object parsed from the input string
// Returns an error if the input string has an invalid format.
// Valid format of the input string is "A-B", where A and B are ints.
// Example "123-456"
func rangeFromString(str string) (*Range, error) {
	strs := strings.Split(str, "-")
	if len(strs) != 2 {
		return nil, fmt.Errorf("Invalid range string: %s", str)
	}
	min, err := strconv.Atoi(strs[0])
	if err != nil {
		return nil, fmt.Errorf("Invalid number in min value of range in string: %s", str)
	}
	max, err := strconv.Atoi(strs[1])
	if err != nil {
		return nil, fmt.Errorf("Invalid number in max value of range in string: %s", str)
	}
	r := new(Range)
	r.min = min
	r.max = max
	return r, nil
}

// Parses the input file
// Returns a list of all pairs of ranges in the input
// Exits the program if the format of one of the lines is invalid
func parseInput(fileName string) []RangePair {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pairs []RangePair
	for scanner.Scan() {
		line := scanner.Text()
		pairstr := strings.Split(line, ",")
		if len(pairstr) != 2 {
			log.Fatalf("Unknown pair string: %s\n", line)
			break
		}
		r1, err := rangeFromString(pairstr[0])
		if err != nil {
			log.Fatal(err)
			break
		}
		r2, err := rangeFromString(pairstr[1])
		if err != nil {
			log.Fatal(err)
			break
		}

		var pair RangePair
		pair.r1 = *r1
		pair.r2 = *r2
		pairs = append(pairs, pair)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return pairs
}

// Returns the number of pairs in the list where one of the ranges fully contain the other
func getNumPairsFullyContaining(pairs []RangePair) int {
	sum := 0
	for _, pair := range pairs {
		if pair.r1.FullyContains(pair.r2) || pair.r2.FullyContains(pair.r1) {
			sum++
		}
	}
	return sum
}

// Returns the number of pairs in the list where the ranges overlap each other
func getNumPairsOverlapping(pairs []RangePair) int {
	sum := 0
	for _, pair := range pairs {
		if pair.r1.Overlaps(pair.r2) || pair.r2.Overlaps(pair.r1) {
			sum++
		}
	}
	return sum
}

func main() {
	pairs := parseInput("input.txt")

	numPairFullyContaining := getNumPairsFullyContaining(pairs)
	fmt.Printf("Part 1 - Number of pairs that fully contain the other: %d\n", numPairFullyContaining)

	numPairOverlapping := getNumPairsOverlapping(pairs)
	fmt.Printf("Part 2 - Number of pairs that overlaps the other: %d\n", numPairOverlapping)
}
