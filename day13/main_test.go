package main

import "testing"

func TestExamplePart1(t *testing.T) {
	pairs := ParseInputPairsFile("input_test.txt")
	metric := GetSumOfOrderedPairIndices(*pairs)

	actual := metric
	expected := 13
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

func TestExamplePair1(t *testing.T) {
	pairs := ParseInputPairs("[1,1,3,1,1]\n[1,1,5,1,1]")

	actual := (*pairs)[0].IsOrdered()
	expected := true
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair2(t *testing.T) {
	pairs := ParseInputPairs("[[1],[2,3,4]]\n[[1],4]")

	actual := (*pairs)[0].IsOrdered()
	expected := true
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair3(t *testing.T) {
	pairs := ParseInputPairs("[9]\n[[8,7,6]]")

	actual := (*pairs)[0].IsOrdered()
	expected := false
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair4(t *testing.T) {
	pairs := ParseInputPairs("[[4,4],4,4]\n[[4,4],4,4,4]")

	actual := (*pairs)[0].IsOrdered()
	expected := true
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair5(t *testing.T) {
	pairs := ParseInputPairs("[7,7,7,7]\n[7,7,7]")

	actual := (*pairs)[0].IsOrdered()
	expected := false
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair6(t *testing.T) {
	pairs := ParseInputPairs("[]\n[3]")

	actual := (*pairs)[0].IsOrdered()
	expected := true
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair7(t *testing.T) {
	pairs := ParseInputPairs("[[[]]]\n[[]]")

	actual := (*pairs)[0].IsOrdered()
	expected := false
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}

func TestExamplePair8(t *testing.T) {
	pairs := ParseInputPairs("[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]")

	actual := (*pairs)[0].IsOrdered()
	expected := false
	if actual != expected {
		t.Errorf("expected %t, actual %t", expected, actual)
	}
}
