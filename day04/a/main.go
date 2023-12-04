// https://adventofcode.com/2023/day/4
// Day 4: Scratchcards
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

type Card struct {
	cardnumber     int         // Unnessesary card number
	winningNumbers map[int]int // We'll only be using the index
	numbers        map[int]int // This could be a slice, but who knows part 2
	winners        []int       // The winners on this card
	points         int         // The number of points this card scored
}

// whichWinningNumbers checks it's own numbers and stores them in winners
func (c Card) whichWinningNumbers() Card {
	for n, _ := range c.numbers {
		_, ok := c.winningNumbers[n]
		if ok {
			c.winners = append(c.winners, n)
		}
	}
	// fmt.Printf("Card %d:\nWinners: %v\nNumbers: %v \nWinners: %v\n\n", c.cardnumber, c.winningNumbers, c.numbers, c.winners)
	return c
}

// tallyPoints counts the winners and then adds the points to points
func (c Card) tallyPoints() Card {
	if len(c.winners) > 0 {
		if len(c.winners) == 1 {
			c.points = 1
		} else {
			points := 1
			for n := 1; n < len(c.winners); n++ {
				points *= 2
			}
			c.points = points
		}
	}
	// fmt.Printf("Card %d:\nWinners: %v\nNumbers: %v \nWinners: %v\nPoints: %d\n\n", c.cardnumber, c.winningNumbers, c.numbers, c.winners, c.points)
	return c
}

var Cards []Card

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initCards(input)
	fmt.Printf("Winner, Winner, Chicken-dinner with %d points!\n", countPoints(Cards))

}

func initCards(input []string) {
	// Get the input in the cards
	for n, line := range input {
		winners := make(map[int]int)
		numbers := make(map[int]int)
		step1 := strings.Split(line, ":")     // Card 1 <> 20 35 42 | 25 24 56 54
		step2 := strings.Split(step1[1], "|") // 20 35 42 <> 25 24 56 54
		step3 := strings.Split(step2[0], " ") // Winners with double spaces
		for _, number := range step3 {
			if number != "" {
				winners[libaoc.SilentAtoi(number)] = 1
			}
		}
		step4 := strings.Split(step2[1], " ") // Numbers with double spaces
		for _, number := range step4 {
			if number != "" {
				numbers[libaoc.SilentAtoi(number)] = 1
			}
		}

		card := Card{ // Build the card!
			cardnumber:     n + 1,
			winningNumbers: winners,
			numbers:        numbers,
		}
		card = card.whichWinningNumbers()
		card = card.tallyPoints()
		// fmt.Printf("Card: %#v\n", card)
		Cards = append(Cards, card)
	}
	// Init the winners and Init the points (this can be in one run)

	return
}

func countPoints(Cards []Card) int {
	var points int
	for _, card := range Cards {
		points += card.points
	}
	return points
}
