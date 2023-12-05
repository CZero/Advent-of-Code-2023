// https://adventofcode.com/2023/day/5
// Day 5: If You Give A Seed A Fertilizer
package main

import (
	"aoc/libaoc"
	"fmt"
	"strings"
)

type SourceDestinationMap [3]int                             // Contains: Destination, Source, Length
type SourceDestinationMaps map[string][]SourceDestinationMap // Contains: mapsnames with their rangemaps

// addRange adds ranges to a map
func (s SourceDestinationMaps) addRange(destination, source, length int, sdmap string) {
	thisMap := SourceDestinationMap{destination, source, length}
	s[sdmap] = append(s[sdmap], thisMap)
}

// Resolve the value
func (s SourceDestinationMaps) resolveValue(source int, mapName string) int {
	for _, mapRange := range s[mapName] {
		if source >= mapRange[1] && source <= mapRange[1]+mapRange[2] {
			posInRange := source - mapRange[1]
			source = mapRange[0] + posInRange
			return source
		}
	}
	return source
}

var (
	sourceDestinationMaps = SourceDestinationMaps{ // Add the map-entries and give them the maps
		"SeedToSoil":            []SourceDestinationMap{},
		"SoilToFertilizer":      []SourceDestinationMap{},
		"FertilizerToWater":     []SourceDestinationMap{},
		"WaterToLight":          []SourceDestinationMap{},
		"LightToTemperature":    []SourceDestinationMap{},
		"TemperatureToHumidity": []SourceDestinationMap{},
		"HumidityToLocation":    []SourceDestinationMap{},
	}

	translateNames = map[string]string{ // A dictionary so we can translate the names of the maps from the input
		"seed-to-soil":            "SeedToSoil",
		"soil-to-fertilizer":      "SoilToFertilizer",
		"fertilizer-to-water":     "FertilizerToWater",
		"water-to-light":          "WaterToLight",
		"light-to-temperature":    "LightToTemperature",
		"temperature-to-humidity": "TemperatureToHumidity",
		"humidity-to-location":    "HumidityToLocation",
	}

	seeds []int // The initial seeds
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	initRanges(input)

	lowestLocation, lowestSeed := findLowestLocation()
	fmt.Printf("The lowest location = %d, with seed %d\n", lowestLocation, lowestSeed)
}

func initRanges(input []string) {
	// First we handle the first line: Seeds
	step1Seeds := strings.Split(input[0], " ")
	for i := 1; i < len(step1Seeds); i++ {
		seeds = append(seeds, libaoc.SilentAtoi(step1Seeds[i]))
	}
	fmt.Printf("Seeds: %v\n", seeds)

	// Now come the maps
	newmap := true
	mapname := ""
	for i := 2; i < len(input); i++ {
		if newmap == true { // This will be the start of a map, first naming it
			step1Newmap := strings.Split(input[i], " ")
			mapname = step1Newmap[0]
			newmap = false
			continue
		}
		if input[i] == "" { // Empty line, new map coming up
			newmap = true
			continue
		}

		// If we got here, it's a range
		step1Range := strings.Split(input[i], " ")
		dest, src, length := libaoc.SilentAtoi(step1Range[0]), libaoc.SilentAtoi(step1Range[1]), libaoc.SilentAtoi(step1Range[2])
		sourceDestinationMaps.addRange(dest, src, length, translateNames[mapname])
	}
}

// resolveSeed returns the location of a seed
func resolveSeed(processValue int) int {
	doThese := []string{"SeedToSoil", "SoilToFertilizer", "FertilizerToWater", "WaterToLight", "LightToTemperature", "TemperatureToHumidity", "HumidityToLocation"}
	for _, thisMap := range doThese {
		processValue = sourceDestinationMaps.resolveValue(processValue, thisMap)
	}
	return processValue
}

func findLowestLocation() (lowestloc, lowestseed int) {
	lowestseed = seeds[0]
	lowestloc = resolveSeed(seeds[0])
	for _, seed := range seeds {
		if resolveSeed(seed) < lowestloc {
			lowestloc = resolveSeed(seed)
			lowestseed = seed
		}
	}
	return lowestloc, lowestseed
}
