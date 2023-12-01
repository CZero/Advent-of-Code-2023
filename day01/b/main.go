// https://adventofcode.com/2023/day/1
// Day 1: Trebuchet?!
package main

import (
	"fmt"
	"oac/libaoc"
	"strconv"
)

func main() {
	// input, err := libaoc.ReadLines("example.txt")
	input, err := libaoc.ReadLines("input.txt")
	if err != nil {
		panic("No input!")
	}
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
	for n, p := range input {
		_, errFirst := strconv.Atoi(string(p))
		if errFirst == nil {
			return string(p)
		} else {
			written := checkForWrittenNumber(input, n)
			if written != "" {
				return written
			}
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
		} else {
			written := checkForWrittenNumber(input, i)
			if written != "" {
				return written
			}
		}
	}
	return "0"
}

func checkForWrittenNumber(input string, pos int) string {
	// fmt.Println(input, pos)
	switch string(input[pos]) {
	case "o": // one?
		if len(input)-1 < pos+2 {
			return ""
		}
		if string(input[pos+1]) == "n" {
			if string(input[pos+2]) == "e" {
				return "1"
			}
		}
	case "t": // two / three ?
		if len(input)-1 < pos+2 { // 8-1 (7) < 6+2
			return ""
		}
		if string(input[pos+1]) == "w" {
			if string(input[pos+2]) == "o" {
				return "2"
			}
		}
		if len(input)-1 < pos+4 {
			return ""
		}
		if string(input[pos+1]) == "h" {
			if string(input[pos+2]) == "r" {
				if string(input[pos+3]) == "e" {
					if string(input[pos+4]) == "e" {
						return "3"
					}
				}
			}
		}
	case "f": // four / five ?
		if len(input)-1 < pos+3 {
			return ""
		}
		if string(input[pos+1]) == "o" {
			if string(input[pos+2]) == "u" {
				if string(input[pos+3]) == "r" {
					return "4"
				}
			}
		}
		if string(input[pos+1]) == "i" {
			if string(input[pos+2]) == "v" {
				if string(input[pos+3]) == "e" {
					return "5"
				}
			}
		}
	case "s": // six / seven ?
		if len(input)-1 < pos+2 {
			return ""
		}
		if string(input[pos+1]) == "i" {
			if string(input[pos+2]) == "x" {
				return "6"
			}
		}
		if len(input)-1 < pos+4 {
			return ""
		}
		if string(input[pos+1]) == "e" {
			if string(input[pos+2]) == "v" {
				if string(input[pos+3]) == "e" {
					if string(input[pos+4]) == "n" {
						return "7"
					}
				}
			}
		}
	case "e": // eight ?
		if len(input)-1 < pos+4 {
			return ""
		}
		if string(input[pos+1]) == "i" {
			if string(input[pos+2]) == "g" {
				if string(input[pos+3]) == "h" {
					if string(input[pos+4]) == "t" {
						return "8"
					}
				}
			}
		}
	case "n": // nine?
		if len(input)-1 < pos+3 {
			return ""
		}
		if string(input[pos+1]) == "i" {
			if string(input[pos+2]) == "n" {
				if string(input[pos+3]) == "e" {
					return "9"
				}
			}
		}
	}
	return ""
}
