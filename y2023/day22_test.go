package y2023

import (
	"testing"
)

type OverlapsTest struct {
	A, B     Brick
	overlaps bool
}

func TestOverlaps(t *testing.T) {
	tests := []OverlapsTest{
		{
			A:        Brick{[2]Point{{0, 0, 0}, {1, 0, 0}}},
			B:        Brick{[2]Point{{1, 0, 0}, {1, 1, 0}}},
			overlaps: true,
		},
		{
			A:        Brick{[2]Point{{0, 0, 0}, {1, 0, 0}}},
			B:        Brick{[2]Point{{1, 0, 0}, {2, 0, 0}}},
			overlaps: true,
		},
		{
			A:        Brick{[2]Point{{0, 0, 0}, {1, 0, 0}}},
			B:        Brick{[2]Point{{0, 1, 0}, {1, 1, 0}}},
			overlaps: false,
		},
		{
			A:        Brick{[2]Point{{2, 4, 0}, {6, 4, 0}}},
			B:        Brick{[2]Point{{5, 3, 0}, {5, 5, 0}}},
			overlaps: true,
		},
	}

	for _, test := range tests {
		if test.A.Overlaps(test.B) != test.overlaps {
			t.Errorf("%s and %s should have overlap == %v", test.A, test.B, test.overlaps)
		}
	}
}

func TestDay22(t *testing.T) {
	input := `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

	result := Day22(input)

	if result != 5 {
		t.Errorf("expected 5 got %d", result)
	}
}
