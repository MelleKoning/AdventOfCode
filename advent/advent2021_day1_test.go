package advent

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1AmountLargerThanPrevious(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/inputday1.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	var numbers []int

	for _, line := range fileLines {
		intVar, err := strconv.Atoi(line)
		if err != nil {
			t.Fatalf("shit %v", err)
		}
		numbers = append(numbers, intVar)
	}

	increments := GetIncrements(numbers)

	fmt.Printf(" we are seeing %d increments\n", increments)

	assert.Equal(t, 1215, increments)
}

func TestDay1SlidingAverage(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/inputday1.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	var numbers []int

	for _, line := range fileLines {
		intVar, err := strconv.Atoi(line)
		if err != nil {
			t.Fatalf("shit %v", err)
		}
		numbers = append(numbers, intVar)
	}

	// determine averages in a new list of averageResults
	// average is the sliding average of the number,
	// and the next two numbers
	var averageResults []int
	for idx, n := range numbers {

		avg := n
		if idx < len(numbers)-2 {
			avg = avg + numbers[idx+1]
			avg = avg + numbers[idx+2]
		}
		averageResults = append(averageResults, avg)
	}

	increments := GetIncrements(averageResults[0 : len(averageResults)-2])
	assert.Equal(t, 1150, increments)

}
