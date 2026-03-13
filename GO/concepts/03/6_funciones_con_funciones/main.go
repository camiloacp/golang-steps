package main

import "fmt"

// func main() {
// 	nums := []int{2, 5, 12, 23, 92, 21, 79}

// 	result := filter(nums, func(num int) bool { return num > 10 })
// 	fmt.Println(result)
// }

// func filter(nums []int, callback func(int) bool) []int {
// 	result := make([]int, 0, len(nums))

// 	for _, num := range nums {
// 		if callback(num) {
// 			result = append(result, num)
// 		}
// 	}

// 	return result
// }

// otra forma de implementarlo
func main() {
	nums := []int{2, 5, 12, 23, 92, 21, 79}

	result := filter(nums, greatherInt) // lessInt
	fmt.Println(result)
}

func greatherInt(num int) bool {
	return num > 50
}

func lessInt(num int) bool {
	return num < 10
}

func filter(nums []int, callback func(int) bool) []int {
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}

	return result
}
