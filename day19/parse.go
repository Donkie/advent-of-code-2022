package main

import (
	"advent-of-code-2022/lib"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(?m)(\d+) ore.+?(\d+) ore.+?(\d+) ore.+?(\d+) clay.+?(\d+) ore.+?(\d+) obsidian`)

func ParseBlueprint(line string) (*Blueprint, error) {
	blueprint := new(Blueprint)

	match := re.FindStringSubmatch(line)
	oreRobotOreCost, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}
	blueprint.oreRobotOreCost = oreRobotOreCost
	clayRobotOreCost, err := strconv.Atoi(match[2])
	if err != nil {
		return nil, err
	}
	blueprint.clayRobotOreCost = clayRobotOreCost
	obsidianRobotOreCost, err := strconv.Atoi(match[3])
	if err != nil {
		return nil, err
	}
	blueprint.obsidianRobotOreCost = obsidianRobotOreCost
	obsidianRobotClayCost, err := strconv.Atoi(match[4])
	if err != nil {
		return nil, err
	}
	blueprint.obsidianRobotClayCost = obsidianRobotClayCost
	geodeRobotOreCost, err := strconv.Atoi(match[5])
	if err != nil {
		return nil, err
	}
	blueprint.geodeRobotOreCost = geodeRobotOreCost
	geodeRobotObsidianCost, err := strconv.Atoi(match[6])
	if err != nil {
		return nil, err
	}
	blueprint.geodeRobotObsidianCost = geodeRobotObsidianCost
	return blueprint, nil
}

func ParseBlueprints(fileName string) (blueprints []Blueprint) {
	lib.ParseInputByLine(fileName, func(line string) error {
		blueprint, err := ParseBlueprint(line)
		if err != nil {
			return err
		}
		blueprints = append(blueprints, *blueprint)
		return nil
	})
	return
}
