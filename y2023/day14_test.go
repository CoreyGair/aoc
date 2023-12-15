package y2023

import (
	"reflect"
	"strings"
	"testing"
)

type RotateTest struct {
	input  []string
	output []string
}

var (
	rotateTests = []RotateTest{
		{
			output: []string{
				"O...",
				"....",
				"....",
				"....",
			},
			input: []string{
				"....",
				"....",
				"....",
				"O...",
			},
		},
		{
			output: []string{
				"...O",
				"#...",
			},
			input: []string{
				"O.",
				"..",
				"..",
				".#",
			},
		},
	}
)

func TestRotate(t *testing.T) {
	for _, test := range rotateTests {
		r := RotateClockwise(test.input)

		if !reflect.DeepEqual(r, test.output) {
			t.Errorf("expected:\n%s\ngot:\n%s\n", strings.Join(test.output, "\n"), strings.Join(r, "\n"))
		}
	}
}

func TestDay14(t *testing.T) {
	input := `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	result := Day14(input)

	if result != 136 {
		t.Errorf("expected 136 got %d", result)
	}

	result = Day14Part2(input)

	if result != 64 {
		t.Errorf("expected 64 got %d", result)
	}
}
