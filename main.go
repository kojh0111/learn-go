package main

import (
	"fmt"
	"time"
)

func main() {
	go helloCount("jh")
	go helloCount("koh")
	time.Sleep(time.Second * 5)
}

func helloCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", person, i)
		time.Sleep(time.Second)
	}
}
