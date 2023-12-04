package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	wins []int
	ours []int
}

func (c *Card) countWins() (result int) {
	for _, ourNum := range c.ours {
		for _, winNum := range c.wins {
			if ourNum == winNum {
				result++
				break
			}
		}
	}

	return
}

func (c *Card) GetValue() int {
	if c.countWins() == 0 {
		return 0
	}
	return 1 << (c.countWins() - 1)
}

func Day4(input string) (result int) {
	var cards []Card
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		split := strings.Split(strings.Split(line, ":")[1], "|")
		cards = append(cards, Card{
			wins: parseIntSlice(split[0]),
			ours: parseIntSlice(split[1]),
		})
	}

	for _, card := range cards {
		result += card.GetValue()
	}

	return
}

func parseIntSlice(input string) (result []int) {
	matches := regexp.MustCompile("[0-9]+").FindAllString(input, -1)
	for _, match := range matches {
		result = append(result, Must1(strconv.Atoi(match)))
	}

	return
}
