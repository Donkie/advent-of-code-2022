package main

import "log"

type Operator uint8

const (
	None Operator = iota
	Add
	Sub
	Mul
	Div
	Eq
)

type Monkey struct {
	name            string
	isSolved        bool
	isValueMonkey   bool
	value           int
	dependency1Name string
	dependency2Name string
	dependency1     *Monkey
	dependency2     *Monkey
	op              Operator
}

type Riddle struct {
	monkieNames map[string]*Monkey
}

func (r *Riddle) GetMonkey(name string) *Monkey {
	return r.monkieNames[name]
}

// Gets the value of a specific monkey by resolving the dependency tree
func (r *Riddle) GetValue(monk *Monkey) int {
	if monk.isSolved {
		return monk.value
	}

	monk1Val := r.GetValue(monk.dependency1)
	monk2Val := r.GetValue(monk.dependency2)

	switch monk.op {
	case Add:
		monk.value = monk1Val + monk2Val
	case Sub:
		monk.value = monk1Val - monk2Val
	case Mul:
		monk.value = monk1Val * monk2Val
	case Div:
		monk.value = monk1Val / monk2Val
	}
	monk.isSolved = true

	return monk.value
}

func dependencyGraphContainsHuman(monk *Monkey) bool {
	if monk.name == "humn" {
		return true
	} else if monk.isValueMonkey {
		return false
	}

	return dependencyGraphContainsHuman(monk.dependency1) || dependencyGraphContainsHuman(monk.dependency2)
}

// Solves the X in "X <op> second == result"
func solveFirst(op Operator, second int, result int) int {
	switch op {
	case Add:
		return result - second
	case Sub:
		return result + second
	case Mul:
		return result / second
	case Div:
		return result * second
	case Eq:
		return second
	}
	log.Panic("Invalid operator")
	return 0
}

// Solves the X in "first <op> X == result"
func solveSecond(op Operator, first int, result int) int {
	switch op {
	case Add:
		return result - first
	case Sub:
		return first - result
	case Mul:
		return result / first
	case Div:
		return first / result
	case Eq:
		return first
	}
	log.Panic("Invalid operator")
	return 0
}

func (r *Riddle) solveForOther(monk *Monkey, targetValue int) {
	if monk.name == "humn" {
		monk.value = targetValue
		return
	}

	if (monk.dependency1.name != "humn" && monk.dependency1.isSolved) || dependencyGraphContainsHuman(monk.dependency2) {
		// We can get the value of dependency1
		dep1Value := r.GetValue(monk.dependency1)
		dep2TargetValue := solveSecond(monk.op, dep1Value, targetValue)

		r.solveForOther(monk.dependency2, dep2TargetValue)
	} else if (monk.dependency2.name != "humn" && monk.dependency2.isSolved) || dependencyGraphContainsHuman(monk.dependency1) {
		// We can get the value of dependency2
		dep2Value := r.GetValue(monk.dependency2)
		dep1TargetValue := solveFirst(monk.op, dep2Value, targetValue)

		r.solveForOther(monk.dependency1, dep1TargetValue)
	} else {
		log.Panic("We shouldn't get here, perhaps both edges contains the human?")
	}
}

func (r *Riddle) GetNumberToYell() int {
	root := r.monkieNames["root"]
	root.op = Eq

	r.solveForOther(root, 0) // The 0 here is not gonna be used for anything.

	return r.monkieNames["humn"].value
}
