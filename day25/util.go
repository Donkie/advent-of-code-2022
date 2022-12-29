package main

import "advent-of-code-2022/lib"

func SNAFUToNumberList(snafus []string) []int {
	numbers := make([]int, len(snafus))
	for i, snafu := range snafus {
		numbers[i] = SNAFUToNumber(snafu)
	}
	return numbers
}

func ParseInputAsLines(fileName string) (strings []string) {
	lib.ParseInputByLine(fileName, func(line string) error {
		strings = append(strings, line)
		return nil
	})
	return
}

func SumList(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}
