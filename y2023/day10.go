package y2023

import (
	"slices"
	"strings"
)

type Cardinal string

const (
	North Cardinal = "n"
	South Cardinal = "s"
	East  Cardinal = "e"
	West  Cardinal = "w"
)

var (
	cardinalReverse = map[Cardinal]Cardinal{
		North: South,
		South: North,

		East: West,
		West: East,
	}
)

func (c Cardinal) Reverse() Cardinal {
	return cardinalReverse[c]
}

func (c Cardinal) Move(start Location) Location {
	switch c {
	case North:
		start.row--
	case South:
		start.row++
	case East:
		start.col++
	case West:
		start.col--
	}
	return start
}

type PipeDirections []Cardinal

var (
	runeToPipeDirections = map[rune]PipeDirections{
		'|': {North, South},
		'-': {East, West},
		'J': {North, West},
		'L': {North, East},
		'F': {East, South},
		'7': {West, South},
		'.': {},
	}
)

func (pd PipeDirections) Contains(other Cardinal) bool {
	return slices.Contains(pd, other)
}

func LookupLocation(location Location, rows []string) (rune, bool) {
	if location.row < 0 || location.row >= len(rows) {
		return 0, false
	}
	row := rows[location.row]

	if location.col < 0 || location.col >= len(row) {
		return 0, false
	}
	return []rune(row)[location.col], true
}

func Day10(input string) int {
	rows := strings.Split(input, "\n")

	var start Location
	for i, r := range rows {
		for j, letter := range r {
			if letter == 'S' {
				start = Location{
					row: i,
					col: j,
				}
				goto done
			}
		}
	}

done:
	for _, startDir := range []Cardinal{North, East, South, West} {
		currLocation := startDir.Move(start)
		prevDirection := startDir
		loopLength := 0
		foundLoop := false
		for {
			if currLetter, ok := LookupLocation(currLocation, rows); ok {
				loopLength++
				if currLetter == 'S' {
					foundLoop = true
					break
				}

				pipeDirs := runeToPipeDirections[currLetter]
				if !pipeDirs.Contains(prevDirection.Reverse()) {
					break
				}

				nextDir := pipeDirs[0]
				if nextDir == prevDirection.Reverse() {
					nextDir = pipeDirs[1]
				}

				currLocation = nextDir.Move(currLocation)
				prevDirection = nextDir
			} else {
				break
			}
		}

		if foundLoop {
			return loopLength / 2
		}
	}

	panic("oops")
}

func Clear(x [][]int) {
	for i := range x {
		for j := range x[i] {
			x[i][j] = -2
		}
	}
}

func Day10Part2(input string) (result int) {
	rows := strings.Split(input, "\n")

	var start Location
	for i, r := range rows {
		for j, letter := range r {
			if letter == 'S' {
				start = Location{
					row: i,
					col: j,
				}
				goto done
			}
		}
	}

done:
	indexMap := make([][]int, len(rows))
	for i := range indexMap {
		indexMap[i] = make([]int, len(rows[0]))
	}

	var lastLoopIndex int
	for _, startDir := range []Cardinal{North, East, South, West} {
		currentLoopIndex := 0

		Clear(indexMap)
		indexMap[start.row][start.col] = 0

		currLocation := startDir.Move(start)
		prevDirection := startDir
		foundLoop := false
		for {
			currentLoopIndex++
			if currLetter, ok := LookupLocation(currLocation, rows); ok {
				if currLetter == 'S' {
					foundLoop = true
					break
				}
				lastLoopIndex = currentLoopIndex

				indexMap[currLocation.row][currLocation.col] = currentLoopIndex

				pipeDirs := runeToPipeDirections[currLetter]
				if !pipeDirs.Contains(prevDirection.Reverse()) {
					break
				}

				nextDir := pipeDirs[0]
				if nextDir == prevDirection.Reverse() {
					nextDir = pipeDirs[1]
				}

				currLocation = nextDir.Move(currLocation)
				prevDirection = nextDir
			} else {
				break
			}
		}

		if foundLoop {
			break
		}
	}

	for i, row := range rows {
		crosings := 0
		for j := range row {
			loopIndex := indexMap[i][j]

			if loopIndex < 0 {
				if crosings != 0 {
					result++
					indexMap[i][j] = -5
				}
				continue
			}

			if i < len(rows)-1 {
				loopIndexUnder := indexMap[i+1][j]

				if AbsInt(loopIndexUnder-loopIndex) == 1 {
					if loopIndexUnder > loopIndex {
						crosings++
					} else {
						crosings--
					}
				} else if AbsInt(loopIndexUnder-loopIndex) == lastLoopIndex {
					if loopIndexUnder < loopIndex {
						crosings++
					} else {
						crosings--
					}
				}
			}
		}
	}

	return
}
