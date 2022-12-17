package main

import "log"

func main() {
	jetStream := ParseJetStream("input.txt")
	world := makeRockFallingWorld(*jetStream)
	world.Simulate(2022)
	height := world.GetHighestPoint()

	log.Printf("Part 1 - Tower height: %d", height)

	jetStream = ParseJetStream("input.txt")
	world = makeRockFallingWorld(*jetStream)
	world.Simulate(1_000_000_000_000)
	height = world.GetHighestPoint()

	log.Printf("Part 2 - Tower height: %d", height)
}
