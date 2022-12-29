package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"math"
	"strings"
)

type Map [][]bool

var ENTRY_X = 1
var ENTRY_Y = 0

type World struct {
	blizzards *Blizzards
	entryX    int
	entryY    int
	exitX     int
	exitY     int
	mapw      int
	maph      int
}

func makeWorldFromInput(mapstr string) (w World) {
	lines := strings.Split(strings.Trim(mapstr, "\n"), "\n")

	w.blizzards = newBlizzardsFromInput(lines)
	w.mapw = w.blizzards.mapw
	w.maph = w.blizzards.maph
	w.entryX = 1
	w.entryY = 0
	w.exitX = w.blizzards.mapw
	w.exitY = w.blizzards.maph + 1
	return w
}

func (w *World) IsOOB(x int, y int) bool {
	if (x == w.entryX && y == w.entryY) ||
		(x == w.exitX && y == w.exitY) ||
		(x > 0 && y > 0 && x <= w.mapw && y <= w.maph) {
		return false
	}
	return true
}

func (w *World) Print(px int, py int, time int) {
	for y := 0; y < w.maph+2; y++ {
		for x := 0; x < w.mapw+2; x++ {
			if w.IsOOB(x, y) {
				fmt.Print("#")
			} else if x == px && y == py {
				fmt.Print("E")
			} else {
				blizzChar := w.blizzards.GetChar(x, y, time)
				if blizzChar != ' ' {
					fmt.Print(string(blizzChar))
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (w *World) PrintDuration(maxTime int) {
	fmt.Println("Initial state:")
	w.Print(-1, -1, 0)

	for t := 1; t <= maxTime; t++ {
		fmt.Printf("Minute %d:\n", t)
		w.Print(-1, -1, t)
	}
}

type State struct {
	px    int
	py    int
	time  int
	phase uint8
}

func (w *World) FindShortestPath(part bool) int {
	queue := make([]State, 1)
	queue[0] = State{
		px:    w.entryX,
		py:    w.entryY,
		time:  0,
		phase: 0,
	}

	cache := make(map[State]struct{})

	blizz := w.blizzards

	bestTime := math.MaxInt

	var X State
	for len(queue) > 0 {
		// BFS
		X, queue = queue[0], queue[1:]
		// DFS
		// n := len(queue) - 1
		// X, queue = queue[n], queue[:n]

		// Memoization
		Xperiodic := X
		Xperiodic.time %= blizz.CycleTime
		_, alreadyDone := cache[Xperiodic]
		if alreadyDone {
			continue
		}
		cache[Xperiodic] = struct{}{}

		nextPhase := X.phase

		// Check if we're done
		if part == false {
			if X.px == w.exitX && X.py == w.exitY {
				bestTime = lib.Min(bestTime, X.time)
				continue
			}
		} else {
			switch X.phase {
			case 0:
				// First reach goal
				if X.px == w.exitX && X.py == w.exitY {
					nextPhase++
				}
			case 1:
				// Go back to start
				if X.px == w.entryX && X.py == w.entryY {
					nextPhase++
				}
			case 2:
				// Reach goal again
				if X.px == w.exitX && X.py == w.exitY {
					bestTime = lib.Min(bestTime, X.time)
					continue
				}
			}
		}

		nextT := X.time + 1

		if nextT > bestTime {
			continue
		}

		if !w.IsOOB(X.px, X.py-1) && !blizz.Get(X.px, X.py-1, nextT) {
			// Move north
			queue = append(queue, State{
				px:    X.px,
				py:    X.py - 1,
				time:  nextT,
				phase: nextPhase,
			})
		}
		if !w.IsOOB(X.px-1, X.py) && !blizz.Get(X.px-1, X.py, nextT) {
			// Move west
			queue = append(queue, State{
				px:    X.px - 1,
				py:    X.py,
				time:  nextT,
				phase: nextPhase,
			})
		}
		if !blizz.Get(X.px, X.py, nextT) {
			// Stay
			queue = append(queue, State{
				px:    X.px,
				py:    X.py,
				time:  nextT,
				phase: nextPhase,
			})
		}
		if !w.IsOOB(X.px, X.py+1) && !blizz.Get(X.px, X.py+1, nextT) {
			// Move south
			queue = append(queue, State{
				px:    X.px,
				py:    X.py + 1,
				time:  nextT,
				phase: nextPhase,
			})
		}
		if !w.IsOOB(X.px+1, X.py) && !blizz.Get(X.px+1, X.py, nextT) {
			// Move east
			queue = append(queue, State{
				px:    X.px + 1,
				py:    X.py,
				time:  nextT,
				phase: nextPhase,
			})
		}
	}

	return bestTime
}
