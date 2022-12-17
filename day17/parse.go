package main

import (
	"log"
	"os"
)

func ParseJetStream(fileName string) *JetStream {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	dirs := make([]Direction, len(bytes))
	for i, c := range bytes {
		if c == '<' {
			dirs[i] = Left
		} else {
			dirs[i] = Right
		}
	}
	js := new(JetStream)
	js.dir = dirs
	return js
}
