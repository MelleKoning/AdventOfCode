package advent

import (
	"fmt"
	"testing"
)

func TestFindFourCharacters(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day6.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Define the datastream buffer
	buffer := fileLines[0]

	// Loop over the characters in the buffer
	for i := 0; i < len(buffer)-3; i++ {
		// Check if the last four characters are all different
		if buffer[i] != buffer[i+1] && buffer[i] != buffer[i+2] && buffer[i] != buffer[i+3] &&
			buffer[i+1] != buffer[i+2] && buffer[i+1] != buffer[i+3] && buffer[i+2] != buffer[i+3] {
			// Print the number of characters processed
			fmt.Printf("%s%s%s%s: %d", string(buffer[i]), string(buffer[i+1]), string(buffer[i+2]), string(buffer[i+3]), i+4)
			break
		}
	}

	// answer is 1766
}

func TestFind14MessageMarker(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day6.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Define the datastream buffer
	buffer := fileLines[0]

	// Loop over the characters in the buffer
	for i := 0; i < len(buffer)-13; i++ {
		// Check if the last 14 characters are all different
		different := true
		for j := 0; j < 14; j++ {
			for k := j + 1; k < 14; k++ {
				if buffer[i+j] == buffer[i+k] {
					different = false
					break
				}
			}
			if !different {
				break
			}
		}
		if different {
			// Print the number of characters processed
			fmt.Printf("%s: %d", string(buffer[i:i+14]), i+14)
			// puzzle answer was 2383
			break
		}
	}
}
