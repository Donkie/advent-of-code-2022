package main

import "testing"

func TestExamplePart1(t *testing.T) {
	forest := ParseForest("input_test.txt")
	visMap := forest.ComputeVisibilityMap()
	numVis := CountTrue(*visMap)

	actual := numVis
	expected := 21
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart12(t *testing.T) {
	forest := ParseForest("input_test2.txt")
	visMap := forest.ComputeVisibilityMap()
	numVis := CountTrue(*visMap)

	actual := numVis
	expected := 30
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePart2(t *testing.T) {
	forest := ParseForest("input_test.txt")
	scores := forest.ComputeScenicScores()
	bestScore := Max(*scores)

	actual := bestScore
	expected := 8
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
