package main

import "fmt"

func main() {
	// primera forma de hacer un for
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// segunda forma de hacer un for
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for forever
	// i := 1
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for con break
	// i := 1
	// for {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// ejemplo con slices decomida
	// food := []string{"🍎", "🍌", "🍇", "🍓", "🍊"}
	// for i, v := range food {
	// 	fmt.Println(i, v)
	// }

	// ejemplo con slice de comida
	// for i, v := range []string{"🍎", "🍌", "🍇", "🍓", "🍊"} {
	// 	fmt.Println(i, v)
	// }

	// ejemplo con slice de numeros
	// numbers := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for i := range numbers {
	// 	numbers[i] *= 2
	// }
	// fmt.Println(numbers)

	// ejemplo con map de comida
	// food := map[string]string{
	// 	"apple":      "🍎",
	// 	"banana":     "🍌",
	// 	"orange":     "🍊",
	// 	"pineapple":  "🍍",
	// 	"strawberry": "🍓",
	// 	"watermelon": "🍉",
	// }

	// for k, v := range food {
	// 	fmt.Println(k, v)
	// }

	for i, v := range "Hello World!" {
		fmt.Println("index:", i, "value:", string(v))
	}
}
