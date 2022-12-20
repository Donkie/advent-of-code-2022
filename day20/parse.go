package main

import (
	"advent-of-code-2022/lib"
	"strconv"
)

func ParseNumbers(fileName string) (nmbrs []int) {
	lib.ParseInputByLine(fileName, func(line string) error {
		n, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		nmbrs = append(nmbrs, n)
		return nil
	})
	return
}
