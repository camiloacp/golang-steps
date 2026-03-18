package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 3.2))
}

type Number interface {
	~int | ~float64 | ~float32 | ~uint
}

func sum[T Number](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
