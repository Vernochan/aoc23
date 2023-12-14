package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	lines := ReadFile("test.txt")
	field := parseLines(lines)
	tiltedField := field // tiltWest(field)

	//sum := countLoad(tiltedField)
	//lastcycle := 0
	for i := 0; i < 2000; i++ {
		tiltedField = tiltNorth(tiltedField)
		tiltedField = tiltWest(tiltedField)
		tiltedField = tiltSouth(tiltedField)
		tiltedField = tiltEast(tiltedField)

		sum := countLoad(tiltedField)
		// if sum == 100467 {
		// 	// repeats every 39 lines
		// 	fmt.Println("100467 at cycle", i, " diff: ", i-lastcycle)
		// 	lastcycle = i
		// }
		if i%39 == 1000000000%39 {
			fmt.Println(i, "  Load: ", sum)
		}
		//fmt.Println(i, "  Load: ", sum)
	}

	// fmt.Println("Load: ", sum)
	for _, row := range tiltedField {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Print("\n")
	}
}

func countLoad(field [][]rune) int {
	sum := 0
	for row := range field {
		for col := range field[row] {
			if field[row][col] == 'O' {
				sum += len(field) - row
			}
		}
	}
	return sum
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

func tiltNorth(field [][]rune) [][]rune {

	for col := 0; col < len(field[0]); col++ {
		for row := 1; row < len(field); row++ {
			prevRow := row - 1
			if field[row][col] == 'O' {
				//move up
				for prevRow >= 0 && field[prevRow][col] == '.' {
					field[prevRow][col] = 'O'
					field[prevRow+1][col] = '.'
					prevRow--
				}
			}
		}
	}

	return field
}

func tiltSouth(field [][]rune) [][]rune {

	for col := 0; col < len(field[0]); col++ {
		for row := len(field) - 2; row >= 0; row-- {
			prevRow := row + 1
			if field[row][col] == 'O' {
				//move up
				for prevRow < len(field) && field[prevRow][col] == '.' {
					field[prevRow][col] = 'O'
					field[prevRow-1][col] = '.'
					prevRow++
				}
			}
		}
	}

	return field
}

func tiltEast(field [][]rune) [][]rune {

	for row := 0; row < len(field); row++ {
		for col := len(field[row]) - 2; col >= 0; col-- {
			prevCol := col + 1
			if field[row][col] == 'O' {
				//move up
				for prevCol < len(field[row]) && field[row][prevCol] == '.' {
					field[row][prevCol] = 'O'
					field[row][prevCol-1] = '.'
					prevCol++
				}
			}
		}
	}

	return field
}

func tiltWest(field [][]rune) [][]rune {

	for row := 0; row < len(field); row++ {
		for col := 1; col < len(field[0]); col++ {
			prevCol := col - 1
			if field[row][col] == 'O' {
				//move up
				for prevCol >= 0 && field[row][prevCol] == '.' {
					field[row][prevCol] = 'O'
					field[row][prevCol+1] = '.'
					prevCol--
				}
			}
		}
	}

	return field
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
