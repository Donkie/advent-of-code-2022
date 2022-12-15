package main

import (
	"advent-of-code-2022/lib"
)

type Reading struct {
	sensorPos   lib.Vector2
	beaconPos   lib.Vector2
	sensorRange int
}

func makeReading(sensorX int, sensorY int, beaconX int, beaconY int) (reading Reading) {
	reading.sensorPos = lib.Vector2{
		X: sensorX,
		Y: sensorY,
	}
	reading.beaconPos = lib.Vector2{
		X: beaconX,
		Y: beaconY,
	}
	reading.sensorRange = reading.sensorPos.ManhattanDistance(reading.beaconPos)
	return
}

func (r *Reading) GetBorderPositions() chan lib.Vector2 {
	c := make(chan lib.Vector2)
	go func() {
		topX, topY := r.sensorPos.X, r.sensorPos.Y-r.sensorRange
		rightX, rightY := r.sensorPos.X+r.sensorRange, r.sensorPos.Y
		bottomX, bottomY := r.sensorPos.X, r.sensorPos.Y+r.sensorRange
		leftX, leftY := r.sensorPos.X-r.sensorRange, r.sensorPos.Y

		// ..1...
		// ..#...
		// .###2.
		// ##X##3
		// .###..
		// ..#4..
		// First loop we want to iterate from 1 -> 2
		// Second loop, 3 -> 4, etc
		for i := 0; i <= r.sensorRange; i++ {
			x, y := topX+i, topY-1+i
			c <- lib.Vector2{X: x, Y: y}
		}
		for i := 0; i <= r.sensorRange; i++ {
			x, y := rightX+1-i, rightY+i
			c <- lib.Vector2{X: x, Y: y}
		}
		for i := 0; i <= r.sensorRange; i++ {
			x, y := bottomX-i, bottomY+1-i
			c <- lib.Vector2{X: x, Y: y}
		}
		for i := 0; i <= r.sensorRange; i++ {
			x, y := leftX-1+i, leftY-i
			c <- lib.Vector2{X: x, Y: y}
		}
		close(c)
	}()
	return c
}

func (r *Reading) GetRange() int {
	return r.sensorRange
}

func (r *Reading) IsCoveredBySensor(test lib.Vector2) bool {
	return r.sensorPos.ManhattanDistance(test) <= r.sensorRange
}

type World struct {
	readings []Reading
}

func (w *World) GetXBounds() (minx int, maxx int) {
	minx = w.readings[0].sensorPos.X
	maxx = minx
	for _, reading := range w.readings {
		minx = lib.Min(minx, reading.sensorPos.X-reading.sensorRange)
		maxx = lib.Max(maxx, reading.sensorPos.X+reading.sensorRange)
	}
	return
}

func (w *World) GetUncoveredPosition(minx int, maxx int, miny int, maxy int) (position lib.Vector2) {
	// The position we're looking for has to be on the outside border of one of the sensors range
	// So to find it, we just have to look through all those outside borders if any other sensor is covering it

	for i, testReading := range w.readings {
		for testPoint := range testReading.GetBorderPositions() {
			if testPoint.X < minx || testPoint.X > maxx || testPoint.Y < miny || testPoint.Y > maxy {
				continue
			}

			isCovered := false
			for j, reading := range w.readings {
				if j == i {
					continue
				}
				if reading.IsCoveredBySensor(testPoint) {
					isCovered = true
					break
				}
			}
			if !isCovered {
				return testPoint
			}
		}
	}
	return
}

func (w *World) GetPositionsCoveredWithNoBeacon(yLevel int) (sum int) {
	minx, maxx := w.GetXBounds()
	for x := minx; x <= maxx; x++ {
		covered := false
		testPoint := lib.Vector2{X: x, Y: yLevel}
		for _, reading := range w.readings {
			if reading.beaconPos.Equal(testPoint) {
				covered = false
				break
			}
			if reading.IsCoveredBySensor(testPoint) {
				covered = true
				break
			}
		}
		if covered {
			sum++
		}
	}
	return
}

func GetTuningFrequency(pos lib.Vector2) int {
	return pos.X*4000000 + pos.Y
}
