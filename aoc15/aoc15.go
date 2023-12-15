package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
}

type Box struct {
	LensLabels map[string]int
	Lenses     []Lens
}

func RemoveIndex(s []Lens, index int) []Lens {
	ret := make([]Lens, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func createEmptyBox() Box {
	return Box{Lenses: make([]Lens, 0), LensLabels: make(map[string]int)}
}
func main() {

	line := ReadFile("test.txt")

	codes := strings.Split(line, ",")
	sum := 0
	for _, code := range codes {
		sum += getHash(code)
	}

	boxes := make([]Box, 256)

	for i := 0; i <= 255; i++ {
		boxes[i] = createEmptyBox()
	}

	for _, code := range codes {
		strings := strings.Split(code, "=")

		if len(strings) == 2 {
			// must contain =
			label := strings[0][:len(strings[0])]
			boxNumber := getHash(label)

			focalLength, _ := strconv.Atoi(strings[1])
			boxIndex, ok := boxes[boxNumber].LensLabels[label]
			if ok {
				if boxIndex < 0 {
					// add
					boxes[boxNumber].Lenses = append(boxes[boxNumber].Lenses, Lens{Label: label, FocalLength: focalLength})
					boxes[boxNumber].LensLabels[label] = len(boxes[boxNumber].Lenses) - 1
				} else {
					//replace
					boxes[boxNumber].Lenses[boxes[boxNumber].LensLabels[label]] = Lens{Label: label, FocalLength: focalLength}
				}

			} else {
				boxes[boxNumber].Lenses = append(boxes[boxNumber].Lenses, Lens{Label: label, FocalLength: focalLength})
				boxes[boxNumber].LensLabels[label] = len(boxes[boxNumber].Lenses) - 1
			}
		} else {
			//len(strings) == 1
			// contains -, remove from box
			label := strings[0][:len(strings[0])-1]
			boxNumber := getHash(label)
			boxIndex, ok := boxes[boxNumber].LensLabels[label]
			if ok && boxIndex >= 0 {
				boxes[boxNumber].Lenses = RemoveIndex(boxes[boxNumber].Lenses, boxIndex)
				boxes[boxNumber].LensLabels[label] = -1
				for i := boxIndex; i < len(boxes[boxNumber].Lenses); i++ {
					boxes[boxNumber].LensLabels[boxes[boxNumber].Lenses[i].Label]--
				}
			}
		}
	}

	sum2 := 0
	for idx := range boxes {
		for i := range boxes[idx].Lenses {
			tmp := (idx + 1) * (i + 1) * boxes[idx].Lenses[i].FocalLength
			sum2 += tmp
		}
	}

	fmt.Println("Puzzle 1: ", sum)
	fmt.Println("Puzzle 2: ", sum2)
}

func getHash(code string) int {
	currentValue := 0
	for _, char := range code {
		currentValue += int(char)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func ReadFile(fName string) string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := string(f)
	return content

}
