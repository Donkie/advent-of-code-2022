package main

import (
	"advent-of-code-2022/lib"
	"log"
	"strconv"
	"strings"
)

func parseVertex(str string) (x int, y int, err error) {
	split := strings.Split(str, ",")
	x, err = strconv.Atoi(split[0])
	if err != nil {
		return
	}
	y, err = strconv.Atoi(split[1])
	return
}

func populateRockLine(world *World, line string) {
	vertices := strings.Split(line, " -> ")
	if len(vertices) == 1 {
		x, y, err := parseVertex(vertices[0])
		if err != nil {
			log.Panic(err)
			return
		}
		world.SetBlock(x, y, Rock)
	} else {
		for i := 0; i < len(vertices)-1; i++ {
			v1 := vertices[i]
			v2 := vertices[i+1]
			x1, y1, err := parseVertex(v1)
			if err != nil {
				log.Panic(err)
				return
			}
			x2, y2, err := parseVertex(v2)
			if err != nil {
				log.Panic(err)
				return
			}
			world.SetBlockLine(x1, y1, x2, y2, Rock)
		}
	}
}

func ParseWorld(fileName string) World {
	world := makeWorld()

	lib.ParseInputByLine(fileName, func(line string) error {
		populateRockLine(&world, line)
		return nil
	})

	return world
}
