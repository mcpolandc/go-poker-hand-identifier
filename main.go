package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func processHand(cards []string) (hand string, err error) {
	fmt.Printf("cards: %v\n", cards)
	return "", nil
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
		_, _ = processHand(strings.Split(hand, " "))
	}

}
