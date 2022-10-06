package main

import "fmt"

type bot interface{
	getGreeting() string
	getSound(int) (string,int)
	
}
type user struct{
	name string
}

type nit interface{
	respondToUser(user) string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb:= englishBot{}
	// sb:= spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)

	

}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
	fmt.Println(b.getSound(12))
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func (englishBot) getSound(number int) (string,int){
	return "EEE",number
}

