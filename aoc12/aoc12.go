package main

import (
	"os"
	"strconv"
	"strings"
)

// type Point struct {
// 	x int
// 	y int
// }

func main() {
	lines := ReadFile("test.txt")

	sum := 0
	for _, x := range lines {
		sum += getTotalConfigurations(x)
	}
}

func getTotalConfigurations(line string) int {
	springs := strings.Fields(line)[0]
	numbersString := strings.Fields(line)[1]
	numbersStrings := strings.Split(numbersString, ",")
	numbers := make([]int, 0)

	for _, x := range numbersStrings {
		i, _ := strconv.Atoi(x)
		numbers = append(numbers, i)
	}

	count := getConfigurations(springs, numbers)

	return count
}

func getConfigurations(line string, nums []int) int {
	if len(line) == 0 || len(nums) == 0 {
		return 0
	}
	retVal := 0
	potentialIdx := strings.Index(line, "?")
	springLen := 0
	numIndex := 0
	for i := 0; i < potentialIdx; i++ {

		if line[i] == '#' {
			springLen++
			if springLen == nums[numIndex] {
				numIndex++
			}
		}

		if line[i] == '?' {
			//lookahead and count #
			retVal = 1 // decoy so path is not empty
		}

	}

	return retVal
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
