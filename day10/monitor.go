package main

import (
	"fmt"
)

type CRT struct {
}

func makeCRT() (crt CRT) {
	return
}

func (crt CRT) Draw(machine *Machine) {
	drawXPos := ((machine.Cycle - 1) % 40)
	spriteLeft := machine.RegX - 1
	spriteRight := machine.RegX + 1

	spriteVisible := drawXPos >= spriteLeft && drawXPos <= spriteRight
	if spriteVisible {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}

	if drawXPos == 39 {
		fmt.Println()
	}
}
