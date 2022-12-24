package advent

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// puzzle https://adventofcode.com/2022/day/18

type CubyDay18 struct {
	X, Y, Z int
}

type CubeListDay18 struct {
	Cubes []*CubyDay18
}

func TestCubesDay18Example(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day18Example.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(fileLines)

	cubeList := &CubeListDay18{}

	totalExposedArea := 0

	for _, line := range fileLines {
		cubeString := strings.Split(line, ",")

		lineCube := &CubyDay18{GetNumberFromString(cubeString[0]),
			GetNumberFromString(cubeString[1]),
			GetNumberFromString(cubeString[2])}

		cubeList.Cubes = append(cubeList.Cubes, lineCube)

		exposed := cubeList.AddCube(lineCube)
		totalExposedArea += 6
		totalExposedArea -= (6 - exposed) * 2
	}
	fmt.Printf("Exposed: %d\n", totalExposedArea)
}

func TestFind18_Example(t *testing.T) {

	// the input gives us all cubes in an x,y,z grid that we can layout
	// in the grid.
	// suppose their are two cubes adjacent to each other at
	// 1,1,1 and 2,1,1, these two cubes are adjacent and each individual
	// cube is exposed on all other (five) sides, giving a surface area of 10.

	// How to determine a surface area for any given number of cubes?
	// We could keep track of every individual cube we add to the grid.
	// Once a side gets covered from an already existing cube, we subtract
	// that surface area.
	// So add 1,1,1: six surface added and visible
	// Now add 2,1,1: add six to surface area again
	// also: inspect the six surrounding cubes of this new cube, and subtract one surface area
	// in case that grid position next to that side is already filled.
	// or rather, subtract two, because both cubesides are covered.
	// To do so we should keep track of all added x,y,z so that we can check:
	// (x-1,y,z) , (x+1, y, z), (x, y-1, z), (x, y+2, z), (x, y, z-1), (x, y, z+1)
	// if we just keep the cubes in a list, this should not be too difficult?

	// let's try with only two cubes first.
	cubeList := &CubeListDay18{}
	cubeList.Cubes = append(cubeList.Cubes, &CubyDay18{1, 1, 1})

	exposed := cubeList.AddCube(&CubyDay18{2, 1, 1})
	assert.Equal(t, 5, exposed)
}

func TestCubesDay18Puzzle1(t *testing.T) {
	fileLines, err := GetFileLines("inputdata/input2022day18.txt")
	// 3610 is the right answer!
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(fileLines)

	cubeList := &CubeListDay18{}

	totalExposedArea := 0

	for _, line := range fileLines {
		cubeString := strings.Split(line, ",")

		lineCube := &CubyDay18{GetNumberFromString(cubeString[0]),
			GetNumberFromString(cubeString[1]),
			GetNumberFromString(cubeString[2])}

		cubeList.Cubes = append(cubeList.Cubes, lineCube)

		exposed := cubeList.AddCube(lineCube)
		totalExposedArea += 6
		totalExposedArea -= (6 - exposed) * 2
	}
	fmt.Printf("Exposed: %d\n", totalExposedArea)
}

func TestDay18Part2(t *testing.T) {
	// for part two, examining the entire outside of the surface
	// can be done with a BreathFirstSearch. The droplets are
	// marked as occupied spaces of the x,y,z grid. When walking
	// right next to the surface and peeking around in all six areas
	// we can see if the air we are in, outside of the lavachunk, is
	// next to a drop. If yes, we count + 1 side, if not, we push
	// the found airgap to be inspected for the BFS.
	fileLines, err := GetFileLines("inputdata/input2022day18.txt")
	// 3610 is the right answer!
	if err != nil {
		t.Fatalf("%v", err)
	}
	fmt.Println(fileLines)

	cubeList := &CubeListDay18{}

	totalExposedArea := 0

	for _, line := range fileLines {
		cubeString := strings.Split(line, ",")

		lineCube := &CubyDay18{GetNumberFromString(cubeString[0]),
			GetNumberFromString(cubeString[1]),
			GetNumberFromString(cubeString[2])}

		cubeList.Cubes = append(cubeList.Cubes, lineCube)

		exposed := cubeList.AddCube(lineCube)
		totalExposedArea += 6
		totalExposedArea -= (6 - exposed) * 2
	}
	// We should examine the entire area from the minimal x,y,z to max x,y,z
	// so we have to determine those. That is, we assume the entire lavachunk
	// is inside points 0,0,0 and maxX, maxY, maxZ.
	// We simply mark all points as false (air) and then mark
	// all droplets as true (lavadrop)
	maxX, maxY, maxZ := cubeList.MaxXYZ()

	fmt.Printf("maxX,maxY,maxZ:%d,%d,%d\n", maxX, maxY, maxZ)
	fmt.Printf("number of cubes: %d\n", len(cubeList.Cubes))
	sides := 0 // number of sides hit against

	transforms := []CubyDay18{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

	QtoExplore := []*CubyDay18{&CubyDay18{maxX, maxY, maxZ}} // start BFS search
	Seen := []*CubyDay18{}                                   // keep list of cubes already seen
	var v *CubyDay18
	for {
		if len(QtoExplore) == 0 {
			break
		}
		v, QtoExplore = QtoExplore[0], QtoExplore[1:] // pop from Q
		if AlreadySeen(Seen, v) {
			continue
		}
		Seen = append(Seen, v) // add inspected cube to seenlist
		//fmt.Printf("inspect %d,%d,%d\n", v.X, v.Y, v.Z)
		for _, moves := range transforms {
			xp := v.X + moves.X
			yp := v.Y + moves.Y
			zp := v.Z + moves.Z
			if xp < -1 || yp < -1 || zp < -1 || xp > maxX || yp > maxY || zp > maxZ {
				continue // out of lavachunk bounds, but still walking at level 0!
			}
			if cubeList.LavaAt(xp, yp, zp) {
				sides++
				continue
			}

			QtoExplore = append(QtoExplore, &CubyDay18{xp, yp, zp}) // no lava there, air found is to explore

		}
	}

	fmt.Printf("Found sides %d\n", sides) // 2082 sides found!
}

func AlreadySeen(seen []*CubyDay18, inspect *CubyDay18) bool {
	for _, c := range seen {
		if c.X == inspect.X &&
			c.Y == inspect.Y &&
			c.Z == inspect.Z {
			return true
		}
	}
	return false
}

/*
const runPart2 = drops => {
   drops = drops.map(([x, y, z]) => [x + 1, y + 1, z + 1]);
   const maxX = Math.max(...drops.map(d => d[0])) + 1;
   const maxY = Math.max(...drops.map(d => d[1])) + 1;
   const maxZ = Math.max(...drops.map(d => d[2])) + 1;
   const map = range(maxX + 1).map(() => range(maxY + 1).map(() => range(maxZ + 1).map(() => false)));
   drops.forEach(([x, y, z]) => {
      map[x][y][z] = true;
   });

   let sides = 0;
   const transforms = [[-1, 0, 0], [1, 0, 0], [0, -1, 0], [0, 1, 0], [0, 0, -1], [0, 0, 1]];
   const seen = map.map(slice => slice.map(row => row.map(() => false)));

   const toExpore = [[0, 0, 0]];
   while (toExpore.length > 0) {
      const [x, y, z] = toExpore.pop();
      if (seen[x][y][z]) {
         continue;
      }

      seen[x][y][z] = true;
      for (const [dx, dy, dz] of transforms) {
         const xp = x + dx;
         const yp = y + dy;
         const zp = z + dz;
         if (xp < 0 || yp < 0 || zp < 0 || xp > maxX || yp > maxY || zp > maxZ) {
            continue;
         }

         if (map[xp][yp][zp]) {
            sides++;
            continue;
         }

         toExpore.push([xp, yp, zp]);
      }
   }

   return sides;
*/

// Addcube returns the exposed number of areas for the
// added cube
func (c *CubeListDay18) AddCube(cube *CubyDay18) int {
	exposed := 6
	for _, cubie := range c.Cubes {
		if IsAdjacent(cube, cubie) {
			exposed -= 1
		}
	}

	return exposed
}

func IsAdjacent(a, b *CubyDay18) bool {
	if (a.X == b.X && a.Y == b.Y && (a.Z == b.Z+1 || a.Z == b.Z-1)) ||
		(a.Y == b.Y && a.Z == b.Z && (a.X == b.X+1 || a.X == b.X-1)) ||
		(a.Z == b.Z && a.X == b.X && (a.Y == b.Y+1 || a.Y == b.Y-1)) {
		return true
	}
	return false
}

func (c *CubeListDay18) LavaAt(x, y, z int) bool {
	for _, cubie := range c.Cubes {
		if x == cubie.X &&
			y == cubie.Y &&
			z == cubie.Z {
			return true
		}
	}
	return false
}

func (c *CubeListDay18) MaxXYZ() (int, int, int) {
	maxX := 0
	maxY := 0
	maxZ := 0
	for _, cubie := range c.Cubes {
		if cubie.X > maxX {
			maxX = cubie.X
		}
		if cubie.Y > maxY {
			maxY = cubie.Y
		}
		if cubie.Z > maxZ {
			maxZ = cubie.Z
		}

	}

	return maxX + 2, maxY + 2, maxZ + 2
}
