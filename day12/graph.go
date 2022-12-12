package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"math"
)

func posToStr(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func smallestNodeInSet(set lib.PtrSet[Node]) *Node {
	var smallestNode *Node
	for node := range set.Data {
		if smallestNode == nil || smallestNode.dist > node.dist {
			smallestNode = node
		}
	}
	return smallestNode
}

type Node struct {
	neighbours []*Node
	dist       int
	prev       *Node
	isEnd      bool
}

func makeNode() (node Node) {
	return
}

type Graph struct {
	nodes     map[string]*Node
	startNode *Node
}

func makeGraph() (graph Graph) {
	graph.nodes = make(map[string]*Node)
	return
}

func (g *Graph) GetNode(xPos int, yPos int) *Node {
	key := posToStr(xPos, yPos)
	node := g.nodes[key]
	if node == nil {
		g.nodes[key] = new(Node)
	}
	return g.nodes[key]
}

func (g Graph) FindShortestPath() int {
	unvisited := lib.MakePtrSet[Node]()
	for _, node := range g.nodes {
		unvisited.Add(node)
		if node == g.startNode {
			node.dist = 0
		} else {
			node.dist = math.MaxInt
		}
	}

	var endNodeDist int

	for unvisited.Len() > 0 {
		curNode := smallestNodeInSet(unvisited)
		unvisited.Remove(curNode)

		if curNode.isEnd {
			endNodeDist = curNode.dist
			break
		}

		if curNode.dist == math.MaxInt {
			// Node is unreachable
			continue
		}

		for _, neighbor := range curNode.neighbours {
			if unvisited.Contains(neighbor) {
				newDist := curNode.dist + 1
				if newDist < neighbor.dist {
					neighbor.dist = newDist
					neighbor.prev = curNode
				}
			}
		}
	}

	return endNodeDist
}
