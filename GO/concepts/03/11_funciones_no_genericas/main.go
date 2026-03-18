package main

import "fmt"

func main() {
	PrintList("apple", "banana", "cherry")
	PrintList(1, 2, 3, 4, 5)
}

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
