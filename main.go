package main

import (
	"fmt"
)

func canIDrink(age int) bool {
	koreanAge := age + 2
	switch {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 80:
		return false
	}
	return true
}

func main() {
	nums := [6]int{14, 16, 18, 76, 78, 80}
	for _, num := range nums {
		fmt.Println(num, canIDrink(num))
	}
}
