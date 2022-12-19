package main

import "log"

type Blueprint struct {
	oreRobotOreCost        uint8
	clayRobotOreCost       uint8
	obsidianRobotOreCost   uint8
	obsidianRobotClayCost  uint8
	geodeRobotOreCost      uint8
	geodeRobotObsidianCost uint8
}

type State struct {
	ore       uint8
	clay      uint8
	obsidian  uint8
	geodes    uint8
	oreBots   uint8
	clayBots  uint8
	obsBots   uint8
	geodeBots uint8
}

var maxTime = 24

type Operation uint8

const (
	Nothing Operation = iota
	MakeOreBot
	MakeClayBot
	MakeObsidianBot
	MakeGeodeBot
)

func stateFun(blueprint Blueprint, X State, op Operation) (State, bool) {
	X.ore += X.oreBots
	X.clay += X.clayBots
	X.obsidian += X.obsBots
	X.geodes += X.geodeBots

	switch op {
	case MakeOreBot:
		if X.ore < blueprint.oreRobotOreCost {
			return State{}, false
		}
		X.oreBots++
		X.ore -= blueprint.oreRobotOreCost
	case MakeClayBot:
		if X.ore < blueprint.clayRobotOreCost {
			return State{}, false
		}
		X.clayBots++
		X.ore -= blueprint.clayRobotOreCost
	case MakeObsidianBot:
		if X.ore < blueprint.obsidianRobotOreCost || X.clay < blueprint.obsidianRobotClayCost {
			return State{}, false
		}
		X.obsBots++
		X.ore -= blueprint.obsidianRobotOreCost
		X.clay -= blueprint.obsidianRobotClayCost
	case MakeGeodeBot:
		if X.ore < blueprint.geodeRobotOreCost || X.obsidian < blueprint.geodeRobotObsidianCost {
			return State{}, false
		}
		X.geodeBots++
		X.ore -= blueprint.geodeRobotOreCost
		X.obsidian -= blueprint.geodeRobotObsidianCost
	}

	return X, true
}

func getNumberOfGeodesSingle(blueprint Blueprint) int {
	// Backwards pass
	totCosts := make([]map[State]uint32, maxTime)
	bestOpMap := make([]map[State]Operation, maxTime)

	for t := uint8(22); t >= 0; t-- {
		totCosts[t] = make(map[State]uint32)
		bestOpMap[t] = make(map[State]Operation)

		log.Printf("t: %d", t)
		for ore := uint8(0); ore <= 8; ore++ {
			for clay := uint8(0); clay <= 30; clay++ {
				for obsidian := uint8(0); obsidian <= 30; obsidian++ {
					for geodes := uint8(0); geodes <= 20; geodes++ {
						for oreBots := uint8(1); oreBots <= 10; oreBots++ {
							for clayBots := uint8(0); clayBots <= 10; clayBots++ {
								for obsBots := uint8(0); obsBots <= 10; obsBots++ {
									for geodeBots := uint8(0); geodeBots <= 5; geodeBots++ {
										X := State{
											ore:       ore,
											clay:      clay,
											obsidian:  obsidian,
											geodes:    geodes,
											oreBots:   oreBots,
											clayBots:  clayBots,
											obsBots:   obsBots,
											geodeBots: geodeBots,
										}

										var bestOp Operation
										var bestOpCost uint32
										for opi := uint8(0); opi < 5; opi++ {
											op := Operation(opi)
											Xnext, feasible := stateFun(blueprint, X, op)
											if feasible {
												var totCost uint32
												if t == uint8(maxTime)-1 {
													totCost = uint32(Xnext.geodes)
												} else {
													totCost = uint32(Xnext.geodes) + totCosts[t+1][Xnext]
												}
												if totCost > bestOpCost {
													bestOp = op
													bestOpCost = totCost
												}
											}
										}
										bestOpMap[t][X] = bestOp
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func GetNumberOfGeodes(blueprints []Blueprint) (geodes []int) {
}

func GetQualitySum(geodes []int) (qly int) {

}
