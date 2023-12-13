package y2023

import "testing"

func TestDay12(t *testing.T) {
	input := `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

	result := Day12(input)

	if result != 21 {
		t.Errorf("expected 21 got %d", result)
	}

	result = Day12Part2(input)

	if result != 525152 {
		t.Errorf("expected 525152 got %d", result)
	}
}
