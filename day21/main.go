package main

import "log"

func part1() {
	monkeyRiddle := ParseMonkeyRiddle("input.txt")
	number := monkeyRiddle.GetValue(monkeyRiddle.GetMonkey("root"))

	log.Printf("Part 1 - Number: %d", number)
}

func main() {
	part1()
}
