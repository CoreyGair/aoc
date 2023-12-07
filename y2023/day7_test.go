package y2023

import "testing"

func TestDay7(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	result := Day7(input)

	if result != 5905 {
		t.Errorf("expected 5905 got %d", result)
	}
}
