package main

import "log"

func main() {
	valveGraph := ParseValveGraph("input.txt")
	releasedPressure := valveGraph.GetOptimalPressureReleaseAmount(30)

	log.Printf("Part 1 - Released pressure: %d", releasedPressure)

	releasedPressure2 := valveGraph.GetOptimalPressureReleaseAmountTwoOperators(26)

	log.Printf("Part 2 - Released pressure: %d", releasedPressure2)
}
