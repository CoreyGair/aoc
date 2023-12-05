package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	partNumberRegex     = regexp.MustCompile("[0-9]+")
	isAllNonSymbolRegex = regexp.MustCompile("^[.0-9]+$")
)

type Location struct {
	row int
	col int
}

type PartLocation struct {
	partNum int

	row int

	// column start
	start int
	// column end
	end int
}

// IsAdjacentTo returns true if this part is adjacent to (row,col)
func (p *PartLocation) IsAdjacentTo(l Location) bool {
	return l.row >= p.row-1 && l.row <= p.row+1 && l.col >= p.start-1 && l.col <= p.end
}

type PartList []PartLocation

func (l *PartList) FilterActualParts(lines []string) (result PartList) {
	for _, candidatePart := range *l {
		if candidatePart.row > 0 {
			// search row above
			searchRow := lines[candidatePart.row-1]
			search := searchRow[max(0, candidatePart.start-1):min(len(searchRow)-1, candidatePart.end+1)]
			if !isAllNonSymbolRegex.MatchString(search) {
				result = append(result, candidatePart)
				continue
			}
		}

		{
			// search row of
			searchRow := lines[candidatePart.row]
			search := searchRow[max(0, candidatePart.start-1):min(len(searchRow)-1, candidatePart.end+1)]
			if !isAllNonSymbolRegex.MatchString(search) {
				result = append(result, candidatePart)
				continue
			}
		}

		if candidatePart.row < len(lines)-1 {
			// search row below
			searchRow := lines[candidatePart.row+1]
			search := searchRow[max(0, candidatePart.start-1):min(len(searchRow)-1, candidatePart.end+1)]
			if !isAllNonSymbolRegex.MatchString(search) {
				result = append(result, candidatePart)
				continue
			}
		}
	}
	return
}

func Day3(input string) (result int) {
	lines := strings.Split(input, "\n")

	candidateParts := make(PartList, 0)
	for row, line := range lines {
		partNumberLocations := partNumberRegex.FindAllStringIndex(line, -1)
		for _, partNumberLocation := range partNumberLocations {
			candidateParts = append(candidateParts, PartLocation{
				partNum: Must1(strconv.Atoi(line[partNumberLocation[0]:partNumberLocation[1]])),
				row:     row,
				start:   partNumberLocation[0],
				end:     partNumberLocation[1],
			})
		}
	}

	parts := candidateParts.FilterActualParts(lines)
	for _, p := range parts {
		result += p.partNum
	}

	return
}

func Day3Part2(input string) (result int) {
	lines := strings.Split(input, "\n")

	var parts []PartLocation
	{
		candidateParts := make(PartList, 0)
		for row, line := range lines {
			partNumberLocations := partNumberRegex.FindAllStringIndex(line, -1)
			for _, partNumberLocation := range partNumberLocations {
				candidateParts = append(candidateParts, PartLocation{
					partNum: Must1(strconv.Atoi(line[partNumberLocation[0]:partNumberLocation[1]])),
					row:     row,
					start:   partNumberLocation[0],
					end:     partNumberLocation[1],
				})
			}
		}

		parts = candidateParts.FilterActualParts(lines)
	}

	var candidateGears []Location
	for row, line := range lines {
		gearLocations := regexp.MustCompile("\\*").FindAllStringIndex(line, -1)
		for _, g := range gearLocations {
			candidateGears = append(candidateGears, Location{
				row: row,
				col: g[0],
			})
		}
	}

	for _, g := range candidateGears {
		var adjacentParts PartList
		for _, p := range parts {
			if p.IsAdjacentTo(g) {
				adjacentParts = append(adjacentParts, p)
			}
		}
		if len(adjacentParts) != 2 {
			continue
		}
		result += adjacentParts[0].partNum * adjacentParts[1].partNum
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
