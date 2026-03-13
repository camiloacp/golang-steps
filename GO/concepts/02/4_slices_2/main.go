package main

import "fmt"

func main() {
	// len(): # de elementos en el slice
	// cap(): # de elementos del array origen, a partir del indice donde se creo el slice

	/*
		animals := [5]string{"🐢", "🐶", "🐈", "🦁", "🦉"}
		pets := animals[1:3]
		pets = append(pets, "🦜", "🐠", "🐝")

		// Array[4]{"🐶", "🐈", "🦁", "🦉"}
		// New Array[8]{"🐶", "🐈", "🦜", "🐠", "🐝"}

		fmt.Println(animals)
		fmt.Println(pets)
		fmt.Println("Tamaño de mascotas:", len(pets))
		fmt.Println("Capacidad:", cap(pets))
	*/

	//pets := []string{"🐶", "🐈"}

	pets := make([]string, 0, 3)
	pets = append(pets, "🦜", "🐠", "🐝")

	fmt.Println(pets)
	fmt.Println("Tamaño de mascotas:", len(pets))
	fmt.Println("Capacidad:", cap(pets))

}
