package advent

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// to store the map cells  ...
type CellDay22 struct {
	X, Y       int    // coordinates of the cell
	Block      string // the contents of this cell
	WallyMoved WallyFacing
}

type GridDay22 struct {
	Cells      []*CellDay22 // X columns, Y rows, starting at 0,0
	Rows, Cols int          // amount of Rows(Y) and Cols(X) in the grid
	Wally      *Wally       // location of wally
}

func (g *GridDay22) ReadGrid(lines []string, lineLength int) {
	row := 0
	g.Cols = lineLength

	// Determine the startLocation, this is the first . of the first line
	g.Wally = &Wally{
		X:      strings.Index(lines[0], "."),
		Y:      0,
		Facing: 0, // initially facing right
	}
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		items := strings.Split(line, "")
		col := 0
		// make a row
		for _, block := range items {
			if block == "" {
				block = " "
			}
			cell := &CellDay22{Block: block, X: col, Y: row, WallyMoved: WALLY_NONE}
			g.Cells = append(g.Cells, cell)
			col = col + 1
		}

		// append empty cells at end of row
		for cFiller := col; cFiller < lineLength; cFiller++ {
			block := " "
			cell := &CellDay22{Block: block, X: cFiller, Y: row, WallyMoved: WALLY_NONE}
			g.Cells = append(g.Cells, cell)
		}
		row = row + 1
	}
	g.Rows = row
	// we have now read all lines into the grid
	// and the movement instruction is still left
}

func (g *GridDay22) PrintGrid() {
	fmt.Printf("Rows: %d Cols: %d\n", g.Rows, g.Cols)

	for y := 0; y < g.Rows; y++ {
		fmt.Printf("|")
		for x := 0; x < g.Cols; x++ {
			chr := g.CellAtPoint(x, y).Block
			if chr == "." {
				if g.Wally.X == x && g.Wally.Y == y {
					fmt.Printf("W")
					continue
				}
				wallyMoved := g.CellAtPoint(x, y).WallyMoved
				if wallyMoved != WallyFacing(-1) {
					switch wallyMoved {
					case RIGHT:
						fmt.Printf(">")
					case LEFT:
						fmt.Printf("<")
					case UP:
						fmt.Printf("^")
					case DOWN:
						fmt.Printf("v")
					}
					continue
				}
			}
			fmt.Printf("%s", chr)
		}
		fmt.Printf("|\n")
	}
	fmt.Println()
	fmt.Printf("Wally at (%d,%d) facing: %v\n", g.Wally.X+1, g.Wally.Y+1, FacingString(WallyFacing(g.Wally.Facing)))
}

func FacingString(facing WallyFacing) string {
	switch facing {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case RIGHT:
		return "RIGHT"
	case LEFT:
		return "LEFT"
	}
	return "FACING?"
}
func (g *GridDay22) CellAtPoint(X int, Y int) *CellDay22 {
	valueInList := X + Y*g.Cols
	return g.Cells[valueInList]
}

func LongestLine(lines []string) int {
	max := 0
	for idx, line := range lines {
		if idx > len(lines)-3 {
			break
		}
		if len(line) > max {
			max = len(line)
		}
	}
	return max
}

type WalkInstruction struct {
	TurnLeft  bool
	TurnRight bool
	Steps     int
}

func ParseWalkInstruction(input string) []WalkInstruction {

	var instructions []WalkInstruction
	var currentInstruction WalkInstruction
	var currentSteps string

	for _, r := range input {
		if r == 'L' {
			// If we have any steps left, set the Steps field and append the current instruction
			instruction := convertSteps(&currentSteps)
			if instruction != nil {
				instructions = append(instructions, *instruction)
			}

			// Set the TurnLeft field to true and append the current instruction to the slice
			currentInstruction.TurnLeft = true
			instructions = append(instructions, currentInstruction)

			// Reset the current instruction
			currentInstruction = WalkInstruction{}
		} else if r == 'R' {
			// If we have any steps left, set the Steps field and append the current instruction
			instruction := convertSteps(&currentSteps)
			if instruction != nil {
				instructions = append(instructions, *instruction)
			}

			// Set the TurnRight field to true and append the current instruction to the slice
			currentInstruction.TurnRight = true
			instructions = append(instructions, currentInstruction)

			// Reset the current instruction
			currentInstruction = WalkInstruction{}
		} else if r >= '0' && r <= '9' {
			// If the character is a digit, add it to the currentSteps string
			currentSteps += string(r)
		}
	}

	// If we have any steps left, set the Steps field and append the current instruction
	instruction := convertSteps(&currentSteps)
	if instruction != nil {
		instructions = append(instructions, *instruction)
	}

	return instructions
}

func convertSteps(steps *string) *WalkInstruction {
	if *steps != "" {
		stepCount, _ := strconv.Atoi(*steps)
		instruction := WalkInstruction{
			Steps: stepCount,
		}
		*steps = ""
		return &instruction
	}
	return nil
}

// to track Wally
type Wally struct {
	X, Y   int // location of wally
	Facing int // wally is facing Facing is 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^).
}
type WallyFacing int

const WALLY_NONE WallyFacing = -1
const RIGHT WallyFacing = 0
const DOWN WallyFacing = 1
const LEFT WallyFacing = 2
const UP WallyFacing = 3

func (w *Wally) TurnLeft() {
	switch WallyFacing(w.Facing) {
	case RIGHT:
		w.Facing = int(UP)
	case DOWN:
		w.Facing = int(RIGHT)
	case LEFT:
		w.Facing = int(DOWN)
	case UP:
		w.Facing = int(LEFT)
	}
}

func (w *Wally) TurnRight() {
	switch WallyFacing(w.Facing) {
	case RIGHT:
		w.Facing = int(DOWN)
	case DOWN:
		w.Facing = int(LEFT)
	case LEFT:
		w.Facing = int(UP)
	case UP:
		w.Facing = int(RIGHT)
	}
}

func (g *GridDay22) FirstOnRowAvailable() bool {
	// move all the way to the first "." on this row
	for n := 0; n < g.Cols; n++ {
		if g.CellAtPoint(n, g.Wally.Y).Block == "." {
			g.Wally.X = n
			return true
		}
		if g.CellAtPoint(n, g.Wally.Y).Block == "#" {
			// ouch there is a wall at the other end
			return false
		}
	}
	return false
}

func (g *GridDay22) LastOnRowAvailable() bool {
	// move all the way to the first "." on this row
	for n := g.Cols - 1; n >= 0; n-- {
		if g.CellAtPoint(n, g.Wally.Y).Block == "." {
			g.Wally.X = n
			return true
		}
		if g.CellAtPoint(n, g.Wally.Y).Block == "#" {
			// ouch there is a wall at the other end
			return false
		}
	}
	return false
}

func (g *GridDay22) FirstOnColumnAvailable() bool {
	// move all the way to the first "." on this row
	for n := 0; n < g.Rows; n++ {
		if g.CellAtPoint(g.Wally.X, n).Block == "." {
			g.Wally.Y = n
			return true
		}
		if g.CellAtPoint(g.Wally.X, n).Block == "#" {
			// ouch there is a wall at the other end
			return false
		}
	}
	return false
}

func (g *GridDay22) LastOnColumnAvailable() bool {
	// move all the way to the first "." on this row
	for n := g.Rows - 1; n >= 0; n-- {
		if g.CellAtPoint(g.Wally.X, n).Block == "." {
			g.Wally.Y = n
			return true
		}
		if g.CellAtPoint(g.Wally.X, n).Block == "#" {
			// ouch there is a wall at the other end
			return false
		}
	}
	return false
}

func (g *GridDay22) NextWallyStep() bool {
	if g.Wally.Facing == int(RIGHT) {
		x := g.Wally.X + 1
		if x >= g.Cols {
			return g.FirstOnRowAvailable() // sets wally to first on row or returns false
		}
		if g.CellAtPoint(x, g.Wally.Y).Block == " " { // walking off the board
			return g.FirstOnRowAvailable() // sets wally to first on row or returns false
		}
		if g.CellAtPoint(x, g.Wally.Y).Block == "#" {
			// not possible to move here
			return false
		}
		g.Wally.X = x
		return true // there must be a "." on that cell..
	}
	if g.Wally.Facing == int(LEFT) {
		x := g.Wally.X - 1
		if x < 0 {
			return g.LastOnRowAvailable() // sets wally to first on row or returns false
		}
		if g.CellAtPoint(x, g.Wally.Y).Block == " " { // walking off the board
			return g.LastOnRowAvailable() // sets wally to first on row or returns false
		}
		if g.CellAtPoint(x, g.Wally.Y).Block == "#" {
			// not possible to move here
			return false
		}
		g.Wally.X = x
		return true // there must be a "." on that cell..
	}
	if g.Wally.Facing == int(DOWN) {
		y := g.Wally.Y + 1
		if y >= g.Rows {
			return g.FirstOnColumnAvailable() // sets wally there or returns false
		}
		if g.CellAtPoint(g.Wally.X, y).Block == " " { // walking off the board
			return g.FirstOnColumnAvailable() // sets wally to first on row or returns false
		}
		if g.CellAtPoint(g.Wally.X, y).Block == "#" {
			// not possible to move here
			return false
		}
		g.Wally.Y = y
		return true // there must be a "." on that cell..
	}
	if g.Wally.Facing == int(UP) {
		y := g.Wally.Y - 1
		if y < 0 {
			return g.LastOnColumnAvailable() // sets wally there or returns false
		}
		if g.CellAtPoint(g.Wally.X, y).Block == " " { // walking off the board
			return g.LastOnColumnAvailable() // sets wally to first on row or returns false
		}
		if g.CellAtPoint(g.Wally.X, y).Block == "#" {
			// not possible to move here
			return false
		}
		g.Wally.Y = y
		return true // there must be a "." on that cell..
	}
	return false // uncovered case
}

// move wally over the grid according to the walkInstruction
func (g *GridDay22) WalkWally(instruction WalkInstruction) {
	if instruction.TurnLeft {
		g.Wally.TurnLeft()
	}
	if instruction.TurnRight {
		g.Wally.TurnRight()
	}
	if instruction.Steps > 0 {
		// depending on the direction Wally has to move
		// till hitting a wall or going round the board,
		// it could be round the board has a wall then stay
		// in position
		steps := 0
		for {
			stepSuccess := g.NextWallyStep()
			if !stepSuccess {
				break // no use walking on
			}
			steps += 1
			if steps >= instruction.Steps {
				break
			}
			g.CellAtPoint(g.Wally.X, g.Wally.Y).WallyMoved = WallyFacing(g.Wally.Facing)
		}
	}
}

func (g *GridDay22) NextWallyStepCube() bool {
	var x, y int
	newFacing := g.Wally.Facing
	switch g.Wally.Facing {
	case int(RIGHT):
		{
			x = g.Wally.X + 1
			y = g.Wally.Y
			// edge checking moving right
			// over edges 1 and 5: nothing special
			if x == 100 && g.Wally.Y > 49 && g.Wally.Y <= 99 {
				// over edge of side 3 into bottom of side 2
				y = 49
				x = g.Wally.Y + 50 // flip y of edge 3 onto edge 2 50..99 to 100.149
				newFacing = int(UP)
			}
			if x == 100 && g.Wally.Y > 99 && g.Wally.Y <= 149 {
				// over edge of side 4 into right of side 2
				y = 149 - g.Wally.Y // should flip from 100..149 onto 49..0
				x = 149
				newFacing = int(LEFT)
			}
			if x == 150 {
				// walking over edge of side 2 into right side of 4
				y = 149 - g.Wally.Y // y flip from 0..49 onto 149..100
				x = 99              // right edge of 4
				newFacing = int(LEFT)
			}
			if x == 50 && g.Wally.Y > 149 { // do not take edge 5
				// walking over edge of side 6 into..bottom of side 4
				y = 149
				x = g.Wally.Y - 100 // 150..199..to 50..99
				newFacing = int(UP)
			}
		}
	case int(LEFT):
		{
			x = g.Wally.X - 1
			y = g.Wally.Y
			if x == 49 && g.Wally.Y < 50 {
				// over edge of side 1 into left of side 5
				x = 0
				y = 149 - g.Wally.Y // flip 0..49 becomes 149..100
				newFacing = int(RIGHT)
			} else if x == 49 && g.Wally.Y < 100 {
				// over edge of side 3 into top of side 5
				y = 100
				x = g.Wally.Y - 50 // flip 50..99 to 0..49
				newFacing = int(DOWN)
			} else if x == -1 && g.Wally.Y < 150 {
				// over left edge of side 5 into left side of side 1 to the right
				x = 50
				y = 149 - g.Wally.Y // flip 149..100 to 0..49
				newFacing = int(RIGHT)
			} else if x == -1 && g.Wally.Y < 200 {
				// over left edge of side 6 into top of side 1 !!
				y = 0
				x = g.Wally.Y - 100 // map 150..199 to 50..99
				newFacing = int(DOWN)
			}
		}
	case int(UP):
		{
			y = g.Wally.Y - 1
			x = g.Wally.X
			if y == -1 && g.Wally.X < 100 {
				// from side 1 up into left side 6
				y = g.Wally.X + 100 // 50..99 to 150..199
				x = 0
				newFacing = int(RIGHT)
			} else if y == -1 && g.Wally.X < 150 {
				// from side 2 into bottom of side 6
				y = 199
				x = g.Wally.X - 100 // 100..149 onto 0..49
				newFacing = int(UP)
			} else if y == 99 && g.Wally.X < 50 {
				// face 5 up into left side of 3 to the right
				x = 50
				y = g.Wally.X + 50 // flip 0..49 to 50..99
				newFacing = int(RIGHT)
			}
		}
	case int(DOWN):
		{
			y = g.Wally.Y + 1
			x = g.Wally.X
			if y == 50 && g.Wally.X > 99 {
				// from side 2 down into right side of side 3
				y = g.Wally.X - 50 // from 100..149 to 50..99
				x = 99
				newFacing = int(LEFT)
			} else if y == 150 && g.Wally.X > 49 {
				// from side 4 down into right side of side 6
				y = g.Wally.X + 100 // from 50..99 onto 150.199
				x = 49
				newFacing = int(LEFT)
			} else if y == 200 && g.Wally.X < 50 {
				// from side 6 to top of side 2
				y = 0
				x = 149 - g.Wally.X // flip 00..49 ont 149..100
				newFacing = int(DOWN)
			} else if y == 200 {
				// over the edge is not allowed
				return false // but this should not have occurred at all..
			}
		}
	}
	if g.CellAtPoint(x, y).Block == "#" { // not possible to move here
		return false
	}
	if g.CellAtPoint(x, y).Block == " " {
		return false // we should not have arrived here?
	}
	g.Wally.X = x
	g.Wally.Y = y
	g.Wally.Facing = newFacing
	return true

}

func (g *GridDay22) WalkWallyCube(instruction WalkInstruction) {
	/* Each side 1..6 connects on each edge to
		   another side... if each cell can
		   tell us in what side they exist and when we go over
		   an edge (>49, >99, >149) (<0,<50,<100), <then we can
		   decide to what edge they go and what direction Wally faces

	                   ^ 6     6^
	                 <5 1 -> <-2-> 4
	   					^      v
						v      3
					5<	3> 2<  ^
	            3       v
				v
			1<	5	<>	4  2<
				v       ^
				^	    6
			<1  6 >4
	            v
				2

	*/

	if instruction.TurnLeft {
		g.Wally.TurnLeft()
	}
	if instruction.TurnRight {
		g.Wally.TurnRight()
	}
	if instruction.Steps > 0 {
		// depending on the direction Wally has to move
		// till hitting a wall or going off edge of cube
		steps := 0
		for {
			stepSuccess := g.NextWallyStepCube()
			if !stepSuccess {
				break // no use walking on
			}
			steps += 1
			if steps >= instruction.Steps {
				break
			}
			g.CellAtPoint(g.Wally.X, g.Wally.Y).WallyMoved = WallyFacing(g.Wally.Facing)
		}
	}
}

func TestParseWalk(t *testing.T) {
	instructions := ParseWalkInstruction("10R5L5R10L4R5L5")
	assert.Equal(t, 13, len(instructions))
}

func TestMapDay22_Task1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day22.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	grid := &GridDay22{}
	grid.ReadGrid(fileLines, LongestLine(fileLines))
	grid.PrintGrid()

	WalkInstructions := ParseWalkInstruction(fileLines[len(fileLines)-1])
	fmt.Printf("%+v", WalkInstructions)

	for _, instr := range WalkInstructions {
		fmt.Printf("Instruction: %v\n", instr)
		grid.WalkWally(instr)
	}
	grid.PrintGrid()

	fmt.Printf("Password:%d\n", (grid.Wally.Y+1)*1000+(grid.Wally.X+1)*4+grid.Wally.Facing)
}

func TestMapDay22_Task2(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day22.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	grid := &GridDay22{}
	grid.ReadGrid(fileLines, LongestLine(fileLines))
	grid.PrintGrid()

	WalkInstructions := ParseWalkInstruction(fileLines[len(fileLines)-1])
	//fmt.Printf("%+v", WalkInstructions)

	for _, instr := range WalkInstructions {
		grid.WalkWallyCube(instr)
		//grid.PrintGrid()
		fmt.Printf("Instruction was: TurnLeft: %v, TurnRight: %v, Steps: %v\n", instr.TurnLeft, instr.TurnRight, instr.Steps)
	}
	grid.PrintGrid()

	fmt.Printf("Password:%d\n", (grid.Wally.Y+1)*1000+(grid.Wally.X+1)*4+grid.Wally.Facing)
	// 197160 too high
	// 46290 too low

	//Wally at (54,142) facing: 3
	// Password:142219 - but nothing walking into bottom of side 2 - check UP

	//Wally at (98,52) facing: 3
	//Password:52395 - answer is TOO LOW - walking edges 2 and 4 should be flipping the Y value

	//Wally at (100,90) facing: 1
	//Password:90401 -- WRONG

	//Wally at (139,18) facing: 1
	//Password:18557 -- too low..

	// Wally at (74,117) facing: RIGHT
	//Password:117296 -- WRONG

	// Draw a cube
	// Wally at (16,145) facing: DOWN
	// Password:145065 -> That's the right answer! *
}
