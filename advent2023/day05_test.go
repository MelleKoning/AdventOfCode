package advent2023

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

type GardenerMap struct {
	// 52 50 48
	// source 50 with length 48 (50..98) maps to (52..99)
	DestinationStart uint64
	SourceStart      uint64
	RangeLenth       uint64
}

func TestDay05Task1(t *testing.T) {

	lines := GetLinesFromFile("input/day5input.txt")

	GardenerMaps := make(map[string][]GardenerMap, 0)

	// reading in the data...
	var seedlist []uint64
	for idx, l := range lines {
		if strings.Contains(l, "seeds:") {
			seedlist = getSeedList(l)
		}
		if strings.Contains(l, "seed-to-soil map:") {
			GardenerMaps["seed-to-soil"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "soil-to-fertilizer map:") {
			GardenerMaps["soil-to-fertilizer"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "fertilizer-to-water map:") {
			GardenerMaps["fertilizer-to-water"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "water-to-light map:") {
			GardenerMaps["water-to-light"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "light-to-temperature map:") {
			GardenerMaps["light-to-temperature"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "temperature-to-humidity map:") {
			GardenerMaps["temperature-to-humidity"] = GetMapValuesFromLines(idx+1, lines)
		}
		if strings.Contains(l, "humidity-to-location map:") {
			GardenerMaps["humidity-to-location"] = GetMapValuesFromLines(idx+1, lines)
		}
	}

	fmt.Printf("seedlist: %v\nGardenerMaps:\n%v\n", seedlist, GardenerMaps)

	lowestlocation := uint64(math.MaxUint64)
	for _, seed := range seedlist {
		// lookup the soil number for this seed
		// but initialize with its own number in case no map was found
		// soil := seed
		soil := GetMappedNumber(seed, GardenerMaps["seed-to-soil"])
		fmt.Printf("seed %d, soil %d,", seed, soil)
		fertilizer := GetMappedNumber(soil, GardenerMaps["soil-to-fertilizer"])
		fmt.Printf("fertilizer %d,", fertilizer)
		water := GetMappedNumber(fertilizer, GardenerMaps["fertilizer-to-water"])
		fmt.Printf("water %d,", water)
		light := GetMappedNumber(water, GardenerMaps["water-to-light"])
		fmt.Printf("light %d,", light)
		temperature := GetMappedNumber(light, GardenerMaps["light-to-temperature"])
		fmt.Printf("temperature %d,", temperature)
		humidity := GetMappedNumber(temperature, GardenerMaps["temperature-to-humidity"])
		fmt.Printf("humidity %d,", humidity)
		location := GetMappedNumber(humidity, GardenerMaps["humidity-to-location"])
		fmt.Printf("location %d,", location)

		if location < lowestlocation {
			lowestlocation = location
		}
		fmt.Printf("\n")
	}

	fmt.Printf("lowest location is %d\n", lowestlocation)
}

func GetMappedNumber(origin uint64, valuemap []GardenerMap) uint64 {
	// val := origin
	for _, m := range valuemap {
		if origin >= m.SourceStart && origin <= m.SourceStart+m.RangeLenth {
			// we know how to map this:
			if m.RangeLenth > m.SourceStart {
				return m.DestinationStart - m.SourceStart + origin
			}
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
			RangeLenth:       getNumbers[2],
		}
		rangemaps = append(rangemaps, gardernerMap)
	}
	return rangemaps
}
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
