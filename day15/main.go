package main

import "log"

func main() {
	world := ParseSensorReadings("input.txt")
	num := world.GetPositionsCoveredWithNoBeacon(2000000)

	log.Printf("Part 1 - Positions covered: %d", num)
}
