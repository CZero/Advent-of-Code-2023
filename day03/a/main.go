// https://adventofcode.com/2023/day/3
// Day 3: Gear Ratios
package main

import (
	"aoc/libaoc"
	"fmt"
	"strconv"
)

type Coord struct {
	c int //column
	r int //row
}

// Schematic is a map of coordinates containing the engine schematics
type Schematic map[Coord]string

// isSymbol checks if a position (coord) has a Symbol (not being a number)
func (s Schematic) isSymbol(coord Coord) bool {
	_, ok := s[coord]
	if ok {
		_, err := strconv.Atoi(s[coord])
		if err != nil { // if error == nil it was a number, that's not a symbol.
			return true
		}
	}
	return false
}

// surrounded checks if there's a symbol around a position (coord)
func (s Schematic) surrounded(coord Coord) bool {
	var (
		viableCoords []Coord
		left         bool
		right        bool
		above        bool
		below        bool
	)
	if coord.c > 0 {
		left = true
	}
	if coord.c < gridwidth {
		right = true
	}
	if coord.r > 0 {
		above = true
	}
	if coord.r < gridheight {
		below = true
	}
	if left && above {
		viableCoords = append(viableCoords, Coord{coord.c - 1, coord.r - 1})
	}
	if above {
		viableCoords = append(viableCoords, Coord{coord.c, coord.r - 1})
	}
	if right && above {
		viableCoords = append(viableCoords, Coord{coord.c + 1, coord.r - 1})
	}
	if left {
		viableCoords = append(viableCoords, Coord{coord.c - 1, coord.r})
	}
	if right {
		viableCoords = append(viableCoords, Coord{coord.c + 1, coord.r})
	}
	if left && below {
		viableCoords = append(viableCoords, Coord{coord.c - 1, coord.r + 1})
	}
	if below {
		viableCoords = append(viableCoords, Coord{coord.c, coord.r + 1})
	}
	if right && below {
		viableCoords = append(viableCoords, Coord{coord.c + 1, coord.r + 1})
	}
	// fmt.Printf("Coord: %v\nViable:\n%#v\n", coord, viableCoords)
	// Now we should have a list of all viable coords. Let's check them all.
	for _, checkCoord := range viableCoords {
		if s.isSymbol(checkCoord) {
			return true
		}
	}
	return false
}

// getPartnumbers get's all valid partnumbers (numbers surrounded by a symbol)
func (s Schematic) getPartnumbers() []int {
	var (
		found      []int
		number     string
		surrounded bool
	)
	for r := 0; r <= gridheight; r++ {
		for c := 0; c <= gridwidth; c++ {
			if isNumber(s[Coord{c, r}]) { // We've found the start of a number
				number, surrounded = "", false // Reset
				for ; c <= gridwidth && isNumber(s[Coord{c, r}]); c++ {
					number = number + s[Coord{c, r}]
					if s.surrounded(Coord{c, r}) {
						surrounded = true
					}
				}
				if surrounded {
					found = append(found, libaoc.SilentAtoi(number))
				}
			}
		}
	}
	// fmt.Printf("%#v\n", found)
	return found
}

var (
	schematic  = make(Schematic)
	gridheight int
	gridwidth  int
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
	// input done

	// Setting the stage
	gridheight = len(input) - 1
	gridwidth = len(input[0]) - 1
	buildSchematic(input)

	// Do stuff
	partnumbers := schematic.getPartnumbers()
	fmt.Printf("Partnumbers are: %v\nSum = %d", partnumbers, libaoc.SumSlice(partnumbers))
}

func buildSchematic(input []string) {
	for rpos, row := range input {
		for cpos, col := range row {
			if string(col) != "." {
				schematic[Coord{cpos, rpos}] = string(col)
			}
		}
	}
}

func isNumber(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
