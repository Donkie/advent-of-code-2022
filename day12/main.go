package main

import "log"

func main() {
	graph := ParseHeightMap("input.txt")
	numSteps := graph.FindShortestPath()

	log.Printf("Part 1 - Number of steps from start to end: %d", numSteps)
}
