package main

import (
	"advent-of-code-2022/lib"
	"log"
)

type Blizzards struct {
	northbound Map
	southbound Map
	westbound  Map
	eastbound  Map
	mapw       int
	maph       int
	CycleTime  int
}

func newBlizzardsFromInput(lines []string) (b *Blizzards) {
	b = new(Blizzards)

	b.mapw = len(lines[0]) - 2
	b.maph = len(lines) - 2

	b.CycleTime = lib.LCM(b.mapw, b.maph)

	b.northbound = make(Map, b.maph)
	b.southbound = make(Map, b.maph)
	b.westbound = make(Map, b.maph)
	b.eastbound = make(Map, b.maph)

	for y := 1; y < len(lines)-1; y++ {
		b.northbound[y-1] = make([]bool, b.mapw)
		b.southbound[y-1] = make([]bool, b.mapw)
		b.westbound[y-1] = make([]bool, b.mapw)
		b.eastbound[y-1] = make([]bool, b.mapw)

		for x := 1; x < len(lines[y])-1; x++ {
			switch lines[y][x] {
			case '^':
				b.northbound[y-1][x-1] = true
			case 'v':
				b.southbound[y-1][x-1] = true
			case '<':
				b.westbound[y-1][x-1] = true
			case '>':
				b.eastbound[y-1][x-1] = true
			}
		}
	}
	return
}

// Returns true if the tile has a blizzard bound to a specific direction on it
func (b *Blizzards) getByDir(x int, y int, dir Direction, time int) bool {
	switch dir {
	case North:
		y = lib.Mod(y+time, b.maph)
		return b.northbound[y][x]
	case South:
		y = lib.Mod(y-time, b.maph)
		return b.southbound[y][x]
	case West:
		x = lib.Mod(x+time, b.mapw)
		return b.westbound[y][x]
	case East:
		x = lib.Mod(x-time, b.mapw)
		return b.eastbound[y][x]
	}
	log.Panic("Invalid direction")
	return false
}

// Returns true if the tile has a blizzard on it
func (b *Blizzards) Get(x int, y int, time int) bool {
	x--
	y--
	if y < 0 || x < 0 || x >= b.mapw || y >= b.maph {
		return false
	}
	return b.getByDir(x, y, North, time) || b.getByDir(x, y, South, time) || b.getByDir(x, y, West, time) || b.getByDir(x, y, East, time)
}

// Returns how many blizzards a tile has on it
func (b *Blizzards) GetNum(x int, y int, time int) (num int) {
	if b.getByDir(x, y, North, time) {
		num++
	}
	if b.getByDir(x, y, South, time) {
		num++
	}
	if b.getByDir(x, y, West, time) {
		num++
	}
	if b.getByDir(x, y, East, time) {
		num++
	}
	return
}

func (b *Blizzards) GetChar(x int, y int, time int) byte {
	x--
	y--
	if y < 0 || x < 0 {
		return ' '
	}
	num := b.GetNum(x, y, time)
	if num == 0 {
		return ' '
	} else if num == 1 {
		if b.getByDir(x, y, North, time) {
			return '^'
		}
		if b.getByDir(x, y, South, time) {
			return 'v'
		}
		if b.getByDir(x, y, West, time) {
			return '<'
		}
		if b.getByDir(x, y, East, time) {
			return '>'
		}
		log.Panic("Shouldn't happen")
		return ' '
	} else {
		return '0' + uint8(num)
	}
}
