package main

import "fmt"

func main(){

	colors := map[string]string{
		"red":"#ff0000",
		"green":"#4bf321",
		"white":"#ffffff",
	}
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// fmt.Println(colors)
	 trees  := make(map[string]string)
	trees["vn"] = "asa"
	trees["japan"] = "DUong xi"
	delete(trees,"vn")
	fmt.Println(trees)
	printMap(colors)
}

func printMap(c map[string]string){
	for color,hex := range c{
		fmt.Printf("The color %v has the hex: %v",color,hex)
	}
}