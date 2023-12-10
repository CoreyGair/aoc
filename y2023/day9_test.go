package y2023

import "testing"

func TestDay9(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

	result := Day9(input)

	if result != 114 {
		t.Errorf("expected 114 got %d", result)
	}

	result = Day9Part2(input)

	if result != 2 {
		t.Errorf("expected 2 got %d", result)
	}
}
