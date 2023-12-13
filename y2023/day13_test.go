package y2023

import "testing"

func TestSingleBitSet(t *testing.T) {
	if SingleBitSet(0) {
		t.Error("zero does not have single bit set")
	}
	if !SingleBitSet(1) {
		t.Error("1 does have single bit set")
	}
	if !SingleBitSet(2) {
		t.Error("2 does have single bit set")
	}

	if !SingleBitSet(64) {
		t.Error("64 does have single bit set")
	}
	if SingleBitSet(65) {
		t.Error("65 does not have single bit set")
	}
}

func TestDay13(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	result := Day13(input)

	if result != 405 {
		t.Errorf("expected 405 got %d", result)
	}

	result = Day13Part2(input)

	if result != 400 {
		t.Errorf("expected 400 got %d", result)
	}
}
