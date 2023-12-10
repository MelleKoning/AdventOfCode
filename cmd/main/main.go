package main

import (
	"fmt"

	"github.com/MelleKoning/AdventOfCode/advent"
	"github.com/MelleKoning/AdventOfCode/advent2023"
)

func main() {
	fileLines, err := advent.GetFileLines("advent2023/input/day05input.txt")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for _, line := range fileLines {
		fmt.Printf("%s", line)
	}

	lowestLocation := Year2023Day05Task02(fileLines)
	fmt.Printf("\nlowest location is %d\n", lowestLocation)

}

func Year2023Day05Task02(lines []string) uint64 {
	lowest := advent2023.Year2023Day05Task02(lines)

	//lowest location is 104070863
	return lowest

}
