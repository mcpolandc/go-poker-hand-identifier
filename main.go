package main

import (
	"sort"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// card data structure
type card struct {
	rank string
	suit string
}

// used as a way to order by a numerical value
// as face cards would not get ordered correctly
// a - ace low, A - ace high
var rankOrder = map[string]int{
	"a": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 5,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

var suits = "HSCD"

// isValidSuit - Returns true if list of valid suits contains passed in value
func isValidSuit(suit string) bool {
	return strings.Contains(suits, string(suit))
}

// extractCards - takes a string representing a hand and returns array of type cards
func extractCards(hand string) (cards [] card, err error) {
	cardData := strings.Split(hand, " ")

	for _, crd := range cardData {

		rank := string(crd[0])
		suit := string(crd[1])

		if !isValidSuit(suit) {
			err = fmt.Errorf("Invalid suit '%s' received", suit)
			return
		}

		cards = append(
							cards,
							card {
								rank: rank,
								suit: suit,
							})
	}
	return
}

// getCounts - takes an array of cards and returns a map with a count of each card appearance
func getCounts(cards []card) (counts map[string]int) {

	counts = make(map[string]int)

	for _, crd := range cards {

		_, match := counts[crd.rank]

		if match {
			counts[crd.rank] += 1
		} else {
			counts[crd.rank] = 1
		}
	}

	return counts
}

// sortByRank - Sorts cards in order of their rank value
func sortByRank(cards []card) {
	sort.SliceStable(cards, func(i, j int) bool { return rankOrder[cards[i].rank] < rankOrder[cards[j].rank] })
}

// isFlush - Returns true if all cards in hand are of the same suit
func isFlush(cards []card) bool {
	suit := cards[0].suit

	for _, crd := range cards {
		if crd.suit != suit {
			return false
		}
	}
	return true
}

// isSequence - Returns true if cards are in sequence
func isSequence(cards []card) bool {

	// return false if adding 1 to current value does not equal next value
	for i := 0; i < len(cards) - 1; i++ {
		currentValPlusOne := rankOrder[cards[i].rank] + 1
		nextVal := rankOrder[cards[i + 1].rank]

		if currentValPlusOne != nextVal {
			return false
		}
	} 
	return true
}

// flipAce - returns array of cards with ace rank flipped
func flipAce(cards [] card) (aceLow []card) {
	copy(aceLow, cards)
	aceLow = append([]card{ card{ rank: "a", suit: cards[len(cards) - 1].suit }}, aceLow...)
	aceLow = aceLow[:len(aceLow) - 1]

	return
}

// getHandName - Returns the name of the hand found when given an array of cards
func getHandName(cards []card) string {

	// Royal flush (sort & check suit & check ace high)
	// Straight flush (sort & check suit)
	// Flush (check all same suit)
	// Straight (sort and check in sequence)

	sortByRank(cards)

	if isFlush(cards) {

		if isSequence(cards) {
			if cards[len(cards) - 1].rank == "A" {
				return "Royal flush"
			} else {
				return "Straight flush"
			}
		} else {
			// may still be straight flush
			// flip ace and check if in sequence again

			if cards[len(cards) - 1].rank == "A" {
				aceLow := flipAce(cards)

				if isSequence(aceLow) {
					return "Straight flush"
				}
			}
		}

		return "Flush"

	} else if isSequence(cards) {
		return "Straight"
	}

	// Four of a kind (rank count is 2, one rank is 4)
	// Full house (rank count is 2, one rank is 3)
	// Three of a kind (rank count 3, one rank is 3)
	// Two pair (rank count is 3, 2 ranks are 2)
	// One pair (rank count is 4, 1 rank is 2)

	counts := getCounts(cards)

	switch (len(counts)) {

	case 2:

		for _, count := range counts {
			if (count == 3) {
				return "Full house"
			}
			if (count == 4) {
				return "Four of a kind"
			}
		}

	case 3:
		for _, count := range counts {
			if (count == 3) {
				return "Three of a kind"
			}
		}
		return "Two pair"

	case 4:
		return "One pair"
	}

	// if no other hand found then we deem it as a high card hand
	return "High card"
}

func main() {

	// return early if no file supplied
	if len(os.Args) < 2 {
		fmt.Println("No file supplied")
		os.Exit(1)
	}

	file := os.Args[1]
	data, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("There was a problem reading the file")
		os.Exit(1)
	}

	hands := strings.Split(string(data), "\n")

	for _, hand := range hands {

		cards, err := extractCards(hand)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	
		fmt.Printf("%s - %s\n", hand, getHandName(cards))
	}
}
