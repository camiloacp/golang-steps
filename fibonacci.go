package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Calculando Fibonacci(45) en Go...")

	// Medimos el tiempo de inicio
	startTime := time.Now()

	// Ejecutamos la función
	resultado := fibonacci(45)

	// Calculamos el tiempo transcurrido
	elapsedTime := time.Since(startTime)

	fmt.Printf("Resultado: %d\n", resultado)
	fmt.Printf("Tiempo de ejecución: %s\n", elapsedTime)
}
