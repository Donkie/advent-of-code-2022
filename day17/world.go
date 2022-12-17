package main

import (
	"advent-of-code-2022/lib"
	"fmt"
)

// Coordinate system:
// ^
// |
// y
// x ---->

type Space bool

const (
	Empty Space = false
	Rock  Space = true
)

type Shape []lib.Vector2

type ShapeType uint8

var shapes = map[ShapeType]Shape{
	0: {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},               // -
	1: {{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 2}}, // +
	2: {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}}, // _|
	3: {{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},               // |
	4: {{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},               // o
}

func vecToIdx(v lib.Vector2) int {
	return v.X + v.Y<<5 // Don't need to shift more than 5 since the world width is quite narrow
}

var dirVec = map[Direction]lib.Vector2{
	Left:  {X: -1, Y: 0},
	Right: {X: 1, Y: 0},
}

type RockFallingWorld struct {
	jetStream    JetStream
	field        map[int]Space
	curShape     ShapeType
	highestPoint int
}

func makeRockFallingWorld(jetStream JetStream) (world RockFallingWorld) {
	world.jetStream = jetStream
	world.field = make(map[int]Space)
	return
}

func (w *RockFallingWorld) get(v lib.Vector2) Space {
	s, ok := w.field[vecToIdx(v)]
	if !ok {
		return Empty
	} else {
		return s
	}
}

func (w *RockFallingWorld) rockFits(pos lib.Vector2, shape Shape) bool {
	// newPos := pos.Add(dirVec[dir])
	for _, piece := range shape {
		piecePos := pos.Add(piece)
		if piecePos.X < 0 || piecePos.X > 6 || piecePos.Y < 0 || w.get(piecePos) != Empty {
			return false
		}
	}
	return true
}

func (w *RockFallingWorld) solidifyRock(pos lib.Vector2, shape Shape) {
	for _, piece := range shape {
		piecePos := pos.Add(piece)
		w.field[vecToIdx(piecePos)] = Rock
		w.highestPoint = lib.Max(w.highestPoint, piecePos.Y+1)
	}
}

func (w *RockFallingWorld) simulateRock() {
	pos := lib.Vector2{X: 2, Y: w.GetHighestPoint() + 3}
	shape := shapes[w.curShape]
	for true {
		// w.PrintWithRock(pos, w.curShape)

		// Move sideways by jet
		dir := w.jetStream.PopDir()
		newPos := pos.Add(dirVec[dir])
		if w.rockFits(newPos, shape) {
			pos = newPos
		}

		// w.PrintWithRock(pos, w.curShape)

		// Move downwards
		newPos = pos.Add(lib.Vector2{X: 0, Y: -1})
		if w.rockFits(newPos, shape) {
			pos = newPos
		} else {
			w.solidifyRock(pos, shape)
			break
		}
	}

	w.curShape = (w.curShape + 1) % 5
}

func (w *RockFallingWorld) Simulate(numRocks int) {
	for i := 0; i < numRocks; i++ {
		if i%10000000000 == 0 {
			fmt.Printf("%d%%", 100*i/1000000000000)
		}
		w.simulateRock()
		// if i < 10 {
		// 	fmt.Println()
		// 	w.Print()
		// }
	}
}

func (w *RockFallingWorld) GetHighestPoint() int {
	return w.highestPoint
}

func rockOccupiesPos(testPos lib.Vector2, rockPos lib.Vector2, shape ShapeType) bool {
	for _, piece := range shapes[shape] {
		if testPos.Equal(rockPos.Add(piece)) {
			return true
		}
	}
	return false
}

func (w *RockFallingWorld) PrintWithRock(pos lib.Vector2, shape ShapeType) {
	for y := w.GetHighestPoint() + 3; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x <= 6; x++ {
			testPos := lib.Vector2{X: x, Y: y}
			if rockOccupiesPos(testPos, pos, shape) {
				fmt.Print("@")
			} else {
				switch w.get(testPos) {
				case Rock:
					fmt.Print("#")
				default:
					fmt.Print(".")
				}
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}

func (w *RockFallingWorld) Print() {
	for y := w.GetHighestPoint() + 1; y >= 0; y-- {
		fmt.Print("|")
		for x := 0; x <= 6; x++ {
			switch w.get(lib.Vector2{X: x, Y: y}) {
			case Rock:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}
