package main

import (
	"fmt"

	"github.com/kojh0111/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "Foremost in position, rank, or importance"}
	def, err := dictionary.Search("firts")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	def2, err2 := dictionary.Search("first")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(def2)
	}
}
