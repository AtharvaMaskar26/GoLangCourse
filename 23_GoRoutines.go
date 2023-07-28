package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // These are pointers

func GoRoutines() {
	fmt.Printf("Go Routines in Go!\n")

	// Here first a couple of times hello is being printed then a couple of times world

	// To make it a go routine is adding go keyword
	// Inthe below line we will only get printed world because we are not waiting for the hello to be printed
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	// While running this program you might notics that the response is very slow, what we can do is create go routines and fire them parallely
	for _, web := range websiteList {
		// Creates multiple threads and fires it
		go getStatusCode(web)

		// Keeps a track of how many go routines are out there
		wg.Add(1)
	}

	// Does not end to code as long as the go routine is returned
	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	// returns a signal done after returning
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Problem in endpoint")
	}
	fmt.Printf("%d Status code for website %s\n", res.StatusCode, endpoint)

}
