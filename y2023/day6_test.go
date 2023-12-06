package y2023

import "testing"

func TestDay6(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`

	result := Day6(input)

	if result != 288 {
		t.Errorf("expected 288 got %v", result)
	}

	result = Day6Part2(input)

	if result != 71503 {
		t.Errorf("expected 71503 got %v", result)
	}
}
