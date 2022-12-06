package main

import (
	"testing"
)

func TestExamplePart1(t *testing.T) {
	var tests = []struct {
		message   string
		markerPos int
	}{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for i, test := range tests {
		actual := GetMarkerPosition(test.message, 4)
		if actual != test.markerPos {
			t.Errorf("Case %d: expected: %d, actual: %d\n", i, test.markerPos, actual)
		}
	}
}

func TestExamplePart2(t *testing.T) {
	var tests = []struct {
		message   string
		markerPos int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for i, test := range tests {
		actual := GetMarkerPosition(test.message, 14)
		if actual != test.markerPos {
			t.Errorf("Case %d: expected: %d, actual: %d\n", i, test.markerPos, actual)
		}
	}
}
