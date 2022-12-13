package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func findIndexOfMatchingBracket(line string, startBracketIdx int) int {
	level := 0
	for i := startBracketIdx; i < len(line); i++ {
		c := line[i]

		if c == '[' {
			level++
		} else if c == ']' {
			level--

			if level == 0 {
				return i
			}
		}
	}
	return -1
}

func findItemEnd(line string, itemStart int) int {
	level := 0
	for i := itemStart; i < len(line); i++ {
		c := line[i]

		if c == '[' {
			level++
		} else if c == ']' {
			level--
		} else if c == ',' && level == 0 {
			return i
		}
	}
	return len(line)
}

func isNumStr(line string) bool {
	if len(line) == 0 {
		return false
	}
	for _, c := range line {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func parseItemLine(line string) *Item {
	// Remove start and end brackets
	if line[0] == '[' {
		line = line[1 : len(line)-1]
	} else if isNumStr(line) {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Print(err)
			return nil
		}
		return newItem(val)
	}

	var children []*Item

	if len(line) > 0 {
		startIdx := 0
		for true {
			itemEnd := findItemEnd(line, startIdx)
			child := parseItemLine(line[startIdx:(itemEnd)])
			children = append(children, child)

			if itemEnd == len(line) {
				break
			} else {
				startIdx = itemEnd + 1
			}
		}
	}

	item := newItem(-1)
	item.children = children

	return item
}

func ParseInputPairsFile(fileName string) *[]Pair {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return ParseInputPairs(string(bytes))
}

func ParseInputPairs(pairText string) *[]Pair {
	lines := strings.Split(pairText, "\n")

	var pairs []Pair

	for i := 0; i < len(lines); i += 3 {
		p1str := lines[i]
		p2str := lines[i+1]

		p1 := parseItemLine(p1str)
		p2 := parseItemLine(p2str)

		pairs = append(pairs, Pair{
			p1: *p1,
			p2: *p2,
		})
	}

	return &pairs
}
