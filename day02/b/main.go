// https://adventofcode.com/2023/day/2
// Day 2: Cube Conundrum
package main

import (
	"fmt"
	"oac/libaoc"
	"strings"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	gamePowers := getGamePowers(input)
	fmt.Printf("Game powers:\n%+v\n\nSum: %d", gamePowers, libaoc.SumSlice(gamePowers))
}

func getGamePowers(input []string) (powers []int) {
	for _, game := range input {
		powers = append(powers, gamePower(game))
	}
	return powers
}

func gamePower(game string) int {
	highestcubes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	step1 := strings.Split(game, ":") // step1 contains game x, rgb
	// fmt.Printf("Step1: %#v\n", step1)
	step2 := strings.Split(step1[1], ";") // step 2 contains "3 blue, 4 red"; "3 green, 5 blue"
	// fmt.Printf("Step2: %#v\n", step2)
	for _, round := range step2 {
		inputs := strings.Split(round, ",") // 3 blue, 4 red
		for _, singleinput := range inputs {
			values := strings.Split(strings.TrimSpace(singleinput), " ") // 3, blue
			// fmt.Printf("Only values: %#v\n", values)
			if libaoc.SilentAtoi(values[0]) > highestcubes[values[1]] {
				highestcubes[values[1]] = libaoc.SilentAtoi(values[0])
			}
		}
	}
	// fmt.Printf("Game %s = %d\n", game, highestcubes["red"]*highestcubes["green"]*highestcubes["blue"])
	return highestcubes["red"] * highestcubes["green"] * highestcubes["blue"]
}
