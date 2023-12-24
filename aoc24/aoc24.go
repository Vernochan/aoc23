package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const epsilon = 1e-9

func IsEqualFloat64(f1 float64, f2 float64) bool {
	return math.Abs(f1-f2) <= epsilon
}

type Point3d struct {
	X float64
	Y float64
	Z float64
}

func (p Point3d) String() string {
	return fmt.Sprintf("(%f,%f,%f)", p.X, p.Y, p.Z)
}

func (p Point3d) IsIn(points ...Point3d) bool {
	for _, x := range points {
		if x == p {
			return true
		}
	}
	return false
}

type Direction3d struct {
	X float64
	Y float64
	Z float64
}

func (d Direction3d) String() string {
	return fmt.Sprintf("(%f,%f,%f)", d.X, d.Y, d.Z)
}

type Line struct {
	Origin    Point3d
	Direction Direction3d
}

func (l Line) IntersectsXY(l2 Line) (Point3d, bool) {
	p2l := move(l.Origin, l.Direction, 1)

	ml := (p2l.Y - l.Origin.Y) / (p2l.X - l.Origin.X)
	// ml := l.Direction.X / l.Direction.Y
	bl := l.Origin.Y - ml*l.Origin.X

	//fmt.Println(l.Origin)

	//formula for line 1: y = ml * x + bl

	p22 := move(l2.Origin, l2.Direction, 1)

	ml2 := (p22.Y - l2.Origin.Y) / (p22.X - l2.Origin.X)

	//ml2 := l2.Direction.X / l2.Direction.Y
	bl2 := l2.Origin.Y - ml2*l2.Origin.X

	// lines are either equal or parallel
	if IsEqualFloat64(ml, ml2) {
		return Point3d{}, false
	}

	newX := (bl2 - bl) / (ml - ml2)
	newY := ml*newX + bl

	return Point3d{X: newX, Y: newY}, true
}

func (l Line) IsInFuture(p Point3d) bool {
	if l.Direction.X > 0 {
		if p.X > l.Origin.X {
			return true
		}
	} else {
		if p.X < l.Origin.X {
			return true
		}
	}

	return false
}

// var Left = Direction3d{X: 0, Y: -1}
// var Right = Direction3d{X: 0, Y: +1}
// var Up = Direction3d{X: -1, Y: 0}
// var Down = Direction3d{X: 1, Y: 0}

func main() {
	lines := ReadFile("test.txt")

	// since i use a simple/naive approach, it takes a while!

	fmt.Println("Puzzle1: ", puzzle1(lines))
	// Puzzle1:  2202
	fmt.Println("Puzzle2: ", puzzle2(lines))
	// Puzzle2:  6226
}
func puzzle1(lines []string) int {
	hail := parseLines(lines)

	count := 0
	testMin := 200000000000000.0
	testMax := 400000000000000.0

	for i := 0; i < len(hail); i++ {
		h1 := hail[i]

		for i2 := i; i2 < len(hail); i2++ {
			h2 := hail[i2]
			if h1 == h2 {
				continue
			}

			intersect, ok := h1.IntersectsXY(h2)
			if !ok {
				continue
			}

			if intersect.X > testMin && intersect.X < testMax &&
				intersect.Y > testMin && intersect.Y < testMax {
				if h1.IsInFuture(intersect) && h2.IsInFuture(intersect) {
					count++
				}
			}

		}
	}

	return count
}

func puzzle2(lines []string) int {
	//field := parseLines(lines)
	// could not figure out how to use z3 in go properly..
	fmt.Println(`puzzle 2 no solution in go, could not figure out how to use z3 in go`)
	fmt.Println(`please use "python ./aoc24.py | bc"`)
	fmt.Println(`Make sure z3 is installed! https://github.com/Z3Prover/z3?tab=readme-ov-file#python`)

	return 0
}

func move(p Point3d, d Direction3d, steps float64) Point3d {
	return Point3d{X: p.X + (d.X)*steps, Y: p.Y + (d.Y)*steps, Z: p.Z + (d.Z)*steps}
}

func parseLines(lines []string) []Line {
	parsedLines := make([]Line, 0, len(lines))
	for _, line := range lines {
		newLine := Line{}
		nums := strings.Split(line, " @ ")

		// num[0] == coord

		tmp := strings.Split(nums[0], ", ")

		pX, _ := strconv.Atoi(strings.TrimSpace(tmp[0]))
		pY, _ := strconv.Atoi(strings.TrimSpace(tmp[1]))
		pZ, _ := strconv.Atoi(strings.TrimSpace(tmp[2]))

		newLine.Origin.X = float64(pX)
		newLine.Origin.Y = float64(pY)
		newLine.Origin.Z = float64(pZ)

		// num[1] == velocity

		tmp = strings.Split(nums[1], ",")

		vX, _ := strconv.Atoi(strings.TrimSpace(tmp[0]))
		vY, _ := strconv.Atoi(strings.TrimSpace(tmp[1]))
		vZ, _ := strconv.Atoi(strings.TrimSpace(tmp[2]))

		newLine.Direction.X = float64(vX)
		newLine.Direction.Y = float64(vY)
		newLine.Direction.Z = float64(vZ)

		parsedLines = append(parsedLines, newLine)
	}
	return parsedLines
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
