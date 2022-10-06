package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}
type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName)


	var alex person
	alex.firstName = "Sony"
	alex.lastName = "Nguyen"
	// fmt.Printf("%+v",alex)

	jim := person{
		firstName :"Jim",
		lastName: "Harry",
		contactInfo : contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		}, 

	}

	// jimPointer := &jim;
	// fmt.Println(" jimpointer is: ", jimPointer)
	jim.updateName("Sony","Nguyen")
	jim.printF()

	var x int = 5748
	var p *int
	p = &x
	fmt.Println("Value stored in x = ", x)
    fmt.Println("Address of x = ", &x)
    fmt.Println("Value stored in variable p = ", *p)


	 greeting := []string{"Hello","everyone"}
	updateString(greeting)
	fmt.Println(greeting)
	name := "bill"
 
 namePointer := &name
 
 fmt.Println(namePointer)
 printPointer(namePointer)
}

func (p person) printF(){
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string, newLastName string){
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}
func updateString(text []string){
	text[0] = "Bye"
}
func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
   }