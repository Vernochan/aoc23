package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type raceInformation struct {
	time     int
	distance int
}

func main() {
	lines := ReadFile("test.txt")

	timesString := (lines[0])[5:]
	distancesString := (lines[1][9:])

	timesString = standardizeSpaces(timesString)
	distancesString = standardizeSpaces(distancesString)

	times := strings.Split(timesString, " ")
	distances := strings.Split(distancesString, " ")

	singleTimeString := strings.Replace(timesString, " ", "", -1)
	singleDistanceString := strings.Replace(distancesString, " ", "", -1)

	singleTime, _ := strconv.Atoi(singleTimeString)
	singleDistance, _ := strconv.Atoi(singleDistanceString)

	singleRace := raceInformation{distance: singleDistance, time: singleTime}

	races := make([]raceInformation, 0)

	for idx := range times {
		distance, _ := strconv.Atoi(distances[idx])
		time, _ := strconv.Atoi(times[idx])

		tmp := raceInformation{distance: distance, time: time}
		races = append(races, tmp)
	}

	total := 1
	for _, race := range races {
		_, _, totalTimes := getMaxWaysToWin(race)
		total *= totalTimes
	}

	fmt.Println("Total: ", total)

	_, _, singleTotalTimes := getMaxWaysToWin(singleRace)
	fmt.Println("TotalSingle: ", singleTotalTimes)

}

func getMaxWaysToWin(r raceInformation) (int, int, int) {
	minTimeToBeat := 99999999
	maxTimeToBeat := 0
	waysToWin := 0

	for timeToCharge := 0; timeToCharge < r.time; timeToCharge++ {
		timeToRace := r.time - timeToCharge
		dist := timeToCharge * timeToRace
		//fmt.Println("charge: ", timeToCharge, " race: ", timeToRace, " distance: ", dist)
		if dist > r.distance {
			waysToWin++
			if timeToCharge < minTimeToBeat {
				minTimeToBeat = timeToCharge
			}

			if timeToCharge > maxTimeToBeat {
				maxTimeToBeat = timeToCharge
			}
		}

	}
	return minTimeToBeat, maxTimeToBeat, waysToWin
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
