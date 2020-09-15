package main

import "fmt"

func main() {
	var side float32

	fmt.Println("Ingrese el lado del cuadrado:")
	fmt.Scan(&side)
	fmt.Println("Area del cuadrado: ", side*side)
}
