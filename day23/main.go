package main

import "log"

func part1() {
	world := ParseWorld("input.txt")
	world.Simulate(10)
	emptyTiles := world.GetEmptySquaresInBoundingRegion()

	log.Printf("Part 1 - Rectangle contains %d empty ground tiles", emptyTiles)
}

func part2() {
	world := ParseWorld("input.txt")
	round := world.SimulateUntilEnd()

	log.Printf("Part 2 - Ended at round %d", round)
}

func main() {
	part1()
	part2()
}
