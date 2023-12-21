package y2023

import "testing"

func TestDay21(t *testing.T) {
	input := `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

	result := Day21Tester(input)

	if result != 16 {
		t.Errorf("expected 16 got %d", result)
	}
}
