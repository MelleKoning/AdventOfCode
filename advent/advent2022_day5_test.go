package advent

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMoveTowers(t *testing.T) {
	// fileLines := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8", "12-23,5-24"}
	fileLines, err := GetFileLines("inputdata/input2022day5.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	towers := make(map[int]string, 9)
	towers[1] = "PFMQWGRT" // bottom to top
	towers[2] = "HFR"
	towers[3] = "PZRVGHSD"
	towers[4] = "QHPBFWG"
	towers[5] = "PSMJH"
	towers[6] = "MZTHSRPL"
	towers[7] = "PTHNML"
	towers[8] = "FDQR"
	towers[9] = "DSCNLPH"

	fmt.Printf("%v\n", towers)
	for _, line := range fileLines {
		if !strings.Contains(line, "move") {
			continue
		}

		fmt.Println(line)
		lc := strings.Split(line, " ")

		amount, err := strconv.Atoi(lc[1])
		if err != nil {
			t.Fatalf("shit %v", err)
		}

		from, err := strconv.Atoi(lc[3])
		if err != nil {
			t.Fatalf("shit %v", err)
		}

		to, err := strconv.Atoi(lc[5])
		if err != nil {
			t.Fatalf("shit %v", err)
		}

		for n := 1; n <= amount; n++ {
			// take top crate from
			topfrom := towers[from][len(towers[from])-1]
			// adjust from by reducing slice
			towers[from] = towers[from][0 : len(towers[from])-1]
			// add to
			towers[to] = towers[to] + string(topfrom)
		}

		fmt.Printf("%v\n", towers)
	}

	fmt.Printf("END: %v\n", towers)
	// TPGVQPFDH

}

func TestMoveTowersMultipleAtonce(t *testing.T) {
	// fileLines := []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8", "12-23,5-24"}
	fileLines, err := GetFileLines("inputdata/input2022day5.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	towers := map[int]string{
		1: "PFMQWGRT", // bottom to top
		2: "HFR",
		3: "PZRVGHSD",
		4: "QHPBFWG",
		5: "PSMJH",
		6: "MZTHSRPL",
		7: "PTHNML",
		8: "FDQR",
		9: "DSCNLPH"}

	fmt.Printf("%v\n", towers)
	for _, line := range fileLines {
		if !strings.Contains(line, "move") {
			continue
		}

		fmt.Println(line)
		lc := strings.Split(line, " ")

		amount := GetNumberFromString(lc[1])
		from := GetNumberFromString(lc[3])
		to := GetNumberFromString(lc[5])

		crates := towers[from][len(towers[from])-amount:]
		towers[to] = towers[to] + crates
		towers[from] = towers[from][:len(towers[from])-amount]

		// take top crates from
		//topfrom := towers[from][len(towers[from])-amount : len(towers[from])]
		// adjust from by reducing slice
		//towers[from] = towers[from][0 : len(towers[from])-amount]
		// add to
		//towers[to] = towers[to] + string(topfrom)

		fmt.Printf("%v\n", towers)
	}

	fmt.Printf("END: %v\n", towers) // DMRDFRHHH

}
