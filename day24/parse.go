package main

import (
	"log"
	"os"
)

func ParseWorld(fileName string) *World {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	world := makeWorldFromInput(string(bytes))
	return &world
}
