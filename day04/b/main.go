// https://adventofcode.com/2023/day/4
// Day 4: Scratchcards
package main

import (
	"fmt"
	"oac/libaoc"
	"strings"
)

type Card struct {
	cardnumber     int         // Unnessesary card number
	winningNumbers map[int]int // We'll only be using the index
	numbers        map[int]int // This could be a slice, but who knows part 2
	winners        []int       // The winners on this card
	copies         int         // How many copies?
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

// procesWinnings handles the winning and card copying
func (c Card) procesWinnings() {
	if len(c.winners) > 0 {
		copyCard := c.cardnumber
		copies := c.copies
		for n := 0; n <= len(c.winners)-1; n++ {
			Cards[copyCard+n].copies += copies
			// fmt.Printf("Card %d heeft nu %d copies\n", copyCard, Cards[copyCard+n].copies)
		}
	}
}

var Cards []Card

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initCards(input)
	// for n := 0; n < len(Cards); n++ {
	// 	fmt.Printf("%d - %#v - Copies: %d\n", Cards[n].cardnumber, Cards[n].winners, Cards[n].copies)
	// }
	fmt.Printf("There is a total of %d cards!\n", countCards())

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
			copies:         1,
		}
		// fmt.Printf("Card: %#v\n", card)
		Cards = append(Cards, card)
	}
	// Init the winners
	for n, card := range Cards {
		Cards[n] = card.whichWinningNumbers()
	}

	for n := 0; n < len(Cards); n++ { // Volgorde telt, maps zijn random!
		Cards[n].procesWinnings()
	}
	// card = card.tallyPoints()

	return
}

func countCards() (totalCards int) {
	for _, card := range Cards {
		totalCards += card.copies
	}
	return totalCards
}
