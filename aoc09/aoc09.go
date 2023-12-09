package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := ReadFile("test.txt")
	numberSequences := make([][]int, 0)

	for _, line := range lines {
		numberSequences = append(numberSequences, parseNumberSequence(line))
	}

	// Puzzle 1
	sumNextValues := 0
	for _, numberSequence := range numberSequences {
		sumNextValues += getNextValue(numberSequence)
	}

	// Puzzle 2
	sumPreviousValues := 0
	for _, numberSequence := range numberSequences {
		sumPreviousValues += getPreviousValue(numberSequence)
	}

	fmt.Println("Puzzle 1: ", sumNextValues)
	fmt.Println("Puzzle 2: ", sumPreviousValues)
}

func getPreviousValue(sequence []int) int {
	differences := getDifferenceSequence(sequence)

	for idx := range differences {
		if differences[idx] != 0 {
			val := getPreviousValue(differences)
			return sequence[0] - val
		}
	}

	return sequence[0]
}

func getNextValue(sequence []int) int {
	differences := getDifferenceSequence(sequence)

	for idx := range differences {
		if differences[idx] != 0 {
			val := getNextValue(differences)
			return sequence[len(sequence)-1] + val
		}
	}

	return sequence[len(sequence)-1]
}

func getDifferenceSequence(sequence []int) []int {
	differenceSequence := make([]int, 0)
	for i := 0; i < len(sequence)-1; i++ {
		val1 := sequence[i]
		val2 := sequence[i+1]
		difference := val2 - val1
		differenceSequence = append(differenceSequence, difference)
	}
	return differenceSequence
}

func parseNumberSequence(line string) []int {
	sequence := make([]int, 0)
	for _, numberString := range strings.Split(line, " ") {
		value, _ := strconv.Atoi(numberString)
		sequence = append(sequence, value)
	}
	return sequence
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
