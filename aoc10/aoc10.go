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

var pipeMapRaw [][]rune

func main() {
	lines := ReadFile("test.txt")

	pipeMapRaw = parsePipeMap(lines)

	var startPos Point

	for pipeLine := range pipeMapRaw {
		for pipeRow := range pipeMapRaw[pipeLine] {
			if pipeMapRaw[pipeLine][pipeRow] == 'S' {
				startPos = Point{x: pipeLine, y: pipeRow}
			}
		}
	}

	fmt.Printf("StartPos: (%3d,%3d): %c\n", startPos.x, startPos.y, pipeMapRaw[startPos.x][startPos.y])

	loop := findLoop(startPos, pipeMapRaw)

	trueStartSymbol := getTrueStartSymbol(loop, pipeMapRaw)

	pipeMapRaw[startPos.x][startPos.y] = trueStartSymbol

	countInside := 0

	for row := range pipeMapRaw {

		isInside := false
		lastFlipRune := '0'
		for column := range pipeMapRaw[row] {
			val := pipeMapRaw[row][column]
			p := Point{x: row, y: column}

			if !isPartOfLoop(p, loop) {
				if isInside {
					countInside++
				}
				continue
			}

			// check if next till will be inside or not

			// |      : always change side
			// F----7 : Edge along the line: Change side 2x (result: do not change)
			// F----J : diagonal line: Change side
			// L----J : Edge along the line: Change side 2x (result: do not change)
			// L----7 : diagonal line: Change side
			// only works, because - is not counted anyway, number of - can vary (even be 0)

			if val == '|' {
				isInside = !isInside
				lastFlipRune = val
			}

			if isIn(val, 'L', 'F') {
				isInside = !isInside
				lastFlipRune = val
			}

			if val == '7' && lastFlipRune == 'F' {
				isInside = !isInside
				lastFlipRune = val
			}

			if val == 'J' && lastFlipRune == 'L' {
				isInside = !isInside
				lastFlipRune = val
			}

		}
		//fmt.Print("\n")
	}

	fmt.Println("Puzzle 1", len(loop)/2)
	fmt.Println("Puzzle 2: ", countInside)

}

func getTrueStartSymbol(loop []Point, pipeMap [][]rune) rune {
	next := loop[1]
	previous := loop[len(loop)-1]
	var retVal rune

	if pipeMap[next.x][next.y] == pipeMap[previous.x][previous.y] {
		retVal = pipeMap[next.x][next.y]
		return retVal
	}

	if next.x > previous.x {
		//left is below
		if next.y > previous.y {
			retVal = '7'
		} else {
			retVal = 'F'
		}
	} else {
		//left is above
		if next.y > previous.y {
			retVal = 'J'
		} else {
			retVal = 'L'
		}
	}

	return retVal
}

func findLoop(startPos Point, pipeMap [][]rune) []Point {
	loop := make([]Point, 0)
	loop = append(loop, startPos)

	connections := getConnections(startPos, pipeMapRaw)

	prevPos := startPos
	curPos := connections[0]

	for curPos != startPos {
		loop = append(loop, curPos)
		nextPositions := getConnections(curPos, pipeMapRaw)
		if nextPositions[0] == prevPos {
			prevPos = curPos
			curPos = nextPositions[1]

		} else { // nextConnections[1] == prePos
			prevPos = curPos
			curPos = nextPositions[0]
		}

	}

	return loop
}

func isPartOfLoop(p Point, loop []Point) bool {
	for _, x := range loop {
		if x == p {
			return true
		}
	}
	return false
}

func isIn(r rune, runes ...rune) bool {
	for _, x := range runes {
		if x == r {
			return true
		}
	}
	return false
}

func getConnections(p Point, pipeMap [][]rune) []Point {
	connections := make([]Point, 0)
	row := p.x
	column := p.y

	curPipe := pipeMap[row][column]

	//up
	if row > 0 {
		nextPipe := pipeMap[row-1][column]
		if isIn(curPipe, '|', 'J', 'L', 'S') && (isIn(nextPipe, '|', 'F', '7') || nextPipe == 'S') {
			connections = append(connections, Point{x: row - 1, y: column})
		}
	}

	// down
	if row < (len(pipeMap) - 1) {
		nextPipe := pipeMap[row+1][column]
		if isIn(curPipe, '|', 'F', '7', 'S') && (isIn(nextPipe, '|', 'J', 'L') || nextPipe == 'S') {
			connections = append(connections, Point{x: row + 1, y: column})
		}
	}

	// left
	if column > 0 {
		nextPipe := pipeMap[row][column-1]
		if isIn(curPipe, '-', 'J', '7', 'S') && (isIn(nextPipe, '-', 'F', 'L') || nextPipe == 'S') {
			connections = append(connections, Point{x: row, y: column - 1})
		}
	}

	// right
	if column < (len(pipeMap[row]) - 1) {
		nextPipe := pipeMap[row][column+1]
		if isIn(curPipe, '-', 'F', 'L', 'S') && (isIn(nextPipe, '-', 'J', '7') || nextPipe == 'S') {
			connections = append(connections, Point{x: row, y: column + 1})
		}
	}

	return connections
}

func parsePipeMap(lines []string) [][]rune {
	newPipeMap := make([][]rune, 0)

	for _, line := range lines {
		pipeMapRow := make([]rune, 0)
		for _, pipe := range line {
			pipeMapRow = append(pipeMapRow, pipe)
		}
		newPipeMap = append(newPipeMap, pipeMapRow)
	}
	return newPipeMap
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
