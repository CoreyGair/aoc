package y2023

import (
	"slices"
	"strconv"
	"strings"
)

type HandType = int

const (
	FiveOfAKind  HandType = 6
	FourOfAKind  HandType = 5
	FullHouse    HandType = 4
	ThreeOfAKind HandType = 3
	TwoPair      HandType = 2
	OnePair      HandType = 1
	HighCard     HandType = 0
)

func DetermineHandType(cards [5]int) HandType {
	// map[card]count
	counts := make(map[int]int, len(charToValue))
	for i := 0; i < len(charToValue); i++ {
		counts[i] = 0
	}

	jokers := 0

	for _, c := range cards {
		if c == 0 {
			jokers++
			continue
		}
		counts[c]++
	}

	// heuristic: always best for all jokers to pretend to be highest count card
	max := 0
	maxCard := -1
	for card, count := range counts {
		if count > max {
			max = count
			maxCard = card
		}
	}
	counts[maxCard] += jokers

	// map[count]num cards of that count
	numCounts := make(map[int]int)
	for _, count := range counts {
		numCounts[count]++
	}

	if numCounts[5] == 1 {
		return FiveOfAKind
	}
	if numCounts[4] == 1 {
		return FourOfAKind
	}
	if numCounts[3] == 1 && numCounts[2] == 1 {
		return FullHouse
	}
	if numCounts[3] == 1 {
		return ThreeOfAKind
	}
	if numCounts[2] == 2 {
		return TwoPair
	}
	if numCounts[2] == 1 {
		return OnePair
	}
	if numCounts[1] == 5 {
		return HighCard
	}
	panic("uh oh")
}

var (
	charToValue = map[rune]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}
)

type Hand struct {
	cards    [5]int
	handType HandType
	bid      int
}

func NewHandFromLine(line string) Hand {
	cardsString := strings.Split(line, " ")[0]
	var cards [5]int
	for i, r := range cardsString {
		cards[i] = charToValue[r]
	}

	return Hand{
		cards:    cards,
		handType: DetermineHandType(cards),
		bid:      Must1(strconv.Atoi(strings.Split(line, " ")[1])),
	}
}

func (h *Hand) Cmp(other Hand) int {
	if v := other.handType - h.handType; v != 0 {
		return v
	}

	for i := range h.cards {
		if h.cards[i] != other.cards[i] {
			return other.cards[i] - h.cards[i]
		}
	}
	panic("oops")
}

func Day7(input string) (result int) {
	lines := strings.Split(input, "\n")
	var hands []Hand
	for _, line := range lines {
		if line == "" {
			continue
		}
		hands = append(hands, NewHandFromLine(line))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return b.Cmp(a)
	})

	for i, h := range hands {
		result += h.bid * (i + 1)
	}
	return
}
