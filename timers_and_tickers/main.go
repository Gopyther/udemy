package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go tickCounter(ticker, done)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(10 * time.Second)
	fmt.Println("Exting...")
}

func tickCounter(ticker *time.Ticker, done chan bool) {
	i := 0
Loop:
	for {
		select {
		case t := <-ticker.C:
			i++
			fmt.Println("Count ", i, " at ", t)
		case <-done:
			fmt.Println("done signal")
			break Loop
		}
	}
	fmt.Println("Exiting the tick counter")
}

// func main() {
// 	go tickCounter(1)
// 	time.Sleep(5 * time.Second)
// }

// func tickCounter(n int) {
// 	ticker := time.NewTicker(time.Duration(n) * time.Second)
// 	i := 0
// 	for t := range ticker.C {
// 		i++
// 		fmt.Println("Count ", i, " at ", t)
// 	}
// }

// func main() {

// 	nc := make(chan int)
// 	stopc := make(chan bool)

// 	go SlowCounter(1, nc, stopc)
// 	time.Sleep(5 * time.Second)

// 	nc <- 2
// 	time.Sleep(6 * time.Second)
// 	stopc <- true
// 	time.Sleep(1 * time.Second)
// }

// func SlowCounter(n int, nc chan int, stopc chan bool) {
// 	i := 0
// 	// create a duration of n seconds
// 	d := time.Duration(n) * time.Second

// 	for {
// 		select {
// 		// Use time.After channel to wait for a time period
// 		case <-time.After(d):
// 			i++
// 			fmt.Println(i)
// 		case n = <-nc:
// 			fmt.Println("Timer duration changed to", n)
// 			d = time.Duration(n) * time.Second
// 		case <-stopc:
// 			fmt.Println("Timer stopped")
// 			break
// 		}
// 	}
// 	fmt.Println("Existing Slow Counter")
// }
