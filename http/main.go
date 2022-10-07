package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


type logWrite struct{}



func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil{
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWrite{}
 
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	
	// fmt.Println(string(bs))


	io.Copy(lw, resp.Body)
} 

// custome Write function to match Write interface
func (logWrite) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}