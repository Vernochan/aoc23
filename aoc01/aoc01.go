package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var numbers map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	lines := ReadFile("test.txt")
	var sum, sumText int
	for _, line := range lines {
		sum += getCoordinate(line)
		sumText += getTextCoordinate(line)

	}

	fmt.Println("Sum for numbers: ", sum)
	fmt.Println("Sum for numbers with text: ", sumText)
}

func getTextCoordinate(line string) int {
	var indexFirst, indexLast int
	var first, last int
	indexFirst = math.MaxInt
	for numberText, numberValue := range numbers {
		currentIndexFirst := strings.Index(line, numberText)
		currentIndexLast := strings.LastIndex(line, numberText)
		if currentIndexFirst >= 0 && currentIndexFirst <= indexFirst {
			indexFirst = currentIndexFirst
			first = numberValue
		}
		if currentIndexLast >= 0 && currentIndexLast >= indexLast {
			indexLast = currentIndexLast
			last = numberValue
		}

	}

	return (first*10 + last)
}

func getCoordinate(line string) int {
	var first, last int
	for _, c := range line {

		i, err := strconv.Atoi(string(c))

		if err != nil {
			continue
		}

		if first == 0 {
			first = i
		}

		last = i
	}
	return (first*10 + last)
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		os.Exit(1)
	}
	content := strings.Split(string(f), "\n")
	return content

}
