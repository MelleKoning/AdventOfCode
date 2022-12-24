package advent

import (
	"strings"
	"testing"
)

type Cell14 int

const CELL_EMPTY Cell14 = 0
const CELL_ROCK Cell14 = 1
const CELL_SAND Cell14 = 2

type Rocks struct {
	rocksMap map[Point]Cell14
}

func (r *Rocks) ParseLine(line string) {
	datapoints := strings.Split(line, " -> ")
	numberString := strings.Join(datapoints, ",")
	numbers := strings.Split(numberString, ",")

	entryX := 0
	entryY := 1
	nextX := GetNumberFromString(numbers[entryX])
	nextY := GetNumberFromString(numbers[entryY])
	numbers = numbers[2:]

	for {
		prevX := nextX
		prevY := nextY
		nextX := GetNumberFromString(numbers[entryX])
		nextY := GetNumberFromString(numbers[entryY])
		// create rocks from prevX..nextX prevY..nextY
		if nextX != prevX {
			XInc := 1
			if nextX < prevX {
				XInc = -1
			}
			for n := prevX; n <= nextX; n += XInc {
				r.rocksMap[Point{n, prevY}] = CELL_ROCK
			}
		} else {
			YInc := 1
			if nextY < prevY {
				YInc = -1
			}
			for n := prevY; n <= nextY; n += YInc {
				r.rocksMap[Point{prevX, n}] = CELL_ROCK
			}
		}
		// remove entries from slice
		numbers = numbers[2:]
		if len(numbers) == 0 {
			break
		}

	}

}

func TestSandFalling_Task1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day14example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	Grove := &Rocks{
		rocksMap: make(map[Point]Cell14),
	}
	for _, line := range fileLines {
		Grove.ParseLine(line)
	}
}
