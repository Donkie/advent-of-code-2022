package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	foodItems []int
}

func (elf Elf) totalCalories() int {
	var sum int = 0
	for _, item := range elf.foodItems {
		sum += item
	}
	return sum
}

func parseInput(fileName string) []Elf {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfs := make([]Elf, 0)
	var items []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && items != nil && len(items) > 0 {
			var elf Elf
			elf.foodItems = items
			elfs = append(elfs, elf)
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

	return elfs
}

func main() {
	elfs := parseInput("input.txt")
	var highestCalories int = 0
	for _, elf := range elfs {
		totCal := elf.totalCalories()
		if totCal > highestCalories {
			highestCalories = totCal
		}
	}
	fmt.Printf("The elf carrying the most calories is carrying %d calories.\n", highestCalories)
}
