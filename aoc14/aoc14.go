package main

import (
	"fmt"
	"os"
	"strings"
)

func findLongestRepeatingSequence(numbers []int) []int {
	for length := len(numbers); length > 0; length-- {
		for start := 0; start <= len(numbers)-length; start++ {
			sequence := numbers[start : start+length]
			for i := start + 1; i <= len(numbers)-length; i++ {
				if sequenceEqual(sequence, numbers[i:i+length]) {
					return sequence
				}
			}
		}
	}
	return nil
}
func sequenceEqual(seq1, seq2 []int) bool {
	if len(seq1) != len(seq2) {
		return false
	}
	for i, value := range seq1 {
		if value != seq2[i] {
			return false
		}
	}
	return true
}

func main() {

	lines := ReadFile("test.txt")
	field := parseLines(lines)
	tiltedField := field // tiltWest(field) (puzzle 1)

	sums := make([]int, 0)
	// add 0 to maintain correct position
	sums = append(sums, 0)

	seenBoards := make(map[string]int, 0)
	board := ""

	cycleLength := 0 //39
	i := 1
	for i < 1000000000 {
		tiltedField = tiltNorth(tiltedField)
		tiltedField = tiltWest(tiltedField)
		tiltedField = tiltSouth(tiltedField)
		tiltedField = tiltEast(tiltedField)

		sum := countLoad(tiltedField)

		sums = append(sums, sum)

		board = generateHash(field)

		if _, ok := seenBoards[board]; ok {
			break
		}

		seenBoards[board] = i
		i++
	}

	repeatedIndex := seenBoards[board]
	cycleLength = i - repeatedIndex
	idx := repeatedIndex + (1000000000-repeatedIndex)%cycleLength

	fmt.Println("Cycle: ", cycleLength)
	fmt.Println("Sum: ", sums[idx])
}

func generateHash(field [][]rune) string {
	result := ""
	for _, row := range field {
		result += string(row)
	}
	return result
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
