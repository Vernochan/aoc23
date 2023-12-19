package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}
type Direction struct {
	X int
	Y int
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

	polygon := make([]Point, 0)

	// set start point
	currentPoint := Point{X: 0, Y: 0}

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		direction := getDirectionFromString(lineSplit[0])
		steps, _ := strconv.Atoi(lineSplit[1])
		currentPoint = move(currentPoint, direction, steps)
		polygon = append(polygon, currentPoint)

	}
	return getAreaFromPolygon(polygon)
}

func puzzle2(lines []string) int {

	startingPoint := Point{X: 0, Y: 0}
	polygon := make([]Point, 0)
	currentPoint := startingPoint
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		hexString := lineSplit[2]
		hexString = hexString[1 : len(hexString)-1]
		numberString := hexString[1 : len(hexString)-1]

		direction := getDirectionFromString(hexString[6:7])
		steps, _ := strconv.ParseInt(numberString, 16, 64)

		currentPoint = move(currentPoint, direction, int(steps))
		polygon = append(polygon, currentPoint)

	}
	return getAreaFromPolygon(polygon) // + sumSteps/2 + 1
}

func getDistance(p1 Point, p2 Point) int {

	x := p2.X - p1.X
	y := p2.Y - p1.Y
	return int(math.Sqrt(float64(x*x) + float64(y*y)))
}

// calculate the full area for any polygon
// https://en.wikipedia.org/wiki/Shoelace_formula
// since shoelace only counts field inside of the polygon,
// half of the perimeter (+1) is added as well
func getAreaFromPolygon(polygon []Point) int {

	sum := 0.0
	sumPerimeter := 0
	for i := 0; i < len(polygon); i++ {

		idx2 := (i + 1) % len(polygon) // wrap around and compare last with first
		p1 := polygon[i]
		p2 := polygon[idx2]

		sumPerimeter += getDistance(p1, p2)

		sum += float64(p1.X*p2.Y - p1.Y*p2.X)
	}
	return int(math.Abs((sum / 2))) + sumPerimeter/2 + 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move(p Point, d Direction, steps int) Point {
	return Point{X: p.X + (d.X)*steps, Y: p.Y + (d.Y)*steps}
}

func getDirectionFromString(input string) Direction {
	switch input {
	case "L":
		fallthrough
	case "2":
		fallthrough
	case "l":
		return Left
	case "R":
		fallthrough
	case "0":
		fallthrough
	case "r":
		return Right
	case "U":

		fallthrough
	case "3":
		fallthrough
	case "u":
		return Up
	case "D":
		fallthrough
	case "1":
		fallthrough
	case "d":
		return Down
	}

	return Direction{X: 0, Y: 0}
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
