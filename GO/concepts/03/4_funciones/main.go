package main

import (
	"fmt"
	"strings"
)

func main() {
	//great("Camilo", "Cortes")
	name := "Camilo"
	toUpper(&name)
	fmt.Println(name)
}

// func great(firstName, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

func toUpper(text *string) {
	*text = strings.ToUpper(*text)
}
