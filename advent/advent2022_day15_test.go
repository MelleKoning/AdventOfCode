package advent

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorsAndBeaconsExampleTask1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day15example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	SensorBeaconPair := NewSensorBeaconPair(fileLines)

	for _, pair := range SensorBeaconPair.Pairs {
		fmt.Printf("%v\n", pair)
	}

	minBeaconX, maxBeaconX := SensorBeaconPair.MinMaxBeaconX()
	fmt.Printf("minBeaconX:%d, maxBeaconX:%d\n", minBeaconX, maxBeaconX)

	minRowCoverage, maxRowCoverage := SensorBeaconPair.CoveredLine(10)
	fmt.Printf("minX, maxX covered: %d,%d\n", minRowCoverage, maxRowCoverage)

	fmt.Printf("There are %d positions where a beacon can not be present", Abs(maxRowCoverage-minRowCoverage))
}

func TestSensorsAndBeacons_InputTask1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day15.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	SensorBeaconPair := NewSensorBeaconPair(fileLines)

	for _, pair := range SensorBeaconPair.Pairs {
		fmt.Printf("%v\n", pair)
	}

	minBeaconX, maxBeaconX := SensorBeaconPair.MinMaxBeaconX()
	fmt.Printf("minBeaconX:%d, maxBeaconX:%d\n", minBeaconX, maxBeaconX)

	minRowCoverage, maxRowCoverage := SensorBeaconPair.CoveredLine(2000000)
	fmt.Printf("minX, maxX covered: %d,%d\n", minRowCoverage, maxRowCoverage)

	// The answer here is 5166077
	fmt.Printf("There are %d positions where a beacon can not be present", Abs(maxRowCoverage-minRowCoverage))
}

func TestSearchDistressSignal_Task2(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day15.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	SensorBeaconPair := NewSensorBeaconPair(fileLines)

	for _, pair := range SensorBeaconPair.Pairs {
		// fmt.Printf("%v\n", pair)
		cov, min, max := pair.CoverageOfRowFromTo(11)
		fmt.Printf("Coverage %t between %d,%d\n", cov, min, max)
	}

	minBeaconX, maxBeaconX := SensorBeaconPair.MinMaxBeaconX()
	fmt.Printf("minBeaconX:%d, maxBeaconX:%d\n", minBeaconX, maxBeaconX)

	minRowCoverage, maxRowCoverage := SensorBeaconPair.CoveredLine(10)
	fmt.Printf("minX, maxX covered: %d,%d\n", minRowCoverage, maxRowCoverage)

	for row := 0; row <= 4000000; row++ { // This takes far too loong.. there must be a smarter way than this
		findGap, gapNumber := SensorBeaconPair.FindGapInRow(row, 4000000)
		if findGap {
			fmt.Printf("Gap found %t at row %d, x:%d\n", findGap, row, gapNumber)
		}
	}

	//fmt.Printf("Gap found %t at %d\n", findGap, gapNumber)
}

func TestSearchDistressSignalExampleTask2(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day15example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	SensorBeaconPair := NewSensorBeaconPair(fileLines)

	for _, pair := range SensorBeaconPair.Pairs {
		// fmt.Printf("%v\n", pair)
		cov, min, max := pair.CoverageOfRowFromTo(11)
		fmt.Printf("Coverage %t between %d,%d\n", cov, min, max)
	}

	minBeaconX, maxBeaconX := SensorBeaconPair.MinMaxBeaconX()
	fmt.Printf("minBeaconX:%d, maxBeaconX:%d\n", minBeaconX, maxBeaconX)

	minRowCoverage, maxRowCoverage := SensorBeaconPair.CoveredLine(10)
	fmt.Printf("minX, maxX covered: %d,%d\n", minRowCoverage, maxRowCoverage)

	findGap, gapNumber := SensorBeaconPair.FindGapInRow(11, 20)

	fmt.Printf("Gap found %t at %d\n", findGap, gapNumber)

	assert.Equal(t, 14, gapNumber) // gap should be at X 14
}

func TestManhattenDistance(t *testing.T) {
	sbPair := &SensorBeaconPair{
		SensorX: 1,
		SensorY: 1,
		BeaconX: 2,
		BeaconY: 2,
	}

	assert.Equal(t, 2, sbPair.ManhattenDistance())

	sbPair = &SensorBeaconPair{
		SensorX: 4,
		SensorY: 4,
		BeaconX: 10,
		BeaconY: 12,
	}

	assert.Equal(t, 14, sbPair.ManhattenDistance())
}

func TestRowCoverageOfPair(t *testing.T) {
	// 7 -> 19
	// 6 -> (9)*2+1 -(2*distance) ->  17
	// 5 -> 9*2+1 - (2*2) -> 19-4 = 15
	// 4 -> 9*2+1 - (2*3) -> 19-6 = 13
	// 0 -> 9*2+1 - (2*7) -> 19-14 = 5
	sbPair := &SensorBeaconPair{
		SensorX: 8,
		SensorY: 7,
		BeaconX: 2,
		BeaconY: 10,
	}
	assert.Equal(t, 19, sbPair.CoverageOnRow(7))
	assert.Equal(t, 15, sbPair.CoverageOnRow(5))
	assert.Equal(t, 13, sbPair.CoverageOnRow(4))
	assert.Equal(t, 5, sbPair.CoverageOnRow(0))
	assert.Equal(t, 1, sbPair.CoverageOnRow(-2))
	assert.Equal(t, 1, sbPair.CoverageOnRow(16))
	assert.Equal(t, 0, sbPair.CoverageOnRow(17))
	assert.Equal(t, 0, sbPair.CoverageOnRow(-3))

	// coverageFromTo
	b, from, to := sbPair.CoverageOfRowFromTo(7)
	assert.True(t, b)
	assert.Equal(t, -1, from)
	assert.Equal(t, 17, to)

	b, from, to = sbPair.CoverageOfRowFromTo(5)
	assert.True(t, b)
	assert.Equal(t, 1, from)
	assert.Equal(t, 15, to)

	b, from, to = sbPair.CoverageOfRowFromTo(10)
	assert.True(t, b)
	assert.Equal(t, 2, from)
	assert.Equal(t, 14, to)

	b, _, _ = sbPair.CoverageOfRowFromTo(17)
	assert.False(t, b)
}
