package main

import "sort"

type Monkey struct {
	inspections   int
	items         []int
	op            func(worry int) int
	divisibleTest int
	testTrue      int
	testFalse     int
}

type MonkeyTroop struct {
	monkies []Monkey
}

func makeMonkeyTroop() (troop MonkeyTroop) {
	troop.monkies = make([]Monkey, 0)
	return
}

func (troop MonkeyTroop) performMonkeyRound(monkey *Monkey, worryLevelsDividedByThree bool, combinedDivisor int) {
	for _, worry := range monkey.items {
		monkey.inspections++

		worry = monkey.op(worry)
		if worryLevelsDividedByThree {
			worry /= 3
		}

		var targetMonkey *Monkey
		if worry%monkey.divisibleTest == 0 {
			targetMonkey = &troop.monkies[monkey.testTrue]
		} else {
			targetMonkey = &troop.monkies[monkey.testFalse]
		}

		worry = worry % combinedDivisor

		targetMonkey.items = append(targetMonkey.items, worry)
	}
	monkey.items = monkey.items[:0]
}

func (troop *MonkeyTroop) PerformRounds(times int, worryLevelsDividedByThree bool) {
	combinedDivisor := 1
	for _, monkey := range troop.monkies {
		combinedDivisor *= monkey.divisibleTest
	}

	for round := 0; round < times; round++ {
		for i := 0; i < len(troop.monkies); i++ {
			troop.performMonkeyRound(&troop.monkies[i], worryLevelsDividedByThree, combinedDivisor)
		}
	}
}

func (troop MonkeyTroop) GetMonkeyBusinessLevel() int {
	inspections := make([]int, len(troop.monkies))
	for i, monkey := range troop.monkies {
		inspections[i] = monkey.inspections
	}

	sort.SliceStable(inspections, func(i, j int) bool { return inspections[i] > inspections[j] })

	return inspections[0] * inspections[1]
}
