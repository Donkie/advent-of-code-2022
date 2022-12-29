package main

import "testing"

func TestExamplePart1(t *testing.T) {
	snafus := ParseInputAsLines("input_test.txt")
	numbers := SNAFUToNumberList(snafus)
	sum := SumList(numbers)
	output := NumberToSNAFU(sum)

	actual := output
	expected := "2=-1=0"
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
