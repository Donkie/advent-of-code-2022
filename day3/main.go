package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Rucksack string
type Compartment string

func (rucksack Rucksack) GetFirstCompartment() Compartment {
	l := len(rucksack)
	return Compartment(rucksack[0 : l/2])
}

func (rucksack Rucksack) GetSecondCompartment() Compartment {
	l := len(rucksack)
	return Compartment(rucksack[l/2 : l])
}

func (comp Compartment) GetDuplicateItem(otherComp Compartment) byte {
	l := len(comp)
	if len(otherComp) != l {
		log.Printf("Comp1: %s\n", comp)
		log.Printf("Comp2: %s\n", otherComp)
		log.Panicln("The two compartments aren't of equal length!")
		return 0
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if comp[i] == otherComp[j] {
				return comp[i]
			}
		}
	}
	log.Panicln("Couldn't find any duplicate items between the two compartments")
	return 0
}

func getItemPriority(item byte) int {
	if item >= 97 {
		return int(item) - 96
	} else {
		return int(item) - (65 - 27)
	}
}

func parseInput(fileName string) []Rucksack {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rucksacks []Rucksack
	for scanner.Scan() {
		rucksack := Rucksack(scanner.Text())
		rucksacks = append(rucksacks, rucksack)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rucksacks
}

func getSummedDuplicatePriorities(rucksacks []Rucksack) int {
	sum := 0
	for _, rucksack := range rucksacks {
		sum += getItemPriority(rucksack.GetFirstCompartment().GetDuplicateItem(rucksack.GetSecondCompartment()))
	}
	return sum
}

func main() {
	rucksacks := parseInput("input.txt")

	sum := getSummedDuplicatePriorities(rucksacks)
	fmt.Printf("Priority sum: %d\n", sum)
}
