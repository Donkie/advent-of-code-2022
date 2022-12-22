package main

import "log"

func part1() {
	world := ParseWorld("input.txt")
	world.Traverse()
	pw := world.GetPassword()

	log.Printf("Part 1 - Password: %d", pw)
}

func part2() {
	world := ParseWorld("input.txt")
	world.InterpretCube(50, [][]Face{
		{FNone, FTop, FRight},
		{FNone, FFront, FNone},
		{FLeft, FBottom, FNone},
		{FBack, FNone, FNone},
	}, map[Face]int{
		FTop:    0,
		FFront:  0,
		FBottom: 0,
		FRight:  -1,
		FLeft:   -1,
		FBack:   -1,
	})
	world.Traverse()
	pw := world.GetPassword()

	log.Printf("Part 2 - Password: %d", pw)
}

func main() {
	// part1()
	part2()
}
