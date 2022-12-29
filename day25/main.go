package main

import "fmt"

func part1() {
	snafus := ParseInputAsLines("input.txt")
	numbers := SNAFUToNumberList(snafus)
	sum := SumList(numbers)
	output := NumberToSNAFU(sum)

	fmt.Printf("Part 1 - Fuel sum in SNAFU is: %s\n", output)
}

func part2() {
}

func main() {
	part1()
	part2()
}
