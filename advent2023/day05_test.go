package advent2023

import (
	"fmt"
	"math"
	"testing"
)

func TestDay05Task1(t *testing.T) {

	lines := GetLinesFromFile("input/day05test.txt")

	// reading in the data...
	almanac := &Almanac{}
	almanac.ReadAlmanacSeeds(lines)

	lowestlocation := uint64(math.MaxUint64)
	for _, seed := range almanac.Seedlist {
		// lookup the soil number for this seed
		// but initialize with its own number in case no map was found
		// soil := seed
		location := almanac.SeedToLocation(seed)
		fmt.Printf("location %d,", location)

		if location < lowestlocation {
			lowestlocation = location
		}
		fmt.Printf("\n")
	}

	fmt.Printf("lowest location is %d\n", lowestlocation)
}

func TestDay05Task2(t *testing.T) {

	lines := GetLinesFromFile("input/day05test.txt")

	lowest := Year2023Day05Task02(lines)

	//lowest location is 104070863 104070862
	fmt.Printf("\nlowest location is %d\n", lowest)
}
