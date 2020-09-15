package main

import "fmt"

func main() {
	var base float32
	var height float32

	fmt.Println("Ingrese la base del triangulo:")
	fmt.Scan(&base)
	fmt.Println("Ingrese la altura del triangulo:")
	fmt.Scan(&height)
	fmt.Println("Area del triangulo: ", (base*height)/2)
}
