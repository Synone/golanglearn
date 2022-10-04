package main

import (
	"fmt"
	"strings"
)

// Create a new type of "deck"
//  which is a slice of strings
type Deck []string

// Create and return list of playing cards
func newDeck() Deck{
	cards := Deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues :=[]string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","King","Queen"}

	for _, suit := range cardSuits{ // _ let go know that we dont need the index varialbe
		for _, value := range cardValues{
			cards = append(cards, suit + " of " + value)
		}
	}
	return cards
}

func (d Deck) print(){           // Only instance of Deck could call print() function
	for i, card := range d{
		fmt.Println(i,card)
	}
}


// Test zone
/*
Test zone
*/
type Size int;

func (s Size) returnSize() Size{
	return s
}

// end of test zone


 
func deal(d Deck, handSize int) (Deck, Deck) {   // return 2 value type Deck
	return d[:handSize], d[handSize:]
}	

func (d Deck) toString() string{
	return strings.Join([]string(d)," ,")
	
}