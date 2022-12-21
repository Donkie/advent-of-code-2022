package main

type Operator uint8

const (
	None Operator = iota
	Add
	Sub
	Mul
	Div
)

type Monkey struct {
	name            string
	isSolved        bool
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

	// stack := []*Monkey{monk}

	// for len(stack) > 0 {
	// 	// Pop item from stack
	// 	stack, monk = stack[:len(stack)-1], stack[len(stack)-1]

	// }

}
