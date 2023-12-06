// https://adventofcode.com/2023/day/6
// Day 6: Wait For It
package main

import (
	"fmt"
	"log"
	"oac/libaoc"
	"strings"
	"time"
)

type RaceRounds struct {
	time           []int
	dist           []int
	winningOptions []int
}
type PossibleRatios struct {
	button     []int
	traveltime []int
}

func (r RaceRounds) getWinningOptions() {
	for i := 0; i < len(r.time); i++ { // Loops all races
		var wins int
		goldilocks := false // The goldilocks zone is where we are winning and having options
		for n := 1; n < r.time[i]; n++ {
			if r.dist[i] >= n*(r.time[i]-n) && goldilocks { // Goldlocks was set, but this was no winner. Passed the goldilocks.
				break
			}
			if r.dist[i] < n*(r.time[i]-n) {
				goldilocks = true // We started winning and entered the goldilocks zone.
				wins++
			}
		}
		raceRounds.winningOptions = append(raceRounds.winningOptions, wins)
	}
	return
}

var (
	raceRounds RaceRounds
)

func main() {
	start := time.Now()
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}

	// Do stuff
	readRounds(input)
	raceRounds.getWinningOptions()
	fmt.Printf("%#v\n%#v\n%#v\n\n", raceRounds.time, raceRounds.dist, raceRounds.winningOptions)
	fmt.Printf("Options multiplied: %d\n\n", multiplyOptions())
	// Print time taken
	elapsed := time.Since(start)
	log.Printf("Program took %s", elapsed)
}

func readRounds(input []string) {
	var timesBuilding string
	step1times := strings.Split(input[0], " ")
	for i := 1; i < len(step1times); i++ {
		if step1times[i] != "" {
			timesBuilding = timesBuilding + step1times[i]
		}
	}
	raceRounds.time = append(raceRounds.time, libaoc.SilentAtoi(timesBuilding))

	var distBuilding string
	step1dist := strings.Split(input[1], " ")
	for i := 1; i < len(step1dist); i++ {
		if step1dist[i] != "" {
			distBuilding = distBuilding + step1dist[i]
		}
	}
	raceRounds.dist = append(raceRounds.dist, libaoc.SilentAtoi(distBuilding))
	fmt.Printf("%#v\n", raceRounds)
}

func multiplyOptions() (total int) {
	for _, n := range raceRounds.winningOptions {
		if total < 1 {
			total = n
		} else {
			total *= n
		}
	}
	return total
}
