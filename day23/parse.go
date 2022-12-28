package main

import "advent-of-code-2022/lib"

func ParseWorld(fileName string) (world World) {
	world.tiles = make(Map)
	world.elves = make([]*ElfObj, 0)

	y := 0
	lib.ParseInputByLine(fileName, func(line string) error {
		for x, c := range line {
			if c == '#' {
				elf := new(ElfObj)
				elf.pos = lib.Vector2{X: x, Y: y}

				world.tiles[elf.pos] = elf
				world.elves = append(world.elves, elf)
			}
		}

		y++
		return nil
	})

	return
}
