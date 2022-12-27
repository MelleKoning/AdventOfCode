package advent

import (
	"math"
	"strings"
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

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func (sb *SensorBeaconPair) ManhattenDistance() int {
	// 0,0 1,1 : 2
	// 0,0,1,2 : 3
	// 0,0,2,2 : 4
	// 0,0,1,3 : 4
	// 0,0,1,4 : 5
	// 0,0,3,3 : 6
	// 0,0,4,4 : 8
	return (Abs(sb.BeaconX - sb.SensorX)) + (Abs(sb.BeaconY - sb.SensorY))
}

func (sb *SensorBeaconPair) CoverageOnRow(row int) int {
	/*
		               1    1    2    2
		     0    5    0    5    0    5
		-2 ..........#.................
		-1 .........###................
		 0 ....S...#####...............
		 1 .......#######........S.....
		 2 ......#########S............
		 3 .....###########SB..........
		 4 ....#############...........
		 5 ...###############..........
		 6 ..#################.........
		 7 .#########S#######S#........
		 8 ..#################.........
		 9 ...###############..........
		10 ....B############...........
		11 ..S..###########............
		12 ......#########.............
		13 .......#######..............
		14 ........#####.S.......S.....
		15 B........###................
		16 ..........#SB...............
		17 ................S..........B
		18 ....S.......................
		19 ............................
		20 ............S......S........
		21 ............................
		22 .......................B....
	*/
	// The sensor at 8,7 defines the point where coverage is widthest,
	// while the closest beacon at 2,10 gives it a manhattan distance of
	// (8-2)+(10-7) is 6+3=9. This value 9 provides the distance from the
	// sensors row (7) to calculate the coverage for a row as:
	// 9*2+1 - Abs(row-distance) -> 9 * 2+1 - 0 = 19
	// 7 -> 19
	// 6 -> (9)*2+1 -(2*distance) ->  17
	// 5 -> 9*2+1 - (2*2) -> 19-4 = 15
	// 4 -> 9*2+1 - (2*3) -> 19-6 = 13
	// 0 -> 9*2+1 - (2*7) -> 19-14 = 5

	mandistance := sb.ManhattenDistance()

	// if we are too far from the requested row we should return 0
	if (sb.SensorY - mandistance) > row {
		return 0
	}
	if (sb.SensorY + mandistance) < row {
		return 0
	}
	return mandistance*2 + 1 - (2 * Abs(sb.SensorY-row))

}

func (sb *SensorBeaconPair) CoverageOfRowFromTo(row int) (bool, int, int) {
	// the starting coverage can be calculated by subtracting
	// SensorY from the starting point of the coverage which is simply
	// SensorX - (SensorX-row)
	// the ending coverage by adding that difference
	// If there is no coverage we should return that fact
	if sb.CoverageOnRow(row) == 0 {
		return false, 0, 0
	}
	mhDistance := sb.ManhattenDistance()
	return true, sb.SensorX - mhDistance + Abs(sb.SensorY-row),
		sb.SensorX + mhDistance - Abs(sb.SensorY-row)
}

func (a *AllSensorBeaconPairs) MinMaxBeaconX() (int, int) {
	// determine the lowest/highest X beacon points
	minBeaconX, maxBeaconX := math.MaxInt, math.MinInt
	for _, beacon := range a.Pairs {
		if beacon.BeaconX < minBeaconX {
			minBeaconX = beacon.BeaconX
		}
		if beacon.BeaconX > maxBeaconX {
			maxBeaconX = beacon.BeaconX
		}
	}
	return minBeaconX, maxBeaconX
}

func (a *AllSensorBeaconPairs) FindGapInRow(row int, maxX int) (bool, int) {

	// We can OR all results together to form a bitmask..

	mask := make(map[int]bool)
	for _, pair := range a.Pairs {
		covered, min, max := pair.CoverageOfRowFromTo(row)
		if covered {

			for n := min; n <= max; n++ {
				mask[n] = true
			}
		}
	}

	// now see if there is a gap between maxA and minB
	for n := 0; n <= maxX; n++ {
		_, ok := mask[n]
		if !ok {
			return true, n
		}

	}
	return false, -1
}
func (a *AllSensorBeaconPairs) CoveredLine(row int) (int, int) {
	// determine coverage of a line by going over all the pairs
	// and determine the area covered.

	// determine the lowest/highest X beacon points
	//minBeaconX, maxBeaconX := a.MinMaxBeaconX()

	// determine the minimum and maximum Coverage points
	minRowCoverage := math.MaxInt
	maxRowCoverage := math.MinInt

	for _, pair := range a.Pairs {
		covered, min, max := pair.CoverageOfRowFromTo(row)
		if covered {
			if min < minRowCoverage {
				minRowCoverage = min
			}
			if max > maxRowCoverage {
				maxRowCoverage = max
			}
		}
	}

	return minRowCoverage, maxRowCoverage
}

func NewSensorBeaconPair(lines []string) *AllSensorBeaconPairs {
	allpairs := &AllSensorBeaconPairs{}

	for _, line := range lines {
		allpairs.Parse(line)
	}

	return allpairs
}
