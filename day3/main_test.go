package main

import (
	"fmt"
	"testing"
)

func TestFindCommonItem(t *testing.T) {
	var tests = []struct {
		rucksacks  []Rucksack
		commonItem byte
	}{
		{[]Rucksack{"abcDd", "hjjgfc"}, 'c'},
		{[]Rucksack{"JJS", "KKSs"}, 'S'},
		{[]Rucksack{"JJS", "KKSs", "ABCS"}, 'S'},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("Test %d", i)
		t.Run(testname, func(t *testing.T) {
			ans, err := FindCommonItem(tt.rucksacks)
			if err != nil {
				t.Error(err)
			}
			if ans != tt.commonItem {
				t.Errorf("Got %c, want %c", ans, tt.commonItem)
			}
		})
	}
}
