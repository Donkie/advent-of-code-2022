package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Represents an elf and the food that it carries
type Elf struct {
	foodItems []int // List of food item calories
}

// Calculates the total number of calories of all food items that the elf is carrying.
func (elf Elf) totalCalories() int {
	sum := 0
	for _, item := range elf.foodItems {
		sum += item
	}
	return sum
}

// Parses the elf food item calories input file and returns a list of elves
func parseInput(fileName string) []Elf {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elves := make([]Elf, 0)
	var items []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && items != nil && len(items) > 0 {
			var elf Elf
			elf.foodItems = items
			elves = append(elves, elf)
			items = nil
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(fmt.Sprintf("Failed converting \"%s\" to integer", line))
			}
			items = append(items, calories)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elves
}

// Calculates the total number of calories that the numHighest elves that are carrying the most are carrying.
// For example, numHighest = 1 will return the total number of calories that the most encumbered elf is carrying.
// numHighest = 3 will return the total number of calories that the 3 most encumbered elves are carrying.
//
// Side effects: The input elves array will be sorted in the process.
func getCaloriesOfHighestElves(elves []Elf, numHighest int) int {
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].totalCalories() > elves[j].totalCalories()
	})

	sum := 0
	for i := 0; i < numHighest; i++ {
		sum += elves[i].totalCalories()
	}
	return sum
}

func main() {
	elves := parseInput("input.txt")

	highestCalories := getCaloriesOfHighestElves(elves, 1)
	fmt.Printf("Part 1: The elf carrying the most calories is carrying %d calories.\n", highestCalories)

	highestCalories = getCaloriesOfHighestElves(elves, 3)
	fmt.Printf("Part 2: The three elves carrying the most calories is carrying %d calories.\n", highestCalories)
}
