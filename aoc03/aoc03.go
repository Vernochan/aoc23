package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// Intermediate format used while parsing Parts
type partNumber struct {
	number     int
	indexLeft  int
	indexRight int
}

// All information about an engine part, including it's coordinates (row,column)
type Part struct {
	partNumbers []int
	symbol      byte
	row         int
	column      int
}

func (p Part) String() string {
	return fmt.Sprintf("{%c, (%d,%d) %v}", p.symbol, p.row, p.column, p.partNumbers)
}

func main() {
	//lines := ReadFile("testFull.txt")
	lines := ReadFile("test2.txt")

	parts := parseParts(lines)

	sum := 0
	sum2 := 0

	for _, v := range parts {
		ratio := 1 // assignment defines the ratio as a product, starting with 1 as a neutral element for products

		for _, n := range v.partNumbers {
			sum += n
			ratio *= n
		}

		if v.symbol == '*' && len(v.partNumbers) == 2 {
			sum2 += ratio
		}
	}
	//fmt.Println(parts)
	fmt.Println("Sum of parts: ", sum)
	fmt.Println("Sum of gear ratios: ", sum2)
}

func isNumber(r byte) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

// Reads all contents of lines and generates a slice of parts with their associated part numbers
func parseParts(lines []string) []Part {

	parts, partNumbersByRows := readGlyphs(lines)

	for i := 0; i < len(parts); i++ {
		part := &parts[i] // reference is needed, since it will be modified

		//previous row, only applicable when part is not in row 0
		if part.row > 0 {
			for _, v := range partNumbersByRows[part.row-1] {
				if part.column >= v.indexLeft-1 &&
					part.column <= v.indexRight+1 {

					part.partNumbers = append(part.partNumbers, v.number)
				}
			}
		}
		//current row
		for _, v := range partNumbersByRows[part.row] {
			if v.indexLeft == part.column+1 ||
				v.indexRight == part.column-1 {

				part.partNumbers = append(part.partNumbers, v.number)
			}
		}

		// next row, only applicable when part is not in last row
		if part.row < len(partNumbersByRows) {
			for _, v := range partNumbersByRows[part.row+1] {
				if part.column >= v.indexLeft-1 &&
					part.column <= v.indexRight+1 {

					part.partNumbers = append(part.partNumbers, v.number)
				}
			}
		}
	}

	return parts
}

// Reads all contents of lines and generates parts and partNumbers.
// Numbers will be returned as a slice for every row in lines
// Parts are not matched with partNumbers!
//
// This function is used as an intermediate function for parseParts
func readGlyphs(lines []string) ([]Part, [][]partNumber) {
	parts := make([]Part, 0)

	partNumbersByRows := make([][]partNumber, 0)

	for indexRow, line := range lines {
		partNumbers := make([]partNumber, 0)

		numberCount := 0
		value := 0

		var glyph byte

		//for indexColumn, glyph := range line {
		for indexColumn := len(line) - 1; indexColumn >= 0; indexColumn-- {
			glyph = line[indexColumn]

			if !isNumber(glyph) {
				if numberCount > 0 {
					tmpPart := partNumber{number: value, indexLeft: indexColumn + 1, indexRight: indexColumn + numberCount}

					// append to the left, since we are traversing backwards and want to maintain order
					partNumbers = append([]partNumber{tmpPart}, partNumbers...)

					// reset
					numberCount = 0
					value = 0
				}
				if glyph == '.' {
					continue
				}
				newPart := Part{row: indexRow, column: indexColumn, symbol: glyph, partNumbers: make([]int, 0)}
				parts = append(parts, newPart)

				continue
			}

			value += int(glyph-'0') * int(math.Pow10(numberCount))
			numberCount++
		}

		// finish edge case (number starts at first byte)
		if numberCount > 0 {
			tmpPart := partNumber{number: value, indexLeft: 0, indexRight: numberCount - 1}

			// append to the left, since we are traversing backwards and want to maintain order
			partNumbers = append([]partNumber{tmpPart}, partNumbers...)
		}

		partNumbersByRows = append(partNumbersByRows, partNumbers)
	}

	return parts, partNumbersByRows
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)

	if err != nil {
		panic("error reading file")
	}

	content := strings.Split(string(f), "\n")

	return content
}
