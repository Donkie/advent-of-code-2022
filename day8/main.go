package main

import (
	"advent-of-code-2022/lib"
	"log"
)

func reverse[V any](arr []V) *[]V {
	reversed := make([]V, len(arr))
	for i := 0; i < len(arr); i++ {
		reversed[i] = arr[len(arr)-i-1]
	}
	return &reversed
}

func iterateSlice[V any](arr []V, reversed bool, cb func(item V) bool) {
	if reversed {
		for i := len(arr) - 1; i >= 0; i-- {
			if cb(arr[i]) {
				break
			}
		}
	} else {
		for i := 0; i < len(arr); i++ {
			if cb(arr[i]) {
				break
			}
		}
	}
}

// Forest represents a 2D collection of trees with different heights
type Forest struct {
	heightMap [][]byte
}

// Width returns the width of the forest
func (f Forest) Width() int {
	return len(f.heightMap)
}

// Height returns the height of the forest
func (f Forest) Height() int {
	return len(f.heightMap[0])
}

// Row returns a column of trees in the forest
func (f Forest) Row(i int) *[]byte {
	return &f.heightMap[i]
}

// Column returns a column of trees in the forest
func (f Forest) Column(i int) *[]byte {
	col := make([]byte, f.Height())
	for j := 0; j < f.Height(); j++ {
		col[j] = f.heightMap[j][i]
	}
	return &col
}

// ComputeVisibilityMap returns a boolean 2D map where each value is true if that tree is visible from the edge of the forest
func (f Forest) ComputeVisibilityMap() *[][]bool {
	visMap := make([][]bool, f.Height())
	for i := 0; i < f.Height(); i++ {
		visMap[i] = make([]bool, f.Width())
	}

	// Loop through top-to-bottom and check visibility left<->right
	for i := 0; i < f.Height(); i++ {
		row := f.heightMap[i]

		// Left -> right
		var curHeight byte = 0
		for j := 0; j < f.Width(); j++ {
			if row[j] > curHeight {
				curHeight = row[j]
				visMap[i][j] = true
			}
		}

		// Right -> left
		curHeight = 0
		for j := f.Width() - 1; j >= 0; j-- {
			if row[j] > curHeight {
				curHeight = row[j]
				visMap[i][j] = true
			}
		}
	}

	// Loop through from left to right and check visibility top<->bottom
	for j := 0; j < f.Width(); j++ {
		// Top -> Bottom
		var curHeight byte = 0
		for i := 0; i < f.Height(); i++ {
			if f.heightMap[i][j] > curHeight {
				curHeight = f.heightMap[i][j]
				visMap[i][j] = true
			}
		}

		// Bottom -> Top
		curHeight = 0
		for i := f.Height() - 1; i >= 0; i-- {
			if f.heightMap[i][j] > curHeight {
				curHeight = f.heightMap[i][j]
				visMap[i][j] = true
			}
		}
	}

	return &visMap
}

func getVisibleTrees(arr []byte, thisHeight byte) (num int) {
	for _, height := range arr {
		num++
		if height >= thisHeight {
			break
		}
	}
	return
}

// computeScenicScore returns the scenic score for a specific tree in the forest
func (f Forest) computeScenicScore(rowi int, coli int) int {
	thisHeight := f.heightMap[rowi][coli]
	row := f.Row(rowi)
	col := f.Column(coli)

	rightVis := getVisibleTrees((*row)[coli+1:], thisHeight)
	leftVis := getVisibleTrees(*reverse((*row)[:coli]), thisHeight)
	downVis := getVisibleTrees((*col)[rowi+1:], thisHeight)
	upVis := getVisibleTrees(*reverse((*col)[:rowi]), thisHeight)

	return rightVis * leftVis * downVis * upVis
}

// ComputeScenicScores returns a map of the scenic score for each tree in the forest
func (f Forest) ComputeScenicScores() *[][]int {
	scores := make([][]int, f.Height())
	for i := 0; i < f.Height(); i++ {
		scores[i] = make([]int, f.Width())

		for j := 0; j < f.Width(); j++ {
			scores[i][j] = f.computeScenicScore(i, j)
		}
	}
	return &scores
}

// CountTrue returns the number of true values in the 2D map
func CountTrue(arr [][]bool) (sum int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] {
				sum++
			}
		}
	}
	return
}

// Max returns the highest value in the 2D map
func Max(arr [][]int) (out int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > out {
				out = arr[i][j]
			}
		}
	}
	return
}

// numberLineToSlice takes in a string of numbers and converts it to a slice of bytes
func numberLineToSlice(line string) *[]byte {
	arr := make([]byte, len(line))
	for i := 0; i < len(line); i++ {
		// +1 to fix the fact that the shortest tree can be 0, which makes the other algorithms more tricky since a byte is unsigned.
		// By adding 1, we can represent 0 as a non-tree, which is used on the edge of the forest.
		arr[i] = line[i] - '0' + 1
	}
	return &arr
}

// ParseForest parses the input text into a Forest type
func ParseForest(fileName string) (forest Forest) {
	var heightMap [][]byte
	lib.ParseInputByLine(fileName, func(line string) error {
		heightMap = append(heightMap, *numberLineToSlice(line))
		return nil
	})
	forest.heightMap = heightMap
	return forest
}

func main() {
	forest := ParseForest("input.txt")
	visMap := forest.ComputeVisibilityMap()
	numVis := CountTrue(*visMap)
	log.Printf("Part 1 - Number of visible trees: %d", numVis)

	scores := forest.ComputeScenicScores()
	bestScore := Max(*scores)
	log.Printf("Part 2 - Score of best tree: %d", bestScore)
}
