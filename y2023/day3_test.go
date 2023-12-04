package y2023

import "testing"

func TestDay3(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	result := Day3(input)

	if result != 4361 {
		t.Errorf("expected 4361 got %v", result)
	}
}
