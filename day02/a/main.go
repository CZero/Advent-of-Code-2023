// https://adventofcode.com/2023/day/2
// Day 2: Cube Conundrum
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

var Gamelimit = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	validGames := gamesWithinLimits(input)
	fmt.Printf("Valid games:\n%+v\n\nSum: %d", validGames, libaoc.SumSlice(validGames))
}

func gamesWithinLimits(input []string) (possibleGames []int) {
	for n, game := range input {
		if gameWithinLimits(game) {
			possibleGames = append(possibleGames, n+1)
		}
	}
	return possibleGames
}

func gameWithinLimits(game string) bool {
	step1 := strings.Split(game, ":") // step1 contains game x, rgb
	// fmt.Printf("Step1: %#v\n", step1)
	step2 := strings.Split(step1[1], ";") // step 2 contains "3 blue, 4 red"; "3 green, 5 blue"
	// fmt.Printf("Step2: %#v\n", step2)
	for _, round := range step2 {
		inputs := strings.Split(round, ",") // 3 blue, 4 red
		for _, singleinput := range inputs {
			values := strings.Split(strings.TrimSpace(singleinput), " ") // 3, blue
			// fmt.Printf("Only values: %#v\n", values)
			if Gamelimit[values[1]] < libaoc.SilentAtoi(values[0]) {
				return false
			}
		}
	}
	return true
}
