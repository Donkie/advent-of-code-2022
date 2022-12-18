package main

import "log"

func main() {
	cubes := ParseCubes("input.txt")
	area := GetSurfaceArea(cubes)

	log.Printf("Part 1 - Surface area: %d", area)

	exteriorArea := GetExteriorSurfaceArea(cubes)

	log.Printf("Part 2 - Exterior surface area: %d", exteriorArea)
}
