package main

import (
	"fmt"
	"os"
	"strings"
)

type mapEntry struct {
	directions map[byte]string
	location   string
}

var mapEntries map[string]mapEntry
var steps string

func main() {
	lines := ReadFile("test.txt")

	steps = lines[0]

	mapEntries = getMapEntries(lines[2:])

	// Puzzle 1
	fmt.Println("Puzzle 1: ", countStepsToReachDestination(mapEntries["AAA"], "ZZZ"))

	// Puzzle2
	startingLocations := make(map[string]mapEntry)

	for _, entry := range mapEntries {
		if entry.location[2] == 'A' {
			startingLocations[entry.location] = entry
		}
	}

	stepCounts := make([]int, 0)

	for _, v := range startingLocations {
		currentStepCount := countStepsToReachDestination(v, "Z")
		stepCounts = append(stepCounts, currentStepCount)
	}

	fmt.Println("Puzzle 2: ", leastCommonMultiple(stepCounts...))
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(integers ...int) int {
	if len(integers) < 2 {
		panic("not enough numbers for LCM")
	}
	result := integers[0] * integers[1] / greatestCommonDivisor(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}

func countStepsToReachDestination(startLocation mapEntry, destSuffix string) int {
	currentStep := 1
	totalSteps := 1
	destination := mapEntries[startLocation.location].directions[steps[0]]
	// for destination != "ZZZ" {
	for !strings.HasSuffix(destination, destSuffix) {
		destination = mapEntries[destination].directions[steps[currentStep]]
		currentStep++
		totalSteps++
		if currentStep >= len(steps) {
			currentStep = 0
		}
	}
	return totalSteps
}

func getMapEntries(lines []string) map[string]mapEntry {
	result := make(map[string]mapEntry)

	for _, line := range lines {
		newMapEntry := mapEntry{}
		newMapEntry.location = string(line[0:3])

		newMap := make(map[byte]string, 2)

		newMapEntry.directions = newMap
		newMapEntry.directions['L'] = string(line[7:10])
		newMapEntry.directions['R'] = string(line[12:15])
		result[newMapEntry.location] = newMapEntry
	}

	return result
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
