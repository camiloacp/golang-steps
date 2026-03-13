package main

import (
	"fmt"
	"os"
)

// bastante util para cerrar archivos, conexiones, controladores de recursos, bases de datos, etc.

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello Gophers"))
	if err != nil {
		fmt.Println("Error al escribir en el archivo", err)
		return
	}
}
