package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
)

type Block int

const (
	Air Block = iota
	Sand
	Rock
)

func blockToChar(blk Block) byte {
	switch blk {
	case Air:
		return '.'
	case Sand:
		return 'o'
	case Rock:
		return '#'
	default:
		return ' '
	}
}

func xyToIdx(x int, y int) int {
	return x<<16 + y
}

type Bounds struct {
	minx int
	miny int
	maxx int
	maxy int
}

type World struct {
	squares map[int]Block
	bounds  *Bounds
}

func makeWorld() (world World) {
	world.squares = make(map[int]Block)
	return
}

func (w *World) GetBlock(x int, y int) Block {
	blk, ok := w.squares[xyToIdx(x, y)]
	if !ok {
		return Air
	} else {
		return blk
	}
}

func (w *World) UpdateBounds(x int, y int) {
	if w.bounds == nil {
		w.bounds = &Bounds{
			minx: x,
			miny: y,
			maxx: x,
			maxy: y,
		}
	} else {
		w.bounds.minx = lib.Min(w.bounds.minx, x)
		w.bounds.miny = lib.Min(w.bounds.miny, y)
		w.bounds.maxx = lib.Max(w.bounds.maxx, x)
		w.bounds.maxy = lib.Max(w.bounds.maxy, y)
	}
}

func (w *World) Print() {
	if w.bounds == nil {
		log.Panic("Can't print yet, bounds not determined")
		return
	}
	for y := w.bounds.miny; y <= w.bounds.maxy; y++ {
		for x := w.bounds.minx; x <= w.bounds.maxx; x++ {
			fmt.Print(fmt.Sprintf("%c", blockToChar(w.GetBlock(x, y))))
		}
		fmt.Println()
	}
}

func (w *World) SetBlock(x int, y int, blk Block) {
	w.squares[xyToIdx(x, y)] = blk
	w.UpdateBounds(x, y)
}

func (w *World) SetBlockLine(x1 int, y1 int, x2 int, y2 int, blk Block) {
	if x1 == x2 {
		// Vertical line
		if y1 > y2 {
			// Swap around so y2 is always >= y1
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			w.SetBlock(x1, y, blk)
		}
	} else if y1 == y2 {
		// Horizontal line
		if x1 > x2 {
			// Swap around so y2 is always >= y1
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			w.SetBlock(x, y1, blk)
		}
	} else {
		log.Panicf("Can't set line, it is not straight!")
	}
}

func (w *World) SimulateSand(debugPrint bool) (restingSand int) {
	wentIntoAbyss := false
	for !wentIntoAbyss {
		restingSand++
		curX, curY := 500, 0
		for true {
			if w.GetBlock(curX, curY+1) == Air {
				// Empty spot down
				curY++
				if curY >= w.bounds.maxy {
					wentIntoAbyss = true
					restingSand--
					break
				}
			} else if w.GetBlock(curX-1, curY+1) == Air {
				// Empty spot down to the left
				curX--
				curY++
			} else if w.GetBlock(curX+1, curY+1) == Air {
				// Empty spot down to the right
				curX++
				curY++
			} else {
				// No empty spot, rest sand
				w.SetBlock(curX, curY, Sand)

				if debugPrint &&
					(restingSand == 1 || restingSand == 2 || restingSand == 5 || restingSand == 22 || restingSand == 24) {
					w.Print()
				}
				break
			}
		}
	}
	return
}
