package y2023

import (
	"fmt"
	"regexp"
	"strings"
)

func Day13(input string) (result int) {
	patterns := GetAshRockPatterns(input)

	for _, pattern := range patterns {
		rows, columns := GetRowsAndColumnHashes(pattern)

		for i, c := range rows {
			if c == 0 {
				fmt.Printf("%s row%d is 0\n", pattern, i)
			}
		}

		for _, idx := range FindReflectionIndices(rows) {
			result += 100 * (idx + 1)
		}
		for _, idx := range FindReflectionIndices(columns) {
			result += idx + 1
		}
	}
	return
}

func Day13Part2(input string) (result int) {
	patterns := GetAshRockPatterns(input)

	for _, pattern := range patterns {
		rows, columns := GetRowsAndColumnHashes(pattern)

		for i, c := range rows {
			if c == 0 {
				fmt.Printf("%s row%d is 0\n", pattern, i)
			}
		}

		for _, idx := range FindSmudgedReflectionIndices(rows) {
			result += 100 * (idx + 1)
			// observed only one reflection per pattern
			continue
		}
		for _, idx := range FindSmudgedReflectionIndices(columns) {
			result += idx + 1
		}
	}
	return
}

func GetAshRockPatterns(input string) []string {
	return regexp.MustCompile("([#.]+\n?)+").FindAllString(input, -1)
}

func GetRowsAndColumnHashes(input string) (rowHashes []int, colHashes []int) {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	rowHashes = make([]int, len(lines))
	colHashes = make([]int, len(lines[0]))

	for i, line := range lines {
		for j, letter := range line {
			if letter == '#' {
				rowHashes[i] += 1 << j
				colHashes[j] += 1 << i
			}
		}
	}
	return
}

func FindReflectionIndices(hashes []int) (indices []int) {
	for i := 0; i < len(hashes)-1; i++ {
		isReflection := true
		for j := 0; j <= i; j++ {
			if i+j+1 >= len(hashes) {
				break
			}

			if hashes[i-j] != hashes[i+j+1] {
				isReflection = false
				break
			}
		}
		if isReflection {
			indices = append(indices, i)
		}
	}
	return
}

func FindSmudgedReflectionIndices(hashes []int) (indices []int) {
	for i := 0; i < len(hashes)-1; i++ {
		isReflection := true
		doneSmudge := false
		for j := 0; j <= i; j++ {
			lower := i - j
			higher := i + j + 1

			if higher >= len(hashes) {
				break
			}

			if hashes[lower] != hashes[higher] {
				diff := hashes[lower] ^ hashes[higher]
				if SingleBitSet(diff) && !doneSmudge {
					doneSmudge = true
				} else {
					isReflection = false
					break
				}
			}
		}
		if isReflection && doneSmudge {
			indices = append(indices, i)
			// observed only one reflection per pattern
			break
		}
	}
	return
}

func SingleBitSet(x int) bool {
	return (x != 0) && !((x & (x - 1)) != 0)
}
