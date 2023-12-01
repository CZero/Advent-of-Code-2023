// https://adventofcode.com/2023/day/1
// Day 1: Trebuchet?!
package main

import (
	"aoc/libaoc"
	"fmt"
	"strconv"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	// fmt.Println(input)
	sum := libaoc.SumSlice(getNumbers(input))
	fmt.Printf("Uitkomst: %d\n", sum)
}

func getNumbers(input []string) (numbers []int) {
	var allNumbers []int
	for _, row := range input {
		num1 := first(row)
		num2 := last(row)
		// fmt.Printf("Num1: %s, Num2: %s\n", num1, num2)
		number := num1 + num2
		allNumbers = append(allNumbers, libaoc.SilentAtoi(number))
	}
	return allNumbers
}
func first(input string) string {
	for _, p := range input {
		_, errFirst := strconv.Atoi(string(p))
		if errFirst == nil {
			return string(p)
		}
	}
	return "0"
}
func last(input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		p := input[i]
		_, errFirst := strconv.Atoi(string(p))
		if errFirst == nil {
			return string(p)
		}
	}
	return "0"
}
