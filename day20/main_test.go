package main

import "testing"

func TestExamplePart1(t *testing.T) {
	values := ParseNumbers("input_test.txt")
	numbers := makeNumbers(values)
	numbers.Mix()
	zeroIdx := numbers.GetIndexOfZero()
	c1, c2, c3 := numbers.Get(zeroIdx+1000), numbers.Get(zeroIdx+2000), numbers.Get(zeroIdx+3000)
	cSum := c1 + c2 + c3

	if c1 != 4 {
		t.Errorf("First coordinate should be %d, it was %d though", 4, c1)
	}
	if c2 != -3 {
		t.Errorf("First coordinate should be %d, it was %d though", -3, c2)
	}
	if c3 != 2 {
		t.Errorf("First coordinate should be %d, it was %d though", 2, c3)
	}
	if cSum != 3 {
		t.Errorf("The coordinate sum should be %d, it was %d though", 3, cSum)
	}
}

func TestExamplePart2(t *testing.T) {
	values := ParseNumbers("input_test.txt")
	numbers := makeNumbers(values)
	numbers.ApplyDecryptionKey(811589153)
	for i := 0; i < 10; i++ {
		numbers.Mix()
	}
	zeroIdx := numbers.GetIndexOfZero()
	c1, c2, c3 := numbers.Get(zeroIdx+1000), numbers.Get(zeroIdx+2000), numbers.Get(zeroIdx+3000)
	cSum := c1 + c2 + c3

	if c1 != 811589153 {
		t.Errorf("First coordinate should be %d, it was %d though", 4, c1)
	}
	if c2 != 2434767459 {
		t.Errorf("First coordinate should be %d, it was %d though", -3, c2)
	}
	if c3 != -1623178306 {
		t.Errorf("First coordinate should be %d, it was %d though", 2, c3)
	}
	if cSum != 1623178306 {
		t.Errorf("The coordinate sum should be %d, it was %d though", 3, cSum)
	}
}
