package main

import "fmt"

func main() {
	// Slice: Son operadoresa un array, no poseen datos
	things := [7]string{"💚", "❤️", "💜", "🌳", "🐦‍🔥", "🦁", "🐢"} // un slice apunta a un array
	hearts := things[:3]                                      // "💚", "❤️", "💜"
	nature := things[3:]                                      // "🌳", "🐦‍🔥", "🦁", "🐢"
	nature[1] = "🦋"

	fmt.Println("Things:", things)
	fmt.Println("Hearts:", hearts)
	fmt.Println("Nature:", nature)

	fmt.Println("Hearts[0]:", hearts[0])
	fmt.Println("Nature[0]:", nature[0])
}
