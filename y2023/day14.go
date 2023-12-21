package y2023

import (
	"fmt"
	"strings"
)

func Day14(input string) (result int) {
	platform := strings.Split(input, "\n")
	platform = FilterNewlines(platform)

	platform = RollRocksNorth(platform)

	rocks := LocateRocks(platform)

	for _, r := range rocks {
		result += len(platform) - r.row
	}
	return
}

func Day14Part2(input string) (result int) {
	platform := strings.Split(input, "\n")
	platform = FilterNewlines(platform)

	rocksHistory := [][]Location{LocateRocks(platform)}

	loopStart := 0
	loopLength := 0
	for i := 1; i < 10000; i++ {
		for j := 0; j < 4; j++ {
			platform = RollRocksNorth(platform)
			platform = RotateClockwise(platform)
		}

		newRocks := LocateRocks(platform)

		for k, oldRocks := range rocksHistory {
			if LocationListsEqual(newRocks, oldRocks) {
				loopStart = k
				loopLength = (i - k)
				goto end
			}
		}

		rocksHistory = append(rocksHistory, newRocks)
	}
end:
	lastIndex := loopStart + ((1000000000 - loopStart) % loopLength)

	lastRocks := rocksHistory[lastIndex]

	for _, r := range lastRocks {
		result += len(platform) - r.row
	}
	return
}

func FilterNewlines(strs []string) (result []string) {
	result = make([]string, 0, len(strs))
	for _, s := range strs {
		if s == "" {
			continue
		}
		result = append(result, s)
	}
	return
}

func RollRocksNorth(platform []string) []string {
	newPlatform := make([]string, 0, len(platform))
	for _, row := range platform {
		newPlatform = append(newPlatform, strings.Clone(row))
	}

	for i, row := range platform {
		for j, letter := range row {
			if letter == 'O' {
				rolled := false
				for k := i - 1; k >= 0; k-- {
					newLetter := []rune(newPlatform[k])[j]
					if newLetter != '.' {
						rolled = true
						newPlatform[i] = replaceAt(newPlatform[i], j, '.')
						newPlatform[k+1] = replaceAt(newPlatform[k+1], j, 'O')
						break
					}
				}
				if !rolled {
					newPlatform[i] = replaceAt(newPlatform[i], j, '.')
					newPlatform[0] = replaceAt(newPlatform[0], j, 'O')
				}
			}
		}
	}

	return newPlatform
}

func RotateClockwise(platform []string) (newPlatform []string) {
	newPlatform = make([]string, len(platform[0]))

	for _, row := range platform {
		for j, letter := range row {
			newPlatform[j] = string(letter) + newPlatform[j]
		}
	}
	return
}

func replaceAtNoCopy(s string, i int, r rune) string {
	out := []rune(s)
	out[i] = r
	return string(out)
}

func LocateRocks(platform []string) (result []Location) {
	for i, row := range platform {
		for j, letter := range row {
			if letter == 'O' {
				result = append(result, Location{
					row: i,
					col: j,
				})
			}
		}
	}
	return
}

func LocationListsEqual(l1, l2 []Location) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i, x := range l1 {
		if !l2[i].Equals(x) {
			return false
		}
	}
	return true
}

func (l Location) Equals(other Location) bool {
	return l.row == other.row && l.col == other.col
}

func PrintStringList(xs []string) {
	for _, x := range xs {
		fmt.Println(x)
	}
}
