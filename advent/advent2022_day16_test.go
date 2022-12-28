package advent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Valve struct {
	ID       string
	LeadTo   []*Valve
	FlowRate int
	Open     bool
	Explored bool           // used in shortestPath
	Parent   *Valve         // used in shortestPath
	MoveCost map[*Valve]int // cost of moving to other valve in steps/minutes
}

type ValveMove struct {
	Minute          int   // what minute this move
	Released        int   // released this minute
	StandingAtValve Valve // standing at valve
	OpeningValve    Valve // opening a valve

}

func (v *Valve) AddToLeadTo(otherValve *Valve) {
	for _, searchValve := range v.LeadTo {
		if searchValve.ID == otherValve.ID {
			// already in list
			return
		}
	}
	v.LeadTo = append(v.LeadTo, otherValve)
}

type Day16Search struct {
	Valves                      []*Valve
	StartValve                  *Valve
	Released                    int // accumulate relased pressure
	ValvesWithFlowRateAboveZero int
}

func (ds16 *Day16Search) DetermineShortestPaths() {
	// the idea is to keep a list of shortest paths between the startValve and valves
	// that have a flowrate>0, as well as the shortest paths between each of those
	// valves with a flowrate>0, so that
	// we only have to inspect walking through these paths. This is because
	// we only need to go through the tunnels to open valves that are still closed
	// and we want to do so in a quick manner withoug wasting time walking through
	// tunnels without any cause....
	for _, v := range ds16.Valves {
		v.MoveCost = make(map[*Valve]int)
		for _, otherV := range ds16.Valves {
			if (v == ds16.StartValve && otherV.FlowRate > 0 ||
				(v.FlowRate > 0 && otherV.FlowRate > 0)) &&
				v != otherV {
				// determine shortest path to other valve
				cost := ds16.ShortestRouteToValve(v, otherV)
				v.MoveCost[otherV] = cost
			}
		}
	}
}

func (ds16 *Day16Search) ReleasedPressure() int {
	released := 0
	for _, v := range ds16.Valves {
		if v.Open {
			released += v.FlowRate
		}
	}
	return released
}

func (ds16 *Day16Search) OpenValve(valve *Valve) error {
	if valve.Open {
		return fmt.Errorf("Valve %s was already open", valve.ID)
	}
	valve.Open = true
	return nil
}

// ShortestRouteToValve returns the cost to reach searchPoint or -1 if no route found
func (ds16 *Day16Search) ShortestRouteToValve(currentPos *Valve, searchPoint *Valve) int {
	for _, valveReset := range ds16.Valves {
		valveReset.Explored = false
		valveReset.Parent = nil
	}
	var Q []*Valve // Queue to search the graph
	var v *Valve
	Q = append(Q, currentPos) // start
	currentPos.Explored = true
	for {
		if len(Q) == 0 {
			break // no route found?
		}
		v, Q = Q[0], Q[1:] // pop from queue
		if v == searchPoint {
			currentPoint := v
			steps := 0
			for {
				// fmt.Printf("step %d via (%s\n", steps, currentPoint.ID)
				currentPoint = currentPoint.Parent
				if currentPoint == nil {
					break
				}
				steps += 1
			}
			return steps // basically the cost to arrive from currentPos to searchPoint
		}
		for _, tunnelMove := range v.LeadTo {
			// if move not labeled as explored already
			if !tunnelMove.Explored {
				tunnelMove.Explored = true
				tunnelMove.Parent = v

				Q = append(Q, tunnelMove)
			}
		}
	}
	return -1
}

// return max found Released, initially to be called
// with the first minute (1) and stops at minutes  (30)
// minute: is the minute we are in
func (ds16 *Day16Search) OpenValvesRecursive(minute int, OpenedValves int, MoveToValve *Valve) int {
	releasedThisMinute := ds16.ReleasedPressure()
	if minute == 30 { // stopcondition
		return releasedThisMinute
	}

	if OpenedValves == ds16.ValvesWithFlowRateAboveZero { // all valves already open
		released := ds16.OpenValvesRecursive(minute+1, OpenedValves, MoveToValve)
		return releasedThisMinute + released // no alternative moves to make just wait it out
	}

	// From the current valve we can open it, if the valve was not already opened before
	max := 0
	OpenAction := false
	if !MoveToValve.Open && MoveToValve.FlowRate > 0 { // initial valve AA has flowRate of 0
		MoveToValve.Open = true // this is the action of this minute
		OpenAction = true
		OpenedValves += 1
	}
	// or if valve is (already) open, we can move to other valves
	for travelValve, cost := range MoveToValve.MoveCost {
		// if we can't open that Valve, because it was already opened, there really is no need
		// to visit that valve, so skip it
		if travelValve.Open {
			continue
		}
		// in case the cost is higher than we have time for we should not go there
		// but just return the releaseAmount that can be reached
		if minute+cost > 30 {
			released := ds16.OpenValvesRecursive(minute+1, OpenedValves, MoveToValve) // just pass value to stay
			return releasedThisMinute + released                                      // no alternative moves to make just wait it out
		}
		releasedWalkingThere := ds16.ReleasedPressure() * (cost - 1) // this includes CURRENT!
		released := ds16.OpenValvesRecursive(minute+cost, OpenedValves, travelValve)
		// have to add the movecost ReleasedPressue without CURRENT as it is already include
		released += releasedWalkingThere

		if released > max {
			max = released
		}
	}
	if OpenAction {
		released := ds16.OpenValvesRecursive(minute+1, OpenedValves, MoveToValve)
		MoveToValve.Open = false // close after recursive call
		return released + releasedThisMinute
	}
	return releasedThisMinute + max
}

func MovePathCircleDetected(movePath []*ValveMove) bool {
	// AA -> DD -> AA -> DD should not be ok
	if len(movePath) < 3 {
		return false
	}
	// only inspect the last 4 items
	lastFourMoves := movePath[:4]
	if lastFourMoves[0].StandingAtValve.ID == lastFourMoves[2].StandingAtValve.ID &&
		lastFourMoves[1].StandingAtValve.ID == lastFourMoves[3].StandingAtValve.ID {
		return true
	}
	return false
}

func (ds16 *Day16Search) FindValveOrCreate(valve string) *Valve {
	for _, v := range ds16.Valves {
		if v.ID == valve { // found!
			return v
		}
	}

	newValve := &Valve{
		ID: valve,
	}

	ds16.Valves = append(ds16.Valves, newValve)
	return newValve
}

func (ds16 *Day16Search) SetValveFlowRate(valve string, flowrate int) *Valve {

	// search the valve in the list, if it exist return it
	for _, v := range ds16.Valves {
		if v.ID == valve { // found!
			v.FlowRate = flowrate
			return v

		}
	}
	newValve := &Valve{
		ID:       valve,
		FlowRate: flowrate,
	}
	ds16.Valves = append(ds16.Valves, newValve)
	return newValve
}
func NewDay16Search(lines []string) *Day16Search {
	ds16 := &Day16Search{}
	for _, line := range lines {
		r := strings.NewReader(line)
		var valve string
		var flowrate int
		fmt.Fscanf(r, "Valve %s has flow rate=%d; tunnels lead to valves ", &valve, &flowrate)
		leadsto := line[strings.Index(line, "to valve")+8:]
		if string(leadsto[0]) == `s` {
			leadsto = leadsto[2:] // cut the `s `` that is there from "to valves "
		}
		v := ds16.SetValveFlowRate(valve, flowrate)
		// lets position ourselves at the first valve read, as puzzle
		// description NOT too clear on this? It says has to be AA but AA not part
		// of the puzzle input..
		if ds16.StartValve == nil {
			ds16.StartValve = v
		}
		for _, l := range strings.Split(leadsto, ",") {
			leadto := strings.Trim(l, " ")
			leadtoValve := ds16.FindValveOrCreate(leadto)
			v.AddToLeadTo(leadtoValve)
		}

	}
	abovezero := 0
	for _, valve := range ds16.Valves {
		if valve.FlowRate > 0 {
			abovezero += 1
		}
	}
	ds16.ValvesWithFlowRateAboveZero = abovezero

	// it would be nice if we could also keep a list of the shortest routes to the valves
	// that have a FlowRate. If we have such a list, we only need to worry about
	// moving through the tunnels of those particular paths.

	// first, get a list of all the valves that have a flowrate.
	ds16.DetermineShortestPaths()

	return ds16
}

func (ds16 *Day16Search) PrintValves() {
	fmt.Printf("valves read %d", len(ds16.Valves))

	for _, v := range ds16.Valves {
		fmt.Printf("%s with rate %d leads to..\n", v.ID, v.FlowRate)
		for _, l := range v.LeadTo {
			fmt.Printf("     %s with rate %d\n", l.ID, l.FlowRate)
		}
		for flowValve, cost := range v.MoveCost {
			fmt.Printf("    %s with cost %d\n", flowValve.ID, cost)
		}
	}

}
func TestDay16Task1(t *testing.T) {
	fileLines, err := GetFileLines("input2022day16Example.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}

	ds16 := NewDay16Search(fileLines)
	ds16.PrintValves()

	released := ds16.OpenValvesRecursive(1, 0, ds16.StartValve)
	fmt.Printf("max released %d\n", released)

	assert.Equal(t, 1651, released)
	/*for _, move := range movepath {
		fmt.Printf("At: %s, minute: %d, Released %d\n", move.StandingAtValve.ID, move.Minute, move.Released)
		if move.OpeningValve.ID != "" {
			fmt.Printf("Opening %s\n", move.OpeningValve.ID)
		}
	}*/
}
