package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	partNumberRegex = regexp.MustCompile("[0-9]+")
	nonSymbolRegex  = regexp.MustCompile("[.0-9]")
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

func (p *PartLocation) GetLocalSearchLocations() (result []Location) {
	for i := p.start - 1; i <= p.end+1; i++ {
		for j := -1; j <= 1; j++ {
			result = append(result, Location{
				row: p.row + j,
				col: i,
			})
		}
	}

	return
}

func Day3(input string) (result int) {
	lines := strings.Split(input, "\n")

	candidateParts := make([]PartLocation, 0)
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

	for _, candicandidatePart := range candidateParts {
		searchLocations := candicandidatePart.GetLocalSearchLocations()
		for _, loc := range searchLocations {
			if loc.row < 0 || loc.row > len(lines)-1 {
				continue
			}
			row := lines[loc.row]

			if loc.col < 0 || loc.col > len(row)-1 {
				continue
			}
			symbol := []rune(row)[loc.col]

			if !nonSymbolRegex.MatchString(string(symbol)) {
				result += candicandidatePart.partNum
				break
			}
		}
	}

	return
}
