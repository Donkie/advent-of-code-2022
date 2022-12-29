package main

import "fmt"

func part1() {
	world := ParseWorld("input.txt")
	bestTime := world.FindShortestPath()

	fmt.Printf("Part 1 - Shortest path: %d", bestTime)
}

func part2() {

}

func main() {
	part1()
	part2()
}
