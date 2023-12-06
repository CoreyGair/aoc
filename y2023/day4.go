package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	num  int
	wins []int
	ours []int
}

func (c *Card) CountWins() (result int) {
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
	if c.CountWins() == 0 {
		return 0
	}
	return 1 << (c.CountWins() - 1)
}

func Day4(input string) (result int) {
	var cards []Card
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(strings.Split(line, ":")[1], "|")
		cards = append(cards, Card{
			num:  Must1(strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(strings.Split(line, ":")[0]))) - 1,
			wins: parseIntSlice(split[0]),
			ours: parseIntSlice(split[1]),
		})
	}

	for _, card := range cards {
		result += card.GetValue()
	}

	return
}

func Day4Part2(input string) (result int) {
	var cards []Card
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(strings.Split(line, ":")[1], "|")
		cards = append(cards, Card{
			wins: parseIntSlice(split[0]),
			ours: parseIntSlice(split[1]),
		})
	}

	// memoise total number of cards from playing card [i] (incl. copies and their copies...)
	cumulativeNumberOfCards := make([]int, len(cards))
	// last one cannot make any copies
	cumulativeNumberOfCards[len(cumulativeNumberOfCards)-1] = 1
	result = 1

	for i := len(cards) - 2; i >= 0; i-- {
		// count this card
		cumulativeNumberOfCards[i] = 1

		// count cards from won copies
		numCards := cards[i].CountWins()
		for j := i + 1; j <= i+numCards; j++ {
			cumulativeNumberOfCards[i] += cumulativeNumberOfCards[j]
		}

		result += cumulativeNumberOfCards[i]
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
