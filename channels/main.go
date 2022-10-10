package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	 
	fmt.Println("this is under the for loop, let's see if it is blocked")
	
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for{ // create an infinite loop
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		go func (link string)  {
			time.Sleep(5*time.Second)
			checkLink(link,c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	time.Sleep(5*time.Second)
	_, err := http.Get(link)
	if err != nil{
		fmt.Println(link + "might be down!!")
		// c <- link + "Might be down I think"
		c <- link
		return 
	}

	fmt.Println(link, "is up")
	// c <- "Yup it's up"
	c <- link
}