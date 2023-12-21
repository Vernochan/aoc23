package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Direction struct {
	X int
	Y int
}

func (p Direction) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

var Left = Direction{X: 0, Y: -1}
var Right = Direction{X: 0, Y: +1}
var Up = Direction{X: -1, Y: 0}
var Down = Direction{X: 1, Y: 0}

func main() {
	lines := ReadFile("test.txt")

	fmt.Println("Puzzle1: ", puzzle1(lines))
	fmt.Println("Puzzle2: ", puzzle2(lines))
}
func puzzle1(lines []string) int {
	var startingPosition Point
	field := parseLines(lines)
	graph := generateGraph(field)
	for x := 0; x < len(field); x++ {
		for y := 0; y < len(field); y++ {
			if field[x][y] == 'S' {
				startingPosition = Point{X: x, Y: y}
			}
		}
	}

	maxSteps := 64

	currentTargets := graph[startingPosition]
	//getTargets(field, startingPosition)
	// currentPosition := startingPosition

	for i := 1; i < maxSteps; i++ {
		upcomingPositions := make(map[Point]rune, 0)
		for target := range currentTargets {
			tmpTargets := getTargets(field, target) //
			for tmpTarget, v := range tmpTargets {
				upcomingPositions[tmpTarget] = v
			}

		}
		currentTargets = upcomingPositions
	}

	//fmt.Println(currentTargets)
	return len(currentTargets)
}

func puzzle2(lines []string) int {
	stepTarget := 26501365
	stepValues := make([]int, 0)

	field := parseLines(lines)

	// this is the max amount of steps we really need to go to find sufficient values
	maxSteps := len(lines)*2 + len(lines)/2

	var startingPosition Point

	field2 := generateGridOfFields(field, (maxSteps/len(field))+1)
	graph := generateGraph(field2)

	// There are no rocks in the axis where the starting point is
	// After all the Steps, we get a maximum in all 4 cardinal directions
	// in the limit, this forms a diamond/rhombus shape (kind of like a tilted square)
	// That leads in the direction of a quadratic formula
	//
	// Another observation:
	// 26501365 = 202300 * len(field) + 65
	// => Starting Point + multiple of original length
	//
	// To get a sufficiently precise estimate for the quadratic formula,
	// we need fitting values.
	// Values are taken after having taking enough steps to get a count for
	// 1 field,        ( len/2           )
	// 9 fields (3x3)  ( len/2 +     len )
	// 25 fields (5x5) ( len/2 + 2 * len )
	// Getting more values would be preferable, but it takes enough time as it is
	// also, the result is accurate enough
	//
	// for manual checking
	// numbers for step:
	// step 65           : 3755
	// step 65 + 131     : 33494
	// step 65 + 131 * 2 : 92811
	//
	// Found Quadratic formula that fits:
	// 14789*x^2 + 14950*x + 3755
	// see https://www.wolframalpha.com/input?i=quadratic+fit+calculator&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3x%22%7D+-%3E%22%7B0%2C+1%2C+2%7D%22&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3y%22%7D+-%3E%22%7B3755%2C+33494%2C+92811%7D%22
	//

	//starting position is always exactly in the center
	startingPosition = Point{X: (len(field2) - 1) / 2, Y: (len(field2[0]) - 1) / 2}

	currentTargets := graph[startingPosition]
	getTargets(field2, startingPosition)

	// could be faster with concurrency
	for i := 1; i < maxSteps; i++ {
		if (i == len(lines)/2) || (i == (len(lines) * 3 / 2)) {
			stepValues = append(stepValues, len(currentTargets))
		}
		upcomingPositions := make(map[Point]rune, 0)

		for target := range currentTargets {
			tmpTargets := graph[target] //getTargets(field2, target) //
			for tmpTarget, v := range tmpTargets {
				upcomingPositions[tmpTarget] = v
			}

		}
		currentTargets = upcomingPositions

	}
	stepValues = append(stepValues, len(currentTargets))

	// targetSteps is multiple oflen(field) + 65
	// so we need to evaluate the new quadratic formula for
	targetX := stepTarget / len(field) // int division

	data := []struct{ X, Y float64 }{
		{float64(0), float64(stepValues[0])},
		{float64(1), float64(stepValues[1])},
		{float64(2), float64(stepValues[2])},
	}

	// find coefficients for quadratic formula
	f2, f1, f0 := Quadratic(data)

	return int(f2)*targetX*targetX + int(f1)*targetX + int(f0)

}

// https://pkg.go.dev/github.com/soniakeys/meeus/v3@v3.0.1/fit
// i know i could have imported it, but i wanted to have it all in one place

// Quadratic fits y = ax² + bx + c to sample data.
//
// Argument p is a list of data points.  Results a, b, and c are coefficients
// of the best fit quadratic y = ax² + bx + c.
func Quadratic(p []struct{ X, Y float64 }) (a, b, c float64) {
	var P, Q, R, S, T, U, V float64
	for i := range p {
		x := p[i].X
		y := p[i].Y
		x2 := x * x
		P += x
		Q += x2
		R += x * x2
		S += x2 * x2
		T += y
		U += x * y
		V += x2 * y
	}
	N := float64(len(p))
	// (4.5) p. 43
	D := N*Q*S + 2*P*Q*R - Q*Q*Q - P*P*S - N*R*R
	// (4.6) p. 43
	a = (N*Q*V + P*R*T + P*Q*U - Q*Q*T - P*P*V - N*R*U) / D
	b = (N*S*U + P*Q*V + Q*R*T - Q*Q*U - P*S*T - N*R*V) / D
	c = (Q*S*T + Q*R*U + P*R*V - Q*Q*V - P*S*U - R*R*T) / D
	return
}

func move(p Point, d Direction, steps int) Point {
	return Point{X: p.X + (d.X)*steps, Y: p.Y + (d.Y)*steps}
}

func generateGraph(field [][]rune) map[Point]map[Point]rune {
	g := make(map[Point]map[Point]rune)

	for x := 0; x < len(field); x++ {
		// assume square
		for y := 0; y < len(field[0]); y++ {
			key := Point{X: x, Y: y}
			targets := getTargets(field, key)

			g[key] = targets

		}
	}
	return g
}

func generateGridOfFields(field [][]rune, copies int) [][]rune {
	length := copies*2 + 1
	newField := make([][]rune, length*len(field))

	for x := 0; x < len(newField); x++ {
		newField[x] = make([]rune, len(field[0])*length)
		for y := 0; y < len(newField[x]); y++ {
			newField[x][y] = field[x%len(field)][y%len(field[0])]

		}

	}
	return newField
}

func getTargets(field [][]rune, p Point) map[Point]rune {
	targets := make(map[Point]rune, 0)
	if p.X > 0 {
		key := move(p, Up, 1)
		if field[key.X][key.Y] != '#' {
			targets[key] = field[key.X][key.Y]
		}
	}

	if p.Y > 0 {
		key := move(p, Left, 1)
		if field[key.X][key.Y] != '#' {
			targets[key] = field[key.X][key.Y]
		}
	}

	if p.X < len(field)-1 {
		key := move(p, Down, 1)
		if field[key.X][key.Y] != '#' {
			targets[key] = field[key.X][key.Y]
		}
	}

	if p.Y < len(field[0])-1 {
		key := move(p, Right, 1)
		if field[key.X][key.Y] != '#' {
			targets[key] = field[key.X][key.Y]
		}
	}
	return targets
}

func parseLines(lines []string) [][]rune {
	runes := make([][]rune, 0, len(lines))
	for _, line := range lines {
		newRunes := make([]rune, 0, len(line))
		for _, r := range line {
			newRunes = append(newRunes, r)
		}
		runes = append(runes, newRunes)
	}
	return runes
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
