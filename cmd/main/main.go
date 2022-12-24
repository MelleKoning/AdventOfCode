package main

import (
	"fmt"

	"github.com/MelleKoning/AdventOfCode/advent"
)

func main() {
	fileLines, _ := advent.GetFileLines("inputdata/inputday3.txt")
	for _, line := range fileLines {
		fmt.Printf("%s", line)
	}
}
