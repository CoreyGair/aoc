package y2023

import (
	"strings"
)

type Sequence []int

func (s Sequence) IsZero() bool {
	for _, x := range s {
		if x != 0 {
			return false
		}
	}
	return true
}

func (s Sequence) CalculateDifferences() (result Sequence) {
	result = make(Sequence, len(s)-1)
	for i := 0; i < len(s)-1; i++ {
		result[i] = s[i+1] - s[i]
	}
	return
}

func Extrapolate(sequences []Sequence) (result int) {
	for _, s := range sequences {
		diffs := []Sequence{s}
		for !diffs[len(diffs)-1].IsZero() {
			diffs = append(diffs, diffs[len(diffs)-1].CalculateDifferences())
		}

		v := 0
		for i := len(diffs) - 2; i >= 0; i-- {
			v += diffs[i][len(diffs[i])-1]
		}

		result += v
	}
	return
}

func Day9(input string) int {
	lines := strings.Split(input, "\n")

	sequences := make([]Sequence, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		sequences = append(sequences, parseIntSlice(l))
	}

	return Extrapolate(sequences)
}

func (s Sequence) Reverse() Sequence {
	reversed := make(Sequence, len(s))
	for i, x := range s {
		reversed[len(reversed)-1-i] = x
	}
	return reversed
}

func Day9Part2(input string) int {
	lines := strings.Split(input, "\n")

	sequences := make([]Sequence, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		sequences = append(sequences, parseIntSlice(l))
	}

	revSequences := make([]Sequence, len(sequences))
	for i, s := range sequences {
		revSequences[i] = s.Reverse()
	}

	return Extrapolate(revSequences)
}
