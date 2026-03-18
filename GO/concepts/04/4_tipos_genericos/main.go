package main

import "fmt"

type Product[T uint | string] struct {
	ID          T
	Description string
	Pice        float64
}

func main() {
	product1 := Product[uint]{ID: 1, Description: "Zapatos", Pice: 100.0}
	product2 := Product[string]{ID: "12hjsd-28sjjd-jasbdmv", Description: "Zapatos", Pice: 100.0}
	fmt.Println(product1)
	fmt.Println(product2)
}
