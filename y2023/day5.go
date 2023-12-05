package y2023

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type MapRange struct {
	sourceStart int
	sourceEnd   int
	destStart   int
}

func (m *MapRange) IsInSource(sourceValue int) bool {
	return sourceValue >= m.sourceStart && sourceValue <= m.sourceEnd
}

func (m *MapRange) Map(sourceValue int) int {
	dist := sourceValue - m.sourceStart
	return m.destStart + dist
}

type MapStage struct {
	ranges []MapRange
}

func (m *MapStage) Map(sourceValue int) int {
	for _, r := range m.ranges {
		if r.IsInSource(sourceValue) {
			return r.Map(sourceValue)
		}
	}
	return sourceValue
}

func Day5(input string) int {
	seedStrings := regexp.MustCompile("[0-9]+").FindAllString(strings.Split(input, "\n")[0], -1)

	seeds := make([]int, len(seedStrings))
	for i, s := range seedStrings {
		seeds[i] = Must1(strconv.Atoi(s))
	}

	mapStrings := regexp.MustCompile("map:\n(([0-9]+ )+[0-9]+\n)+").FindAllString(input, -1)

	mapStages := make([]MapStage, len(mapStrings))
	for i, s := range mapStrings {
		lines := strings.Split(s, "\n")[1:]

		ranges := make([]MapRange, len(lines))
		for i, line := range lines {
			if line == "" {
				continue
			}
			ints := parseIntSlice(line)
			ranges[i] = MapRange{
				sourceStart: ints[1],
				sourceEnd:   ints[1] + ints[2],
				destStart:   ints[0],
			}
		}

		mapStages[i] = MapStage{
			ranges: ranges,
		}
	}

	lowest := math.MaxInt
	for _, seed := range seeds {
		x := seed

		steps := []int{seed}
		for _, stage := range mapStages {
			x = stage.Map(x)
			steps = append(steps, x)
		}
		if x < lowest {
			lowest = x
		}
	}
	return lowest
}
