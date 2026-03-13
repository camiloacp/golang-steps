package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	tech := map[string]string{
		"computer": "💻",
		"mouse":    "🖱️",
	}

	fmt.Println(tech)

	//eliminar elementos
	delete(tech, "computer")
	fmt.Println(tech)

	fmt.Println(music["violin"])
	fmt.Println(music["fake"])

	content, ok := music["fake"]
	fmt.Println(content, ok)
}
