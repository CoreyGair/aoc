package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	wordDigitMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
)

func Day1(input string) (result int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// digits := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine").FindString(line)
		// if len(digits) == 0 {
		// 	continue
		// }

		first := regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine").FindString(line)
		if first == "" {
			continue
		}
		var v int
		var ok bool
		if v, ok = wordDigitMap[first]; !ok {
			v = Must1(strconv.Atoi(first))
		}
		result += 10 * v

		// work around go non-overlapping regex matches...
		var last string
		{
			for i := range line {
				l := line[len(line)-i-1:]
				last = regexp.MustCompile("[0-9]|one|two|three|four|five|six|seven|eight|nine").FindString(l)
				if last != "" {
					break
				}
			}
		}
		if v, ok = wordDigitMap[last]; !ok {
			v = Must1(strconv.Atoi(last))
		}
		result += v
	}

	return
}
