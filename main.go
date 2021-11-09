package main

import "fmt"

func main() {
	const name string = "jh"
	fmt.Println(name)
	myname := "jh" // var myname string = "jh" 같은 의미!
	myname = "koh jh"
	fmt.Println(myname)
}
