package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`Sensor at x=([\d-]+), y=([\d-]+): closest beacon is at x=([\d-]+), y=([\d-]+)`)

func ParseSensorReadings(fileName string) (world World) {
	lib.ParseInputByLine(fileName, func(line string) error {
		match := re.FindStringSubmatch(line)
		if match == nil {
			return fmt.Errorf("Regex didn't match line")
		}

		sx, err := strconv.Atoi(match[1])
		if err != nil {
			return err
		}

		sy, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}

		bx, err := strconv.Atoi(match[3])
		if err != nil {
			return err
		}

		by, err := strconv.Atoi(match[4])
		if err != nil {
			return err
		}

		world.readings = append(world.readings, makeReading(sx, sy, bx, by))
		return nil
	})
	return
}
