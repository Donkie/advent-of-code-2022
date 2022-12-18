package main

import (
	"advent-of-code-2022/lib"
	"strconv"
	"strings"
)

func ParseCube(line string) (*Cube, error) {
	components := strings.Split(line, ",")
	x, err := strconv.Atoi(components[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(components[1])
	if err != nil {
		return nil, err
	}
	z, err := strconv.Atoi(components[2])
	if err != nil {
		return nil, err
	}

	return &Cube{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func ParseCubes(fileName string) (cubes []Cube) {
	lib.ParseInputByLine(fileName, func(line string) error {
		cube, err := ParseCube(line)
		if err != nil {
			return err
		}
		cubes = append(cubes, *cube)
		return nil
	})
	return
}
