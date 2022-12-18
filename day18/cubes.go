package main

import (
	"advent-of-code-2022/lib"
	"math"
)

type Cube lib.Vector3

func hashVec(v lib.Vector3) int {
	return v.X<<16 + v.Y<<8 + v.Z
}

func hashVec3(x int, y int, z int) int {
	return x<<16 + y<<8 + z
}

var cardinals = []lib.Vector3{
	{X: 1, Y: 0, Z: 0},
	{X: -1, Y: 0, Z: 0},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 0, Y: 0, Z: 1},
	{X: 0, Y: 0, Z: -1},
}

func GetSurfaceArea(cubes []Cube) (area int) {
	cubeMap := make(map[int]struct{})
	for _, cube := range cubes {
		cubeMap[hashVec(lib.Vector3(cube))] = struct{}{}
	}

	for _, cube := range cubes {
		for _, dir := range cardinals {
			_, hasCube := cubeMap[hashVec(lib.Vector3(cube).Add(dir))]
			if !hasCube {
				area++
			}
		}
	}
	return
}

func GetExteriorSurfaceArea(cubes []Cube) (area int) {
	cubeMap := make(map[int]struct{})
	minx, miny, minz := math.MaxInt, math.MaxInt, math.MaxInt
	maxx, maxy, maxz := math.MinInt, math.MinInt, math.MinInt
	for _, cube := range cubes {
		cubeMap[hashVec(lib.Vector3(cube))] = struct{}{}

		minx = lib.Min(minx, cube.X)
		maxx = lib.Max(maxx, cube.X)
		miny = lib.Min(miny, cube.Y)
		maxy = lib.Max(maxy, cube.Y)
		minz = lib.Min(minz, cube.Z)
		maxz = lib.Max(maxz, cube.Z)
	}
	minx--
	miny--
	minz--
	maxx++
	maxy++
	maxz++

	// Spread the air inwards.
	airMap := make(map[int]struct{})
	for i := 0; i < 10; i++ { // Do this a couple of times since we will only reach 1 step inwards in the opposite direction per pass
		for x := minx; x <= maxx; x++ {
			for y := miny; y <= maxy; y++ {
				for z := minz; z <= maxz; z++ {
					v := lib.Vector3{X: x, Y: y, Z: z}
					vHash := hashVec3(x, y, z)

					// If we're on the border then we're sure that it's open air
					if x == minx || x == maxx || y == miny || y == maxy || z == minz || z == maxz {
						airMap[vHash] = struct{}{}
					} else {
						_, hasCube := cubeMap[vHash]
						if !hasCube {
							for _, dir := range cardinals {
								// If the tile nearby is air, and this tile isn't a cube, then this tile must be air too.
								_, hasAir := airMap[hashVec(v.Add(dir))]
								if hasAir {
									airMap[vHash] = struct{}{}
								}
							}
						}
					}
				}
			}
		}
	}

	// Simply check all adjacent tiles if they are open air or not
	for _, cube := range cubes {
		for _, dir := range cardinals {
			_, hasAir := airMap[hashVec(lib.Vector3(cube).Add(dir))]
			if hasAir {
				area++
			}
		}
	}
	return
}
