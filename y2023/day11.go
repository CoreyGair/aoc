package y2023

import (
	"slices"
	"strings"
)

const (
	galaxy = '#'
)

func ExpandEmptyRowsAndCols(rows []string, times int) []string {
	nonEmptyCols := make([]int, 0, len(rows[0])*len(rows))
	emptyRows := make([]int, 0, len(rows))

	for rowIdx, row := range rows {
		rowEmpty := true
		for colIdx, letter := range row {
			if letter == galaxy {
				rowEmpty = false
				nonEmptyCols = append(nonEmptyCols, colIdx)
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, rowIdx)
		}
	}

	emptyCols := make([]int, 0, len(rows[0]))
	for i := range rows[0] {
		if !slices.Contains(nonEmptyCols, i) {
			emptyCols = append(emptyCols, i)
		}
	}

	expandedRows := make([]string, 0, len(rows)+len(emptyRows))
	for rowIdx, row := range rows {
		if slices.Contains(emptyRows, rowIdx) {
			for i := 0; i < times; i++ {
				expandedRows = append(expandedRows, row)
			}
			continue
		}

		newRow := make([]rune, 0, len(row)+len(emptyCols))
		for colIdx, letter := range row {
			if slices.Contains(emptyCols, colIdx) {
				for i := 0; i < times-1; i++ {
					newRow = append(newRow, letter)
				}
			}
			newRow = append(newRow, letter)
		}
		expandedRows = append(expandedRows, string(newRow))
	}

	return expandedRows
}

func LocateGalaxies(rows []string) (result []Location) {
	for rowIdx, row := range rows {
		for colIdx, letter := range row {
			if letter == galaxy {
				result = append(result, Location{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}
	return
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (l Location) ManhattanDistance(other Location) (result int) {
	return AbsInt(l.col-other.col) + AbsInt(l.row-other.row)
}

func Day11(input string) (result int) {
	rows := strings.Split(input, "\n")

	expandedRows := ExpandEmptyRowsAndCols(rows, 2)

	galaxies := LocateGalaxies(expandedRows)

	for i, g1 := range galaxies {
		if i == len(galaxies)-1 {
			continue
		}
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]

			result += g1.ManhattanDistance(g2)
		}
	}

	return
}

func Day11Part2(input string) (result int) {
	rows := strings.Split(input, "\n")

	expandedRows := ExpandEmptyRowsAndCols(rows, 1000000)

	galaxies := LocateGalaxies(expandedRows)

	for i, g1 := range galaxies {
		if i == len(galaxies)-1 {
			continue
		}
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]

			result += g1.ManhattanDistance(g2)
		}
	}

	return
}
