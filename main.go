package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10
	fmt.Println(&a) // memory address
	fmt.Println(*b) // see in memory
	*b = 20
	fmt.Println(a)
}
