package main

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

	cards := newDeck()
	// cards.print()
	// cards[:3].print()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
}
