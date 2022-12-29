package main

import (
	"advent-of-code-2022/lib"
	"math"
)

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

var powCache = []int{}

func powBase5(y int) int {
	if len(powCache) == 0 {
		size := 1000
		powCache = make([]int, size)
		for yi := 0; yi < size; yi++ {
			powCache[yi] = powInt(5, yi)
		}
	}
	return powCache[y]
}

func SNAFUToNumber(snafu string) (number int) {
	n := len(snafu)
	for i, c := range snafu {
		place := n - i - 1
		placeValue := powBase5(place)
		switch c {
		case '-':
			placeValue *= -1
		case '=':
			placeValue *= -2
		default:
			placeValue *= int(c) - '0'
		}
		number += placeValue
	}
	return
}

func NumberToSNAFU(number int) (snafu string) {
	place := 0
	for true {
		placeValue := powBase5(place)
		placeHalf := (placeValue - 1) / 2

		if number <= placeHalf {
			break
		}

		value := lib.Mod(((number+placeHalf)/placeValue)+2, 5) - 2
		switch value {
		case -2:
			snafu = "=" + snafu
		case -1:
			snafu = "-" + snafu
		default:
			snafu = string('0'+byte(value)) + snafu
		}

		place++
	}
	return
}
