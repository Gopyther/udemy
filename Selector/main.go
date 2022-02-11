package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrCrewNotFound = errors.New("Crew member not found")

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
	//Simulate searching
	time.Sleep(time.Duration(rand.Intn(50)) * time.Microsecond)

	if v, ok := scMapping[name]; !ok {
		panic("Crew member not found")
	} else {
		return v, nil
	}
}

// func findSC(name, server string) (int, error) {
// 	//Simulate searching
// 	time.Sleep(time.Duration(rand.Intn(50)) * time.Microsecond)

// 	if v, ok := scMapping[name]; !ok {
// 		return -1, errors.New("Crew member not found")
// 	} else {
// 		return v, nil
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("A panic recovered", err)
// 		}
// 	}()
// 	if clearance, err := findSC("Ruko", "Server 1"); err != nil {
// 		fmt.Println("Error occured while searching ", err)
// 		if v, ok := err.(findError); ok {
// 			fmt.Println("Server name", v.Server)
// 			fmt.Println("Crew member name", v.Name)
// 		}
// 	} else {
// 		fmt.Println("Crew member has security clearance:", clearance)
// 	}
// }

func main() {
	rand.Seed(time.Now().UnixNano())
	clearance, err := findSC("Ruko", "Server 1")
	fmt.Println(" Clearance level found:", clearance, "Error code", err)
}

// func main() {

// 	rand.Seed(time.Now().UnixNano())

// 	c1 := make(chan int)
// 	c2 := make(chan int)

// 	name := "James"

// 	go findSC(name, "Server 1", c1)
// 	go findSC(name, "Server 2", c2)

// 	/*
// 		Do some work...
// 	*/

// 	select {
// 	case sc := <-c1:
// 		fmt.Println(name, "has a security clearance of ", sc, "found in server1 ")
// 	case sc := <-c2:
// 		fmt.Println(name, "has a security clearance of ", sc, "found in server2 ")
// 	default:
// 		fmt.Println("Too slow!!")
// 	}

// }
