package advent2023

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestDay3Task1(t *testing.T) {
	lines := GetLinesFromFile("input/day3input.txt")

	var allNumbers []*PosNumber
	for lineIndex, line := range lines {
		numbersForLine := findNumbers(line)

		fmt.Printf("%s\n", line)

		for _, res := range numbersForLine {
			// record the lineIndex for found numbers
			res.LineIndex = lineIndex
			fmt.Printf("line: %d, cols:(%d-%d), n:%d\n ", res.LineIndex, res.StartPos, res.EndPos, res.Number)
		}
		allNumbers = append(allNumbers, numbersForLine...)
	}

	// Now we have all numbers and their locations in allNumbers.
	// When we inspect the surrounding for symbols we should
	// be able to find out what numbers are valid
	// to use those in the totalSum

	totalSum := 0
	for _, number := range allNumbers {
		if symbolAround(number, lines) {
			totalSum += number.Number
		}
	}

	fmt.Printf("totalSum: %d\n", totalSum)
}

func symbolAround(number *PosNumber, lines []string) bool {
	from := number.StartPos - 1
	if from < 0 {
		from = 0
	}
	to := number.EndPos + 1
	if to > len(lines[0]) {
		to = len(lines[0])
	} // assume all lines have same len

	// loop three lines
	for looplines := number.LineIndex - 1; looplines <= number.LineIndex+1; looplines++ {
		if looplines < 0 || looplines >= len(lines) {
			continue
		}
		// inspect line characters
		symbolFound := strings.IndexAny(lines[looplines][from:to], "+#$*@/=%-&")
		// we know enough already
		if symbolFound > -1 {
			return true
		}
	}

	return false
}

type PosNumber struct {
	StartPos  int
	EndPos    int
	LineIndex int
	Number    int
}

func findNumbers(str string) []*PosNumber {
	// Define the regular expression for finding numbers
	re := regexp.MustCompile(`\d+`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatchIndex(str, -1)

	// Extract numbers along with start and end positions
	result := make([]*PosNumber, len(matches))
	for i, match := range matches {
		start := match[0]
		end := match[1]
		number, _ := strconv.Atoi(str[start:end])

		result[i] = &PosNumber{
			Number:   number,
			StartPos: start,
			EndPos:   end,
		}
	}

	return result
}
