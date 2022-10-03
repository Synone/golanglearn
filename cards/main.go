package main

import "fmt"



func main(){
	
	card:=newCard()
	fmt.Println(card)
	// var numberOfCards = 12
	// printCity()
	

	cards := []string{"Ace of Diamonds",newCard(), newCard()} //  Slice
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)

	for i, card := range cards{
		fmt.Println( card)
		fmt.Println(i)
	}

}
func newCard() string{
	return "Five of Diamonds"
}