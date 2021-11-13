package main

import (
	"fmt"

	"github.com/kojh0111/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "Foremost in position, rank, or importance"}
	def, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	newWord := "hello"
	newMeaning := "Greeting"
	err2 := dictionary.Add(newWord, newMeaning)
	if err2 != nil {
		fmt.Println(err2)
	}
	meaning, _ := dictionary.Search(newWord)
	fmt.Println("found '", newWord, "' definition:", meaning)
	err3 := dictionary.Add(newWord, newMeaning)
	if err3 != nil {
		fmt.Println(err3)
	}
}
