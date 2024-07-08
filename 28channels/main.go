package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang ")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(3)

	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)
	// Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 21
		myCh <- 22
		close(myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {

		val, isChanelOpen := <-myCh

		fmt.Println(isChanelOpen)
		fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	wg.Wait()

	// myCh <- 12
	// fmt.Println(<-myCh)
}
