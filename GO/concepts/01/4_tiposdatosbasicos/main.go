package main

import "fmt"

func main() {
	// bool, string, numeric
	//var a bool = true
	//var a string = "true"
	//var a uint8 = 33
	//var a rune = 'a' //codigo unicode de la letra a
	//var a float32 = 132.3
	var a uint8 = 255
	var b uint16 = 2550

	//c := uint16(a) + b // se puede hacer un cast para realziar operaciones entre variables de diferentes tipos

	_ = uint16(a) + b // el _ me permite ejecutar el programa sin que arroje error por no usar la variable

	//fmt.Printf("Tipo: %T, Valor: %v\n", c, c)
	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)
}

// se puede hacer un cast para realziar operaciones entre variables de diferentes tipos

/*
El valor cero de un string es una cadena vacia ""
El valor cero de un booleano es false
El valor cero de un numero es 0
*/
