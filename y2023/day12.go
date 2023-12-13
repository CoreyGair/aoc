package y2023

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func GetRowsAndGroups(input string) (rows []string, groups [][]int) {
	lines := strings.Split(input, "\n")

	rowRegex := regexp.MustCompile("[\\?.#]+")
	numberListRegex := regexp.MustCompile("([0-9]+,)*[0-9]+")

	for _, line := range lines {
		if line == "" {
			continue
		}
		rows = append(rows, rowRegex.FindString(line))
		groups = append(groups, parseIntSlice(numberListRegex.FindString(line)))
	}
	return
}

func FindUnknownGroupsInLine(line string) [][]int {
	return regexp.MustCompile("\\?+").FindAllStringIndex(line, -1)
}

var (
	countCombinationsMemo = map[string]int{}
)

func CountCombinationsEntry(row string, groups []int) int {
	countCombinationsMemo = map[string]int{}

	return CountCombinations(row, groups)
}

func CountCombinations(row string, groups []int) (result int) {
	if len(row) == 0 {
		if len(groups) > 1 || (len(groups) == 1 && groups[0] != 0) {
			return 0
		}
		return 1
	}

	hashParams := fmt.Sprintf("%s%v", row, groups)

	memo, ok := countCombinationsMemo[hashParams]
	if ok {
		return memo
	}

	spring := row[0]
	rest := row[1:]

	switch spring {
	case '.':
		if len(groups) == 0 {
			result = CountCombinations(rest, slices.Clone(groups))
			goto end
		}
		if groups[0] < 0 {
			result = 0
			goto end
		}

		if groups[0] == 0 {
			groups = groups[1:]
		}
		result = CountCombinations(rest, slices.Clone(groups))
		goto end

	case '#':
		if len(groups) == 0 || groups[0] == 0 {
			result = 0
			goto end
		}

		if groups[0] > 0 {
			groups[0] = -(groups[0] - 1)
		} else {
			groups[0]++
		}
		result = CountCombinations(rest, slices.Clone(groups))
		goto end

	case '?':
		workingRest := replaceAt(row, 0, '.')
		workingCount := CountCombinations(workingRest, slices.Clone(groups))

		brokenRest := replaceAt(row, 0, '#')
		brokenCount := CountCombinations(brokenRest, slices.Clone(groups))

		result = workingCount + brokenCount
		goto end

	default:
		panic("oops")
	}

end:
	countCombinationsMemo[hashParams] = result
	return
}

func replaceAt(s string, i int, r rune) string {
	rs := slices.Clone([]rune(s))
	rs[i] = r
	return string(rs)
}

func Day12(input string) (result int) {
	rows, groups := GetRowsAndGroups(input)

	for i, r := range rows {
		g := groups[i]

		v := CountCombinationsEntry(r, g)

		result += v
	}
	return
}

func Repeat[T any](slice []T, n int) []T {
	newSlice := make([]T, len(slice)*n)
	for i := 0; i < n; i++ {
		copy(newSlice[i*len(slice):], slice)
	}
	return newSlice
}

func Day12Part2(input string) (result int) {
	rows, groups := GetRowsAndGroups(input)

	for i, row := range rows {
		rows[i] = fmt.Sprintf("%s?%s?%s?%s?%s", row, row, row, row, row)
	}
	for i, group := range groups {
		groups[i] = Repeat(group, 5)
	}

	for i, r := range rows {
		g := groups[i]

		v := CountCombinationsEntry(r, g)

		result += v
	}
	return
}
