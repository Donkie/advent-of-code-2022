package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? (.+)`)

func ParseValveDescription(line string, graph *Graph) error {
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return fmt.Errorf("Failed to match regex")
	}

	name := matches[1]

	valve := graph.GetOrAddValve(name)

	flowRate, err := strconv.Atoi(matches[2])
	if err != nil {
		return err
	}
	valve.flowRate = flowRate

	var neighbours []*Valve
	neighbourNames := strings.Split(matches[3], ", ")
	for _, neighbourName := range neighbourNames {
		neighbours = append(neighbours, graph.GetOrAddValve(neighbourName))
	}
	valve.neighbours = neighbours

	return nil
}

func ParseValveGraph(fileName string) Graph {
	graph := makeGraph()
	lib.ParseInputByLine(fileName, func(line string) error {
		return ParseValveDescription(line, &graph)
	})
	return graph
}
