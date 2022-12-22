package main

import "log"

func part1() {
	world := ParseWorld("input.txt")
	world.Traverse()
	pw := world.GetPassword()

	log.Printf("Part 1 - Password: %d", pw)
}

func part2() {
}

func main() {
	part1()
	part2()
}
