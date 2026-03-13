package main

import "fmt"

func main() {

	division(200, 25)
	division(34, 25)
	division(10, 0)
	division(1000, 52)

}

func division(dividendo int, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Me recupero de un panic")
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por 0")
	}
}
