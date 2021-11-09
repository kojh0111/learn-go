package main

import "fmt"

func main() {
	arr := [4]string{"1", "2", "3"}
	fmt.Println(arr)
	arr[3] = "5"
	fmt.Println(arr)
	names := []string{"jh", "sy", "js"}
	names = append(names, "jy")
	fmt.Println(names)
}
