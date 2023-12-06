package y2023

import (
	"strconv"
	"strings"
)

func Day6(input string) (result int) {
	lines := strings.Split(input, "\n")
	times := parseIntSlice(lines[0])
	dists := parseIntSlice(lines[1])

	for i := range times {
		t := times[i]
		d := dists[i]
		waysToWin := 0

		for j := 0; j < t; j++ {
			distForThisStrategy := j * (t - j)
			if distForThisStrategy > d {
				waysToWin++
			}
		}

		if result == 0 {
			result = waysToWin
		} else {
			result *= waysToWin
		}
	}
	return
}

func Day6Part2(input string) int {
	lines := strings.Split(input, "\n")
	time := Must1(strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")))
	dist := Must1(strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")))

	lowest := 0
	for i := 1; i < time; i++ {
		d := i * (time - i)
		if d > dist {
			lowest = i
			break
		}
	}

	highest := 0
	for i := time; i > 0; i-- {
		d := i * (time - i)
		if d > dist {
			highest = i
			break
		}
	}

	return highest - lowest + 1
}
