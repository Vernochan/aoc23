package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)
var cacheHits = 0

func main() {
	lines := ReadFile("test.txt")

	sum := 0
	sum2 := 0
	for _, x := range lines {
		sum += getTotalConfigurations(x)
		sum2 += getTotalConfigurationsExpanded(x)
	}

	fmt.Println("Puzzle 1: ", sum)
	fmt.Println("Puzzle 1: ", sum2, "Cache hits: ", cacheHits)
}

func getTotalConfigurationsExpanded(line string) int {
	springs := strings.Fields(line)[0]

	numbersStrings := strings.Split(strings.Fields(line)[1], ",")
	numbers := make([]int, 0)

	for _, x := range numbersStrings {
		i, _ := strconv.Atoi(x)
		numbers = append(numbers, i)
	}

	expandedNumbers := make([]int, 0)
	for i := 0; i < 5; i++ {
		expandedNumbers = append(expandedNumbers, numbers...)

	}

	expandedStrings := fmt.Sprintf("%[1]s?%[1]s?%[1]s?%[1]s?%[1]s", springs)

	count := getConfigurations(expandedStrings, expandedNumbers)

	return count
}
func getTotalConfigurations(line string) int {
	springs := strings.Fields(line)[0]
	numbersStrings := strings.Split(strings.Fields(line)[1], ",")
	numbers := make([]int, 0)

	for _, x := range numbersStrings {
		i, _ := strconv.Atoi(x)
		numbers = append(numbers, i)
	}

	count := getConfigurations(springs, numbers)

	return count
}

func getConfigurations(springLine string, springLengths []int) int {
	cacheKey := springLine
	for _, num := range springLengths {
		cacheKey += strconv.Itoa(num) + ","
	}

	if cacheValue, ok := cache[cacheKey]; ok {
		cacheHits++
		return cacheValue
	}
	if len(springLine) == 0 {
		if len(springLengths) == 0 {
			return 1
		}
		return 0
	}

	// Fallunterscheidung:
	// 1 => zusammenhängende Spring (? oder #)     => retVal += 1 + getconfugurations(string[nums[0]:], nums[1:])
	// 2 => hängt nicht zusammen/gehört nicht dazu => retVal += getconfigurations(string[1:], nums)

	if strings.HasPrefix(springLine, ".") {
		val := getConfigurations(strings.TrimPrefix(springLine, "."), springLengths)
		cache[cacheKey] = val
		return val
	}

	if strings.HasPrefix(springLine, "?") {
		return getConfigurations(fmt.Sprintf(".%s", springLine[1:]), springLengths) +
			getConfigurations(fmt.Sprintf("#%s", springLine[1:]), springLengths)
	}

	// line must start with # (no other option left)
	if len(springLengths) == 0 {
		cache[cacheKey] = 0
		return 0
	}
	if len(springLine) < springLengths[0] {
		cache[cacheKey] = 0
		return 0
	}
	if strings.Contains(springLine[0:springLengths[0]], ".") {
		cache[cacheKey] = 0
		return 0
	}
	if len(springLengths) == 1 {
		val := getConfigurations(springLine[springLengths[0]:], springLengths[1:])
		cache[cacheKey] = val
		return val
	}

	if len(springLine) < springLengths[0]+1 || string(springLine[springLengths[0]]) == "#" {
		cache[cacheKey] = 0
		return 0
	}
	val := getConfigurations(springLine[springLengths[0]+1:], springLengths[1:])
	cache[cacheKey] = val
	return val

}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
