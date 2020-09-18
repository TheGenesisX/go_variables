package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Println("Area del circulo.\n")
	fmt.Print("Ingresa la medida del radio: ")
	fmt.Scanf("%f", &r)

	rawValue := math.Pi * (math.Pow(r, 2))
	area := math.Floor(rawValue*100) / 100

	fmt.Print("Area del circulo: ", area, "\n")
}
