package main

import "testing"

func assert[V comparable](t *testing.T, expected V, actual V) {
	if expected != actual {
		t.Errorf("Values not equal")
	}
}

func TestDirection(t *testing.T) {
	assert(t, Down, Right.Turn(Right))
	assert(t, Right, Up.Turn(Right))
	assert(t, Up, Right.Turn(Left))
	assert(t, Right, Down.Turn(Left))
}
