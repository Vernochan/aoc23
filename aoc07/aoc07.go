package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type card struct {
	face byte
}
type hand struct {
	cards []byte
	bid   int
	rank  int
}

var facesPuzzle1 map[byte]int = map[byte]int{
	'1': 0x1,
	'2': 0x2,
	'3': 0x3,
	'4': 0x4,
	'5': 0x5,
	'6': 0x6,
	'7': 0x7,
	'8': 0x8,
	'9': 0x9,
	'T': 0xA,
	'J': 0xB,
	'Q': 0xC,
	'K': 0xD,
	'A': 0xE,
}

var facesPuzzle2 map[byte]int = map[byte]int{
	'J': 0x0,
	'1': 0x1,
	'2': 0x2,
	'3': 0x3,
	'4': 0x4,
	'5': 0x5,
	'6': 0x6,
	'7': 0x7,
	'8': 0x8,
	'9': 0x9,
	'T': 0xA,
	'Q': 0xC,
	'K': 0xD,
	'A': 0xE,
}

func main() {
	lines := ReadFile("test.txt")

	hands := make([]hand, 0)

	for _, line := range lines {
		hand := readHand(line)
		hands = append(hands, hand)

	}

	sort.Slice(hands, func(i, j int) bool {
		return getHandValue(hands[i]) < getHandValue(hands[j])
	})

	sum := 0
	for k, v := range hands {
		v.rank = k + 1
		sum += v.bid * v.rank
		// fmt.Println(v, ": ", fmt.Sprintf("%x", getHandValue(v)))
	}
	fmt.Println("Puzzle 1", sum)

	sort.Slice(hands, func(i, j int) bool {
		return handValueWithJokers(hands[i]) < handValueWithJokers(hands[j])
	})
	sum = 0
	for k, v := range hands {
		v.rank = k + 1
		sum += v.bid * v.rank
		// fmt.Println(v, ": ", fmt.Sprintf("%x", handValueWithJokers(v)))
	}
	fmt.Println("Puzzle 2", sum) //249817836
}

// high card: 0, one pair: 0x100000, two pair: 0x200000, three of a kind: 0x300000, fullHouse: 0x400000, Four of a kind: 0x500000, Five of a kind: 0x600000
func getHandValue(h hand) int {

	handValue := 0
	handValue = getHandBaseValue(h)

	handValue += facesPuzzle1[h.cards[0]] * 0x10000
	handValue += facesPuzzle1[h.cards[1]] * 0x1000
	handValue += facesPuzzle1[h.cards[2]] * 0x100
	handValue += facesPuzzle1[h.cards[3]] * 0x10
	handValue += facesPuzzle1[h.cards[4]] * 0x1

	return handValue
}

func getHandBaseValue(h hand) int {
	handStr := string(h.cards)
	var isThree, isOnePair bool
	handValue := 0
	for k := range facesPuzzle1 {
		occurences := strings.Count(handStr, string(k))

		if occurences == 5 {
			handValue = 0x600000
		}

		if occurences == 4 {
			handValue = 0x500000
		}

		if occurences == 3 {
			if isOnePair {
				// Full house
				handValue = 0x400000
			} else {
				handValue = 0x300000
			}
			isThree = true
		}

		if occurences == 2 {
			if isThree {
				//isFullHouse = true
				// fullHouse
				handValue = 0x400000
			} else {
				if isOnePair {
					handValue = 0x200000
					//isTwoPairs = true

				} else {
					handValue = 0x100000
				}
			}

			isOnePair = true
		}

	}

	return handValue
}

// high card: 0, one pair: 0x100000, two pair: 0x200000, three of a kind: 0x300000, fullHouse: 0x400000, Four of a kind: 0x500000, Five of a kind: 0x600000
func handValueWithJokers(h hand) int {
	handStr := string(h.cards)
	handValue := getHandBaseValue(h)
	// occurences := strings.Count(string(h.cards), "J")
	// fmt.Print(occurences)
	for k := range facesPuzzle2 {
		if k == 'J' {
			continue
		}
		newCardString := strings.ReplaceAll(handStr, "J", string(k))
		newHand := readHand(fmt.Sprintf("%s %d", newCardString, h.bid))

		newHandValue := getHandBaseValue(newHand)
		if newHandValue > handValue {
			handValue = newHandValue
		}

	}

	handValue += facesPuzzle2[h.cards[0]] * 0x10000
	handValue += facesPuzzle2[h.cards[1]] * 0x1000
	handValue += facesPuzzle2[h.cards[2]] * 0x100
	handValue += facesPuzzle2[h.cards[3]] * 0x10
	handValue += facesPuzzle2[h.cards[4]] * 0x1

	return handValue
}

func readHand(line string) hand {
	cards := make([]byte, 0)

	for i := 0; i < 5; i++ {
		cards = append(cards, line[i])
	}

	value, _ := strconv.Atoi(line[6:])

	tmp := hand{rank: 0, bid: value, cards: cards}
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
