package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type scratchCard struct {
	cardNumber     int
	winningNumbers []int
	numbers        []int
	matches        int
}

func main() {
	lines := ReadFile("test.txt")

	scratchCards := make([]scratchCard, 0)

	numCopies := make([]int, 0)

	for _, line := range lines {
		scratchCards = append(scratchCards, readScratchCard(line))

		numCopies = append(numCopies, 1)
	}

	sumByCards := 0
	for _, card := range scratchCards {
		if card.matches < 1 {
			continue
		}

		sumByCards += 1 << (card.matches - 1)
	}

	for idx := range scratchCards {

		// numCopies always has the same number of elements than scratchCards and has the same order
		copies := numCopies[idx]

		for i := 0; i < copies; i++ {
			for match := 1; match <= scratchCards[idx].matches; match++ {
				if len(numCopies) > (idx + match) {
					numCopies[idx+match]++
				}
			}
		}
	}

	sum2 := 0
	for _, v := range numCopies {
		sum2 += v
	}

	fmt.Println("Sum for numbers: ", sumByCards)
	fmt.Println("Count Numbers: ", sum2)
}

func readScratchCard(line string) scratchCard {
	cardPart := strings.Split(line, ":")

	gameNumberString := strings.Replace(cardPart[0], "Card", "", -1)

	cardNumber, _ := strconv.Atoi(gameNumberString)

	gamePart := strings.Split(cardPart[1], "|")

	winningNumbersStrings := strings.Split(strings.TrimSpace(gamePart[0]), " ")

	gameNumbersString := strings.Split(strings.TrimSpace(gamePart[1]), " ")

	winningNumbers := make([]int, 0)

	gameNumbers := make([]int, 0)

	for _, v := range winningNumbersStrings {
		val, _ := strconv.Atoi(v)
		if val != 0 {
			winningNumbers = append(winningNumbers, val)
		}

	}

	for _, v := range gameNumbersString {
		val, _ := strconv.Atoi(v)
		if val != 0 {
			gameNumbers = append(gameNumbers, val)
		}

	}

	matches := 0
	for _, v := range gameNumbers {
		if slices.Contains(winningNumbers, v) {
			matches++
		}

	}

	tmp := scratchCard{cardNumber: cardNumber, winningNumbers: winningNumbers, numbers: gameNumbers, matches: matches}

	return tmp
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
