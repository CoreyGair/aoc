package y2023

import "testing"

func TestDay11(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

	result := Day11(input)

	if result != 374 {
		t.Errorf("expected 374 got %d", result)
	}
}
