package main

import "fmt"

// Create a new type of "deck"
//  which is a slice of strings
type Deck []string

func newDeck() Deck{
	cards := Deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues :=[]string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits{ // _ let go know that we dont need the index varialbe
		for _, value := range cardValues{
			cards = append(cards, suit + "of " + value)
		}
	}
	return cards
}

func (d Deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}



type Size int;

func (s Size) returnSize() Size{
	return s
}