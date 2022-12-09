package main

import (
	"advent-of-code-2022/lib"
	"fmt"
)

// Recorder is an object that records unique vector2 values
type Recorder struct {
	records map[string]bool
}

// RecordPosition records the input position
func (rec *Recorder) RecordPosition(v lib.Vector2) {
	if rec.records == nil {
		rec.records = make(map[string]bool)
	}

	s := fmt.Sprintf("%d-%d", v.X, v.Y)
	rec.records[s] = true
}

// GetNumOfRecordedPositions returns the number of unique recorded positions
func (rec Recorder) GetNumOfRecordedPositions() int {
	return len(rec.records)
}
