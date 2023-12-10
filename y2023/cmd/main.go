package main

import (
	"fmt"

	"github.com/CoreyGair/aoc/y2023"
)

type AOCSolution = func(string) int

var (
	days = [][]AOCSolution{
		{y2023.Day1},
		{y2023.Day2, y2023.Day2Part2},
		{y2023.Day3, y2023.Day3Part2},
		{y2023.Day4, y2023.Day4Part2},
		{y2023.Day5, y2023.Day5Part2},
		{y2023.Day6, y2023.Day6Part2},
		{y2023.Day7},
		{y2023.Day8, y2023.Day8Part2},
		{y2023.Day9, y2023.Day9Part2},
	}
)

func main() {
	for i, fs := range days {
		if len(fs) == 0 {
			continue
		}

		day := i + 1

		file := fmt.Sprintf("/home/corey/Documents/aoc-input/day%d.txt", day)
		input := y2023.ReadFromFile(file)

		var output []int
		for _, f := range fs {
			output = append(output, f(input))
		}
		fmt.Printf("Day %d: %d\n", day, output)
	}
}
