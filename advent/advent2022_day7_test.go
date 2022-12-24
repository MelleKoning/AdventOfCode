package advent

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// puzzle https://adventofcode.com/2022/day/7

// a File has a name and size
type FileContents struct {
	name string
	size int
}

// a dir can have multiple other dirs
// and also files
type DirContents struct {
	name      string
	dirs      []*DirContents
	files     []*FileContents
	parentDir *DirContents
	TotalSize int // total found size of files for this sole dircontents
	DirSize   int
}

func TestFind7_1(t *testing.T) {

	fileLines, err := GetFileLines("inputdata/input2022day7.txt")
	if err != nil {
		t.Fatalf("%v", err)
	}

	// let's keep track of the directory structure
	disk := &DirContents{
		parentDir: nil,
	}

	currentDir := disk
	for _, line := range fileLines {
		if string(line[0]) == "$" {
			fmt.Printf("command detected: %s\n", line)
			command := strings.Split(line, " ") // 0 is $, 1 is cd, 2 is the dir
			if command[1] == "cd" {
				if command[2] == "/" {
					// root is the created "disk" variable
					currentDir.name = "/"
					continue
				}

				if command[2] == ".." {
					// we have to go back up
					currentDir = currentDir.parentDir
					continue
				}
				// the dir should be known, find it and make it active
				for _, subDir := range currentDir.dirs {
					if subDir.name == command[2] {
						currentDir = subDir
						continue
					}
				}
				continue
			}
		} else {
			// we assume disk entries for the "currentDir", just add stuff

			diskentry := strings.Split(line, " ")
			if diskentry[0] == "dir" {
				// add this directory to currentdir
				currentDir.dirs = append(currentDir.dirs, &DirContents{name: diskentry[1], parentDir: currentDir})
				fmt.Printf("new directory %s added to %s\n", diskentry[1], currentDir.name)
			} else {
				// assume its a file
				size := GetNumberFromString(diskentry[0])
				filename := diskentry[1]
				currentDir.files = append(currentDir.files, &FileContents{name: filename, size: size})
				// we can also add the filesize to the totalsize for this dir
				fmt.Printf("added file %s(%d) to dir %s\n", diskentry[1], size, currentDir.name)
				currentDir.TotalSize += size
			}
		}
	}

	// now determine all filesize of a dir including its subfolder sizes

	// execute calculation
	RecursiveSizeOfDir(disk) // calculates DirSizes

	// find all dirs with a size of at most 100000 and add those to the total
	totalLess100k := RecursiveSearchLess100K(disk)
	fmt.Printf("Result: %d", totalLess100k)

	g := 0
	for _, dir := range totalLess100k {
		g = g + dir
	}
	fmt.Printf("Total Result: %d\n", g)

	// SECOND PART OF TEST
	// find closest dir
	// root has 46552309
	// 70 mln minus that is -> 23447691 current free space
	// we need 30 mln so have to free up at least 6552309
	// var alldirs map[int]string

	alldirs := ReturnAllDirs(disk)

	keys := make([]int, 0, len(alldirs))

	for key := range alldirs {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return alldirs[keys[i]] > alldirs[keys[j]]
	})

	for _, k := range keys {
		if k > 6552309 {
			fmt.Printf("just delete this folder to free up %d: %v", k, alldirs[k])
			break
		}
	}

}

func RecursiveSizeOfDir(dir *DirContents) int {
	totalSub := dir.TotalSize
	for _, subDir := range dir.dirs {
		totalSub = totalSub + RecursiveSizeOfDir(subDir)
	}
	dir.DirSize = totalSub
	fmt.Printf("dirsize %s %d\n", dir.name, totalSub)
	return totalSub
}

func RecursiveSearchLess100K(dir *DirContents) []int {
	var sizes []int
	if dir.DirSize <= 100000 {
		sizes = append(sizes, dir.DirSize)
	}
	for _, subDir := range dir.dirs {
		subsizes := RecursiveSearchLess100K(subDir)
		sizes = append(sizes, subsizes...)
	}

	return sizes

}

func ReturnAllDirs(dir *DirContents) map[int]string {
	dirs := make(map[int]string)
	dirs[dir.DirSize] = dir.name
	//sizes = append(sizes, fmt.Sprintf("%s %d", dir.name, dir.DirSize))
	for _, subDir := range dir.dirs {
		subsizes := ReturnAllDirs(subDir)
		for key, value := range subsizes {
			dirs[key] = value
		}
		//dirs = append(dirs, subsizes...)
	}
	return dirs
}
