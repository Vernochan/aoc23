package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// List of colors with max availability for regular games
var colors map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines := ReadFile("test.txt")

	var sum, sumMinPower int

	for _, line := range lines {

		valid, value := checkValidity(line)

		if valid {
			sum += value
		}

		sumMinPower += getMinPower(line)

	}

	fmt.Println("Sum for numbers: ", sum)
	fmt.Println("Sum for numbers with text: ", sumMinPower)
}

func getMinPower(line string) int {

	indexColon := strings.Index(line, ":")

	gamesLine := line[indexColon+1:]

	var minimums map[string]int = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	games := strings.Split(gamesLine, ";")

	for _, game := range games {

		draws := strings.Split(game, ",")

		for _, draw := range draws {

			for color := range colors {

				if !strings.Contains(draw, color) {
					continue
				}

				// remove color from draw and remove whitespace
				numberString := strings.Replace(draw, color, "", -1)
				numberString = strings.TrimSpace(numberString)

				drawNumber, err := strconv.Atoi(numberString)

				if err != nil {
					panic("cannot convert draw number")
				}

				if drawNumber > minimums[color] {
					minimums[color] = drawNumber
				}
			}
		}

	}

	power := 1
	for _, v := range minimums {
		power *= v
	}
	return power
}

func checkValidity(line string) (bool, int) {
	indexColon := strings.Index(line, ":")

	game := line[5:indexColon]
	game = strings.TrimSpace(game)

	gameNumber, err := strconv.Atoi(game)

	if err != nil {
		panic("cannot convert game number")
	}

	gameLine := line[indexColon+1:]

	games := strings.Split(gameLine, ";")

	for _, game := range games {
		draws := strings.Split(game, ",")

		for _, draw := range draws {

			for color, maxValue := range colors {

				if strings.Contains(draw, color) {

					// remove color from draw and remove whitespace
					numberString := strings.Replace(draw, color, "", -1)
					numberString = strings.TrimSpace(numberString)

					val, err := strconv.Atoi(numberString)
					if err != nil {
						panic("cannot convert draw number")
					}

					if val > maxValue {
						return false, gameNumber
					}
				}
			}
		}

	}

	return true, gameNumber
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
