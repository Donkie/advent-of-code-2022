package main

import "fmt"

func part1() {
	world := ParseWorld("input.txt")
	bestTime := world.FindShortestPath(false)

	fmt.Printf("Part 1 - Shortest path: %d\n", bestTime)
}

func part2() {
	world := ParseWorld("input.txt")
	bestTime := world.FindShortestPath(true)

	fmt.Printf("Part 2 - Shortest path: %d\n", bestTime)
}

func main() {
	part1()
	part2()
}
