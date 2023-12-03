package advent2023

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDay2_Task1(t *testing.T) {
	lines := GetLinesFromFile("input/day2input.txt")

	possibleRed := 12
	possibleGreen := 13
	possibleBlue := 14

	totalGameNumber := 0
	for _, line := range lines {
		lineinfo := strings.Split(line, ":")
		gamepart := lineinfo[0]
		grabs := lineinfo[1]
		gamenumber, _ := strconv.Atoi(strings.Split(gamepart, " ")[1])
		fmt.Printf("game: %d\n", gamenumber)

		possible := true
		reveals := strings.Split(grabs, ";")
		for _, reach := range reveals {
			cubenumberandcolors := strings.Split(reach, ",")
			for _, nrcolor := range cubenumberandcolors {
				nrcolor = strings.Trim(nrcolor, " ")
				number, color := strings.Split(nrcolor, " ")[0], strings.Split(nrcolor, " ")[1]
				getal, _ := strconv.Atoi(number)
				switch color {
				case "blue":
					if getal > possibleBlue {
						possible = false
						break
					}
				case "green":
					if getal > possibleGreen {
						possible = false
						break
					}
				case "red":
					if getal > possibleRed {
						possible = false
						break
					}
				}

			}
		}

		if possible {
			totalGameNumber += gamenumber
		}

	}

	fmt.Printf("totals of gamenumbers %d\n", totalGameNumber)

}

func TestDay2_Task2(t *testing.T) {
	lines := GetLinesFromFile("input/day2input.txt")

	totalPower := 0
	for _, line := range lines {

		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		lineinfo := strings.Split(line, ":")
		gamepart := lineinfo[0]
		grabs := lineinfo[1]
		gamenumber, _ := strconv.Atoi(strings.Split(gamepart, " ")[1])
		fmt.Printf("game: %d\n", gamenumber)

		reveals := strings.Split(grabs, ";")
		for _, reach := range reveals {
			cubenumberandcolors := strings.Split(reach, ",")
			for _, nrcolor := range cubenumberandcolors {
				nrcolor = strings.Trim(nrcolor, " ")
				number, color := strings.Split(nrcolor, " ")[0], strings.Split(nrcolor, " ")[1]
				getal, _ := strconv.Atoi(number)
				switch color {
				case "blue":
					if getal > maxBlue {
						maxBlue = getal
						break
					}
				case "green":
					if getal > maxGreen {
						maxGreen = getal
						break
					}
				case "red":
					if getal > maxRed {
						maxRed = getal
						break
					}
				}
			}
		}

		totalPower += maxBlue * maxGreen * maxRed

	}

	fmt.Printf("power of gamenumbers %d\n", totalPower)

}
