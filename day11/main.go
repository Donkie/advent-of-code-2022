package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`(?m)Starting items: ([\d, ]+)[\s]+Operation: new = old ([\w \d+*]+)[\s]+Test: divisible by (\d+)[\s]+If true: throw to monkey (\d+)[\s]+If false: throw to monkey (\d+)`)

func parseItemsStr(itemsStr string) ([]int, error) {
	itemsStrSplit := strings.Split(itemsStr, ", ")
	items := make([]int, len(itemsStrSplit))
	for i := 0; i < len(itemsStrSplit); i++ {
		item, err := strconv.Atoi(itemsStrSplit[i])
		if err != nil {
			return nil, err
		}
		items[i] = item
	}
	return items, nil
}

func parseOpStr(opStr string) (func(worry int) int, error) {
	if opStr == "* old" {
		return func(worry int) int { return worry * worry }, nil
	} else {
		op := opStr[0]
		arg, err := strconv.Atoi(opStr[2:])
		if err != nil {
			return nil, err
		}

		if op == '*' {
			return func(worry int) int { return worry * arg }, nil
		} else { // Assume +
			return func(worry int) int { return worry + arg }, nil
		}
	}
}

func parseMonkey(itemsStr string, opStr string, divisibleTestStr string, testTrueStr string, testFalseStr string) (monkey *Monkey, err error) {
	monkey = new(Monkey)

	items, err := parseItemsStr(itemsStr)
	if err != nil {
		return nil, err
	}
	monkey.items = items

	op, err := parseOpStr(opStr)
	if err != nil {
		return nil, err
	}
	monkey.op = op

	divisibleTest, err := strconv.Atoi(divisibleTestStr)
	if err != nil {
		return nil, err
	}
	monkey.divisibleTest = divisibleTest

	testTrue, err := strconv.Atoi(testTrueStr)
	if err != nil {
		return nil, err
	}
	monkey.testTrue = testTrue

	testFalse, err := strconv.Atoi(testFalseStr)
	if err != nil {
		return nil, err
	}
	monkey.testFalse = testFalse

	return
}

func ParseMonkeyTroop(fileName string) *MonkeyTroop {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	contents := string(bytes)

	monkeyTroop := makeMonkeyTroop()

	for _, match := range re.FindAllStringSubmatch(contents, -1) {
		itemsStr := match[1]
		opStr := match[2]
		divisibleTestStr := match[3]
		testTrueStr := match[4]
		testFalseStr := match[5]

		monkey, err := parseMonkey(itemsStr, opStr, divisibleTestStr, testTrueStr, testFalseStr)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		monkeyTroop.monkies = append(monkeyTroop.monkies, *monkey)
	}

	return &monkeyTroop
}

func main() {
	troop := ParseMonkeyTroop("input.txt")
	troop.PerformRounds(20, true)
	metric := troop.GetMonkeyBusinessLevel()

	log.Printf("Part 1 - Monkey business level: %d", metric)
}
