package main

import (
	"fmt"
	"math"
)

func main() {
	var l float64

	fmt.Println("Area del cuadrado.\n")
	fmt.Print("Ingresa la medida de l: ")
	fmt.Scanf("%f", &l)

	rawValue := l * l
	area := math.Floor(rawValue*100) / 100

	fmt.Print("Area del cuadrado: ", area, "\n")
}
