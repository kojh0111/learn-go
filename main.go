package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 3))
	totalLength, up := lenAndUpper(("koh jh"))
	fmt.Println(totalLength, up)
	repeatMe("a", "b", "c", "d")
}
