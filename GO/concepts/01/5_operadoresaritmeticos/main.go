package main

import "fmt"

func main() {
	// Operadores aritmeticos (), *, /, %, +, -
	var a = (2 + 3) * 5
	fmt.Println(a)

	// Operadores  de asignacion: =, +=, -+, *=, ,%=
	var b int = 5
	b += 2 //b = b + 2 es lo mismo
	fmt.Println(b)

	// Declaracion post-incremento y post-decremento: ++, --
	// (no son una expresion sino una declaracion)
	var c int = 6
	c++
	fmt.Println(c)
}
