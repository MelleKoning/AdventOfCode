package advent2023

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/MelleKoning/AdventOfCode/advent"
)

func GetLinesFromFile(filename string) []string {
	lines, err := advent.GetFileLines(filename)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return lines
}
func TestExampleDay1(t *testing.T) {
	lines, err := advent.GetFileLines("input/puzzleday1.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	total := 0
	for _, line := range lines {
		_, firstDigit := DetermineFirstNumber(line)
		_, lastDigit := DetermineLast(line)
		number := firstDigit*10 + lastDigit

		fmt.Printf("line:%s first:%d last:%d number: %d\n ", line, firstDigit, lastDigit, number)
		total += number
		fmt.Printf("running total:%d\n", total)
	}
	fmt.Printf("Total:%d\n", total)
}

func TestDay1PartTwo(t *testing.T) {
	//lines := GetLinesFromFile("input/day1part2.txt")
	lines := GetLinesFromFile("input/puzzleday1.txt")

	dReader := NewDigitReader()
	total := 0
	for _, line := range lines {
		first := dReader.GetFirstNumber(line)
		last := dReader.GetLastNumber(line)

		number := first*10 + last
		fmt.Printf("line: %s, first %d last %d number %d\n", line, first, last, number)
		total += number
		fmt.Printf("running total:%d\n", total)

	}

	fmt.Printf("Total:%d\n", total)
}

func (d *DigitReader) GetFirstNumber(line string) int {
	posWord, firstDigitWord := d.GetFirstDigit(line)
	posDigit, firstNumber := DetermineFirstNumber(line)

	if posWord == -1 {
		return firstNumber
	}
	if posDigit == -1 {
		return firstDigitWord
	}
	if posWord < posDigit {
		return firstDigitWord
	}
	return firstNumber
}

func (d *DigitReader) GetLastNumber(line string) int {
	posWord, lastDigitWord := d.GetLastDigit(line)
	posDigit, lastNumber := DetermineLast(line)

	if posWord == -1 {
		return lastNumber
	}
	if posDigit == -1 {
		return lastDigitWord
	}
	if posWord > posDigit {
		return lastDigitWord
	}
	return lastNumber
}

type DigitReader struct {
	digits map[int]string
}

// return the position and the found digit
func (d *DigitReader) GetFirstDigit(line string) (int, int) {
	for i := 0; i < len(line); i++ {
		for n := 1; n <= 9; n++ {
			idx := strings.Index(line, d.digits[n])
			if idx > -1 && idx == i {
				return i, n
			}
		}
	}
	return -1, -1
}

func (d *DigitReader) GetLastDigit(line string) (int, int) {
	for i := len(line) - 1; i >= 0; i-- {
		for n := 1; n <= 9; n++ {
			idx := strings.Index(line[i:], d.digits[n])
			if idx > -1 {
				return i, n
			}
		}
	}
	return -1, -1
}

func NewDigitReader() *DigitReader {
	mappy := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	return &DigitReader{
		digits: mappy,
	}
}

func DetermineFirstNumber(line string) (int, int) {
	// based on a line get the first digit and the last
	// and combine those for the return value

	// find first digit
	for xpos, c := range line {
		if i, err := strconv.Atoi(string(c)); err == nil {
			return xpos, i
		}
	}

	return -1, -1 // no number found
}

func DetermineLast(line string) (int, int) {
	for r := len(line) - 1; r >= 0; r-- {

		c := line[r]
		if i, err := strconv.Atoi(string(c)); err == nil {
			return r, i
		}
	}

	return -1, -1 // no number found
}
