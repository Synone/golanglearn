package main

import (
	"fmt"
	"os"
	"strings"
)

type Deck []string

// Create and return list of playing cards
func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

	for _, suit := range cardSuits { // _ let go know that we dont need the index varialbe
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) Deck {
	bs, err := os.ReadFile(filename)
	if err != nil{
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",")
	return Deck(s)
}

func (d Deck) toString() string{
	return strings.Join([]string(d)," ,")
	
}