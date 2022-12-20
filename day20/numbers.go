package main

import (
	"advent-of-code-2022/lib"
	"fmt"
	"log"
	"strings"
)

type Number struct {
	value int
}

type Numbers struct {
	values []*Number
	order  []*Number
	size   int
}

func makeNumbers(values []int) (nmbrs Numbers) {
	nmbrs.size = len(values)
	nmbrs.values = make([]*Number, nmbrs.size)
	nmbrs.order = make([]*Number, nmbrs.size)

	for i, value := range values {
		nmbr := new(Number)
		nmbr.value = value
		nmbrs.values[i] = nmbr
		nmbrs.order[i] = nmbr
	}
	return
}

func (nmbrs *Numbers) Print() {
	strs := make([]string, nmbrs.size)
	for i, nmbr := range nmbrs.values {
		strs[i] = fmt.Sprintf("%d", nmbr.value)
	}
	log.Println(strings.Join(strs, ", "))
}

// Performs the mixing operation on the numbers
func (nmbrs *Numbers) Mix() {
	for _, nmbr := range nmbrs.order {
		curIdx := nmbrs.GetIndexByPointer(nmbr)
		newIdx := lib.Mod(curIdx+nmbr.value, nmbrs.size-1) // Size-1 because the last and first position in the array counts as the same thing
		nmbrs.values = lib.RemoveAtIndex(nmbrs.values, curIdx)
		nmbrs.values = lib.InsertAtIndex(nmbrs.values, newIdx, nmbr)
	}
}

func (nmbrs *Numbers) ApplyDecryptionKey(key int) {
	for _, nmbr := range nmbrs.values {
		nmbr.value *= key
	}
}

// Returns the index of a number by value
func (nmbrs *Numbers) GetIndexByPointer(nmbr *Number) int {
	for i, test := range nmbrs.values {
		if test == nmbr {
			return i
		}
	}
	log.Panic("Failed to find pointer in data")
	return 0
}

// Returns the index of a number by value
func (nmbrs *Numbers) GetIndexOfZero() int {
	for i, test := range nmbrs.values {
		if test.value == 0 {
			return i
		}
	}
	log.Panic("Failed to find 0 in data")
	return 0
}

// Returns the value of a number by index
func (nmbrs *Numbers) Get(idx int) int {
	idx = lib.Mod(idx, nmbrs.size)
	return nmbrs.values[idx].value
}
