package main

import (
	"advent-of-code-2022/lib"
	"log"
)

type Blueprint struct {
	oreRobotOreCost        int
	clayRobotOreCost       int
	obsidianRobotOreCost   int
	obsidianRobotClayCost  int
	geodeRobotOreCost      int
	geodeRobotObsidianCost int
}

type State struct {
	t         int
	ore       int
	clay      int
	obsidian  int
	geodes    int
	oreBots   int
	clayBots  int
	obsBots   int
	geodeBots int
}

func (s *State) Print() {
	log.Printf("t: %d (ore %d, clay %d, obsidian %d, geodes %d) (orebots %d, claybots %d, obsidianbots %d, geodebots %d)",
		s.t,
		s.ore,
		s.clay,
		s.obsidian,
		s.geodes,
		s.oreBots,
		s.clayBots,
		s.obsBots,
		s.geodeBots,
	)
}

type Operation uint8

const (
	Nothing Operation = iota
	MakeOreBot
	MakeClayBot
	MakeObsidianBot
	MakeGeodeBot
)

// Performs an operation on a state and returns the new state
// The second argument is true if this operation is possible to do for the current state, otherwise false (feasibility)
func stateFun(blueprint Blueprint, X State, op Operation) (State, bool) {
	// Start bot fabrication
	switch op {
	case MakeOreBot:
		if X.ore < blueprint.oreRobotOreCost {
			return State{}, false
		}
		X.ore -= blueprint.oreRobotOreCost
	case MakeClayBot:
		if X.ore < blueprint.clayRobotOreCost {
			return State{}, false
		}
		X.ore -= blueprint.clayRobotOreCost
	case MakeObsidianBot:
		if X.ore < blueprint.obsidianRobotOreCost || X.clay < blueprint.obsidianRobotClayCost {
			return State{}, false
		}
		X.ore -= blueprint.obsidianRobotOreCost
		X.clay -= blueprint.obsidianRobotClayCost
	case MakeGeodeBot:
		if X.ore < blueprint.geodeRobotOreCost || X.obsidian < blueprint.geodeRobotObsidianCost {
			return State{}, false
		}
		X.ore -= blueprint.geodeRobotOreCost
		X.obsidian -= blueprint.geodeRobotObsidianCost
	}

	// Gather resources
	X.t++
	X.ore += X.oreBots
	X.clay += X.clayBots
	X.obsidian += X.obsBots
	X.geodes += X.geodeBots

	// Bot finished
	switch op {
	case MakeOreBot:
		X.oreBots++
	case MakeClayBot:
		X.clayBots++
	case MakeObsidianBot:
		X.obsBots++
	case MakeGeodeBot:
		X.geodeBots++
	}

	return X, true
}

func getBestcaseOutcome(X State, maxTime int) int {
	// Returns the maximum possible geodes we could ever hope to achieve from this state
	// That is, assume we can create one new geode bot every minute, how many geodes would we have by the end?
	bots := X.geodeBots
	geodes := X.geodes + bots
	for t := X.t; t < maxTime; t++ {
		bots++
		geodes += bots
	}
	return geodes
}

func getNumberOfGeodesSingle(blueprint Blueprint, maxTime int) (bestGeodes int) {
	startingPoint := State{oreBots: 1}

	// Simple cache so we don't process the same state multiple times
	cache := make(map[State]struct{}, 0)

	// Calculate the highest cost of each resource
	// It's completely useless making more bots than what any production requires, since we can't spend it
	highestOreCost := lib.Max(blueprint.clayRobotOreCost, blueprint.obsidianRobotOreCost, blueprint.oreRobotOreCost, blueprint.geodeRobotOreCost)
	highestClayCost := blueprint.obsidianRobotClayCost
	highestObsidianCost := blueprint.geodeRobotObsidianCost

	// BFS search, initialize the queue
	queue := []State{
		startingPoint,
	}

	// Keep track of the best prediction we've found for every time step
	// If we start on a branch and realize that another branch already achieves better than this, we can cut it away
	bestPrediction := make([]int, maxTime+1)

	var X State
	for len(queue) > 0 {
		// Remove first item from the queue, which makes it a BFS search
		X, queue = queue[0], queue[1:]

		// Check cache for this state, if it doesn't exist, add it
		_, ok := cache[X]
		if ok {
			continue
		}
		cache[X] = struct{}{}

		// If this state can't hope to achieve better than some other state at this time already could, then we can leave
		prediction := getBestcaseOutcome(X, maxTime)
		if prediction < bestPrediction[X.t] {
			continue
		}

		for opi := uint8(0); opi < 5; opi++ {
			// Attempt to perform every operation
			Xnext, feasible := stateFun(blueprint, X, Operation(opi))

			// If the operation is possible, and we don't end up in a resource-wasting state
			if feasible && Xnext.oreBots <= highestOreCost && Xnext.clayBots <= highestClayCost && Xnext.obsBots <= highestObsidianCost {
				// Update best geode prediction
				prediction := getBestcaseOutcome(Xnext, maxTime)
				if prediction > bestPrediction[Xnext.t] {
					bestPrediction[Xnext.t] = prediction
				}

				if Xnext.t >= maxTime {
					// We're finished, report the outcome and leave
					if Xnext.geodes > bestGeodes {
						bestGeodes = Xnext.geodes
					}
				} else {
					// We're not finished yet, this branch needs to go deeper
					queue = append(queue, Xnext)
				}
			}
		}
	}
	return bestGeodes
}

func GetNumberOfGeodes(blueprints []Blueprint, maxTime int) (geodes []int) {
	geodes = make([]int, len(blueprints))
	for i := 0; i < len(blueprints); i++ {
		log.Printf("Processing blueprint %d/%d", i+1, len(blueprints))
		geodes[i] = getNumberOfGeodesSingle(blueprints[i], maxTime)
	}
	return
}

func GetQualitySum(geodes []int) (qly int) {
	for idx, num := range geodes {
		qly += (idx + 1) * num
	}
	return
}
