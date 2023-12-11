package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	lines := ReadFile("test.txt")

	galaxies := parseGalaxies(lines)

	nonEmptyRows, nonEmptyColumns := getNonEmptyRowAndColumns(lines)

	sumDistance := 0
	sumDistance2 := 0

	defaultExpansionFactor := 2
	superExpansionFactor := 1000000

	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			// only compare to galaxies that are in a later row/column

			if galaxy1 == galaxy2 {
				continue
			}
			if galaxy2.x < galaxy1.x {
				continue
			}
			if galaxy1.x == galaxy2.x && galaxy2.y < galaxy1.y {
				continue
			}
			// account for expansion
			//Puzzle 1
			p1 := expandPoint(galaxy1, nonEmptyRows, nonEmptyColumns, defaultExpansionFactor)
			p2 := expandPoint(galaxy2, nonEmptyRows, nonEmptyColumns, defaultExpansionFactor)

			sumDistance += getDistance(p1, p2)

			//Puzzle 2
			p1 = expandPoint(galaxy1, nonEmptyRows, nonEmptyColumns, superExpansionFactor)
			p2 = expandPoint(galaxy2, nonEmptyRows, nonEmptyColumns, superExpansionFactor)

			sumDistance2 += getDistance(p1, p2)
		}
	}

	fmt.Println("Puzzle 1: ", sumDistance)
	fmt.Println("Puzzle 2: ", sumDistance2)
	//fmt.Println(galaxies)
}

func expandPoint(p Point, nonEmptyRows map[int]bool, nonEmptyColumns map[int]bool, expansionFactor int) Point {

	output := Point{x: p.x, y: p.y}
	for i := 0; i < len(nonEmptyRows); i++ {
		if nonEmptyRows[i] {
			continue
		}
		if i < p.x {
			output.x += expansionFactor - 1
		}
	}

	for i := 0; i < len(nonEmptyColumns); i++ {
		if nonEmptyColumns[i] {
			continue
		}
		if i < p.y {
			output.y += expansionFactor - 1
		}
	}
	return output
}

// Distance in this case means "steps to reach destination" without allowing diagonal steps
func getDistance(p1 Point, p2 Point) int {
	x := p2.x - p1.x
	if x < 0 {
		x *= -1
	}
	y := p2.y - p1.y
	if y < 0 {
		y *= -1
	}
	return x + y
}

func getNonEmptyRowAndColumns(lines []string) (map[int]bool, map[int]bool) {
	rows := make(map[int]bool, 0)
	columns := make(map[int]bool, 0)

	for row := range lines {
		rowHasGalaxy := false
		for column := range lines[row] {
			if lines[row][column] != '.' {
				rowHasGalaxy = true
				columns[column] = true
			} else {
				_, x := columns[column]
				if !x {
					columns[column] = false
				}
			}
		}

		rows[row] = rowHasGalaxy
	}

	return rows, columns
}

func parseGalaxies(lines []string) []Point {
	galaxies := make([]Point, 0)

	for row := range lines {
		for column := range lines[row] {
			if lines[row][column] == '#' {
				galaxies = append(galaxies, Point{x: row, y: column})
			}
		}
	}
	return galaxies
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
