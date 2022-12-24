package advent

import (
	"fmt"
	"strings"
	"testing"
)

func TestCycleTube1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day10.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	X := 1 // start value of the X Register
	CycleNumber := 0
	// keep track of signal strength which is CycleNumber times X value
	signalStrength := make(map[int]int) // [stepCycle]Signal
	addValue := 0
	for _, line := range fileLines {
		fmt.Println(line)
		if line == "noop" {
			// the noop takes 1 cycle
			addValue = 0

			CycleNumber += 1
			signalStrength = StoreXRegIfCycleTimeStamp(X, signalStrength, CycleNumber)

			// nothing else at the moment?
		} else {
			cmd := strings.Split(line, " ")
			if cmd[0] == "addx" {
				addValue = GetNumberFromString(cmd[1])
				CycleNumber += 1
				signalStrength = StoreXRegIfCycleTimeStamp(X, signalStrength, CycleNumber)
				CycleNumber += 1
				signalStrength = StoreXRegIfCycleTimeStamp(X, signalStrength, CycleNumber)
			}
		}
		fmt.Printf(" X: %d, CycleNumber: %d AddValue: %d\n", X, CycleNumber, addValue)
		X = X + addValue
	}

	sumStrengths := 0
	for k, v := range signalStrength {
		fmt.Printf("cycle %d, strength %d\n", k, v)
		sumStrengths += v
	}
	fmt.Printf("Sum is %d", sumStrengths) // 17940 for my puzzle input
}

func StoreXRegIfCycleTimeStamp(currentRegister int, signalmap map[int]int, cycle int) map[int]int {
	if cycle == 20 ||
		((cycle-20)%40 == 0 && cycle > 20) {
		fmt.Printf("Storing RegValue: %d * %d (%d) at cycle %d\n", currentRegister, cycle, currentRegister*cycle, cycle)
		signalmap[cycle] = currentRegister * cycle // store this :)
	}

	return signalmap
}

// ZCBAJFJZ
func TestCRTScreenDay10(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day10.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// each line sets the console at a certain point,
	// collect all the commandvs
	var screen string
	screen = "0123456789012345678901234567890123456789\n"
	spritePos := 1
	pixelPos := 0
	addValue := 0
	cycle := 1
	for _, line := range fileLines {
		if line == "noop" {
			// the noop takes 1 cycle
			screen = screen + PrintAtPos(pixelPos, spritePos)
			pixelPos += 1
			addValue = 0

			if cycle%40 == 0 {
				screen = screen + "\n"
				pixelPos = 0
			}
			cycle += 1

		} else {
			cmd := strings.Split(line, " ")
			if cmd[0] == "addx" {
				screen = screen + PrintAtPos(pixelPos, spritePos)
				pixelPos += 1
				if cycle%40 == 0 {
					screen = screen + "\n"
					pixelPos = 0
				}
				cycle += 1

				screen = screen + PrintAtPos(pixelPos, spritePos)
				pixelPos += 1
				if cycle%40 == 0 {
					screen = screen + "\n"
					pixelPos = 0
				}
				cycle += 1

				addValue = GetNumberFromString(cmd[1])

			}
		}
		spritePos = spritePos + addValue
	}

	fmt.Printf("%v", screen)
}

func PrintAtPos(pixelPos, spritePos int) string {
	if pixelPos == spritePos-1 ||
		pixelPos == spritePos ||
		pixelPos == spritePos+1 {
		return "#"
	}
	return " "
}
