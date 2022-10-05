package main

import "fmt"

func main() {
	numbers := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numbers{
		if(num % 2 == 0){
			fmt.Printf("%v is even",num)
		} else {
			fmt.Printf("%v is odd",num)
		}
	}
	
}