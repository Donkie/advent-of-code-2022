package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Represents an elf rucksack
type Rucksack string

// Gets the contents of the first compartment of this rucksack
func (rucksack Rucksack) GetFirstCompartment() Rucksack {
	l := len(rucksack)
	return Rucksack(rucksack[0 : l/2])
}

// Gets the contents of the second compartment of this rucksack
func (rucksack Rucksack) GetSecondCompartment() Rucksack {
	l := len(rucksack)
	return Rucksack(rucksack[l/2 : l])
}

// Finds the common item among a list of rucksacks.
// Returns an error if no common item was found.
// If multiple common items exist, any one of them will be returned.
func FindCommonItem(rucksacks []Rucksack) (byte, error) {
	// Find the common item from a variable number of arrays using the hash-map method
	// Increment the item's value in the "m" map every time it's found in an rucksack
	m := make(map[byte]int)
	for _, rucksack := range rucksacks {
		// Since the rucksacks can contain duplicates, we use a Set-ish map to identify if
		// we've already found this item in this rucksack so we don't count it twice.
		s := make(map[byte]bool, len(rucksack))
		for i := 0; i < len(rucksack); i++ {
			item := rucksack[i]
			if !s[item] {
				m[item]++
				s[item] = true
			}
		}
	}
	for item, count := range m {
		if count == len(rucksacks) {
			return item, nil
		}
	}
	return 0, fmt.Errorf("No common item found.")
}

// Calculates the item priority from its byte value
func getItemPriority(item byte) int {
	if item >= 'a' {
		return int(item) - ('a' - 1)
	} else {
		return int(item) - ('A' - 27)
	}
}

// Parses the input text file into a list of rucksacks
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

// Gets a list of the two compartment's common item for each input rucksack
// Exits the program if any rucksack didn't have a common item in the two compartments
func getCommonItems(rucksacks []Rucksack) []byte {
	var items []byte
	for _, rucksack := range rucksacks {
		item, err := FindCommonItem([]Rucksack{rucksack.GetFirstCompartment(), rucksack.GetSecondCompartment()})
		if err != nil {
			log.Fatalf("No common item found among the two compartments in the rucksack: %s\n", rucksack)
		}
		items = append(items, item)
	}
	return items
}

// Sums the priorities of a list of items
func getItemsSummedPriority(items []byte) int {
	sum := 0
	for _, item := range items {
		sum += getItemPriority(item)
	}
	return sum
}

// Finds the badges among the full list of rucksacks
// Returns the list of badges
// Exits the program if the input list of rucksacks isn't a multiple of 3
// Exits the program if any triplet of elves didn't have a badge
func getBadges(rucksacks []Rucksack) []byte {
	if len(rucksacks)%3 != 0 {
		log.Fatalln("The input rucksack slice isn't a multiple of 3 long.")
		return nil
	}
	badges := make([]byte, len(rucksacks)/3)
	for i := 0; i < len(badges); i++ {
		rucksackIndex := i * 3
		rucksackTriplet := rucksacks[rucksackIndex : rucksackIndex+3]
		badge, err := FindCommonItem(rucksackTriplet)
		if err != nil {
			for _, rucksack := range rucksackTriplet {
				log.Println(rucksack)
			}
			log.Fatalln("No common item found among the rucksacks printed above.")
			break
		}
		badges[i] = badge
	}
	return badges
}

func main() {
	rucksacks := parseInput("input.txt")

	commonItems := getCommonItems(rucksacks)
	commonItemsSum := getItemsSummedPriority(commonItems)
	fmt.Printf("Part 1 - common items sum: %d\n", commonItemsSum)

	badges := getBadges(rucksacks)
	badgesSum := getItemsSummedPriority(badges)
	fmt.Printf("Part 2 - badges sum: %d\n", badgesSum)

}
