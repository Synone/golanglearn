package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// var numberOfCards = 12
	// printCity()
	// hello := Hello()
	// fmt.Println(hello)
	// cards := Deck{"Ace of Diamonds",newCard(), newCard()} //  Slice
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)
	// sizeOfDesk:= Size(52);
	// fmt.Println("Number of cards is: ",sizeOfDesk.returnSize())
	// cards.print()

	// cards := newDeck()
	// cards.print()
	// cards[:3].print()
	// fmt.Println(cards)
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// cards.saveToFile("my_cards.txt")
	cards := newDeckFromFile("my_cards.txt")
	cards.print()

	// byteArray := []byte{'G', 'O', 'L', 'A', 'N', 'G'}
	// anotherByteArr := []byte("Hellooooooo")
	// fmt.Println(byteArray)
	// fmt.Println(anotherByteArr)

	newCards := newDeck()
	newCards.shuffle()
	newCards.print()
	fmt.Println(time.Now().UnixNano())

	 rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(12))
}
