package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	fmt.Println("Ingrese el radio del circulo:")
	fmt.Scan(&radius)
	result := math.Pow(math.Pi*radius, 2)
	fmt.Println("Area del circulo: ", result)
}
