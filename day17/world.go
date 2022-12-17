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

var MAXHEIGHT int = 100

type Space bool

const (
	Empty Space = false
	Rock  Space = true
)

type Shape []lib.Vector2

type ShapeType uint8

// The right-most piece is the first item of the list so we can use that to easily check the width of the piece
var shapes = []Shape{
	0: {{X: 3, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},               // -
	1: {{X: 2, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 2}}, // +
	2: {{X: 2, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}}, // _|
	3: {{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},               // |
	4: {{X: 1, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},               // o
}

var dirVec = []lib.Vector2{
	Left:  {X: -1, Y: 0},
	Right: {X: 1, Y: 0},
}

type RockFallingWorld struct {
	jetStream    JetStream
	field        [][]Space
	spawnOffset  int
	curShape     ShapeType
	highestPoint int
}

func makeRockFallingWorld(jetStream JetStream) (world RockFallingWorld) {
	world.jetStream = jetStream

	world.field = make([][]Space, MAXHEIGHT)
	for i := 0; i < MAXHEIGHT; i++ {
		world.field[i] = make([]Space, 7)
	}
	world.spawnOffset = 0

	return
}

func (w *RockFallingWorld) get(v lib.Vector2) Space {
	return w.field[(v.Y-w.spawnOffset)%MAXHEIGHT][v.X]
}

func (w *RockFallingWorld) set(v lib.Vector2, space Space) {
	w.field[(v.Y-w.spawnOffset)%MAXHEIGHT][v.X] = space
}

func (w *RockFallingWorld) rockFits(pos lib.Vector2, shape Shape) bool {
	if pos.X < 0 || (pos.X+shape[0].X) > 6 || pos.Y < 0 {
		return false
	}

	for _, piece := range shape {
		piecePos := pos.Add(piece)
		if w.get(piecePos) != Empty {
			return false
		}
	}
	return true
}

func (w *RockFallingWorld) solidifyRock(pos lib.Vector2, shape Shape) {
	for _, piece := range shape {
		piecePos := pos.Add(piece)
		w.set(piecePos, Rock)
		w.highestPoint = lib.Max(w.highestPoint, piecePos.Y+1-w.spawnOffset)
	}
}

func (w *RockFallingWorld) shiftField(steps int) {
	for i := 0; i < steps; i++ {
		y := (i - w.spawnOffset) % MAXHEIGHT
		for x := 0; x < 7; x++ {
			w.field[y][x] = Empty
		}
	}
	w.spawnOffset -= steps
}

func (w *RockFallingWorld) simulateRock() {
	// fmt.Println("New block")
	prevHeighest := w.GetHighestPoint()
	pos := lib.Vector2{X: 2, Y: w.GetSpawnY()}
	shape := shapes[w.curShape]
	for true {
		// Move sideways by jet
		dir := w.jetStream.PopDir()
		newPos := pos.Add(dirVec[dir])
		if w.rockFits(newPos, shape) {
			pos = newPos
		}

		// Move downwards
		newPos = pos.Add(lib.Vector2{X: 0, Y: -1})
		if w.rockFits(newPos, shape) {
			pos = newPos
			// w.PrintWithRock(pos, w.curShape)
		} else {
			w.solidifyRock(pos, shape)
			break
		}
	}

	// fmt.Println("Block done")
	// w.Print()

	if (w.GetSpawnY() + 4) > MAXHEIGHT {
		shiftAm := w.GetHighestPoint() - prevHeighest
		w.shiftField(shiftAm)
	}

	w.curShape = (w.curShape + 1) % 5
}

func (w *RockFallingWorld) Simulate(numRocks int) {
	for i := 0; i < numRocks; i++ {
		if i > 0 && i%10_000_000_000 == 0 {
			fmt.Printf("%d%%", 100*i/1_000_000_000_000)
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

func (w *RockFallingWorld) GetSpawnY() int {
	return w.highestPoint + w.spawnOffset + 3
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
	for y := MAXHEIGHT - 1; y >= 0; y-- {
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
	for y := MAXHEIGHT - 1; y >= 0; y-- {
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
