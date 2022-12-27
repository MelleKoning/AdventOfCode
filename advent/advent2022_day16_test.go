package advent

import (
	"fmt"
	"strings"
	"testing"
)

type Valve struct {
	ID       string
	LeadTo   []*Valve
	FlowRate int
	Open     bool
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
	Valves       []*Valve
	StandAtValve *Valve
	Released     int // accumulate relased pressure
	OpenedValves int // count the number of opened valves
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

// return max found Released
func (ds16 *Day16Search) OpenValvesRecursive(minute int) int {
	ds16.Released += ds16.ReleasedPressure()
	if minute == 30 { // stopcondition
		return ds16.Released
	}
	if ds16.OpenedValves == len(ds16.Valves) { // all valves already open
		// no alternative moves to make just wait it out

		return ds16.OpenValvesRecursive(minute + 1)
	}

	// TODO: finish tries and moves

	return 0 // ??
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
		for _, l := range strings.Split(leadsto, ",") {
			leadto := strings.Trim(l, " ")
			leadtoValve := ds16.FindValveOrCreate(leadto)
			v.AddToLeadTo(leadtoValve)
		}

	}
	return ds16
}

func (ds16 *Day16Search) PrintValves() {
	fmt.Printf("valves read %d", len(ds16.Valves))

	for _, v := range ds16.Valves {
		fmt.Printf("%s with rate %d leads to..\n", v.ID, v.FlowRate)
		for _, l := range v.LeadTo {
			fmt.Printf("     %s with rate %d\n", l.ID, l.FlowRate)
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
}
