package main

import "fmt"

func main() {
	jh := map[string]string{"name": "jh", "age": "12"}
	for _, value := range jh {
		fmt.Println(value)
	}
}
