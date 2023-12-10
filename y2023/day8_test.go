package y2023

import "testing"

func TestDay8(t *testing.T) {
	input := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	result := Day8(input)

	if result != 2 {
		t.Errorf("expected 2 got %d", result)
	}

	input = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

	result = Day8Part2(input)

	if result != 6 {
		t.Errorf("expected 6 got %d", result)
	}

}
