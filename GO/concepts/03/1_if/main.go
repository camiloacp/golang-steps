package main

import "fmt"

func main() {
	// "🐒" "🐶" "🐭"

	if character := "🐭"; character == "🐒" { // if character := "🐭" variable temporal que solo existe dentro del if
		fmt.Println("Es un mono")
	} else if character == "🐶" {
		fmt.Println("Es un perro")
	} else {
		fmt.Println("Es un animal desconocido")
	}
}
