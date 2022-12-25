package advent

import (
	"fmt"
	"strings"
	"testing"
)

type AllSensorBeaconPairs struct {
	Pairs []*SensorBeaconPair
}

type SensorBeaconPair struct {
	SensorX int
	SensorY int
	BeaconX int
	BeaconY int
}

func (a *AllSensorBeaconPairs) Parse(line string) {
	//Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	pair := strings.Split(line, ":")
	sensorStr := pair[0]
	beaconStr := pair[1]
	SBPair := &SensorBeaconPair{
		SensorX: GetNumberFromString(sensorStr[strings.Index(sensorStr, "x")+2 : strings.Index(sensorStr, ",")]),
		SensorY: GetNumberFromString(sensorStr[strings.Index(sensorStr, "y")+2:]),
		BeaconX: GetNumberFromString(beaconStr[strings.Index(beaconStr, "x")+2 : strings.Index(beaconStr, ",")]),
		BeaconY: GetNumberFromString(beaconStr[strings.Index(beaconStr, "y")+2:]),
	}
	a.Pairs = append(a.Pairs, SBPair)
}

func (a *AllSensorBeaconPairs) CoveredLine(line int) int {
	// determine coverage of a line by going over all the pairs
	// and determine the area covered.
	return 0
}
func NewSensorBeaconPair(lines []string) *AllSensorBeaconPairs {
	allpairs := &AllSensorBeaconPairs{}

	for _, line := range lines {
		allpairs.Parse(line)
	}

	return allpairs
}
func TestSensorsAndBeaconsExampleTask1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day15example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	SensorBeaconPair := NewSensorBeaconPair(fileLines)

	fmt.Printf("%v", SensorBeaconPair)
}
