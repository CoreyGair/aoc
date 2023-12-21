package y2023

import "testing"

func TestDay16(t *testing.T) {
	input := `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

	result := Day16(input)

	if result != 46 {
		t.Errorf("expected 46 got %d", result)
	}

	result = Day16Part2(input)

	if result != 51 {
		t.Errorf("expected 51 got %d", result)
	}
}
