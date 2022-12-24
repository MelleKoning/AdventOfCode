package advent

import (
	"fmt"
	"strings"
	"testing"
)

type Direction int

const (
	DIR_UP    Direction = 0
	DIR_DOWN  Direction = 1
	DIR_LEFT  Direction = 2
	DIR_RIGHT Direction = 3
)

var (
	DirectionText = map[Direction]string{
		DIR_UP:    "UP",
		DIR_DOWN:  "DOWN",
		DIR_LEFT:  "LEFT",
		DIR_RIGHT: "RIGHT"}
)

type WalkerCommand struct {
	Aim   Direction // Direction of the command
	Steps int       // number of steps for this command
}

type WalkPosition struct {
	X, Y int
}
type Walker struct {
	CurrentPos WalkPosition
	LastPos    WalkPosition
}

func TestTailWalker_1(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day9example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// A Head (H) is walking in a direction given by some commands.
	// A Tail (T) is lagging behind. The tail will only move when it is
	// not adjacent to the head anymore. Adjacent is defined as the
	// coordinates of the Tail are close to the Head.
	// Ccoordinates can only differ by 1, so if the Head is at 1,1 and Tail at 2,1, the tail does not have to move.
	// When the Head moves "left" and reaches 0,1, the tail at 2,1 also has to move. Where does the tail move?
	// Basically to the last position of the head! :-)
	// So the tail in the above example ends up at 1,1...
	// we imagine the grid is having a coordinate of 0,0 (x,y) as start
	// and moving left reduces x, moving right increases x by one
	// moving up increases y by one, moving down decreases y by one

	startPos := WalkPosition{X: 0, Y: 0}
	Head := &Walker{CurrentPos: startPos, LastPos: startPos}
	Tail := &Walker{CurrentPos: startPos, LastPos: startPos}

	// we keep track of all locations the Tail has been with a map
	TailVisited := make(map[WalkPosition]bool)
	TailVisited[Tail.CurrentPos] = true
	for _, line := range fileLines {
		headwalkCmd := GetWalkerCommandFromLine(line)
		fmt.Printf("%v %d\n", DirectionText[headwalkCmd.Aim], headwalkCmd.Steps)
		for steps := 1; steps <= headwalkCmd.Steps; steps++ {

			Head.LastPos = Head.CurrentPos                   // keep track of H lastpos
			Head.CurrentPos = Head.MoveHead(headwalkCmd.Aim) // new position of Head

			Tail.CurrentPos = MoveTailWithHead(Tail, Head)
			TailVisited[Tail.CurrentPos] = true
			fmt.Printf("Head at (%d,%d) Tail at (%d,%d)\n", Head.CurrentPos.X, Head.CurrentPos.Y, Tail.CurrentPos.X, Tail.CurrentPos.Y)
		}

	}

	fmt.Printf("Tail has visited %d unique locations", len(TailVisited)) // 6090
}

func TestTailWalker_2(t *testing.T) {

	//fileLines, err := GetFileLines("inputdata/input2022day9.txt")
	fileLines, err := GetFileLines("inputdata/input2022day9.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// A Head (H) is walking in a direction given by some commands.
	// A Tail (1,2,3,4,5,6,7,8,9) is lagging behind. The tail will only move when it is
	// not adjacent to the head anymore. Adjacent is defined as the
	// coordinates of the Tail are close to the Head.
	// Ccoordinates can only differ by 1, so if the Head is at 1,1 and Tail(1) at 2,1, the tail does not have to move.
	// When the Head moves "left" and reaches 0,1, the tail at 2,1 also has to move. Where does the tail move?
	// Basically to the last position of the head! :-)
	// So the tail(1) in the above example ends up at 1,1...
	// We imagine the grid is having a coordinate of 0,0 (x,y) as start
	// and moving left reduces x, moving right increases x by one
	// moving up increases y by one, moving down decreases y by one
	// However, for the tail that is behind tail(1), the tail(2) does not end up at the
	// last position of the previous tail(1).
	// Where does it end up?
	// Depends on the movement if the previous tail(1) position.
	// If there is "jump" movement (both x and y are different) then tail(2) makes the same
	// jumpmove, otherwise it just moves in the same direction as tail(1)
	// the "Jump" can be determined as the movement of the previous tail x,y difference.

	startPos := WalkPosition{X: 12, Y: 6} // does not matter where we start
	Rope := make(map[int]*Walker)         // have a map for the rope 0..9
	for r := 0; r < 10; r++ {
		fmt.Printf("init %d", r)
		Rope[r] = &Walker{CurrentPos: startPos, LastPos: startPos}
	}

	// we keep track of all locations the Tail has been with a map
	TailVisited := make(map[WalkPosition]bool)
	TailVisited[Rope[9].CurrentPos] = true

	Head := Rope[0]

	for _, line := range fileLines {
		headwalkCmd := GetWalkerCommandFromLine(line)
		fmt.Printf("%v %d\n", DirectionText[headwalkCmd.Aim], headwalkCmd.Steps)
		for steps := 1; steps <= headwalkCmd.Steps; steps++ {

			Head.LastPos = Head.CurrentPos                   // keep track of H lastpos
			Head.CurrentPos = Head.MoveHead(headwalkCmd.Aim) // new position of Head
			fmt.Printf("Head %d at (%d,%d) ", 0, Head.CurrentPos.X, Head.CurrentPos.Y)

			// For the next Rope parts, when not adjacent, it depends how move is done
			for r := 1; r < 10; r++ {
				Rope[r].LastPos = Rope[r].CurrentPos
				Rope[r].CurrentPos = MoveKnotWithPreviousTail(Rope[r], Rope[r-1])
				fmt.Printf("Knot %d at (%d,%d) ", r, Rope[r].CurrentPos.X, Rope[r].CurrentPos.Y)
			}
			fmt.Printf("\n")
			//PrintGrid(Rope)
			TailVisited[Rope[9].CurrentPos] = true

		}
	}

	fmt.Printf("Tail 9 has visited %d unique locations", len(TailVisited))
}

func PrintGrid(snake map[int]*Walker) {
	for y := 20; y >= 0; y-- {
		for x := 0; x < 27; x++ {
			printed := false
			for key, _ := range snake {
				if snake[key].CurrentPos.X == x &&
					snake[key].CurrentPos.Y == y {
					if key == 0 {
						fmt.Print("H")
					} else {
						fmt.Print(key)
					}
					printed = true
					break
				}
			}
			if !printed {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func MoveTailWithHead(tail, head *Walker) WalkPosition {
	// if the tails position is too far from head
	// we move the tail to the previous head position
	if tail.CurrentPos.X > head.CurrentPos.X+1 ||
		tail.CurrentPos.Y > head.CurrentPos.Y+1 ||
		tail.CurrentPos.X < head.CurrentPos.X-1 ||
		tail.CurrentPos.Y < head.CurrentPos.Y-1 {

		return head.LastPos // move to heads previous position
	}
	return tail.CurrentPos // stay at location
}

func MoveKnotWithPreviousTail(knot, prevKnot *Walker) WalkPosition {
	diffX := prevKnot.CurrentPos.X - knot.CurrentPos.X
	diffY := prevKnot.CurrentPos.Y - knot.CurrentPos.Y
	if diffX == 2 || diffX == -2 || diffY == 2 || diffY == -2 {
		if diffX > 0 {
			knot.CurrentPos.X++
		}
		if diffX < 0 {
			knot.CurrentPos.X--
		}
		if diffY > 0 {
			knot.CurrentPos.Y++
		}
		if diffY < 0 {
			knot.CurrentPos.Y--
		}
	}

	return knot.CurrentPos
}

// MoveHead moves the head just one position in the direction
func (w Walker) MoveHead(dir Direction) WalkPosition {
	switch dir {
	case DIR_DOWN:
		w.CurrentPos.Y -= 1
	case DIR_UP:
		w.CurrentPos.Y += 1
	case DIR_LEFT:
		w.CurrentPos.X -= 1
	case DIR_RIGHT:
		w.CurrentPos.X += 1
	}
	return w.CurrentPos
}
func GetWalkerCommandFromLine(line string) WalkerCommand {
	// get instruction from the line
	HeadWalker := WalkerCommand{}
	cmd := strings.Split(line, " ")
	switch cmd[0] {
	case "U":
		{ // upwards direction
			HeadWalker.Aim = DIR_UP
			HeadWalker.Steps = GetNumberFromString(cmd[1])
		}
	case "D":
		{
			HeadWalker.Aim = DIR_DOWN
			HeadWalker.Steps = GetNumberFromString(cmd[1])
		}
	case "L":
		{
			HeadWalker.Aim = DIR_LEFT
			HeadWalker.Steps = GetNumberFromString(cmd[1])
		}
	case "R":
		{
			HeadWalker.Aim = DIR_RIGHT
			HeadWalker.Steps = GetNumberFromString(cmd[1])
		}
	}
	return HeadWalker
}
