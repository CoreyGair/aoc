package main

import (
	"fmt"

	"github.com/CoreyGair/aoc/y2023"
)

var (
	days = [][]func(string) int{
		[y2023.Day1],
		[y2023.Day2,y2023.Day2Part2],
		[y2023.Day3],
		[y2023.Day4],
		[y2023.Day5],
	}
)

func main() {
	for i, f := range days {
		day := i + 1

		file := fmt.Sprintf("/home/corey/Documents/aoc-input/day%d.txt", day)
		input := y2023.ReadFromFile(file)
		output := f(input)

		fmt.Printf("Day %d: %d\n", day, output)
	}
}
