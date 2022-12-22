package main

import "testing"

func TestExamplePart1(t *testing.T) {
	world := ParseWorld("input_test.txt")
	world.Traverse()
	pw := world.GetPassword()

	actual := pw
	expected := 6032
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	world := ParseWorld("input_test.txt")
	world.InterpretCube(4, [][]Face{
		{FNone, FNone, FTop, FNone},
		{FBack, FLeft, FFront, FNone},
		{FNone, FNone, FBottom, FRight},
	}, map[Face]int{
		FTop:    0,
		FFront:  0,
		FBottom: 0,
		FRight:  1,
		FLeft:   0,
		FBack:   0,
	})
	world.Traverse()
	pw := world.GetPassword()

	actual := pw
	expected := 5031
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
