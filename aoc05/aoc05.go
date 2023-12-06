package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type conversionMap struct {
	destinationStart    []int
	sourceStart         []int
	lengths             []int
	sourceMaterial      string
	destinationMaterial string
}

var conversionMaps map[string]conversionMap

func main() {
	lines := ReadFile("test.txt")

	seeds := make([]int, 0)

	seedsText := strings.Split(lines[0][7:], " ")

	lines = lines[2:]

	conversionMaps = make(map[string]conversionMap, 0)
	for len(lines) > 0 {
		conversionMap, mapLength := getMap(lines)
		conversionMaps[conversionMap.sourceMaterial] = conversionMap

		// check if there are more maps
		if len(lines) > mapLength {
			lines = lines[mapLength+1:]
		} else {
			lines = lines[mapLength:]
		}
	}

	firstSeed, _ := strconv.Atoi(seedsText[0])
	lowestDestination := firstSeed

	for idx := 0; idx < len(seedsText); idx += 2 {
		seed := seedsText[idx]
		seedCount := seedsText[idx+1]

		nSeed, _ := strconv.Atoi(seed)
		nSeedCount, _ := strconv.Atoi(seedCount)

		// print something to show that it's still working
		fmt.Println("Adding seeds: ", nSeedCount)

		for i := 0; i < nSeedCount; i++ {
			seedDestination := getSeedLocation(nSeed + i)
			if seedDestination < lowestDestination {
				lowestDestination = seedDestination
			}
		}
	}
	fmt.Println("total seeds", len(seeds))

	fmt.Println("lowest destination: ", lowestDestination)
}

func getSeedLocation(seed int) int {
	destinationNumber := seed
	source := "seed"

	// the last conversion map ends in "location"
	for source != "location" {
		cm := conversionMaps[source]
		newDest := destinationNumber
		for idx := range cm.sourceStart {
			if destinationNumber >= cm.sourceStart[idx] && destinationNumber < cm.sourceStart[idx]+cm.lengths[idx] {
				diff := cm.sourceStart[idx] - cm.destinationStart[idx]
				newDest = destinationNumber - diff
			}
		}

		source = cm.destinationMaterial
		destinationNumber = newDest
	}
	return destinationNumber
}

func getMap(lines []string) (conversionMap, int) {
	rowIndex := 0

	sourceMaterial := lines[0][:strings.Index(lines[0], "-")]
	// line with "sourceMaterial-to-" removed
	tmpString := lines[0][strings.Index(lines[0], "-")+4:]
	destinationMaterial := (tmpString)[:strings.Index(tmpString, " ")]

	sourceRanges := make([]int, 0)
	destinationRanges := make([]int, 0)
	lengths := make([]int, 0)
	for rowIndex = 1; rowIndex < len(lines) && lines[rowIndex] != ""; rowIndex++ {
		valueStrings := strings.Split(lines[rowIndex], " ")

		destinationValue, _ := strconv.Atoi(valueStrings[0])

		destinationRanges = append(destinationRanges, destinationValue)

		sourceValue, _ := strconv.Atoi(valueStrings[1])

		sourceRanges = append(sourceRanges, sourceValue)

		length, _ := strconv.Atoi(valueStrings[2])

		lengths = append(lengths, length)
	}
	tmp := conversionMap{
		sourceStart:         sourceRanges,
		destinationStart:    destinationRanges,
		lengths:             lengths,
		sourceMaterial:      sourceMaterial,
		destinationMaterial: destinationMaterial}

	return tmp, rowIndex
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
