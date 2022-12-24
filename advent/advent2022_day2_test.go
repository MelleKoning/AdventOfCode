package advent

import (
	"fmt"
	"strings"
	"testing"
)

const (
	myHandRock        = "X"
	myHandPaper       = "Y"
	myHandScissors    = "Z"
	theirHandRock     = "A"
	theirHandPaper    = "B"
	theirHandScissors = "C"

	strategyLose       = "X"
	strategyDraw       = "Y"
	strategyWin        = "Z"
	shapeScoreRock     = 1 // A rock oppenent, X rock me
	shapeScorePaper    = 2 // B paper, Y paper
	shapeScoreScissors = 3 // C, Z Scissors

)

func TestStrategyRockPaperScissors(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day2.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	shapeScoreRock := 1     // A rock oppenent, X rock me
	shapeScorePaper := 2    // B paper, Y paper
	shapeScoreScissors := 3 // C, Z Scissors

	totalScore := 0
	//numberofLines := len(fileLines)
	for idx, line := range fileLines {

		play := strings.Split(line, " ")
		switch play[1] {
		case myHandRock:
			totalScore += shapeScoreRock
			totalScore += PlayRPS(play[1], play[0])
		case myHandPaper:
			totalScore += shapeScorePaper
			totalScore += PlayRPS(play[1], play[0])
		case myHandScissors:
			totalScore += shapeScoreScissors
			totalScore += PlayRPS(play[1], play[0])
		default:
			fmt.Printf("line %d, chars: %v", idx, play)
		}
	}
	fmt.Println(totalScore) // 13526 not good?? 13526
}

func TestRockPaperScissorsRound2(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day2.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	totalScore := 0
	//numberofLines := len(fileLines)
	for _, line := range fileLines {

		play := strings.Split(line, " ")
		totalScore += OutComeOnStrategy(play[0], play[1])
	}
	fmt.Println(totalScore) // 16868 is too high
}

func PlayRPS(me string, they string) int {
	if me == myHandRock && they == theirHandRock ||
		me == myHandPaper && they == theirHandPaper ||
		me == myHandScissors && they == theirHandScissors {
		return 3 // draw
	}
	if me == myHandRock && they == theirHandScissors ||
		me == myHandPaper && they == theirHandRock ||
		me == myHandScissors && they == theirHandPaper {
		return 6 // win
	}
	return 0 // loss
}

func OutComeOnStrategy(opponentPlay, strategy string) int {
	theScore := 0
	if strategy == strategyWin {
		theScore = 6
		// we add the score of our choice
		switch opponentPlay {
		case theirHandRock:
			return theScore + shapeScorePaper
		case theirHandPaper:

			return theScore + shapeScoreScissors
		case theirHandScissors:
			return theScore + shapeScoreRock
		}
	}
	if strategy == strategyDraw {
		theScore = 3
		switch opponentPlay {
		case theirHandRock:
			return theScore + shapeScoreRock
		case theirHandPaper:
			return theScore + shapeScorePaper

		case theirHandScissors:
			return theScore + shapeScoreScissors
		}
		return 3
	}

	switch opponentPlay {
	case theirHandRock:
		return shapeScoreScissors // loss
	case theirHandPaper:
		return shapeScoreRock // loss

	case theirHandScissors:
		return shapeScorePaper // loss
	}

	return 0 // strategy loose
}
