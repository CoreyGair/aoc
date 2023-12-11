package y2023

import "testing"

func TestDay10(t *testing.T) {
	input := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

	result := Day10(input)

	if result != 8 {
		t.Errorf("expected 8 got %d", result)
	}

	input = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

	result = Day10Part2(input)

	if result != 4 {
		t.Errorf("expected 4 got %d", result)
	}
}
