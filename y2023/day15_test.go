package y2023

import "testing"

func TestHash(t *testing.T) {
	result := HASH("HASH")
	if result != 52 {
		t.Errorf("expected 52 got %d", result)
	}
}

func TestDay15(t *testing.T) {
	input := `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

	result := Day15(input)

	if result != 1320 {
		t.Errorf("expected 1320 got %d", result)
	}

	result = Day15Part2(input)

	if result != 145 {
		t.Errorf("expected 145 got %d", result)
	}
}
