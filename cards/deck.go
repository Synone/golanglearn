package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
//  which is a slice of strings
type Deck []string

// Create and return list of playing cards
func newDeck() Deck{
	cards := Deck{}
	cardSuits := []string{"Spades","Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","King","Queen"}

	for _, suit := range cardSuits{ // _ let go know that we dont need the index varialbe
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)

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

func (d Deck) saveToFile(filename string) error{
	return	os.WriteFile(filename, []byte(d.toString()), 0666)
	
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

func (d Deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano()) // create a source to make random from

	r := rand.New(source) //New returns a new Rand that uses random values from src to generate other random values.

	for i := range d {
		newPosition := r.Intn(len(d) -1)
		fmt.Println(newPosition)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}