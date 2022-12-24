package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFileLines(file string) ([]string, error) {
	// read the data from the file
	file1, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file with %v", err)
	}

	fileScanner := bufio.NewScanner(file1)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file1.Close()
	return fileLines, nil
}

func GetIncrements(numbers []int) int {
	previousNumber := 999
	increments := 0
	for _, number := range numbers {

		// compare previous with current line
		if number > previousNumber {
			increments++
		}
		previousNumber = number

	}
	return increments
}

type SubmarineCommand int

const (
	Command_Unknown SubmarineCommand = iota
	Command_Forward
	Command_Down
	Command_Up
)

type Command struct {
	command SubmarineCommand
	number  int
}

func GetSubmarinCommand(d string) SubmarineCommand {
	switch d {
	case "forward":
		return Command_Forward
	case "up":
		return Command_Up
	case "down":
		return Command_Down
	}
	return Command_Unknown
}

func ByteStringToDecimal(gammaCode string) int {

	powerOfTwo := 1
	totalGamma := 0
	totalEpsilon := 0
	gammaRunes := []rune(gammaCode)
	for n := len(gammaRunes) - 1; n >= 0; n-- {
		if gammaRunes[n] == '1' {
			totalGamma += powerOfTwo
		} else {
			totalEpsilon += powerOfTwo
		}
		powerOfTwo = powerOfTwo * 2
	}
	return totalGamma
}

func GetGammaCounterForLines(lines []string) []int {
	gammaCounter := make([]int, 12)

	for _, line := range lines {
		runes := []rune(line)
		for pos := 0; pos < len(runes); pos++ {
			intVar, err := strconv.Atoi(string(runes[pos]))
			if err != nil {
				fmt.Printf("shit %v", err)
			}
			if intVar == 1 {
				gammaCounter[pos]++
			}
		}

	}
	return gammaCounter
}

// filters given lines that have rune r on position pos
func GetAllLinesWithRuneAtPos(lines []string, r rune, pos int) []string {
	var returnLines []string
	for _, line := range lines {
		//fmt.Printf("processing line %s", line)
		runes := []rune(line)
		if runes[pos] == r { // this line is ok
			returnLines = append(returnLines, line)
		}
	}
	fmt.Printf("\n%d lines found matching rune %s at position %d\n", len(returnLines), string(r), pos)
	return returnLines
}

func GetNumberFromString(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("could not parse %s\n", str)
		return -1
	}
	return i
}
