package main

import (
	"fmt"
	"math"
)

func main() {
	var b float64
	var h float64

	fmt.Println("Area del triangulo.\n")
	fmt.Print("Ingresa la medida de la base: ")
	fmt.Scanf("%f", &b)
	fmt.Print("Ingresa la medida de la altura: ")
	fmt.Scanf("%f", &h)

	rawValue := (b * h) / 2
	area := math.Floor(rawValue*100) / 100

	fmt.Print("Area del triangulo: ", area, "\n")
}
