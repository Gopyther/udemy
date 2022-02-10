package main

import (
	"fmt"
)

// func main() {
// 	go waitAndSay("World")
// 	fmt.Println("Hello")
// 	time.Sleep(3 * time.Second)
// }

// func waitAndSay(s string) {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println(s)
// }
// func main() {
// 	go func(s string) {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println(s)
// 	}("World")
// 	fmt.Println("Hello")
// 	time.Sleep(3 * time.Second)
// }

// func main() {
// 	c := make(chan bool)
// 	go waitAndSay(c, "world")
// 	fmt.Println("Hello")

// 	//we send a signal to c in order to allow waitAndSay to continue
// 	c <- true

// 	//we wait to receive another signal on c before we exit
// 	<-c
// }

// func waitAndSay(c chan bool, s string) {
// 	if b := <-c; b {
// 		fmt.Println(s)
// 	}
// 	c <- true
// }

// func main() {
// 	ch := make(chan string, 2)

// 	ch <- "Hello"
// 	ch <- "World"
// 	ch <- "Hello again"

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

func main() {
	c := make(chan string)
	go SayHelloMultipleTimes(c, 5)
	for s := range c {
		fmt.Println(s)
	}

	v, ok := <-c
	fmt.Println("Channel close?", !ok, " value ", v)
}

func SayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}
