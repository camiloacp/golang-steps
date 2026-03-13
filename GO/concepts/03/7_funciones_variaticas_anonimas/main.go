package main

import "fmt"

func main() {
	// fmt.Println(sum(2))
	// fmt.Println(sum(2, 3))
	// fmt.Println(sum(2, 3, 12))
	// fmt.Println(sum(2, 3, 12, 14))
	// fmt.Println(sum(2, 3, 12, 14, 25))

	// funcion anonima
	func(name string) {
		fmt.Println("Hello", name)
	}("Gophers")

}

// func sum(nums ...int) (total int) {
// 	//var total int

// 	for _, num := range nums {
// 		total += num
// 	}

// 	return
// }
