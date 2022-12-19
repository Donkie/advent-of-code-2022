package main

import "log"

func main() {
	part := 1

	blueprints := ParseBlueprints("input.txt")
	if part == 1 {
		numGeodesList := GetNumberOfGeodes(blueprints, 24)
		qualitySum := GetQualitySum(numGeodesList)

		log.Printf("Part 1 - Quality Sum: %d", qualitySum)
	} else {
		blueprints = blueprints[:3]
		numGeodesList := GetNumberOfGeodes(blueprints, 32)
		metric := numGeodesList[0] * numGeodesList[1] * numGeodesList[2]

		log.Printf("Part 2 - Metric: %d", metric)
	}
}
