package main

import "fmt"

func main() {
	fmt.Println(Includes([]string{"a", "b", "c"}, "b"))
	fmt.Println(Includes([]int{1, 2, 3}, 2))
	fmt.Println(Includes([]string{"a", "b", "c"}, "d"))
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
