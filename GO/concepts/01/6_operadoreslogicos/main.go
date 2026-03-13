package main

import "fmt"

func main() {
	// Operadores Comparación: ›, ‹, ==, !=, ›=, ‹=
	fmt.Println(4 > 6)

	// Operadores Lógicos &&, ||
	var age uint8 = 33
	fmt.Println("Es adulto?:", age >= 18 && age <= 60)
	fmt.Println("Es niño o anciano?:", age <= 18 || age >= 60)

	// Operador lógico Unario: !
	fmt.Println(!(2+2 == 4))
}
