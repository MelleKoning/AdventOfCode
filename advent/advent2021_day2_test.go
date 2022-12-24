package advent

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2HorizontalPositionAndDepth(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/inputday2.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	var commandList []Command

	for _, line := range fileLines {
		splitted := strings.Split(line, " ")

		direction := GetSubmarinCommand(splitted[0])
		intVar, err := strconv.Atoi(splitted[1])
		if err != nil {
			t.Fatalf("shit %v", err)
		}

		cmd := Command{direction, intVar}

		commandList = append(commandList, cmd)
	}

	assert.Equal(t, 1000, len(commandList))

	forwardposition := 0
	depth := 0
	for _, cmd := range commandList {
		if cmd.command == Command_Unknown {
			t.Fatalf("unknown command found!")
		}
		switch cmd.command {
		case Command_Forward:
			forwardposition += cmd.number

		case Command_Down:
			depth += cmd.number
		case Command_Up:
			depth -= cmd.number
		}
	}

	result := depth * forwardposition
	assert.Equal(t, 1855814, result) // we know this is the right answer
}

func TestDay2AimAndDepth(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/inputday2.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	var commandList []Command

	for _, line := range fileLines {
		splitted := strings.Split(line, " ")

		direction := GetSubmarinCommand(splitted[0])
		intVar, err := strconv.Atoi(splitted[1])
		if err != nil {
			t.Fatalf("shit %v", err)
		}

		cmd := Command{direction, intVar}

		commandList = append(commandList, cmd)
	}

	assert.Equal(t, 1000, len(commandList))

	forwardposition := 0
	depth := 0
	aim := 0
	for _, cmd := range commandList {
		if cmd.command == Command_Unknown {
			t.Fatalf("unknown command found!")
		}
		switch cmd.command {
		case Command_Forward:
			forwardposition += cmd.number
			depth = depth + cmd.number*aim
		case Command_Down:
			aim += cmd.number
		case Command_Up:
			aim -= cmd.number
		}
	}

	result := depth * forwardposition
	assert.Equal(t, 1845455714, result) // we know this is the right answer
}
