package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

type Color = string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

type Game struct {
	id     int
	rounds []Round
}

func (g *Game) Possible() bool {
	for _, r := range g.rounds {
		if !r.Possible() {
			return false
		}
	}
	return true
}

func (g *Game) Power() int {
	highestNumSeen := map[Color]int{
		Red:   0,
		Green: 0,
		Blue:  0,
	}
	for _, r := range g.rounds {
		for k, v := range r.shown {
			if v > highestNumSeen[k] {
				highestNumSeen[k] = v
			}
		}
	}

	return highestNumSeen[Red] * highestNumSeen[Green] * highestNumSeen[Blue]
}

type Round struct {
	shown map[Color]int
}

func (r *Round) Possible() bool {
	return r.shown[Red] <= 12 && r.shown[Green] <= 13 && r.shown[Blue] <= 14
}

func Day2(input string) (result int) {
	games := parseGames(input)

	for _, g := range games {
		if g.Possible() {
			result += g.id
		}
	}

	return
}

func Day2Part2(input string) (result int) {
	games := parseGames(input)

	for _, g := range games {
		result += g.Power()
	}

	return
}

var (
	idRegexp      = regexp.MustCompile("[0-9]+$")
	segmentRegexp = regexp.MustCompile("[0-9]+ [a-z]+")
	colorRegexp   = regexp.MustCompile("[a-z]+")
	valueRegexp   = regexp.MustCompile("[0-9]+")
)

func parseGames(input string) (games []Game) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")
		games = append(games, Game{
			id:     Must1(strconv.Atoi(idRegexp.FindString(split[0]))),
			rounds: parseRounds(split[1]),
		})
	}

	return
}

func parseRounds(input string) (rounds []Round) {
	roundStrings := strings.Split(input, ";")
	for _, round := range roundStrings {
		shown := make(map[Color]int)

		segments := segmentRegexp.FindAllString(round, -1)
		for _, segment := range segments {
			color := colorRegexp.FindString(segment)
			value := Must1(strconv.Atoi(valueRegexp.FindString(segment)))
			shown[color] = value
		}

		rounds = append(rounds, Round{
			shown: shown,
		})
	}

	return
}
