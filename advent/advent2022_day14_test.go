package advent

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

type Cell14 int

const CELL_EMPTY Cell14 = 0
const CELL_ROCK Cell14 = 1
const CELL_SAND Cell14 = 2
const CELL_DROPPOINT Cell14 = 3 // this is the 500,0 point

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
	prevX := nextX
	prevY := nextY

	numbers = numbers[2:]

	for {
		nextX := GetNumberFromString(numbers[entryX])
		nextY := GetNumberFromString(numbers[entryY])
		// create rocks from prevX..nextX prevY..nextY
		XInc := 0
		YInc := 0
		if nextX != prevX {
			XInc = 1
			if nextX < prevX {
				XInc = -1
			}
		} else { // assume there is a diff in Y
			YInc = 1
			if nextY < prevY {
				YInc = -1
			}
		}
		r.rocksMap[Point{prevX, prevY}] = CELL_ROCK
		for {
			prevX = prevX + XInc
			prevY = prevY + YInc
			r.rocksMap[Point{prevX, prevY}] = CELL_ROCK
			if prevX == nextX && prevY == nextY {
				break
			}
		}
		// remove entries from slice
		numbers = numbers[2:]
		if len(numbers) == 0 {
			break
		}

	}

}

func (r *Rocks) MinMaxLocations() (int, int, int, int) {
	xMin := math.MaxInt
	xMax := math.MinInt
	yMin := math.MaxInt
	yMax := math.MinInt
	for elf := range r.rocksMap {
		if elf.X < xMin {
			xMin = elf.X
		}
		if elf.X > xMax {
			xMax = elf.X
		}
		if elf.Y > yMax {
			yMax = elf.Y
		}
		if elf.Y < yMin {
			yMin = elf.Y
		}
	}
	return xMin, xMax, yMin, yMax
}

func (r *Rocks) PrintGrid() {
	minX, maxX, minY, maxY := r.MinMaxLocations()
	for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			block, ok := r.rocksMap[Point{col, row}]
			if ok {
				switch block {
				case CELL_DROPPOINT:
					fmt.Printf("+")
				case CELL_ROCK:
					fmt.Printf("#")
				case CELL_SAND:
					fmt.Printf("O")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	width := maxX - minX + 1
	height := maxY - minY + 1
	fmt.Printf("Cave %d by %d\n", width, height)

}

func NewCave(fileLines []string) *Rocks {
	Cave := &Rocks{
		rocksMap: make(map[Point]Cell14),
	}

	// initialize droppoint :)
	Cave.rocksMap[Point{500, 0}] = CELL_DROPPOINT
	for _, line := range fileLines {
		Cave.ParseLine(line)
		//Cave.PrintGrid()
	}
	return Cave
}

// returns true when sand drops into the VOID
func (r *Rocks) DropSand() bool {
	// Dropping sand in the cave starts at (500,0)
	sandLoc := Point{500, 0}

	// drop till we hit something other then empty,
	// or maybe if we go beyond the current Cove Coordinates
	minX, maxX, _, maxY := r.MinMaxLocations()
	for {
		testPoint := Point{sandLoc.X, sandLoc.Y}
		cell, solid := r.rocksMap[testPoint]
		if solid && cell != CELL_DROPPOINT {
			// we hit something, test if can go left or right
			_, leftSolid := r.rocksMap[Point{sandLoc.X - 1, sandLoc.Y}]
			if !leftSolid && sandLoc.X-1 >= minX { // we can go left
				sandLoc.X = sandLoc.X - 1
				continue
			}
			_, rightSolid := r.rocksMap[Point{sandLoc.X + 1, sandLoc.Y}]
			if !rightSolid && sandLoc.X <= maxX { // sand can go right
				sandLoc.X = sandLoc.X + 1
				continue
			}
			if leftSolid && rightSolid { // we can't go left or right, so sand should rest one up
				r.rocksMap[Point{sandLoc.X, sandLoc.Y - 1}] = CELL_SAND
				break
			}
		}
		// empty as not found so continue dropping, but let's see if we already
		// dropped out of the bounds of the grid
		if testPoint.X < minX || testPoint.X > maxX || testPoint.Y > maxY {
			// THE VOID
			return true
		}
		sandLoc.Y = sandLoc.Y + 1
	}
	return false
}
func TestSandFalling_Example1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day14example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	Cave := NewCave(fileLines)
	Cave.PrintGrid()
	for n := 1; n < 26; n++ {
		void := Cave.DropSand()
		if void {
			fmt.Printf("Sand droplet: %d\n", n)
		}
		Cave.PrintGrid()
	}

}
