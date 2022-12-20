package main

import (
	"fmt"
	"log"
	"strings"
)

// Plain % in Go doesn't handle negative numbers like we expect
func modulus(d, m int) int {
	res := d % m
	if res < 0 {
		return res + m
	}
	return res
}

func insertAtIndex[V any](slice []V, idx int, value V) []V {
	slice = append(slice[:idx+1], slice[idx:]...)
	slice[idx] = value
	return slice
}

func removeAtIndex[V any](slice []V, idx int) []V {
	return append(slice[:idx], slice[idx+1:]...)
}

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

// 1, -3, 2, 3, -2, 0, 4
// curidx 1 newidx 5
// 1, 2, 3, -2, 0, 4

// Performs the mixing operation on the numbers
func (nmbrs *Numbers) Mix() {
	// nmbrs.Print()
	for _, nmbr := range nmbrs.order {
		curIdx := nmbrs.GetIndexByPointer(nmbr)
		newIdx := modulus(curIdx+nmbr.value, nmbrs.size-1)
		nmbrs.values = removeAtIndex(nmbrs.values, curIdx)
		nmbrs.values = insertAtIndex(nmbrs.values, newIdx, nmbr)
		// nmbrs.Print()
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
	idx = modulus(idx, nmbrs.size)
	return nmbrs.values[idx].value
}
