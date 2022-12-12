package main

import (
	"advent-of-code-2022/lib"
	"log"
	"os"
	"strings"
)

type Objective int

const (
	StartToEnd Objective = iota
	AnyToEnd
)

func ParseHeightMap(fileName string) *Graph {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	lines := strings.Split(string(bytes), "\n")

	height := len(lines)
	width := len(lines[0])

	var start lib.Vector2
	var end lib.Vector2

	heights := make([][]int, height)
	for y, line := range lines {
		heights[y] = make([]int, width)

		for x, c := range line {
			if c == 'S' {
				start.X = x
				start.Y = y
				c = 'a'
			} else if c == 'E' {
				end.X = x
				end.Y = y
				c = 'z'
			}
			heights[y][x] = int(byte(c) - 'a')
		}
	}

	graph := makeGraph()

	for y, heightsRow := range heights {
		for x, thisHeight := range heightsRow {
			var neighbours []*Node

			// Check node to the left
			if x > 0 && (thisHeight-heightsRow[x-1]) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x-1, y))
			}
			// Check node to the right
			if x < (width-1) && (thisHeight-heightsRow[x+1]) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x+1, y))
			}
			// Check node up
			if y > 0 && (thisHeight-heights[y-1][x]) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x, y-1))
			}
			// Check node down
			if y < (height-1) && (thisHeight-heights[y+1][x]) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x, y+1))
			}

			graph.GetNode(x, y).neighbours = neighbours
			if x == start.X && y == start.Y {
				graph.GetNode(x, y).isEnd = true
			}
		}
	}

	graph.startNode = graph.GetNode(end.X, end.Y)

	return &graph
}
