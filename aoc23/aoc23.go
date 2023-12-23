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

func (p Point) isIn(points ...Point) bool {
	for _, x := range points {
		if x == p {
			return true
		}
	}
	return false
}

type Direction struct {
	X int
	Y int
}

func (p Direction) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

type Graph map[Point][]Point

var Left = Direction{X: 0, Y: -1}
var Right = Direction{X: 0, Y: +1}
var Up = Direction{X: -1, Y: 0}
var Down = Direction{X: 1, Y: 0}

func main() {
	lines := ReadFile("test.txt")

	// since i use a simple/naive approach, it takes a while!

	fmt.Println("Puzzle1: ", puzzle1(lines))
	// Puzzle1:  2202
	fmt.Println("Puzzle2: ", puzzle2(lines))
	// Puzzle2:  6226
}
func puzzle1(lines []string) int {
	field := parseLines(lines)
	graph := generateGraph(field, false)
	startingPosition := Point{X: 0, Y: 1}
	endPosition := Point{X: len(field) - 1, Y: len(field[0]) - 2}
	visited := make(map[Point]bool)
	return findLongestPath(graph, visited, startingPosition, endPosition)
}

func puzzle2(lines []string) int {
	field := parseLines(lines)
	graph := generateGraph(field, true)
	startingPosition := Point{X: 0, Y: 1}
	endPosition := Point{X: len(field) - 1, Y: len(field[0]) - 2}
	visited := make(map[Point]bool)
	return findLongestPath(graph, visited, startingPosition, endPosition)
}

func findLongestPath(g Graph, visited map[Point]bool, start Point, end Point) int {

	result := 0

	if start == end {
		return len(visited)
	}

	targets := g[start]
	for _, neighbor := range targets {
		if _, ok := visited[neighbor]; ok {
			continue
		}
		visited[neighbor] = true

		tmp := findLongestPath(g, visited, neighbor, end)
		result = max(tmp, result)
		delete(visited, neighbor)
	}

	return result
}

func move(p Point, d Direction, steps int) Point {
	return Point{X: p.X + (d.X)*steps, Y: p.Y + (d.Y)*steps}
}

func contractGraph(g map[Point][]Point, startingPoint Point) Graph {
	targetGraph := make(Graph)
	current := startingPoint
	end := Point{X: 22, Y: 21}
	for current != end {
		currentTargets := g[current]
		if len(currentTargets) == 1 {
			current = currentTargets[0]
		} else if len(currentTargets) == 2 {
			// look at targets of target 0
			current.isIn(g[currentTargets[0]]...)

			// look at targets of target 1
			current.isIn(g[currentTargets[1]]...)

			// if currentTargets[0] == current {
			// 	current = currentTargets[1]
			// } else if currentTargets[1] == current {
			// 	current = currentTargets[0]
			// } else {
			// 	targetGraph[current] = currentTargets
			// }

		}
		// else { // len(currentTargets) must be 3 now

		// }

	}
	return targetGraph
}

func generateGraph(field [][]rune, ignoreslopes bool) Graph {
	g := make(Graph)

	for x := 0; x < len(field); x++ {
		// assume square
		for y := 0; y < len(field[0]); y++ {
			key := Point{X: x, Y: y}
			targets := getTargets(field, key, ignoreslopes)
			if len(targets) > 0 {
				g[key] = targets
			}

		}
	}
	return g
}

func getTargets(field [][]rune, p Point, ignoreSlopes bool) (targets []Point) {
	targets = make([]Point, 0)

	if field[p.X][p.Y] == '#' {
		return targets
	}

	if !ignoreSlopes {
		switch field[p.X][p.Y] {
		case '>':
			targets = append(targets, move(p, Right, 1))
			return
		case '<':
			targets = append(targets, move(p, Left, 1))
			return
		case 'v':
			targets = append(targets, move(p, Down, 1))
			return
		case '^':
			targets = append(targets, move(p, Up, 1))
			return
		}
	}

	if p.X > 0 {
		key := move(p, Up, 1)
		if field[key.X][key.Y] != '#' { //&& field[key.X][key.Y] != 'v' {
			//targets[key] = field[key.X][key.Y]
			targets = append(targets, key)
		}
	}

	if p.Y > 0 {
		key := move(p, Left, 1)
		if field[key.X][key.Y] != '#' { //&& field[key.X][key.Y] != '>' {
			//targets[key] = field[key.X][key.Y]
			targets = append(targets, key)
		}
	}

	if p.X < len(field)-1 {
		key := move(p, Down, 1)
		if field[key.X][key.Y] != '#' { //&& field[key.X][key.Y] != '^' {
			//targets[key] = field[key.X][key.Y]
			targets = append(targets, key)
		}
	}

	if p.Y < len(field[0])-1 {
		key := move(p, Right, 1)
		if field[key.X][key.Y] != '#' { //&& field[key.X][key.Y] != '<' {
			//targets[key] = field[key.X][key.Y]
			targets = append(targets, key)
		}
	}
	return targets
}

func isIn(r rune, runes ...rune) bool {
	for _, x := range runes {
		if x == r {
			return true
		}
	}
	return false
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
