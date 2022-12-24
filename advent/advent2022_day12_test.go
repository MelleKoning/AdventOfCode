package advent

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Point struct {
	X, Y int
}
type Day12Cell struct {
	Point    *Point
	Block    string
	Parent   *Day12Cell
	Explored bool
}

type Day12Grid struct {
	Cells         []*Day12Cell // X columns, Y rows, starting at 0,0
	Rows, Cols    int
	LabeledPoints []*Day12Cell
}

func (g *Day12Grid) FindStartPosition() (*Day12Cell, error) {

	for _, cell := range g.Cells {
		if cell.Block == "S" {
			return cell, nil
		}
	}
	return nil, fmt.Errorf("startposition not found")
}

func (g *Day12Grid) FindStartPositionsFirstColumn() ([]*Day12Cell, error) {

	var startCells []*Day12Cell
	for _, cell := range g.Cells {
		if cell.Block == "S" || (cell.Block == "a" && cell.Point.X == 0) {
			startCells = append(startCells, cell)
		}
	}
	if len(startCells) > 0 {
		return startCells, nil
	}
	return nil, fmt.Errorf("startposition not found")
}

func (g *Day12Grid) FindEndPosition() (*Day12Cell, error) {

	for _, cell := range g.Cells {
		if cell.Block == "E" {
			return cell, nil
		}
	}
	return nil, fmt.Errorf("endposition not found")
}

func TestDay12_1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day12.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// grid of col 0..x to row 0..y
	grid := ReadGrid(fileLines)

	fmt.Printf("Rows: %d Cols: %d", grid.Rows, grid.Cols)
	start, err := grid.FindStartPosition()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Start at %v\n", start)

	end, err := grid.FindEndPosition()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("End at %v\n", end)

	fmt.Printf("%s %s\n", grid.CellAtPoint(Point{0, 0}).Block, grid.CellAtPoint(Point{5, 2}).Block)

	GridWalkerBFS(grid, start, end)
	//GridWalkerRecursive(grid, start, end)
}

// Breath First Search - finds A solution - but not the fastest
func GridWalkerBFS(grid *Day12Grid, currentPos *Day12Cell, searchPoint *Day12Cell) int {
	var Q []*Day12Cell
	var v *Day12Cell
	Q = append(Q, currentPos) // start
	currentPos.Explored = true
	for {
		if len(Q) == 0 {
			break
		}
		v, Q = Q[0], Q[1:] // pop from queue
		//fmt.Printf("Explore: %v\n", v.Block)
		// if the currentPos is endpoint, we are done
		if grid.PointIs(v, "E") {
			fmt.Printf("found endpoint\n") //2770 too highs)
			//walkedPoints = walkedPoints[:len(walkedPoints)-1]
			currentPoint := v
			steps := 0
			for {
				fmt.Printf("step %d %s at (%d, %d)\n", steps, currentPoint.Block, currentPoint.Point.X, currentPoint.Point.Y)
				currentPoint = currentPoint.Parent
				if currentPoint == nil {
					break
				}
				steps += 1

			}
			fmt.Printf("steps: %d", steps)

			return steps
		}

		//for all edges of v try walking left, right, up, down
		for _, move := range grid.PossibleMoves(v, searchPoint) {
			// if move not labeled as explored already
			if !move.Explored {
				move.Explored = true
				move.Parent = v
				Q = append(Q, move)
			}
		}
	}
	return -1
}

func GridWalkerRecursive(grid *Day12Grid, currentPos *Day12Cell, searchPoint *Day12Cell) int {
	if PointIn(currentPos, grid.LabeledPoints) {
		return -1 // already visited that point, this is not ok
	}

	if !PointIn(currentPos, grid.LabeledPoints) {
		grid.LabeledPoints = append(grid.LabeledPoints, currentPos)
	}
	if grid.PointIs(currentPos, "E") {
		fmt.Printf("found endpoint\n") //2770 too highs)
		steps := 1
		for {
			//fmt.Printf("step %d %s at (%d, %d) ", steps, currentPos.Block, currentPos.Point.X, currentPos.Point.Y)
			currentPos = currentPos.Parent
			steps += 1
			if currentPos == nil {
				break
			}
		}
		fmt.Printf("Steps: %d", steps)
		//return len(walkedPoints)
		return 1
	}

	//for all edges of v try walking left, right, up, down
	for _, move := range grid.PossibleMoves(currentPos, searchPoint) {
		move.Parent = currentPos
		GridWalkerRecursive(grid, move, searchPoint)
	}
	// everything done remove again from labeled points
	grid.LabeledPoints = grid.LabeledPoints[:len(grid.LabeledPoints)-1]

	return -1
}

// Returns possible moves for the grid,
// hopefully sorted to the most promising moves
func (g *Day12Grid) PossibleMoves(currentPos *Day12Cell, searchPoint *Day12Cell) []*Day12Cell {
	var PossibleMoves []*Day12Cell
	//try walking left, right, up, down
	tryX := currentPos.Point.X + 1
	if g.XInBounds(tryX) && !PointIn(g.CellAtPoint(Point{X: tryX, Y: currentPos.Point.Y}), g.LabeledPoints) {

		//if currentPos.Block
		blokTry := g.CellAtPoint(Point{X: tryX, Y: currentPos.Point.Y})
		if g.WalkAllowed(g.CellAtPoint(*currentPos.Point), blokTry) {
			PossibleMoves = append(PossibleMoves, blokTry)
		}
	}
	tryX = currentPos.Point.X - 1
	if g.XInBounds(tryX) && !PointIn(g.CellAtPoint(Point{X: tryX, Y: currentPos.Point.Y}), g.LabeledPoints) {
		//if currentPos.Block
		blokTry := g.CellAtPoint(Point{X: tryX, Y: currentPos.Point.Y})
		if g.WalkAllowed(g.CellAtPoint(*currentPos.Point), g.CellAtPoint(*blokTry.Point)) {
			PossibleMoves = append(PossibleMoves, blokTry)
		}
	}
	tryY := currentPos.Point.Y + 1
	if g.YInBounds(tryY) && !PointIn(g.CellAtPoint(Point{X: currentPos.Point.X, Y: tryY}), g.LabeledPoints) {
		//if currentPos.Block
		blokTry := g.CellAtPoint(Point{X: currentPos.Point.X, Y: tryY})
		if g.WalkAllowed(g.CellAtPoint(*currentPos.Point), g.CellAtPoint(*blokTry.Point)) {
			PossibleMoves = append(PossibleMoves, blokTry)
		}
	}

	tryY = currentPos.Point.Y - 1
	if g.YInBounds(tryY) && !PointIn(g.CellAtPoint(Point{X: currentPos.Point.X, Y: tryY}), g.LabeledPoints) {
		//if currentPos.Block
		blokTry := g.CellAtPoint(Point{X: currentPos.Point.X, Y: tryY})
		if g.WalkAllowed(g.CellAtPoint(*currentPos.Point), g.CellAtPoint(*blokTry.Point)) {
			PossibleMoves = append(PossibleMoves, blokTry)
		}
	}
	if len(PossibleMoves) == 1 {
		return PossibleMoves
	}
	// sort PossibleMoves to most promising
	sort.SliceStable(PossibleMoves, func(i, j int) bool {
		if PossibleMoves[i].Block > PossibleMoves[j].Block {
			return true
		}
		if PossibleMoves[j].Block > PossibleMoves[i].Block {
			return false
		}
		return false
	})
	return PossibleMoves
}

func (g *Day12Grid) UnExploreAllCells() {
	for _, c := range g.Cells {
		c.Explored = false
		c.Parent = nil
	}
}
func (g *Day12Grid) PointIs(p *Day12Cell, is string) bool {

	for _, c := range g.Cells {
		if c.Point.X == p.Point.X && c.Point.Y == p.Point.Y {
			if c.Block == is {
				return true
			}
		}
	}
	return false

}

func (g *Day12Grid) CellAtPoint(p Point) *Day12Cell {

	valueInList := p.X + p.Y*g.Cols
	return g.Cells[valueInList]

}

func (g *Day12Grid) XInBounds(x int) bool {
	return x >= 0 && x < g.Cols
}

func (g *Day12Grid) YInBounds(y int) bool {
	return y >= 0 && y < g.Rows
}

func PointIn(p *Day12Cell, labeledPoints []*Day12Cell) bool {
	for _, w := range labeledPoints {
		if w.Point.X == p.Point.X && w.Point.Y == p.Point.Y {
			return true // already visited
		}
	}
	return false
}

func (g *Day12Grid) WalkAllowed(from, to *Day12Cell) bool {
	// walk allowed when to is one higher then from
	// or when to is lower then from...
	// so walking from a to b, b to c, c to d is allowed as well as
	// d to c, d to b, d to a etc.
	valueTo := to.Block
	valueFrom := from.Block
	// The endpoint is special char E
	if valueTo == "E" && (valueFrom == "z" || valueFrom == "y") {
		return true
	}
	// The startpoint is special char S
	if valueFrom == "S" && valueTo == "a" {
		return true
	}
	if valueTo != "E" && valueFrom != "S" {
		if byte(valueFrom[0]) == byte(valueTo[0])-1 ||
			byte(valueFrom[0]) == byte(valueTo[0]) ||
			byte(valueFrom[0]) > byte(valueTo[0]) {
			return true
		}
	}
	return false
}

func ReadGrid(lines []string) *Day12Grid {
	day12Grid := &Day12Grid{}

	row := 0
	for _, line := range lines {
		items := strings.Split(line, "")
		col := 0
		// make a row
		//day12Grid.
		for _, block := range items {
			cell := &Day12Cell{Block: block, Point: &Point{X: col, Y: row}}
			day12Grid.Cells = append(day12Grid.Cells, cell)
			col = col + 1
		}
		day12Grid.Cols = col
		row = row + 1
	}
	day12Grid.Rows = row
	return day12Grid
}

func TestWalkAllowed(t *testing.T) {
	tests := map[string]struct {
		from, to *Day12Cell
		expected bool
	}{
		"Endpoint": {
			from:     &Day12Cell{Block: "z"},
			to:       &Day12Cell{Block: "E"},
			expected: true,
		},
		"FromAtoB": {
			from:     &Day12Cell{Block: "a"},
			to:       &Day12Cell{Block: "b"},
			expected: true,
		},
		"FromBtoA": {
			from:     &Day12Cell{Block: "b"},
			to:       &Day12Cell{Block: "a"},
			expected: true,
		}, "FromStoA": {
			from:     &Day12Cell{Block: "S"},
			to:       &Day12Cell{Block: "a"},
			expected: true,
		}, "FromKtoK": {
			from:     &Day12Cell{Block: "k"},
			to:       &Day12Cell{Block: "k"},
			expected: true,
		}, "FromDtoB": {
			from:     &Day12Cell{Block: "d"},
			to:       &Day12Cell{Block: "b"},
			expected: true,
		}, "FromBtoD": {
			from:     &Day12Cell{Block: "b"},
			to:       &Day12Cell{Block: "d"},
			expected: false,
		}, "FromYoE": { // E is z height, so y to E is allowed
			from:     &Day12Cell{Block: "y"},
			to:       &Day12Cell{Block: "E"},
			expected: true,
		}}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			grid := &Day12Grid{}
			result := grid.WalkAllowed(tt.from, tt.to)
			assert.Equal(t, tt.expected, result)
		})
	}

}
func TestCells(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day12example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// grid of col 0..x to row 0..y
	grid := ReadGrid(fileLines)

	fmt.Printf("Rows: %d Cols: %d", grid.Rows, grid.Cols)

	for y := 0; y < 5; y++ {

		for x := 0; x < 8; x++ {
			fmt.Printf("%s", grid.CellAtPoint(Point{X: x, Y: y}).Block)
		}
		fmt.Println()
	}
}

func TestDay12_2(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day12.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// grid of col 0..x to row 0..y
	grid := ReadGrid(fileLines)

	fmt.Printf("Rows: %d Cols: %d", grid.Rows, grid.Cols)
	starters, err := grid.FindStartPositionsFirstColumn()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Start at %v\n", starters)

	end, err := grid.FindEndPosition()
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("End at %v\n", end)

	type trail struct {
		StartPoint *Day12Cell
		Steps      int
	}
	var trailCollection []*trail

	for _, start := range starters {
		grid.UnExploreAllCells()
		shorty := GridWalkerBFS(grid, start, end)
		trail := &trail{
			StartPoint: start,
			Steps:      shorty,
		}
		trailCollection = append(trailCollection, trail)
	}

	fmt.Println()
	for _, tracks := range trailCollection {
		fmt.Printf("start: %d,%d length: %d\n", tracks.StartPoint.Point.X, tracks.StartPoint.Point.Y, tracks.Steps)
	}
}
