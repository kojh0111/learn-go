package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	jh := person{name: "jh", age: 18, favFood: favFood}
	fmt.Println(jh.name)
}
