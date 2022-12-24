package advent

import (
	"fmt"
	"math"
	"testing"
)

type ElfInCove struct {
	X, Y          int // location of elf
	OldX, OldY    int
	WantX, WantY  int  // wanted location of elf
	CanMove       bool // Elf can't move if another elve wants to move in same location
	MoveDirection MoveDirection
}

func (e *ElfInCove) Move() {
	// you can only move, if no OTHER elf
	// has the potential to move into the exact same space!
	e.OldX = e.X
	e.OldY = e.Y
	e.X = e.WantX
	e.Y = e.WantY
}

func (p *PlantingElves) PotentialOtherElfWantsToMoveThere(thisElf *ElfInCove) bool {
	for _, otherElf := range p.Elves {
		if otherElf != thisElf {
			if otherElf.WantX == thisElf.WantX &&
				otherElf.WantY == thisElf.WantY {
				return true
			}
		}
	}
	return false // looks OK
}

// CanGo checks if elf can move in that direction
func (p *PlantingElves) CanGo(elf *ElfInCove, dir MoveDirection) bool {
	elf.CanMove = false
	elf.WantX = -10000
	elf.WantY = -10000
	// an elf does not want to move if it is already "free" when looking around so
	// let's first check that
	if p.ElfAt(elf.X-1, elf.Y) == nil && p.ElfAt(elf.X-1, elf.Y-1) == nil && p.ElfAt(elf.X-1, elf.Y+1) == nil &&
		p.ElfAt(elf.X+1, elf.Y) == nil && p.ElfAt(elf.X+1, elf.Y-1) == nil && p.ElfAt(elf.X+1, elf.Y+1) == nil &&
		p.ElfAt(elf.X, elf.Y-1) == nil && p.ElfAt(elf.X, elf.Y+1) == nil {
		return false // this elf feels free and does not move
	}
	switch dir {
	case MOVE_NORTH: // check if we can move UP/North
		for x := elf.X - 1; x <= elf.X+1; x++ {
			if p.ElfAt(x, elf.Y-1) != nil {
				// can not go here, as there is an elf here
				return false
			}
		}
		// check was ok, we can potentially move up
		elf.WantX = elf.X
		elf.WantY = elf.Y - 1
		elf.CanMove = true
		elf.MoveDirection = MOVE_NORTH
		return true
	case MOVE_SOUTH: // check if we can move DOWN/SOUTH
		for x := elf.X - 1; x <= elf.X+1; x++ {
			if p.ElfAt(x, elf.Y+1) != nil {
				// can not go here, as there is an elf here
				return false
			}
		}
		elf.WantX = elf.X
		elf.WantY = elf.Y + 1
		elf.CanMove = true
		elf.MoveDirection = MOVE_SOUTH
		return true
	case MOVE_WEST: // check if we can move Left/WEST
		for y := elf.Y - 1; y <= elf.Y+1; y++ {
			if p.ElfAt(elf.X-1, y) != nil {
				// can not go here, as there is an elf here
				return false
			}
		}
		elf.WantX = elf.X - 1
		elf.WantY = elf.Y
		elf.CanMove = true
		elf.MoveDirection = MOVE_WEST
		return true
	case MOVE_EAST: // check if we can move Right/EAST
		for y := elf.Y - 1; y <= elf.Y+1; y++ {
			if p.ElfAt(elf.X+1, y) != nil {
				// can not go here, as there is an elf here
				return false
			}
		}
		elf.WantX = elf.X + 1
		elf.WantY = elf.Y
		elf.CanMove = true
		elf.MoveDirection = MOVE_EAST
		return true
	}
	return false // duh, wrong direction?
}

type PlantingElves struct {
	Elves       []*ElfInCove
	AnyElfMoved bool
	ElvesMap    map[int]*ElfInCove
}

func (p *PlantingElves) AddLine(Y int, line string) {
	col := 0
	for _, char := range line {
		if char == '#' {
			elf := &ElfInCove{
				X:    col,
				Y:    Y,
				OldX: col,
				OldY: Y,
			}
			p.Elves = append(p.Elves, elf)
		}
		col++
	}
}

func (p *PlantingElves) MinMaxLocations() (int, int, int, int) {
	xMin := math.MaxInt
	xMax := math.MinInt
	yMin := math.MaxInt
	yMax := math.MinInt
	for _, elf := range p.Elves {
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

func (p *PlantingElves) AddElveToMap(elf *ElfInCove) {
	p.ElvesMap[elf.X*1000+elf.Y] = elf
}

func (p *PlantingElves) RemoveElfFromMap(elf *ElfInCove) {
	delete(p.ElvesMap, elf.X*1000+elf.Y) // delete key
}

func (p *PlantingElves) ElfAt(col, row int) *ElfInCove {
	//	return p.ElvesMap[col*1000+row]

	for _, elf := range p.Elves {
		if elf.X == col && elf.Y == row {
			return elf
		}
	}
	return nil

}

func (p *PlantingElves) ElfAtFormerLocation(col, row int) *ElfInCove {
	//	return p.ElvesMap[col*1000+row]

	for _, elf := range p.Elves {
		if elf.OldX == col && elf.OldY == row {
			return elf
		}
	}
	return nil

}

func (p *PlantingElves) MoveAllElves(moveOrder []MoveDirection) {
	//minX, maxX, minY, maxY := p.MinMaxLocations()

	// determine if Elves can move
	/*for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			elf := p.ElfAt(col, row)
			if elf != nil {
				for _, dir := range moveOrder {
					if p.CanGo(elf, dir) {
						break
					}
				}
			}
		}
	}*/

	// Clean all elves
	for _, elf := range p.Elves {
		elf.CanMove = false
		elf.WantX = -10000
		elf.WantY = -10000
	}

	for _, elf := range p.Elves {
		for _, dir := range moveOrder {
			if p.CanGo(elf, dir) {
				break
			}
		}
	}

	// this only determined if the potential position to move to was empty,
	// but we can still bounce into another elf..

	p.AnyElfMoved = false
	// set all elves to NewLocation
	for _, elf := range p.Elves {
		if !p.PotentialOtherElfWantsToMoveThere(elf) {
			if elf.CanMove {
				elf.Move()
				p.AnyElfMoved = true
			}
		} else {
			elf.CanMove = false
		}
	}

	/*for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			elf := p.ElfAtFormerLocation(col, row)
			if elf != nil {
				if !p.PotentialOtherElfWantsToMoveThere(elf) {
					if elf.CanMove {
						elf.Move()
						p.AnyElfMoved = true
					}
				} else {
					elf.CanMove = false
				}
			}
		}
	}*/

}
func (p *PlantingElves) PrintGrid() {
	minX, maxX, minY, maxY := p.MinMaxLocations()

	for row := minY; row <= maxY; row++ {
		for col := minX; col <= maxX; col++ {
			elf := p.ElfAt(col, row)
			if elf != nil {
				if elf.CanMove {
					switch elf.MoveDirection {
					case MOVE_NORTH:
						fmt.Printf("N")
					case MOVE_SOUTH:
						fmt.Printf("S")
					case MOVE_EAST:
						fmt.Print("E")
					case MOVE_WEST:
						fmt.Printf("W")
					}

				} else {
					fmt.Printf("#")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}

	width := maxX - minX + 1
	height := maxY - minY + 1
	fmt.Printf("Number of elves:%d\n", len(p.Elves))
	fmt.Printf("Rectangle width %d x height %d: %d\n", width, height, width*height)
	fmt.Printf("number of space covered %d - %d: %d\n", width*height, len(p.Elves), (width*height)-len(p.Elves))
}

type MoveDirection int

const MOVE_NORTH MoveDirection = 0
const MOVE_SOUTH MoveDirection = 1
const MOVE_WEST MoveDirection = 2
const MOVE_EAST MoveDirection = 3

func TestMoveElvesDay23_Task1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day23.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	plain := &PlantingElves{
		ElvesMap: make(map[int]*ElfInCove),
	}
	row := 0
	for _, line := range fileLines {
		plain.AddLine(row, line)
		row++
	}

	moveList := []MoveDirection{MOVE_NORTH, MOVE_SOUTH, MOVE_WEST, MOVE_EAST}
	plain.PrintGrid()
	roundsTillNoElfMoves := 0
	n := 1
	for { // we do ten moves}

		plain.MoveAllElves(moveList)
		if n%10 == 0 {
			plain.PrintGrid()
			fmt.Printf("After round %d\n", n)
		}
		if !plain.AnyElfMoved {
			roundsTillNoElfMoves = n
			break
		}
		// move first item of list to the last
		moveList = append(moveList[1:], moveList[0])
		n++
	}
	plain.PrintGrid()

	fmt.Printf("No elf moved in round %d\n", roundsTillNoElfMoves)

	/* 1004 is too high:
		Number of elves:2810
	Rectangle width 141 x height 141: 19881
	number of space covered 19881 - 2810: 17071
	No elf moved in round 1004
	*/

	/*
		Puzzle answer puzzle 1: 4162
		 Puzzle answer puzzle 2: 986 */

}
