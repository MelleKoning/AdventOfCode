package advent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MathOperation int

const (
	NONE   MathOperation = 0
	ADD    MathOperation = 1
	SUB    MathOperation = 2
	MUL    MathOperation = 3
	DIV    MathOperation = 4
	MULOLD MathOperation = 5
)

var MathOperationText = map[MathOperation]string{
	NONE: "None",
	ADD:  "Add",
	SUB:  "Sub",
	MUL:  "Mul",
	DIV:  "Div",
}

type MonkeyOperation struct {
	MathOp MathOperation
	Value  int
}
type Monkey11 struct {
	Items          []int
	Operation      MonkeyOperation
	TestModuloBy   int
	TrueMonkey     int
	FalseMonkey    int
	InspectedItems int
}

type MonkeyGroup struct {
	Monkeys map[int]*Monkey11
}

func (g *MonkeyGroup) AddMonkeyFoundOnline(idx int, fileLines []string) {
	number := GetNumberFromString(strings.Split(fileLines[idx], " ")[1][:1])
	g.Monkeys[number] = &Monkey11{}
	// get the items
	itmstr := strings.Split(fileLines[idx+1], ":")
	items := strings.Split(strings.TrimSpace(itmstr[1]), ", ")
	for _, i := range items {
		g.Monkeys[number].Items = append(g.Monkeys[number].Items, GetNumberFromString(i))
	}
	g.Monkeys[number].SetMonkeyOperation(fileLines[idx+2])
	modulostr := strings.Split(fileLines[idx+3], " ")

	g.Monkeys[number].TestModuloBy = GetNumberFromString(modulostr[len(modulostr)-1])
	trueMonkeystr := strings.Split(fileLines[idx+4], " ")
	g.Monkeys[number].TrueMonkey = GetNumberFromString(trueMonkeystr[len(trueMonkeystr)-1])
	falseMonkeystr := strings.Split(fileLines[idx+5], " ")
	g.Monkeys[number].FalseMonkey = GetNumberFromString(falseMonkeystr[len(falseMonkeystr)-1])

}

func (g *MonkeyGroup) GreatestCommonMultiple() int {
	multiplyCommonDiv := 1
	for m := 0; m < len(g.Monkeys); m++ {
		multiplyCommonDiv *= g.Monkeys[m].TestModuloBy
		fmt.Printf("m: %d modulo: %d  gcm: %d", m, g.Monkeys[m].TestModuloBy, multiplyCommonDiv)
	}

	return multiplyCommonDiv
}
func TestParseOperation(t *testing.T) {
	sentence := "  Operation: new = old * 19"
	monkey := &Monkey11{}
	monkey.SetMonkeyOperation(sentence)
	assert.Equal(t, 19, monkey.Operation.Value)
	assert.Equal(t, MUL, monkey.Operation.MathOp)
}

// Operate determines to what monkey the item with NewWorryLevel is being thrown
func (m *Monkey11) Operate(item int) (int, int) {
	m.InspectedItems += 1
	newVal := 0
	switch m.Operation.MathOp {
	case MUL:
		newVal = item * m.Operation.Value
	case ADD:
		newVal = item + m.Operation.Value
	case MULOLD:
		newVal = item * item
	}
	newVal = newVal / 3

	if newVal%m.TestModuloBy == 0 {
		return m.TrueMonkey, newVal
	}
	return m.FalseMonkey, newVal
}

// Operate determines to what monkey the item with NewWorryLevel is being thrown
func (m *Monkey11) OperateWorry(item int, gcm int) (int, int) {
	m.InspectedItems += 1
	newVal := 0
	switch m.Operation.MathOp {
	case MUL:
		newVal = item * m.Operation.Value
	case ADD:
		newVal = item + m.Operation.Value
	case MULOLD:
		newVal = item * item
	}

	newVal = newVal % gcm // this is the magic to keep numbers low

	if newVal%m.TestModuloBy == 0 {
		//newVal = newVal / m.TestModuloBy
		return m.TrueMonkey, newVal
	}
	return m.FalseMonkey, newVal
}

func (m *Monkey11) SetMonkeyOperation(line string) {
	items := strings.Split(line, " ")
	number := GetNumberFromString(items[len(items)-1])
	opchar := items[len(items)-2]
	var operation MathOperation
	switch opchar {
	case "*":
		operation = MUL
	case "+":
		operation = ADD
	case "-":
		operation = SUB
	case "/":
		operation = DIV
	}
	if number == -1 {
		operation = MULOLD // old * old
	}
	m.Operation = MonkeyOperation{MathOp: operation, Value: number}

}
func TestDay11Monkeys1(t *testing.T) {

	// the first task is to parse what the Monkeys will do
	fileLines, err := GetFileLines("inputdata/input2022day11.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	MonkeyGroup := &MonkeyGroup{
		Monkeys: make(map[int]*Monkey11), // map of the monkies :),
	}

	for idx, line := range fileLines {
		if strings.Index(line, "Monkey") == 0 {
			MonkeyGroup.AddMonkeyFoundOnline(idx, fileLines)
		}
	}

	for key, m := range MonkeyGroup.Monkeys {
		fmt.Printf("%d, %v\n", key, m)
	}

	greatestCommonMultiple := MonkeyGroup.GreatestCommonMultiple()
	fmt.Printf("Common Divisible: %d\n", greatestCommonMultiple)

	// The Rounds...
	for rounds := 1; rounds <= 10000; rounds++ {
		fmt.Printf("round %d\n", rounds)
		for m := 0; m < len(MonkeyGroup.Monkeys); m++ {
			currentMonkey := MonkeyGroup.Monkeys[m]
			for _, item := range currentMonkey.Items {
				itemGoesTo, withNewWorryLevel := currentMonkey.OperateWorry(item, greatestCommonMultiple)
				MonkeyGroup.Monkeys[itemGoesTo].Items = append(MonkeyGroup.Monkeys[itemGoesTo].Items, withNewWorryLevel)
			}
			// all items thrown so empty the current monkeys list
			currentMonkey.Items = nil
		}
		// items per monkey after the round
		for m := 0; m < len(MonkeyGroup.Monkeys)-1; m++ {
			fmt.Printf("%d: %v\n", m, MonkeyGroup.Monkeys[m].Items)
		}

		// busiest monkeys round 2:
		//Monkey: 2, inspected items: 18493
		//Monkey: 3, inspected items: 128102
		//Monkey: 4, inspected items: 134002
		//Monkey: 5, inspected items: 112542
		//Monkey: 6, inspected items: 133798
		//Monkey: 7, inspected items: 133506
		//Monkey: 0, inspected items: 86264
		//Monkey: 1, inspected items: 14180
		// 134002 * 133798 = 17929199596 is too low?

		// second try:
		//Monkey: 4, inspected items: 152660
		//Monkey: 5, inspected items: 138194
		//Monkey: 6, inspected items: 142807
		//Monkey: 7, inspected items: 123730
		//Monkey: 0, inspected items: 83587
		//Monkey: 1, inspected items: 29059
		//Monkey: 2, inspected items: 9882
		//Monkey: 3, inspected items: 123613
		// 152660 * 142807 = 21800916620
	}

	for idx, monkey := range MonkeyGroup.Monkeys {
		fmt.Printf("Monkey: %d, inspected items: %d\n", idx, monkey.InspectedItems)
	}
}
