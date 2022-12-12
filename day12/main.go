package main

import "log"

func main() {
	graph := ParseHeightMap("input.txt", StartToEnd)
	numSteps := graph.FindShortestPath()

	log.Printf("Part 1 - Number of steps from start to end: %d", numSteps)

	graph = ParseHeightMap("input.txt", AnyToEnd)
	numSteps = graph.FindShortestPath()

	log.Printf("Part 2 - Number of steps from ground to end: %d", numSteps)
}
