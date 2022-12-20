package main

import "log"

func part1() {
	values := ParseNumbers("input.txt")
	numbers := makeNumbers(values)
	numbers.Mix()
	zeroIdx := numbers.GetIndexOfZero()
	c1, c2, c3 := numbers.Get(zeroIdx+1000), numbers.Get(zeroIdx+2000), numbers.Get(zeroIdx+3000)
	cSum := c1 + c2 + c3

	log.Printf("Part 1 - Coordinate sum: %d", cSum)
}

func part2() {
	values := ParseNumbers("input.txt")
	numbers := makeNumbers(values)
	numbers.ApplyDecryptionKey(811589153)
	for i := 0; i < 10; i++ {
		numbers.Mix()
	}
	zeroIdx := numbers.GetIndexOfZero()
	c1, c2, c3 := numbers.Get(zeroIdx+1000), numbers.Get(zeroIdx+2000), numbers.Get(zeroIdx+3000)
	cSum := c1 + c2 + c3

	log.Printf("Part 2 - Coordinate sum: %d", cSum)
}

func main() {
	part1()
	part2()
}
