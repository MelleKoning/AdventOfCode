package advent

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind8_1(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day8.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	var grid [][]int
	for _, line := range fileLines {
		row := strings.Split(line, "")
		var gridRow []int
		for _, s := range row {
			x, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			gridRow = append(gridRow, x)
		}
		grid = append(grid, gridRow)
	}

	fmt.Printf("visible trees: %d", countVisibleTrees(grid))

	fmt.Printf("best scenic score: %d", countScenicScore(grid))
}

// Define a function to count the number of visible trees in a grid.
func countVisibleTrees(grid [][]int) int {
	// Initialize a set to keep track of which trees have been counted.
	counted := make(map[string]bool)

	// Initialize a counter for the number of visible trees.
	count := 0

	// Loop through the rows and columns of the grid.
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			// Skip this tree if it has already been counted.
			if counted[fmt.Sprintf("%d,%d", row, column)] {
				continue
			}

			// Set isVisible to false. It will only be set to true if the tree is visible from any direction.
			isVisible := false

			// Check if the current tree is visible from the left.
			if column == 0 || (column > 0 && grid[row][column] > maxInRow(grid[row], 0, column)) {
				isVisible = true
			}

			// Check if the current tree is visible from the right.
			if column == len(grid[row])-1 || (column < len(grid[row])-1 && grid[row][column] > maxInRow(grid[row], column+1, len(grid[row]))) {
				isVisible = true
			}

			// Check if the current tree is visible from the top.
			if row == 0 || (row > 0 && grid[row][column] > maxInColumn(grid, 0, row, column)) {
				isVisible = true
			}

			// Check if the current tree is visible from the bottom.
			if row == len(grid)-1 || (row < len(grid)-1 && grid[row][column] > maxInColumn(grid, row+1, len(grid), column)) {
				isVisible = true
			}

			// If the current tree is visible, increment the counter and mark it as counted.
			if isVisible {
				count++
				counted[fmt.Sprintf("%d,%d", row, column)] = true
			}
		}
	}

	return count
}

func countScenicScore(grid [][]int) int {
	// Initialize a counter for the number of visible trees.
	maxScenicScore := 0

	// Loop through the rows and columns of the grid.
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {

			l := AmountOfVisibleTreesLeftOf(grid[row], column, grid[row][column])
			r := AmountOfVisibleTreesRightOf(grid[row], column, grid[row][column])
			t := AmountVisibleTreesTop(grid, row, column, grid[row][column])
			b := AmountVisibleTreesBottom(grid, row, column, grid[row][column])

			thisTreeScenicScore := l * r * t * b

			if thisTreeScenicScore > maxScenicScore {
				maxScenicScore = thisTreeScenicScore
				fmt.Printf("row:%d,col:%d, score:%d", row, column, maxScenicScore)
			}
		}
	}

	return maxScenicScore
}

func TestAmountLeft(t *testing.T) {
	scenicLeft := AmountOfVisibleTreesLeftOf([]int{1, 2, 5, 1, 2, 3, 5, 1, 1, 1}, 6, 5)
	assert.Equal(t, 4, scenicLeft)

	scenicRight := AmountOfVisibleTreesRightOf([]int{1, 2, 9, 1, 2, 3, 5, 1, 1, 5}, 6, 5)
	assert.Equal(t, scenicRight, 3)
}

func AmountVisibleTreesTop(grid [][]int, rowpos, column, height int) int {
	// Initialize the maximum value to the first element in the column.
	count := 0

	if rowpos == 0 {
		return 0
	}

	// Loop through the elements in the given range of the column.
	for i := rowpos - 1; i > -1; i-- {
		if grid[i][column] >= height {
			// stop here
			count += 1
			break
		}
		count += 1
	}

	return count
}

func AmountVisibleTreesBottom(grid [][]int, rowpos, column, height int) int {
	// Initialize the maximum value to the first element in the column.
	count := 0

	if rowpos == len(grid[rowpos])-1 {
		return 0
	}
	// Loop through the elements in the given range of the column.
	for i := rowpos + 1; i < len(grid[rowpos]); i++ {
		if grid[i][column] >= height {
			// stop here
			count += 1
			break
		}
		count += 1
	}

	return count
}

func AmountOfVisibleTreesLeftOf(row []int, cpos, height int) int {
	count := 0

	if cpos == 0 {
		return 0
	}
	for i := cpos - 1; i > -1; i-- {
		if row[i] >= height {
			// stop here
			count += 1
			break
		}
		count += 1
	}
	return count
}

func AmountOfVisibleTreesRightOf(row []int, cpos, height int) int {
	count := 0

	if cpos == len(row) {
		return 0
	}
	for i := cpos + 1; i < len(row); i++ {
		if row[i] >= height {
			// stop here
			count += 1
			break
		}
		count += 1
	}
	return count
}

// Define a function to find the maximum value in a given row of a grid.
func maxInRow(row []int, start, end int) int {
	// Initialize the maximum value to the first element in the row.
	max := row[start]

	// Loop through the elements in the given range of the row.
	for i := start + 1; i < end; i++ {
		// Update the maximum value if the current element is larger.
		if row[i] > max {
			max = row[i]
		}
	}

	return max
}

// Define a function to find the maximum value in a given column of a grid.
func maxInColumn(grid [][]int, start, end, column int) int {
	// Initialize the maximum value to the first element in the column.
	max := grid[start][column]

	// Loop through the elements in the given range of the column.
	for i := start + 1; i < end; i++ {
		// Update the maximum value if the current element is larger.
		if grid[i][column] > max {
			max = grid[i][column]
		}
	}

	return max
}

func TestCalculat(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day8.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Parse the grid of trees from the string array.
	var grid [][]int
	for _, line := range fileLines {
		// Parse the current row from the current line.
		row := make([]int, len(line))
		for i, r := range line {
			fmt.Sscanf(string(r), "%d", &row[i])
		}

		// Append the parsed row to the grid.
		grid = append(grid, row)
	}

	// The maximum scenic score and its coordinates.
	var maxScenicScore int
	var maxScenicScoreX, maxScenicScoreY int

	// Find the tree with the maximum scenic score.
	for y, row := range grid {
		for x, tree := range row {
			// The current scenic score.
			var scenicScore int = 0

			// Look up.
			for i := y - 1; i >= 0; i-- {
				scenicScore++
				if grid[i][x] >= tree {
					break
				}

			}

			up := scenicScore
			scenicScore = 0
			// Look down.
			for i := y + 1; i < len(grid); i++ {
				scenicScore++
				if grid[i][x] >= tree {
					break
				}

			}
			down := scenicScore
			scenicScore = 0

			// Look left.
			for i := x - 1; i >= 0; i-- {
				scenicScore++
				if grid[y][i] >= tree {
					break
				}

			}
			left := scenicScore
			scenicScore = 0

			// Look right.
			for i := x + 1; i < len(row); i++ {
				scenicScore++
				if grid[y][i] >= tree {
					break
				}

			}
			right := scenicScore

			calcScore := up * down * left * right
			// Update the maximum scenic score if necessary.
			if calcScore > maxScenicScore {
				maxScenicScore = calcScore
				maxScenicScoreX = x
				maxScenicScoreY = y
			}
		}
	}

	// Print the maximum scenic score and its coordinates.
	fmt.Printf("The maximum scenic score is %d at (%d, %d).\n", maxScenicScore, maxScenicScoreX, maxScenicScoreY)
}
