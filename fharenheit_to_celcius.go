package main

import (
	"fmt"
)

func main() {
	var grades float64

	fmt.Println("Ingrese los grados Fahrenheit a Celcius:")
	fmt.Scan(&grades)
	var result float64 = (grades - 32) * 5 / 9
	fmt.Println("grados Fahrenheit a Celcius: ", result)
}
