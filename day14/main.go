package main

import "log"

func main() {
	world := ParseWorld("input.txt")
	restingSand := world.SimulateSand(false)

	log.Printf("Part 1 - Resting sand particles: %d", restingSand)
}
