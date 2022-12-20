package lib

import (
	"math"
)

// Abs return the absolute value of the input value
func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// Max returns the maximum value of the inputs
func Max(nums ...int) int {
	curMax := math.MinInt
	for _, num := range nums {
		if num > curMax {
			curMax = num
		}
	}
	return curMax
}

// Min returns the minimum value of the inputs
func Min(nums ...int) int {
	curMin := math.MaxInt
	for _, num := range nums {
		if num < curMin {
			curMin = num
		}
	}
	return curMin
}

// greatest common divisor (GCD) via Euclidean algorithm
// From: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// From: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// Plain % in Go doesn't handle negative numbers like we expect
func Mod(d, m int) int {
	res := d % m
	if res < 0 {
		return res + m
	}
	return res
}
