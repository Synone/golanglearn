package main

import (
	"fmt"
	"io"
	"os"
)

type fileStruct struct{}

func main() {
	var fileName = os.Args[1]

	file, err := os.Open(fileName)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	file01 := fileStruct{}
	io.Copy(file01, file)

}


func (fileStruct) Write(bs []byte) (n int, err error){
	fmt.Println(string(bs))

	return len(bs), nil
}