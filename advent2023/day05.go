package advent2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type GardenerMap struct {
	// 52 50 48
	// source 50 with length 48 (50..98) maps to (52..99)
	DestinationStart uint64
	SourceStart      uint64
	RangeLength      uint64
}

type Almanac struct {
	Seedlist   []uint64
	SeedRanges []SeedRange
	AlmanacMap map[string][]GardenerMap
}

func (a *Almanac) ReadAlmanacSeeds(lines []string) {
	a.AlmanacMap = make(map[string][]GardenerMap, 0)

	// reading in the data...
	for idx, l := range lines {
		if strings.Contains(l, "seeds:") {
			a.Seedlist = getSeedList(l)
			a.SeedRanges = getSeedRanges(l)
		}
		if strings.Contains(l, "seed-to-soil map:") {
			a.AlmanacMap["seed-to-soil"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "soil-to-fertilizer map:") {
			a.AlmanacMap["soil-to-fertilizer"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "fertilizer-to-water map:") {
			a.AlmanacMap["fertilizer-to-water"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "water-to-light map:") {
			a.AlmanacMap["water-to-light"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "light-to-temperature map:") {
			a.AlmanacMap["light-to-temperature"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "temperature-to-humidity map:") {
			a.AlmanacMap["temperature-to-humidity"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "humidity-to-location map:") {
			a.AlmanacMap["humidity-to-location"] = GetMapValuesFromLines(idx+1, lines)
		}
	}

	//fmt.Printf("seedlist: %v\nGardenerMaps:\n%v\n", seedlist, Almanac)

	//return Almanac, seedlist
}

func (a *Almanac) SeedToLocation(seed uint64) uint64 {
	// lookup the soil number for this seed
	// but initialize with its own number in case no map was found
	// soil := seed
	soil := GetMappedNumber(seed, a.AlmanacMap["seed-to-soil"])
	//fmt.Printf("seed %d, soil %d,", seed, soil)
	fertilizer := GetMappedNumber(soil, a.AlmanacMap["soil-to-fertilizer"])
	//fmt.Printf("fertilizer %d,", fertilizer)
	water := GetMappedNumber(fertilizer, a.AlmanacMap["fertilizer-to-water"])
	//fmt.Printf("water %d,", water)
	light := GetMappedNumber(water, a.AlmanacMap["water-to-light"])
	//fmt.Printf("light %d,", light)
	temperature := GetMappedNumber(light, a.AlmanacMap["light-to-temperature"])
	//fmt.Printf("temperature %d,", temperature)
	humidity := GetMappedNumber(temperature, a.AlmanacMap["temperature-to-humidity"])
	//fmt.Printf("humidity %d,", humidity)
	location := GetMappedNumber(humidity, a.AlmanacMap["humidity-to-location"])
	//fmt.Printf("location %d,", location)

	return location
}

func GetMappedNumber(origin uint64, valuemap []GardenerMap) uint64 {
	// val := origin
	for _, m := range valuemap {
		if origin >= m.SourceStart && origin < m.SourceStart+m.RangeLength {
			return m.DestinationStart - m.SourceStart + origin

			//return m.SourceStart - m.RangeLenth + origin
		}
	}
	return origin
}

func GetMapValuesFromLines(startIndex int, lines []string) []GardenerMap {
	var getNumbers []uint64
	var rangemaps []GardenerMap
	for idx := startIndex; idx <= len(lines)-1; idx++ {
		if lines[idx] == "" {
			break
		}
		getNumbers = getSeedList(lines[idx])

		gardernerMap := GardenerMap{
			DestinationStart: getNumbers[0],
			SourceStart:      getNumbers[1],
			RangeLength:      getNumbers[2],
		}
		rangemaps = append(rangemaps, gardernerMap)
	}
	return rangemaps
}

// for part 1 we need to read the line as individual seed numbers
func getSeedList(line string) []uint64 {
	parts := strings.Split(line, " ")

	// cut out dangling spaces
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	var numbers []uint64
	for _, part := range parts {
		if part != "" {
			value, err := strconv.ParseUint(part, 10, 64)
			if err == nil { // only when number found
				numbers = append(numbers, value)
			}

		}
	}

	return numbers
}

type SeedRange struct {
	Start       uint64
	RangeLength uint64
}

// for part 2 we need to read the line as a range
func getSeedRanges(line string) []SeedRange {
	parts := strings.Split(line, " ")

	// cut out dangling spaces
	var valueparts []string
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		if parts[i] != "seeds:" && parts[i] != "" {
			valueparts = append(valueparts, parts[i])
		}
	}

	var ranges []SeedRange
	alternate := 0
	start := uint64(0)
	for _, part := range valueparts {
		if part != "" {
			value, err := strconv.ParseUint(part, 10, 64)
			if err == nil { // only when number found
				if alternate%2 == 0 {
					start = value
				} else {
					SeedRange := SeedRange{
						Start:       start,
						RangeLength: value,
					}
					ranges = append(ranges, SeedRange)
				}
				alternate += 1

			}

		}
	}

	return ranges
}

// now how to convert the ranges of the SeedRanges with the almanac?
// suppose we have a range     A........xxxxxxxxxx............
// and need to map              ................vv............
// into                         ....mm........................
// we could slice the unmapped A........xxxxxxxx..............
// and create a new seedrange  B....xx........................
// another mapper would say     ....vvvvvvvv..................
// into                         .....................mmmmmmmm.
// to slice the unmapped   new A............xxxx..............
// and new range               C.........................xxxx.
// to then form A+B+C           ....xx......xxxx.........xxxx.
// it looks like we can always focus on the still unmapped ranges to create new ranges
// and then concatenated the still unmapped with the new ranges..phew..
func GetNewRangesForSeedRanges(seedRanges []SeedRange, gardenerMaps []GardenerMap) []SeedRange {
	var resultSeedRanges []SeedRange

	for _, sr := range seedRanges {
		resultSeedRanges = append(resultSeedRanges, GetNewRangesFor(sr, gardenerMaps)...)
	}

	return resultSeedRanges
}

func GetNewRangesFor(sr SeedRange, gardenerMaps []GardenerMap) []SeedRange {
	var resultSeedRanges []SeedRange
	processedSeedRanges := make(map[SeedRange]bool)

	for _, mapping := range gardenerMaps {
		seedRanges := MapRange(sr, mapping)
		for _, newSeedRange := range seedRanges {
			// Check if the newSeedRange has already been processed
			if _, exists := processedSeedRanges[newSeedRange]; !exists {
				// Append the newSeedRange to the result and mark it as processed
				resultSeedRanges = append(resultSeedRanges, newSeedRange)
				processedSeedRanges[newSeedRange] = true
			}
		}
	}

	return resultSeedRanges
}

// IsInRange checks if a SeedRange falls inside a GardenerMap range
func IsInRange(seedRange SeedRange, gardenerMap GardenerMap) bool {
	// Calculate the end value of the SeedRange
	seedRangeEnd := seedRange.Start + seedRange.RangeLength - 1

	// Check if the SeedRange and GardenerMap ranges overlap
	overlap := seedRangeEnd >= gardenerMap.SourceStart && seedRange.Start <= gardenerMap.SourceStart+gardenerMap.RangeLength-1

	return overlap
}

// MapRange generates an array of SeedRange instances based on GardenerMap
// Should return all mapped values for the seedRange values.
// Suppose seedRange has Start 5 and Length 3 (5,6,7)
// gardernerMap has mapping Start 6, Destination 10, Length 6, the result should become:
// Ranges/Length (5,1 - seed) and (10,2 (for 10 and 11 - mapped))
// ...
// if the seedRange has a longer length like Start 5, length 5 (5,6,7,8,9)
// and mappings exist
// Start: 6, Destination 8, Length: 1 and also
// Start: 8, Destination 1, Length: 2
// we have to keep track of origin an mapped values and get:
// First mapping handled: Ranges/Length (5, 1) then chop 6 to (8,1)mapped and leave (7,3)
// ....xxxxx
// ....x.xXx (overlap at 8, as 8 itself not yet mapped)
// Next map handling:
// xx..x.x.. : (1,2)mapped, (5,1)origin, (7,1)origin cut off and we had (8,1)mapped aside
// Means:
// we have to distinguish the newly mapped ranges as there could be another mapping
// coming by that only manipulates the original values
func MapRange(seedRange SeedRange, gardenerMap GardenerMap) []SeedRange {
	var seedRanges []SeedRange

	// Check if there is no overlap
	//if !IsInRange(seedRange, gardenerMap) {
	// No overlap, return the original SeedRange
	//	return []SeedRange{seedRange}
	//}

	// Calculate the end value of the SeedRange
	seedRangeEnd := seedRange.Start + seedRange.RangeLength - 1

	// Check if the entire SeedRange is overlapped by the GardenerMap
	if seedRange.Start >= gardenerMap.SourceStart && seedRangeEnd <= gardenerMap.SourceStart+gardenerMap.RangeLength-1 {
		// The entire SeedRange is overlapped, create a new SeedRange with mapped values
		newSeedRange := SeedRange{
			Start:       gardenerMap.DestinationStart + (seedRange.Start - gardenerMap.SourceStart),
			RangeLength: seedRange.RangeLength,
		}
		seedRanges = append(seedRanges, newSeedRange)
	} else {
		// The SeedRange is partially overlapped, create SeedRanges for the non-overlapping parts
		if seedRange.Start < gardenerMap.SourceStart {
			// Create SeedRange for the non-overlapping part before the GardenerMap
			nonOverlapBefore := SeedRange{
				Start:       seedRange.Start,
				RangeLength: gardenerMap.SourceStart - seedRange.Start,
			}
			seedRanges = append(seedRanges, nonOverlapBefore)
		}

		// Create SeedRange for the overlapped part mapped according to GardenerMap
		overlapStart := max(seedRange.Start, gardenerMap.SourceStart)
		overlapEnd := min(seedRangeEnd, gardenerMap.SourceStart+gardenerMap.RangeLength-1)
		overlapLength := overlapEnd - overlapStart + 1
		overlapMapped := SeedRange{
			Start:       gardenerMap.DestinationStart + (overlapStart - gardenerMap.SourceStart),
			RangeLength: overlapLength,
		}
		seedRanges = append(seedRanges, overlapMapped)

		if seedRangeEnd > gardenerMap.SourceStart+gardenerMap.RangeLength-1 {
			// Create SeedRange for the non-overlapping part after the GardenerMap
			nonOverlapAfter := SeedRange{
				Start:       gardenerMap.DestinationStart + (gardenerMap.RangeLength + overlapStart - gardenerMap.SourceStart),
				RangeLength: seedRangeEnd - (gardenerMap.SourceStart + gardenerMap.RangeLength) + 1,
			}
			seedRanges = append(seedRanges, nonOverlapAfter)
		}
	}

	return seedRanges
}
func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// The smart way does not seem to work
// so let's go for bruteforcing...
func Year2023Day05Task02(lines []string) uint64 {

	// reading in the data...
	almanac := &Almanac{}
	almanac.ReadAlmanacSeeds(lines)

	fmt.Printf("Almanac: %v\nSeedRanges: %v\n", almanac.AlmanacMap, almanac.SeedRanges)

	lowestlocation := uint64(math.MaxUint64)
	for _, r := range almanac.SeedRanges {
		fmt.Printf("\nSeedrange: %v\n", r)
		from := r.Start
		to := r.Start - 1 + r.RangeLength
		lastpercentage := float32(0)

		for seed := from; seed <= to; seed++ {
			location := almanac.SeedToLocation(seed)

			if location < lowestlocation {
				lowestlocation = location
			}

			perc := (float32(seed-from) / float32(to-from) * 100)
			if perc > lastpercentage+.75 {
				fmt.Printf("\033[2K\r%.6f", perc)
				lastpercentage = perc
			}
		}
		fmt.Printf("\033[2K\r%.6f %d", 100.0, lowestlocation)
	}

	return lowestlocation
}
