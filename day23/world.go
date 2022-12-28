package main

import (
	"advent-of-code-2022/lib"
	"fmt"
)

type ElfObj struct {
	pos        lib.Vector2
	shouldMove bool
	moveTarget lib.Vector2
}

type Bounds struct {
	minx int
	miny int
	maxx int
	maxy int
}

func (b *Bounds) Update(p lib.Vector2) {
	b.minx = lib.Min(b.minx, p.X)
	b.miny = lib.Min(b.miny, p.Y)
	b.maxx = lib.Max(b.maxx, p.X)
	b.maxy = lib.Max(b.maxy, p.Y)
}

type Map map[lib.Vector2]*ElfObj

func (m Map) Get(v lib.Vector2) *ElfObj {
	return m[v]
}

type World struct {
	tiles   Map
	elves   []*ElfObj
	checker Checker
	bounds  Bounds
}

func (w *World) SimulateRound() (anyMove bool) {
	// Step 1: Propose move targets
	// This map contains all proposed destinations, and the first elf that wants to go there
	proposedTargets := make(Map)

	for _, elf := range w.elves {
		moveDir, allDirs := w.checker.GetDirection(w.tiles, elf.pos)
		if moveDir == None || allDirs == 0b1111 {
			// If all directions are available, that means no elf is nearby so we shouldn't move
			continue
		}
		moveTarget := elf.pos.Add(moveDir.ToVec())

		proposedElf, isProposed := proposedTargets[moveTarget]
		if isProposed {
			// Another elf already wants to move here, stop both of them
			proposedElf.shouldMove = false
			continue
		}

		elf.moveTarget = moveTarget
		elf.shouldMove = true
		proposedTargets[elf.moveTarget] = elf
	}

	// Step 2: Perform the move
	for _, elf := range w.elves {
		if elf.shouldMove {
			anyMove = true

			delete(w.tiles, elf.pos)

			elf.pos = elf.moveTarget
			elf.shouldMove = false

			w.tiles[elf.pos] = elf
		}
	}

	w.checker.IncrementStart()
	return
}

var DEBUG = false

func (w *World) Simulate(times int) {
	if DEBUG {
		fmt.Println("== Initial State ==")
		w.Print()
		fmt.Println()
	}
	for i := 0; i < times; i++ {
		w.SimulateRound()
		if DEBUG {
			fmt.Printf("== End of Round %d ==\n", i+1)
			w.Print()
			fmt.Println()
		}
	}
}

func (w *World) SimulateUntilEnd() (round int) {
	for true {
		round++
		anyMove := w.SimulateRound()
		if !anyMove {
			break
		}
	}
	return
}

func (w *World) GetBounds() (b Bounds) {
	for _, elf := range w.elves {
		b.Update(elf.pos)
	}
	return
}

func (w *World) Print() {
	bounds := w.GetBounds()
	for y := bounds.miny; y <= bounds.maxy; y++ {
		for x := bounds.minx; x <= bounds.maxx; x++ {
			if w.tiles[lib.Vector2{X: x, Y: y}] == nil {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (w *World) GetEmptySquaresInBoundingRegion() (sum int) {
	bounds := w.GetBounds()
	for x := bounds.minx; x <= bounds.maxx; x++ {
		for y := bounds.miny; y <= bounds.maxy; y++ {
			if w.tiles[lib.Vector2{X: x, Y: y}] == nil {
				sum++
			}
		}
	}
	return
}
