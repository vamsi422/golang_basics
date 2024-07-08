package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointers
var mut sync.Mutex

func main() {
	// go greeter("tigers ")
	// greeter("cats ")
	websitelist := []string{
		"https://amazon.in",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)

}

func greeter(s string) {
	for i := range 5 {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s, ":", i)
	}
}

func getStatusCode(endpoint string) {

	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint.")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf(" %d status code for %s \n", res.StatusCode, endpoint)

	}
}
