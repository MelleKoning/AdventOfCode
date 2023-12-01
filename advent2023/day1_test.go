package advent2023

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/MelleKoning/AdventOfCode/advent"
)

func TestExampleDay1(t *testing.T) {
	lines, err := advent.GetFileLines("input/puzzleday1.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	total := 0
	for _, line := range lines {
		firstDigit := DetermineNumber(line)
		lastDigit := DetermineLast(line)
		number := firstDigit*10 + lastDigit

		fmt.Printf("line:%s first:%d last:%d number: %d\n ", line, firstDigit, lastDigit, number)
		total += number
		fmt.Printf("running total:%d\n", total)
	}
	fmt.Printf("Total:%d\n", total)
}

func DetermineNumber(line string) int {
	// based on a line get the first digit and the last
	// and combine those for the return value

	// find first digit
	for _, c := range line {
		if i, err := strconv.Atoi(string(c)); err == nil {
			return i
		}
	}
	return 0
}

func DetermineLast(line string) int {
	for r := len(line) - 1; r >= 0; r-- {

		c := line[r]
		if i, err := strconv.Atoi(string(c)); err == nil {
			return i
		}
	}
	panic("no digit found")
	return 0
}
