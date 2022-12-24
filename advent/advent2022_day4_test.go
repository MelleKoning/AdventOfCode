package advent

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestOverlappingSections(t *testing.T) {
	fileLines := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8", "12-23,5-24"}
	//fileLines, err := GetFileLines("inputdata/input2022day4.txt")
	//if err != nil {
	//	t.Fatalf("%v", err)
	//}

	overlappingPairs := 0

	for _, line := range fileLines {
		pair := strings.Split(line, ",")
		range1 := strings.Split(pair[0], "-")
		range2 := strings.Split(pair[1], "-")

		range1_0, err := strconv.Atoi(range1[0])
		if err != nil {
			t.Fatalf("shit %v", err)
		}
		range1_1, err := strconv.Atoi(range1[1])
		if err != nil {
			t.Fatalf("shit %v", err)
		}
		range2_0, err := strconv.Atoi(range2[0])
		if err != nil {
			t.Fatalf("shit %v", err)
		}
		range2_1, err := strconv.Atoi(range2[1])
		if err != nil {
			t.Fatalf("shit %v", err)
		}

		if (range1_0 <= range2_0 && range1_1 >= range2_1) || (range2_0 <= range1_0 && range2_1 >= range1_1) {
			fmt.Printf("overlap for %s, %s\n", range1, range2)
			overlappingPairs++
		} else {
			fmt.Printf("no overlap for %s, %s\n", range1, range2)
		}
	}

	fmt.Println("Number of overlapping pairs:", overlappingPairs)

}

func TestOverlappingPairsFinder(t *testing.T) {
	//fileLines := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8", "12-23,5-24"}
	fileLines, err := GetFileLines("inputdata/input2022day4.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	overlappingPairs := 0

	for _, line := range fileLines {
		pair := strings.Split(line, ",")
		range1 := strings.Split(pair[0], "-")
		range2 := strings.Split(pair[1], "-")

		start1 := GetNumberFromString(range1[0])
		end1 := GetNumberFromString(range1[1])
		start2 := GetNumberFromString(range2[0])
		end2 := GetNumberFromString(range2[1])

		if start1 <= end2 && start2 <= end1 {
			/*		if (start1 >= start2 && start1 <= end2 ||
					end1 >= start2 && end1 <= end2) ||
					(start2 >= start1 && start2 <= end1 ||
						end2 >= start1 && end2 <= end1) {*/
			fmt.Printf("overlap for %s, %s\n", range1, range2)
			overlappingPairs++
		} else {
			fmt.Printf("no overlap for %s, %s\n", range1, range2)
		}
	}

	fmt.Println("Number of overlapping pairs:", overlappingPairs)

}
