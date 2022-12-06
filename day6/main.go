package main

import (
	"advent-of-code-2022/lib"
	"log"
)

// Returns whether all characters in the input string are unique
func allCharactersUnique(s string) bool {
	charMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for _, cnt := range charMap {
		if cnt > 1 {
			return false
		}
	}
	return true
}

// Gets a marker position in the string, by looking for slices of the string where all characters are unique.
// Returns the index of the last character of the marker
// The slice size is determined by the size argument.
func GetMarkerPosition(s string, size int) int {
	if len(s) < size {
		log.Fatal("String can't contain any marker, it's too short.")
		return 0
	}

	for i := size - 1; i < len(s); i++ {
		slice := s[i-size+1 : i+1]
		if allCharactersUnique(slice) {
			return i + 1
		}
	}
	log.Fatal("Failed to find any marker in the string.")
	return 0
}

func main() {
	lib.ParseInputByLine("input.txt", func(line string) error {
		markerPos := GetMarkerPosition(line, 4)
		log.Printf("Part 1 - Marker position: %d", markerPos)

		markerPos = GetMarkerPosition(line, 14)
		log.Printf("Part 2 - Marker position: %d", markerPos)
		return nil
	})
}
