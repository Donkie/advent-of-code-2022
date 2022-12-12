package main

import (
	"advent-of-code-2022/lib"
	"log"
	"os"
	"strings"
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
			if x > 0 && (heightsRow[x-1]-thisHeight) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x-1, y))
			}
			// Check node to the right
			if x < (width-1) && (heightsRow[x+1]-thisHeight) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x+1, y))
			}
			// Check node up
			if y > 0 && (heights[y-1][x]-thisHeight) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x, y-1))
			}
			// Check node down
			if y < (height-1) && (heights[y+1][x]-thisHeight) <= 1 {
				neighbours = append(neighbours, graph.GetNode(x, y+1))
			}

			graph.GetNode(x, y).neighbours = neighbours
		}
	}

	graph.startNode = graph.GetNode(start.X, start.Y)
	graph.endNode = graph.GetNode(end.X, end.Y)

	return &graph
}
