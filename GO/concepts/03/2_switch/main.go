package main

import "fmt"

func main() {

	character := "🐭" // "🐒" "🐶" "🐭"

	// switch character {
	// case "🐒", "🐵":
	// 	fmt.Println("Es un mono")
	// case "🐶", "🐕":
	// 	fmt.Println("Es un perro")
	// default:
	// 	fmt.Println("Es un animal desconocido")
	// }

	// otra forma de hacer el switch
	canSearch := false

	switch {
	case !canSearch:
		fmt.Println("No se puede buscar")
	case character == "🐒" || character == "🐵":
		fmt.Println("Es un mono")
	case character == "🐶" || character == "🐕":
		fmt.Println("Es un perro")
	default:
		fmt.Println("Es un animal desconocido")
	}
}
