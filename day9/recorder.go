package main

import (
	"advent-of-code-2022/lib"
	"fmt"
)

type Recorder struct {
	records map[string]bool
}

func (rec *Recorder) RecordPosition(v lib.Vector2) {
	if rec.records == nil {
		rec.records = make(map[string]bool)
	}

	s := fmt.Sprintf("%d-%d", v.X, v.Y)
	rec.records[s] = true
}

func (rec Recorder) GetNumOfRecordedPositions() int {
	return len(rec.records)
}
