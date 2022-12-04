package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

type RangePair struct {
	r1 Range
	r2 Range
}

func (r Range) FullyContains(r2 Range) bool {
	return r2.min >= r.min && r2.max <= r.max
}

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

func getNumPairsFullyContaining(pairs []RangePair) int {
	sum := 0
	for _, pair := range pairs {
		if pair.r1.FullyContains(pair.r2) || pair.r2.FullyContains(pair.r1) {
			sum++
		}
	}
	return sum
}

func main() {
	pairs := parseInput("input.txt")

	numPairFullyContaining := getNumPairsFullyContaining(pairs)
	fmt.Printf("Part 1 - Number of pairs that fully contain the other: %d\n", numPairFullyContaining)
}
