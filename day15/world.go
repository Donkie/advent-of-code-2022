package main

import "advent-of-code-2022/lib"

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
