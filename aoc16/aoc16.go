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

type Beam struct {
	X         int
	Y         int
	Direction int
}

const (
	Left = iota
	Right
	Up
	Down
)

func main() {
	lines := ReadFile("test.txt")
	field := parseLines(lines)

	// Puzzle 1
	startingPoint := Beam{X: 0, Y: 0, Direction: Right}
	sum := calcBeamLength(field, startingPoint)

	// Puzzle 2
	maxSum := 0

	// calc for all Left and Right Entry Points
	for i := 0; i < len(field); i++ {

		p1 := Beam{X: i, Y: 0, Direction: Right}
		p2 := Beam{X: i, Y: len(field[i]) - 1, Direction: Left}

		num1 := calcBeamLength(field, p1)
		maxSum = max(maxSum, num1)
		num2 := calcBeamLength(field, p2)
		maxSum = max(maxSum, num2)

	}

	// calc for all Up and Down Entry Points
	for i := 0; i < len(field[0]); i++ {

		p1 := Beam{X: 0, Y: i, Direction: Down}
		p2 := Beam{X: len(field) - 1, Y: i, Direction: Up}

		num1 := calcBeamLength(field, p1)
		maxSum = max(maxSum, num1)
		num2 := calcBeamLength(field, p2)
		maxSum = max(maxSum, num2)

	}

	fmt.Println("Puzzle 1: ", sum)
	fmt.Println("Puzzle 2: ", maxSum)

}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func calcBeamLength(field [][]rune, startingBeam Beam) int {

	visitedPoints := make(map[Point]int, 0)
	visitedBeams := make(map[Beam]int, 0)

	visitedPoints[getPoint(startingBeam)]++
	visitedBeams[startingBeam]++

	currentBeams := findNextBeam(field, startingBeam)

	for len(currentBeams) > 0 {
		nextBeams := make([]Beam, 0)
		for i := range currentBeams {
			// if beam staus within field
			if currentBeams[i].X >= 0 && currentBeams[i].Y >= 0 && currentBeams[i].X < len(field) && currentBeams[i].Y < len(field[0]) {
				visitedPoints[getPoint(currentBeams[i])]++
				if _, ok := visitedBeams[currentBeams[i]]; ok {
					// if a beam was already visited, it's a loop. do not count again
					continue
				}
				visitedBeams[currentBeams[i]]++
				nextBeams = append(nextBeams, findNextBeam(field, currentBeams[i])...)
			}
		}
		currentBeams = nextBeams
	}
	return len(visitedPoints)
}

func getPoint(p Beam) Point {
	return Point{X: p.X, Y: p.Y}
}

func getNextBeam(p Beam) Beam {
	var newPoint Beam
	switch p.Direction {
	case Right:
		newPoint = Beam{X: p.X, Y: p.Y + 1, Direction: Right}
	case Left:
		newPoint = Beam{X: p.X, Y: p.Y - 1, Direction: Left}
	case Up:
		newPoint = Beam{X: p.X - 1, Y: p.Y, Direction: Up}
	case Down:
		newPoint = Beam{X: p.X + 1, Y: p.Y, Direction: Down}
	}
	return newPoint
}

func findNextBeam(field [][]rune, p Beam) []Beam {

	nextPoints := make([]Beam, 0, 4)
	if p.X >= len(field) || p.Y >= len(field[0]) || p.X < 0 || p.Y < 0 {
		return nextPoints
	}
	if field[p.X][p.Y] == '.' {
		nextPoints = append(nextPoints, getNextBeam(p))
	}

	if field[p.X][p.Y] == '|' {
		if p.Direction == Right || p.Direction == Left {
			p.Direction = Up
			nextPoints = append(nextPoints, getNextBeam(p))
			p.Direction = Down
			nextPoints = append(nextPoints, getNextBeam(p))
		} else {
			nextPoints = append(nextPoints, getNextBeam(p))
		}
	}

	if field[p.X][p.Y] == '-' {
		if p.Direction == Up || p.Direction == Down {
			p.Direction = Left
			nextPoints = append(nextPoints, getNextBeam(p))
			p.Direction = Right
			nextPoints = append(nextPoints, getNextBeam(p))
		} else {
			nextPoints = append(nextPoints, getNextBeam(p))
		}
	}

	if field[p.X][p.Y] == '/' {
		switch p.Direction {
		case Up:
			p.Direction = Right
			nextPoints = append(nextPoints, getNextBeam(p))
		case Down:
			p.Direction = Left
			nextPoints = append(nextPoints, getNextBeam(p))
		case Left:
			p.Direction = Down
			nextPoints = append(nextPoints, getNextBeam(p))
		case Right:
			p.Direction = Up
			nextPoints = append(nextPoints, getNextBeam(p))
		}
	}
	if field[p.X][p.Y] == '\\' {
		switch p.Direction {
		case Up:
			p.Direction = Left
			nextPoints = append(nextPoints, getNextBeam(p))
		case Down:
			p.Direction = Right
			nextPoints = append(nextPoints, getNextBeam(p))
		case Left:
			p.Direction = Up
			nextPoints = append(nextPoints, getNextBeam(p))
		case Right:
			p.Direction = Down
			nextPoints = append(nextPoints, getNextBeam(p))
		}
	}
	return nextPoints
}

func parseLines(lines []string) [][]rune {
	runes := make([][]rune, 0)
	for _, line := range lines {
		newRunes := make([]rune, 0)
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
