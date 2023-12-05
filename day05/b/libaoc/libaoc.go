package libaoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// SilentAtoi returns an int or panics
func SilentAtoi(input string) int {
	val, err := strconv.Atoi(input)
	if err != nil {
		panic(fmt.Sprintf("SilentAtoi failed with %v as input", val))
	}
	return val
}

// SumSlice returns the sum of a slice
func SumSlice(inOrder []int) (sum int) {
	for _, item := range inOrder {
		sum += item
	}
	return sum
}
