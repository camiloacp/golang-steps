package main

import "fmt"

func main() {
	// Puntero: Variable que almacena la dirección en memoria de un valor
	var color string = "🟥"
	var pointerColor *string
	pointerColor = &color
	*pointerColor = "🟦"

	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s\n", pointerColor, pointerColor, *pointerColor)
}
