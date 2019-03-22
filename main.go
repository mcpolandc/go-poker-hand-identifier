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

// extractCards - takes a string representing a hand returns array of type cards
func extractCards(hand string) (cards [] card) {
	cardData := strings.Split(hand, " ")

	for _, crd := range cardData {
		cards = append(
							cards,
							card {
								rank: string(crd[0]),
								suit:  string(crd[1]),
							})
	}

	return cards
}

// getCounts - takes an array of cards and returns a map with a count of each card appearance
func getCounts(cards []card) (counts map[string]int) {

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

func sortCards(cards []card) {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].rank < cards[j].rank })
}

func processHand(cards []card) string {


	sortCards(cards)

	// counts := getCounts(cards)

	// EXECUTION GROUPS
	//
	// SORTING AND CHECKING SUITS
	// Royal flush (sort & check suit & check last card is ace)
	// Straight flush (sort & check suit)
	// Flush (check all same suit)
	// Straight (sort and check in sequence)

	// COUNTING RANKS CHECKS
	//
	// Four of a kind (rank count is 2, one rank is 4)
	// Full house (rank count is 2, one rank is 3)
	// Three of a kind (rank count 3, one rank is 3)
	// Two pair (rank count is 3, 2 ranks are 2)
	// One pair (rank count is 4, 1 rank is 2)

	return ""
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
	}

	hands := strings.Split(string(data), "\n")

	for _, hand := range hands {

		cards := extractCards(hand)
		_ = processHand(cards)
	}
}
