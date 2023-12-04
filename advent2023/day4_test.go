package advent2023

import (
	"fmt"
	"strings"
	"testing"
)

func TestDay4Task1(t *testing.T) {

	lines := GetLinesFromFile("input/day4input.txt")

	total := 0
	for _, l := range lines {
		count := getWinningCountForLine(l)
		score := convertCountToScore(count)
		fmt.Printf("%s, %d\n", l, score)
		total += score
	}

	// score was 17782
	fmt.Printf("Total %d\n", total)

}

func TestConvertScore(t *testing.T) {
	fmt.Printf("%d", convertCountToScore(3))
}

func convertCountToScore(count int) int {

	score := 0
	for n := 1; n <= count; n++ {
		if n == 1 {
			score = 1
		} else {
			score *= 2
		}
	}
	return score
}
func getWinningCountForLine(l string) int {

	parts := strings.Split(l, "|")

	// get everything after :
	partOne := strings.Split(strings.TrimSpace(parts[0]), ":")

	winningNumbers := strings.Split(strings.TrimSpace(partOne[1]), " ")

	ourNumbers := strings.Split(parts[1], " ")

	count := 0
	// now check..
	for _, winning := range winningNumbers {
		win := strings.TrimSpace(winning)
		if win == "" {
			continue
		}
		for _, ours := range ourNumbers {
			thisNumber := strings.TrimSpace(ours)
			if thisNumber == "" {
				continue
			}
			if thisNumber == win {

				count += 1
			}
		}
	}

	return count
}

func TestDay4Task2(t *testing.T) {

	lines := GetLinesFromFile("input/day4input.txt")

	cards := make([]int, len(lines))

	// initialize
	for idx, _ := range lines {
		cards[idx] = 1
	}

	for idx, l := range lines {
		wincount := getWinningCountForLine(l)
		// increase the count for all the next lines
		counttil := idx + wincount
		// do not count past the end
		if counttil >= len(lines) {
			counttil = len(lines) - 1
		}
		for n := idx + 1; n <= counttil; n++ {
			cards[n] += cards[idx]
		}
	}

	// add up all values
	total := 0
	for idx, _ := range lines {
		total += cards[idx]
	}
	// I got 8477787 for my puzzle input
	fmt.Printf("total:%d\n", total)
}
