package main

import (
	"advent-of-code-2022/lib"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseMap(world []string) [][]Square {
	height := len(world)
	sqrs := make([][]Square, height)

	var width int
	for i := 0; i < len(world); i++ {
		width = lib.Max(width, len(world[i]))
	}

	for y := 0; y < height; y++ {
		row := make([]Square, width)

		for x := 0; x < len(world[y]); x++ {
			switch world[y][x] {
			case '.':
				row[x] = Open
			case '#':
				row[x] = Wall
			default:
				row[x] = Void
			}
		}

		sqrs[y] = row
	}

	return sqrs
}

var re = regexp.MustCompile(`(\d+|[A-Z])`)

func parsePath(path string) (ops []Operation) {
	for _, match := range re.FindAllString(path, -1) {
		var op Operation
		switch match[0] {
		case 'L':
			op.turnDir = Left
		case 'R':
			op.turnDir = Right
		default:
			steps, err := strconv.Atoi(match)
			if err != nil {
				log.Panic(err)
				continue
			}
			op.steps = steps
		}
		ops = append(ops, op)
	}
	return
}

func ParseWorld(fileName string) *World {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	lines := strings.Split(string(bytes), "\n")
	sqrsstr := lines[:len(lines)-2]
	path := lines[len(lines)-1]

	world := new(World)
	world.sqrs = parseMap(sqrsstr)
	world.path = parsePath(path)
	world.stepCache = make(map[Step]lib.Vector2)

	return world
}
