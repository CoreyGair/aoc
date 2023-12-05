package y2023

import "testing"

type Day1Test struct {
	input  string
	result int
}

var (
	day1Tests = []Day1Test{
		{
			input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`,
			result: 142,
		},
		{
			input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
			result: 281,
		},
	}
)

func TestDay1(t *testing.T) {
	for _, test := range day1Tests {
		res := Day1(test.input)
		if res != test.result {
			t.Errorf("expected %v got %v", test.result, res)
		}
	}
}
