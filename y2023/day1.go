package y2023

import (
	"regexp"
	"strings"
)

var (
	digitRegex = regexp.MustCompile("[0-9]")
)

func Day1(input string) (result int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		for _, char := range line {
			if v := runeToInt(char); v != -1 {
				result += 10 * v
				break
			}
		}

		var last int
		for _, char := range line {
			if v := runeToInt(char); v != -1 {
				last = v
			}
		}
		result += last
	}

	return
}

func runeToInt(r rune) int {
	v := int(r - '0')

	if v < 0 || v > 9 {
		return -1
	}
	return v
}
