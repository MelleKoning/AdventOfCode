package advent

import (
	"fmt"
	"strings"
	"testing"
)

// puzzle https://adventofcode.com/2022/day/19

type ElementsDay19 int

const ELM_ORE ElementsDay19 = 1
const ELM_CLAY ElementsDay19 = 2
const ELM_OBSEDIAN ElementsDay19 = 3
const ELM_GEODE ElementsDay19 = 4

type BotCostsDay19 struct {
	Ore      int
	Clay     int
	Obsedian int
	Geodes   int
}
type OwningDay19 struct {
	BluePrint         *BluePrintDay19
	BotAmountOre      int
	BotAmountClay     int
	BotAmountObsedian int
	BotAmountGeode    int
	Ore               int // amount of Ore owned
	Clay              int // amount of Clay owned
	Obsedian          int // amount of obsedian owned
	Geodes            int // amount of Geodes owned
	//MaxOreCost int
}

type BluePrintDay19 struct {
	BotCosts    map[ElementsDay19]BotCostsDay19
	MaxOreCost  int
	MaxClayCost int
}

func (o *OwningDay19) Start() {
	o.BotAmountOre = 1
	o.BotAmountClay = 0
	o.BotAmountObsedian = 0
	o.BotAmountGeode = 0
}

const MaxMinutes = 32

func TestDay19_1(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day19.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	var bluePrint *OwningDay19
	qualityTotal := 0
	for idx, line := range fileLines {
		if idx > 2 {
			break
		}
		fmt.Printf("%s\n", line)
		bluePrint = GoForBlueprint(line)
		overallmax := new(int)
		*overallmax = -1
		geodesDolven := DelveGeodesRecursive(bluePrint, 1, overallmax)

		quality := geodesDolven * (idx + 1)
		fmt.Printf("Geodes opened: %+v, Quality is %d*%d=%d\n", geodesDolven, idx+1, geodesDolven, quality)
		qualityTotal += quality
		fmt.Printf("Qualitytotal so far: %d\n", qualityTotal)
	}
	// 878 is too low!
	// new run 569??
	// 838??
	// 861?

}

func (o *OwningDay19) StopCondition(minute int, overallmax *int) bool {
	// Keep track of the optimal answer so far. Stop if we cannot beat that even
	// if we build a geode robot every minute from here on out.
	// upper bound: currentGeodes + currentGeodes * timeLeft + geodesFromPotentialNewBots (assume +1 bot/min)
	return *overallmax > (o.Geodes + o.Geodes*(MaxMinutes-minute) +
		(((MaxMinutes - minute) * ((MaxMinutes - minute) + 1)) / 2)) // n(n-1)/2
}

// returns the (hopefully max) number of Geodes found
func DelveGeodesRecursive(bluePrint *OwningDay19, minute int, overallmax *int) int {
	if minute == MaxMinutes {
		// last delve!
		bluePrint.Delve()
		return bluePrint.Geodes
	}
	if bluePrint.StopCondition(minute, overallmax) {
		return 0 // negate
	}
	// we inspect the possible moves and return the max found
	max := 0
	for _, move := range bluePrint.Moves() {
		geodes := DelveGeodesRecursive(move, minute+1, overallmax)
		if geodes > max {
			max = geodes
			if max >= *overallmax {
				*overallmax = max
			}
		}
	}
	return max
}

func GoForBlueprint(line string) *OwningDay19 {
	startOwn := &OwningDay19{}
	startOwn.Start()

	bluePrint := &BluePrintDay19{
		// botCosts to be determined by Blueprint recipy
		BotCosts: make(map[ElementsDay19]BotCostsDay19),
	}
	bluePrint.ReadBluePrint(line)

	startOwn.BluePrint = bluePrint

	return startOwn
	//fmt.Printf("line %s gives: \n%v\n", line, startOwn)
}

func (o *OwningDay19) CanCreateBot(bot ElementsDay19) bool {

	// depending on the costs for bots, we could create
	// different type and amount of bots

	switch bot {
	case ELM_GEODE:
		{
			// if we can build a geode robot we should probably just do it
			if o.Obsedian >= o.BluePrint.BotCosts[ELM_GEODE].Obsedian && o.Ore >= o.BluePrint.BotCosts[ELM_GEODE].Ore {
				o.Obsedian -= o.BluePrint.BotCosts[ELM_GEODE].Obsedian
				o.Ore -= o.BluePrint.BotCosts[ELM_GEODE].Ore
				return true
			}
		}
	case ELM_OBSEDIAN:
		{
			// if we can build a obsidian robot we should probably just do it
			if o.Clay >= o.BluePrint.BotCosts[ELM_OBSEDIAN].Clay && o.Ore >= o.BluePrint.BotCosts[ELM_OBSEDIAN].Ore {
				o.Ore -= o.BluePrint.BotCosts[ELM_OBSEDIAN].Ore
				o.Clay -= o.BluePrint.BotCosts[ELM_OBSEDIAN].Clay
				return true
			}
		}
	case ELM_CLAY:
		{
			// if we can build a Clay robot should we do it?? probably yes
			if o.Ore >= o.BluePrint.BotCosts[ELM_CLAY].Ore {
				o.Ore -= o.BluePrint.BotCosts[ELM_CLAY].Ore
				return true
			}
		}
	case ELM_ORE:
		{
			if o.Ore >= o.BluePrint.BotCosts[ELM_ORE].Ore {
				o.Ore -= o.BluePrint.BotCosts[ELM_ORE].Ore
				return true
			}
		}
	}

	return false
}

func (o *OwningDay19) Delve() {
	o.Ore += o.BotAmountOre
	o.Clay += o.BotAmountClay
	o.Obsedian += o.BotAmountObsedian
	o.Geodes += o.BotAmountGeode
}
func CloneBotsMap(bots map[ElementsDay19]int) map[ElementsDay19]int {
	copyMap := make(map[ElementsDay19]int)
	for k, v := range bots {
		copyMap[k] = v
	}
	return copyMap
}
func (o *OwningDay19) Clone() *OwningDay19 {
	clone := &OwningDay19{
		BotAmountOre:      o.BotAmountOre,
		BotAmountClay:     o.BotAmountClay,
		BotAmountObsedian: o.BotAmountObsedian,
		BotAmountGeode:    o.BotAmountGeode,
		Ore:               o.Ore,
		Clay:              o.Clay,
		Obsedian:          o.Obsedian,
		Geodes:            o.Geodes,
		BluePrint:         o.BluePrint,
		//MaxOreCost: o.MaxOreCost,
	}
	return clone
}
func (o *OwningDay19) Moves() []*OwningDay19 {
	var moves []*OwningDay19
	// moves should return multiple possible followups for the state from the given state.
	// so if we own an ore, we can build an ore bot if the costs are ok with that.

	// 1) we first have to determine what bots we could create from the owning elements, if any
	// and we generate any number of moves based on the possibilities
	// the first move would be just no bot building, a copy of the current state

	moveGeo := o.Clone()
	GeoBotCreated := false
	if moveGeo.CanCreateBot(ELM_GEODE) {
		moves = append(moves, moveGeo)
		GeoBotCreated = true
	}
	ObsBotCreated := false
	moveObs := o.Clone()
	if moveObs.CanCreateBot(ELM_OBSEDIAN) {
		moves = append(moves, moveObs)
		ObsBotCreated = true
	}
	ClayBotsCreated := false
	var moveClay *OwningDay19
	// only build another ClayBot if we do not have as many clayBots as maxClayCost
	if o.BotAmountClay < o.BluePrint.MaxClayCost {
		moveClay = o.Clone()
		if moveClay.CanCreateBot(ELM_CLAY) {
			moves = append(moves, moveClay)
			ClayBotsCreated = true
		}
	}
	OreBotsCreated := false
	var moveOre *OwningDay19
	// only build another orebot if we do not have as many as max
	if o.BotAmountOre < o.BluePrint.MaxOreCost {
		moveOre = o.Clone()
		if moveOre.CanCreateBot(ELM_ORE) {
			moves = append(moves, moveOre)
			OreBotsCreated = true
		}
	}

	// if we can create any bot do not consider the no-bot-create branch
	// This reuces the branching tremendously, greatly enhancing the
	// performance but it does NOT provide the right answer
	// for all the blueprints! I found for my puzzle input
	// it gave mostly right answers but not for blueprints 5 and 18
	// Only found out by letting the machine run without these lines.
	// Luckily for me, step 2 (32 minutes) worked fine for the first 1,2,3
	// blueprints....so I just submitted what I had for step 2 and it was accepted.
	if !(OreBotsCreated || ClayBotsCreated || ObsBotCreated || GeoBotCreated) {
		move := o.Clone()
		moves = append(moves, move)

	}

	// 2) the current owned bots can delve, determine what they delve and add
	// to what we own!
	for _, mv := range moves {
		mv.Delve()
	}

	// 3) only after the current adjustments we can add the number of bots
	// for a move
	if GeoBotCreated {
		moveGeo.BotAmountGeode += 1
	}
	if ObsBotCreated {
		moveObs.BotAmountObsedian += 1
	}
	if ClayBotsCreated {
		moveClay.BotAmountClay += 1
	}
	if OreBotsCreated {
		moveOre.BotAmountOre += 1
	}

	return moves
}
func (b *BluePrintDay19) ReadBluePrint(line string) {
	// Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
	blueprint := strings.Split(line, ".")
	oreBotStr := strings.TrimSpace(strings.Split(blueprint[0], ":")[1])
	clayBotStr := strings.TrimSpace(blueprint[1])
	ObsRobotStr := strings.TrimSpace(blueprint[2])
	GeodRobotStr := strings.TrimSpace(blueprint[3])
	//fmt.Printf("Claybotstring: [%s]\n", clayBotStr)
	//fmt.Printf("Obs: [%s]\n", ObsRobotStr)
	//fmt.Printf("Geo: [%s]\n", GeodRobotStr)
	orebotcosts := GetNumberFromString(strings.Split(oreBotStr, " ")[4])
	//fmt.Printf("Orebot costs: [%d]\n", orebotcosts)
	claybotcosts := GetNumberFromString(strings.Split(clayBotStr, " ")[4])
	//fmt.Printf("Claybot costs: [%d]\n", claybotcosts)

	obsbotorecosts := GetNumberFromString(strings.Split(ObsRobotStr, " ")[4])
	obsbotclaycosts := GetNumberFromString(strings.Split(ObsRobotStr, " ")[7])
	//fmt.Printf("ObsBot costs: [%d] ore and [%d] clay \n", obsbotorecosts, obsbotclaycosts)

	geobotorecosts := GetNumberFromString(strings.Split(GeodRobotStr, " ")[4])
	geobotobscosts := GetNumberFromString(strings.Split(GeodRobotStr, " ")[7])
	//fmt.Printf("GeoBot costs: [%d] ore and [%d] obsedian \n", geobotorecosts, geobotobscosts)

	//define bot factory costs
	b.BotCosts[ELM_ORE] = BotCostsDay19{Ore: orebotcosts}
	b.BotCosts[ELM_CLAY] = BotCostsDay19{Ore: claybotcosts}
	b.BotCosts[ELM_OBSEDIAN] = BotCostsDay19{Ore: obsbotorecosts, Clay: obsbotclaycosts}
	b.BotCosts[ELM_GEODE] = BotCostsDay19{Ore: geobotorecosts, Obsedian: geobotobscosts}

	maxorecost := 0
	if orebotcosts > maxorecost {
		maxorecost = orebotcosts
	}
	if claybotcosts > maxorecost {
		maxorecost = claybotcosts
	}
	if obsbotorecosts > maxorecost {
		maxorecost = obsbotorecosts
	}
	if geobotorecosts > maxorecost {
		maxorecost = geobotobscosts
	}
	b.MaxOreCost = maxorecost
	b.MaxClayCost = obsbotclaycosts // we do not have to create claybots if we have this many already

}

/* Results found in run

Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 12 clay. Each geode robot costs 4 ore and 19 obsidian.
Geodes opened: 0, Quality is 1*0=0
Qualitytotal so far: 0
Blueprint 2: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 11 clay. Each geode robot costs 4 ore and 12 obsidian.
Geodes opened: 1, Quality is 2*1=2
Blueprint 3: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 18 clay. Each geode robot costs 2 ore and 11 obsidian.
Geodes opened: 2, Quality is 3*2=6
Blueprint 4: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 4 ore and 20 obsidian.
Geodes opened: 0, Quality is 4*0=0
Blueprint 5: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 16 clay. Each geode robot costs 4 ore and 17 obsidian.
Geodes opened: 1, Quality is 5*1=5
lueprint 6: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 19 clay. Each geode robot costs 2 ore and 12 obsidian.
Geodes opened: 0, Quality is 6*0=0
Qualitytotal so far: 13
Blueprint 7: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 9 clay. Each geode robot costs 2 ore and 10 obsidian.
Geodes opened: 7, Quality is 7*7=49
ualitytotal so far: 62
Blueprint 8: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 5 clay. Each geode robot costs 3 ore and 7 obsidian.
Geodes opened: 9, Quality is 8*9=72
Qualitytotal so far: 134
Blueprint 9: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 11 clay. Each geode robot costs 4 ore and 8 obsidian.
Geodes opened: 3, Quality is 9*3=27
Qualitytotal so far: 161
lueprint 10: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 16 clay. Each geode robot costs 2 ore and 15 obsidian.
Geodes opened: 0, Quality is 10*0=0
Qualitytotal so far: 161
Blueprint 11: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 20 clay. Each geode robot costs 2 ore and 19 obsidian.
Geodes opened: 0, Quality is 11*0=0
Qualitytotal so far: 161
Blueprint 12: Each ore robot costs 3 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 16 clay. Each geode robot costs 3 ore and 20 obsidian.
Geodes opened: 0, Quality is 12*0=0
Blueprint 13: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 3 ore and 14 obsidian.
Geodes opened: 1, Quality is 13*1=13
Qualitytotal so far: 174
Blueprint 14: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 2 ore and 15 obsidian.
Geodes opened: 0, Quality is 14*0=0
Qualitytotal so far: 174
Blueprint 15: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 3 ore and 12 obsidian.
Geodes opened: 1, Quality is 15*1=15
Qualitytotal so far: 189
Blueprint 16: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 17 clay. Each geode robot costs 3 ore and 19 obsidian.
Geodes opened: 2, Quality is 16*2=32
Qualitytotal so far: 221
Blueprint 17: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 4 ore and 9 obsidian.
Geodes opened: 6, Quality is 17*6=102
Qualitytotal so far: 323
Blueprint 18: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 6 clay. Each geode robot costs 3 ore and 16 obsidian.
Geodes opened: 5, Quality is 18*5=90
Qualitytotal so far: 413
Blueprint 19: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 6 clay. Each geode robot costs 2 ore and 14 obsidian.
Geodes opened: 3, Quality is 19*3=57
Qualitytotal so far: 470
Blueprint 20: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 4 ore and 11 clay. Each geode robot costs 3 ore and 15 obsidian.
Geodes opened: 1, Quality is 20*1=20
Qualitytotal so far: 490
Blueprint 21: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 18 clay. Each geode robot costs 4 ore and 19 obsidian.
Geodes opened: 0, Quality is 21*0=0
Qualitytotal so far: 490
Blueprint 22: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 15 clay. Each geode robot costs 2 ore and 20 obsidian.
Geodes opened: 1, Quality is 22*1=22
Qualitytotal so far: 512
Blueprint 23: Each ore robot costs 4 ore. Each clay robot costs 3 ore. Each obsidian robot costs 2 ore and 5 clay. Each geode robot costs 2 ore and 10 obsidian.
Geodes opened: 15, Quality is 23*15=345
Qualitytotal so far: 857
Blueprint 24: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 10 clay. Each geode robot costs 2 ore and 14 obsidian.
Geodes opened: 1, Quality is 24*1=24
Qualitytotal so far: 881
Blueprint 25: Each ore robot costs 4 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 7 clay. Each geode robot costs 4 ore and 13 obsidian.
Geodes opened: 3, Quality is 25*3=75
Qualitytotal so far: 956
Geodes opened: 0, Quality is 26*0=0
Qualitytotal so far: 956
Blueprint 27: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 18 clay. Each geode robot costs 2 ore and 19 obsidian.
Geodes opened: 1, Quality is 27*1=27
Qualitytotal so far: 983
Blueprint 28: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 3 ore and 10 clay. Each geode robot costs 2 ore and 7 obsidian.
Geodes opened: 7, Quality is 28*7=196
Qualitytotal so far: 1179
Blueprint 29: Each ore robot costs 3 ore. Each clay robot costs 4 ore. Each obsidian robot costs 2 ore and 15 clay. Each geode robot costs 3 ore and 7 obsidian.
Geodes opened: 3, Quality is 29*3=87
Qualitytotal so far: 1266
Blueprint 30: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 4 ore and 18 obsidian.
Blueprint 30: Each ore robot costs 2 ore. Each clay robot costs 4 ore. Each obsidian robot costs 4 ore and 20 clay. Each geode robot costs 4 ore and 18 obsidian.
Geodes opened: 0, Quality is 30*0=0
Qualitytotal so far: 1266

*/
