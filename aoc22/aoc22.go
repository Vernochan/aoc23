package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point3d struct {
	X int
	Y int
	Z int
}

type Block [2]Point3d

func (p Point3d) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

type Blocks []Block

func (blocks Blocks) Swap(i, j int) {
	blocks[i][0].X, blocks[i][0].Y, blocks[i][0].Z, blocks[i][1].X, blocks[i][1].Y, blocks[i][1].Z,
		blocks[j][0].X, blocks[j][0].Y, blocks[j][0].Z, blocks[j][1].X, blocks[j][1].Y, blocks[j][1].Z =
		blocks[j][0].X, blocks[j][0].Y, blocks[j][0].Z, blocks[j][1].X, blocks[j][1].Y, blocks[j][1].Z,
		blocks[i][0].X, blocks[i][0].Y, blocks[i][0].Z, blocks[i][1].X, blocks[i][1].Y, blocks[i][1].Z
}

// only Sort by Z-Position
func (blocks Blocks) Less(i, j int) bool {

	if blocks[i][0].Z < blocks[j][0].Z {
		return true
	} else if blocks[i][0].Z == blocks[j][0].Z {
		if blocks[i][0].Y < blocks[j][0].Y {
			return true
		}
	} else if blocks[i][0].Z == blocks[j][0].Z && blocks[i][0].Y == blocks[j][0].Y {
		if blocks[i][0].X < blocks[j][0].X {
			return true
		}
	}
	return false
	//return blocks[i][0].Z < blocks[j][0].Z || blocks[i][0].Y < blocks[j][0].Y || blocks[i][0].X < blocks[j][0].X
}

func (blocks Blocks) Len() int {
	return len(blocks)
}

func main() {
	lines := ReadFile("test2.txt")

	fmt.Println("Puzzle1: ", puzzle1(lines))
	fmt.Println("Puzzle2: ", puzzle2(lines))
}
func puzzle1(lines []string) int {
	blocks := parseLines(lines)

	fmt.Println(blocks)

	sort.Sort(blocks)

	fmt.Println(blocks)

	zBuffer := make(map[string]int, 0)
	for i := 0; i < len(blocks); i++ {
		// look at Z Position,
		// look if there is space underneath
		// move down as far as possible
		p1 := blocks[i][0]
		p2 := blocks[i][1]

		diffX := p2.X - p1.X
		diffY := p2.Y - p1.Y
		diffZ := p2.Z - p1.Z
		// same X Position: Only Cycle Through all Y Positions

		if diffZ != 0 {
			// only check for Z Position
			fmt.Println("DifferentZ", blocks[i])
			x := p1.X
			y := p1.Y
			idx := fmt.Sprintf("(%d,%d)", x, y)
			// zBuffer[idx]
			// if both positions have room underneath
			fmt.Println("Buffer at ", idx, ": ", zBuffer[idx]+1)

		} else if diffX != 0 {
			// check for all X Positions
			fmt.Println("DifferentX", blocks[i])

			for i := 0; i <= diffX; i++ {
				x := p1.X + i
				y := p1.Y
				idx := fmt.Sprintf("(%d,%d)", x, y)
				// zBuffer[idx]
				// if both positions have room underneath
				fmt.Println("Buffer at ", idx, ": ", zBuffer[idx]+1)
				// if zBuffer[idx] {

				// }

			}

		} else if diffY != 0 {
			// check for all Y Positions
			fmt.Println("DifferentY", blocks[i])
			newLowest := p1.Z

			for i := 0; i <= diffY; i++ {
				x := p1.X
				y := p1.Y + i
				idx := fmt.Sprintf("(%d,%d)", x, y)
				bufferHeight := zBuffer[idx]
				fmt.Println("Buffer at ", idx, ": ", bufferHeight)
				if p1.Z > bufferHeight {
					newLowest = bufferHeight
				}

			}
			fmt.Println("new lowest: ", newLowest)

		}

	}
	return len(blocks)
}

func puzzle2(lines []string) int {
	return 0
}

func parseLines(lines []string) Blocks {

	blocks := make(Blocks, 0)
	for _, line := range lines {
		coords := strings.Split(line, "~")
		newBlock := Block{}
		for i, coord := range coords {
			numStrings := strings.Split(coord, ",")
			x, _ := strconv.Atoi(numStrings[0])
			y, _ := strconv.Atoi(numStrings[1])
			z, _ := strconv.Atoi(numStrings[2])
			newBlock[i].X = x
			newBlock[i].Y = y
			newBlock[i].Z = z
		}

		blocks = append(blocks, newBlock)

	}
	return blocks
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
