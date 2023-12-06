package y2023

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	start  int
	length int
}

func NewRange(start int, length int) Range {
	return Range{
		start:  start,
		length: length,
	}
}

func NewRangeFromEnds(start, end int) Range {
	return Range{
		start:  start,
		length: end - start,
	}
}

// End gives the number one-past-the-end of the range
func (r *Range) End() int {
	return r.start + r.length
}

func (r *Range) In(v int) bool {
	return v >= r.start && v < r.End()
}

func (r *Range) Overlaps(other Range) bool {
	return other.start < r.End() && other.End() > r.start
}

type MappedRange struct {
	source Range
	dest   Range
}

func (m *MappedRange) InSourceRange(v int) bool {
	return m.source.In(v)
}

func (m *MappedRange) Map(v int) int {
	return m.dest.start + (v - m.source.start)
}

func (m *MappedRange) OverlapsSourceRange(r Range) bool {
	return m.source.Overlaps(r)
}

// MapRange returns the mapped range, and a list of any unmapped parts (parts of r which are not in the source range)
func (m *MappedRange) MapRange(r Range) (Range, []Range) {
	unmapped := make([]Range, 0, 2)
	if r.start < m.source.start {
		unmapped = append(unmapped, NewRangeFromEnds(r.start, m.source.start))
		r = NewRangeFromEnds(m.source.start, r.End())
	}
	if r.End() > m.source.End() {
		unmapped = append(unmapped, NewRangeFromEnds(m.source.End(), r.End()))
		r = NewRangeFromEnds(r.start, m.source.End())
	}

	mappedStart := m.Map(r.start)
	mappedLast := m.Map(r.End() - 1)
	return NewRangeFromEnds(mappedStart, mappedLast+1), unmapped
}

type MapStage struct {
	ranges []MappedRange
}

func NewMapStageFromString(input string) MapStage {
	lines := strings.Split(input, "\n")[1:]

	ranges := make([]MappedRange, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}

		// [destStart,sourceStart,length]
		ints := parseIntSlice(line)

		ranges[i] = MappedRange{
			source: NewRange(ints[1], ints[2]),
			dest:   NewRange(ints[0], ints[2]),
		}
	}

	return MapStage{
		ranges: ranges,
	}
}

func (m *MapStage) Map(v int) int {
	for _, r := range m.ranges {
		if r.InSourceRange(v) {
			return r.Map(v)
		}
	}
	return v
}

func (m *MapStage) MapRange(r Range) (mappedRanges []Range) {
	rangesToMap := []Range{r}

	for len(rangesToMap) != 0 {
		newRanges := make([]Range, 0)
		for _, r := range rangesToMap {
			overlapped := false
			for _, m := range m.ranges {
				if m.OverlapsSourceRange(r) {
					overlapped = true
					mapped, unmapped := m.MapRange(r)
					mappedRanges = append(mappedRanges, mapped)
					newRanges = append(newRanges, unmapped...)
				}
			}
			if !overlapped {
				mappedRanges = append(mappedRanges, r)
			}
		}
		rangesToMap = newRanges
	}

	return
}

func Day5(input string) int {
	seedStrings := regexp.MustCompile("[0-9]+").FindAllString(strings.Split(input, "\n")[0], -1)

	seeds := make([]int, len(seedStrings))
	for i, s := range seedStrings {
		seeds[i] = Must1(strconv.Atoi(s))
	}

	mapStrings := regexp.MustCompile("map:\n(([0-9]+ )+[0-9]+(\n)?)+").FindAllString(input, -1)

	mapStages := make([]MapStage, len(mapStrings))
	for i, s := range mapStrings {
		mapStages[i] = NewMapStageFromString(s)
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

func Day5Part2(input string) int {
	seedStrings := regexp.MustCompile("[0-9]+").FindAllString(strings.Split(input, "\n")[0], -1)

	seedRanges := make([]Range, len(seedStrings)/2)
	for i := 0; i < len(seedStrings); i += 2 {
		seedRanges[i/2] = NewRange(Must1(strconv.Atoi(seedStrings[i])), Must1(strconv.Atoi(seedStrings[i+1])))
	}

	mapStrings := regexp.MustCompile("map:\n(([0-9]+ )+[0-9]+(\n)?)+").FindAllString(input, -1)

	mapStages := make([]MapStage, len(mapStrings))
	for i, s := range mapStrings {
		mapStages[i] = NewMapStageFromString(s)
	}

	lowestDestStart := math.MaxInt
	for _, seedRange := range seedRanges {
		ranges := []Range{seedRange}
		for _, stage := range mapStages {
			newRanges := make([]Range, 0)
			for _, r := range ranges {
				newRanges = append(newRanges, stage.MapRange(r)...)
			}
			ranges = newRanges
		}

		for _, r := range ranges {
			if r.start < lowestDestStart {
				lowestDestStart = r.start
			}
		}
	}
	return lowestDestStart
}
