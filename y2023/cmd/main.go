package main

import (
	"fmt"
	"os"
	"strconv"

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
		{y2023.Day10, y2023.Day10Part2},
		{}, //{y2023.Day11, y2023.Day11Part2}, // slow
		{y2023.Day12, y2023.Day12Part2},
		{y2023.Day13, y2023.Day13Part2},
		{y2023.Day14, y2023.Day14Part2},
		{y2023.Day15, y2023.Day15Part2},
		{y2023.Day16, y2023.Day16Part2},
		{}, // {y2023.Day17},
		{},
		{},
		{y2023.Day20, y2023.Day20Part2},
	}
)

func main() {
	var day int
	if len(os.Args) > 1 {
		var err error
		day, err = strconv.Atoi(os.Args[1])
		if err != nil {
			day = 0
		}
	}

	for i, fs := range days {
		if day != 0 && i+1 != day {
			continue
		}

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
