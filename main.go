package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [2]string{"jh", "koh"}
	for _, person := range people {
		go helloCount(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func helloCount(person string, c chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 5)
		fmt.Println("Hello ", person)
		c <- true
	}
}
