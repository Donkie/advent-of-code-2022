package main

import "log"

func main() {
	world := ParseSensorReadings("input.txt")
	num := world.GetPositionsCoveredWithNoBeacon(2000000)

	log.Printf("Part 1 - Positions covered: %d", num)

	pos := world.GetUncoveredPosition(0, 4000000, 0, 4000000)
	freq := GetTuningFrequency(pos)

	log.Printf("Part 2 - Tuning frequency: %d", freq)
}
