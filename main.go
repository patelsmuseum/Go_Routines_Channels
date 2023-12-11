package main

import (
	"fmt"
	"net/http"
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

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c) ***************one way we can also use for loop
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)                     ********** 2nd way for doing it so
	// }

	// for {
	// 	go checkLink(<-c, c)                   ********** 3rd way to do so
	// }

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("there se some problem in website", link)
		c <- "Might me down for sometime"
		return
	}

	fmt.Println(link, " is up")
	c <- "Site is up"
}
