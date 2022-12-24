package advent

import (
	"fmt"
	"testing"
)

func TestDay3GammaEpsilonRates(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/inputday3.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	gammarates := make([]int, 12)
	epsilonrates := make([]int, 12)

	numberofLines := len(fileLines)
	for _, line := range fileLines {

		for n := 0; n <= len(line)-1; n++ { // char by char...
			char := line[n]
			if string(char) == "0" {
				epsilonrates[n] += 1
			}
			if string(char) == "1" {
				gammarates[n] += 1
			}
		}
	}

	gammacode := ""
	for n := 0; n <= 11; n++ {
		if gammarates[n] > numberofLines/2 {
			// marker of n is 1, otherwise it is 0
			gammacode += "1"
		} else {
			gammacode += "0"
		}
	}

	fmt.Println(gammacode)
}
