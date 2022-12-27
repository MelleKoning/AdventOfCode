package advent

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindElvesCarriesMostFood(t *testing.T) {

	lines, err := GetFileLines("inputdata/input2022day1.txt")
	if err != nil {
		t.Fatalf("can not read file")
	}

	maxelf := 0
	elfcounter := 0
	for _, l := range lines {

		if l == "" {
			elfcounter = 0
			continue
		}

		intVar, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("%v", intVar)
			t.Fatalf("intvar err %v", err)
		}
		elfcounter += intVar
		if elfcounter > maxelf {
			maxelf = elfcounter
		}
	}

	fmt.Printf("maxelf:%v", maxelf)
	assert.Greater(t, maxelf, -1)
}

func TestFindTop3Elves(t *testing.T) {

	lines, err := GetFileLines("inputdata/input2022day1.txt")
	if err != nil {
		t.Fatalf("can not read file")
	}

	maxelf := make(map[int]int, 3) // keep top 3 elves
	maxelf[0] = 0
	maxelf[1] = 0
	maxelf[2] = 0
	elfnumber := 0
	elfcalorie := 0
	for _, l := range lines {

		if l == "" {
			maxelf[elfnumber] = elfcalorie
			elfcalorie = 0
			elfnumber += 1
			continue
		}

		intVar, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("%v", intVar)
			t.Fatalf("shit %v", err)
		}
		elfcalorie += intVar

	}
	/*
	   	35 66186
	       108 65638
	       209 64980
	*/
	keys := make([]int, 0, len(maxelf))

	for key := range maxelf {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return maxelf[keys[i]] > maxelf[keys[j]]
	})

	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, maxelf[k])
	}

	top3 := 0
	top3count := 0
	for _, k := range keys {
		top3count += maxelf[k]
		top3 += 1
		fmt.Println(top3, top3count)
		if top3 >= 3 {
			break
		}

	}
	fmt.Println(top3count)

	assert.Equal(t, 196804, top3count)

}
