package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"math"
)

type Valve struct {
	name       string
	idx        uint64
	flowRate   int
	neighbours []*Valve
	distances  map[*Valve]int
}

func (v *Valve) computeDistances() {
	// Compute the minimum distance to all other valves using a breadth-first search
	visited := lib.MakePtrSet[Valve]()
	var queue []struct {
		valve *Valve
		depth int
	}
	distances := make(map[*Valve]int)

	for _, neighbour := range v.neighbours {
		queue = append(queue, struct {
			valve *Valve
			depth int
		}{
			valve: neighbour,
			depth: 1,
		})
	}
	visited.Add(v)
	distances[v] = 0

	for len(queue) > 0 {
		testValve := queue[0]
		queue = queue[1:]

		visited.Add(testValve.valve)

		distances[testValve.valve] = testValve.depth

		for _, neighbour := range testValve.valve.neighbours {
			if visited.Contains(neighbour) {
				continue
			}

			queue = append(queue, struct {
				valve *Valve
				depth int
			}{
				valve: neighbour,
				depth: testValve.depth + 1,
			})
		}
	}

	v.distances = distances
}

type Graph struct {
	valves       []*Valve
	usefulValves []*Valve
	nameMap      map[string]*Valve
}

func makeGraph() (graph Graph) {
	graph.valves = make([]*Valve, 0)
	graph.nameMap = make(map[string]*Valve)
	return
}

func (g *Graph) AddValve(valve *Valve) {
	g.valves = append(g.valves, valve)
	g.nameMap[valve.name] = valve
}

func (g *Graph) GetOrAddValve(name string) *Valve {
	valve, ok := g.nameMap[name]
	if ok {
		return valve
	} else {
		valve = new(Valve)
		valve.name = name
		valve.idx = uint64(1 << len(g.valves))
		g.AddValve(valve)
		return valve
	}
}

func (g *Graph) computeDistances() {
	for _, valve := range g.valves {
		valve.computeDistances()
	}
}

func (g *Graph) getWorstValve() (out *Valve) {
	lowest := math.MaxInt
	for _, valve := range g.valves {
		if valve.flowRate > 0 && valve.flowRate < lowest {
			lowest = valve.flowRate
			out = valve
		}
	}
	return
}

func (g *Graph) computeUsefulValves() {
	g.usefulValves = make([]*Valve, 0)
	for _, valve := range g.valves {
		if valve.flowRate > 0 {
			g.usefulValves = append(g.usefulValves, valve)
		}
	}
}

func (g *Graph) getBestPressure(curValve *Valve, timeLeft int, visitedValves lib.PtrSet[Valve]) int {
	visitedValves = visitedValves.Copy()
	visitedValves.Add(curValve)

	if timeLeft <= 1 {
		return 0
	}

	thisValveTotalFlow := 0
	if curValve.flowRate > 0 { // This is just a simple fix to prevent us from ticking down the timer at the first valve
		timeLeft--
		thisValveTotalFlow = timeLeft * curValve.flowRate
	}

	nextBestTotalFlow := 0
	for _, nextValve := range g.valves {
		if nextValve.flowRate > 0 && !visitedValves.Contains(nextValve) {
			nextTimeLeft := timeLeft - curValve.distances[nextValve]
			nextFlow := g.getBestPressure(nextValve, nextTimeLeft, visitedValves)
			if nextFlow > nextBestTotalFlow {
				nextBestTotalFlow = nextFlow
			}
		}
	}
	return thisValveTotalFlow + nextBestTotalFlow
}

func cacheMapKey(curValve1 *Valve, curValve2 *Valve, timeLeft1 int, timeLeft2 int, visitedValves uint64) string {
	if curValve1.name > curValve2.name {
		curValve1, curValve2 = curValve2, curValve1
		timeLeft1, timeLeft2 = timeLeft2, timeLeft1
	}

	return fmt.Sprintf("%s%s%d-%d-%d", curValve1.name, curValve2.name, timeLeft1, timeLeft2, visitedValves)
}

var valueCache = make(map[string]int)

func (g *Graph) getBestPressureTwoOperators(curValve1 *Valve, curValve2 *Valve, timeLeft1 int, timeLeft2 int, visitedValves uint64) int {
	key := cacheMapKey(curValve1, curValve2, timeLeft1, timeLeft2, visitedValves)
	value, ok := valueCache[key]
	if ok {
		return value
	}

	visitedValves |= curValve1.idx | curValve2.idx

	thisValveTotalFlow := 0
	if curValve1.flowRate > 0 { // This is just a simple fix to prevent us from ticking down the timer at the first valve
		timeLeft1--
		timeLeft2--

		if timeLeft1 <= 0 && timeLeft2 <= 0 {
			return 0
		}

		if timeLeft1 > 0 {
			thisValveTotalFlow += timeLeft1 * curValve1.flowRate
		}
		if timeLeft2 > 0 {
			thisValveTotalFlow += timeLeft2 * curValve2.flowRate
		}
	}

	nextBestTotalFlow := 0
	for _, nextValve1 := range g.usefulValves {
		if (visitedValves & nextValve1.idx) == 0 {
			nextTimeLeft1 := timeLeft1 - curValve1.distances[nextValve1]

			for _, nextValve2 := range g.usefulValves {
				if nextValve2 != nextValve1 && (visitedValves&nextValve2.idx) == 0 {
					nextTimeLeft2 := timeLeft2 - curValve2.distances[nextValve2]

					nextFlow := g.getBestPressureTwoOperators(nextValve1, nextValve2, nextTimeLeft1, nextTimeLeft2, visitedValves)
					if nextFlow > nextBestTotalFlow {
						nextBestTotalFlow = nextFlow
					}
				}
			}
		}
	}
	totValue := thisValveTotalFlow + nextBestTotalFlow
	valueCache[key] = totValue
	return totValue
}

func (g *Graph) GetOptimalPressureReleaseAmount(timeLeft int) int {
	g.computeDistances()

	curValve := g.nameMap["AA"]
	visitedValves := lib.MakePtrSet[Valve]()

	return g.getBestPressure(curValve, timeLeft, visitedValves)
}

func (g *Graph) GetOptimalPressureReleaseAmountTwoOperators(timeLeft int) int {
	g.computeDistances()
	g.computeUsefulValves()

	curValve1 := g.nameMap["AA"]
	curValve2 := g.nameMap["AA"]

	return g.getBestPressureTwoOperators(curValve1, curValve2, timeLeft, timeLeft, 0)
}
