package main

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func main() {
	Camilo := Person{
		Name:        "Camilo",
		Age:         29,
		HasChildren: false,
	}

	Andrea := Person{"Andrea", 31, true}

	fmt.Printf("%+v\n", Camilo)
	fmt.Printf("%+v\n", Andrea)
}
