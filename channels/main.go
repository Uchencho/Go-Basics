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
		"http://uchencho.pythonanywhere.com",
		"https://wordcloud-generator-ub.herokuapp.com",
	}

	// Make a channel of type string
	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	// Loop thorough the length of the links
	// For i starting at 0, to len of links, incrementing by 1
	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c, c)
	// 	// fmt.Println(<-c)
	// }

	// Infinite loop
	// for i := 0; i < len(links); i++ {
	// 	go checkLink(<-c)
	// 	// fmt.Println(<-c)
	// }

	// Infinite loop
	for l := range c {
		// Initialize a go routine and a function literal (lambda func)
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
