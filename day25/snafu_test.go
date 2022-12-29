package main

import "testing"

func TestNumberToSnafu(t *testing.T) {
	var n2s = []struct {
		number int
		snafu  string
	}{
		{
			number: 1,
			snafu:  "1",
		},
		{
			number: 2,
			snafu:  "2",
		},
		{
			number: 3,
			snafu:  "1=",
		},
		{
			number: 4,
			snafu:  "1-",
		},
		{
			number: 5,
			snafu:  "10",
		},
		{
			number: 6,
			snafu:  "11",
		},
		{
			number: 7,
			snafu:  "12",
		},
		{
			number: 8,
			snafu:  "2=",
		},
		{
			number: 9,
			snafu:  "2-",
		},
		{
			number: 10,
			snafu:  "20",
		},
		{
			number: 15,
			snafu:  "1=0",
		},
		{
			number: 20,
			snafu:  "1-0",
		},
		{
			number: 2022,
			snafu:  "1=11-2",
		},
		{
			number: 12345,
			snafu:  "1-0---0",
		},
		{
			number: 314159265,
			snafu:  "1121-1110-1=0",
		},
	}

	for _, testCase := range n2s {
		actual := NumberToSNAFU(testCase.number)
		if actual != testCase.snafu {
			t.Errorf("Converting %d to SNAFU failed. Expected: %s, Actual: %s", testCase.number, testCase.snafu, actual)
		}
	}
}

func TestSnafuToNumber(t *testing.T) {
	var s2n = []struct {
		snafu  string
		number int
	}{
		{
			snafu:  "1=-0-2",
			number: 1747,
		},
		{
			snafu:  "12111",
			number: 906,
		},
		{
			snafu:  "2=0=",
			number: 198,
		},
		{
			snafu:  "21",
			number: 11,
		},
		{
			snafu:  "2=01",
			number: 201,
		},
		{
			snafu:  "111",
			number: 31,
		},
		{
			snafu:  "20012",
			number: 1257,
		},
		{
			snafu:  "112",
			number: 32,
		},
		{
			snafu:  "1=-1=",
			number: 353,
		},
		{
			snafu:  "1-12",
			number: 107,
		},
		{
			snafu:  "12",
			number: 7,
		},
		{
			snafu:  "1=",
			number: 3,
		},
		{
			snafu:  "122",
			number: 37,
		},
	}

	for _, testCase := range s2n {
		actual := SNAFUToNumber(testCase.snafu)
		if actual != testCase.number {
			t.Errorf("Converting %s to number failed. Expected: %d, Actual: %d", testCase.snafu, testCase.number, actual)
		}
	}
}
