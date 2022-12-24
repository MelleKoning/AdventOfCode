package advent

import (
	"fmt"
	"strings"
	"testing"
)

func TestRucksacks(t *testing.T) {
	// List of contents from six rucksacks
	fileLines, err := GetFileLines("inputdata/input2022day3.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	rucksacks := fileLines

	// Sum of priorities
	sum := 0

	// Iterate over each rucksack
	for _, rucksack := range rucksacks {
		// Split the rucksack into two compartments
		compartment1 := rucksack[:len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]

		// Iterate over each character in the first compartment
		for _, char := range compartment1 {
			if strings.Contains(compartment2, string(char)) {
				fmt.Printf("found char %s in %s, %s\n", string(char), compartment1, compartment2)
				if char >= 'a' && char <= 'z' {
					sum += int(char - 'a' + 1)
				} else if char >= 'A' && char <= 'Z' {
					sum += int(char - 'A' + 27)
				}
				break
			}
		}

	}
	// Print the sum
	fmt.Println(sum)
}

func TestFindBadge(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day3.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	sum := 0
	for line := 0; line < len(fileLines); line += 3 {
		line1 := fileLines[line]
		line2 := fileLines[line+1]
		line3 := fileLines[line+2]

		// Find the badge of the group
		badge := ""
		for _, c := range line1 {
			if strings.Contains(line2, string(c)) && strings.Contains(line3, string(c)) {
				badge = string(c)
				break
			}
		}

		char := rune(badge[0])
		// Find the priority of the badge
		if char >= 'a' && char <= 'z' {
			sum += int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			sum += int(char - 'A' + 27)
		}

		/*
			// Calculate the priority of the badge
			priority := 0
			if strings.ToLower(badge) == badge {
				priority = int(badge[0]) - 96
			} else {
				priority = int(badge[0]) - 64 + 26
			}
		*/

	}

	fmt.Println("The sum of the priorities of the item types is:", sum)
}
