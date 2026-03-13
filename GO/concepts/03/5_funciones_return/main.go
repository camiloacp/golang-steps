package main

import (
	"fmt"
	"strings"
)

func main() {
	// suma := sum(3, 5)
	// fmt.Println("El resultado de la suma es:", suma)

	lower, upper := convert("MeLi")
	fmt.Println(lower, upper)
}

// func sum(a, b int) int {
// 	return a + b
// }

// forma 1
// func convert(text string) (string, string) {
// 	lower := strings.ToLower(text)
// 	upper := strings.ToUpper(text)

// 	return lower, upper
// }

func convert(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return
}
