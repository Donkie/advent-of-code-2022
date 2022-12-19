package main

import "testing"

func TestExamplePart1(t *testing.T) {
	blueprints := ParseBlueprints("input_test.txt")
	numGeodesList := GetNumberOfGeodes(blueprints, 24)
	qualitySum := GetQualitySum(numGeodesList)

	actual := qualitySum
	expected := 33
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
