package main

import (
	"fmt"
	"math"
)

func main() {
	var fahrenheit float64

	fmt.Println("Conversor de Fahrenheit a Celsius.\n")
	fmt.Print("Ingresa los grados Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)

	rawValue := (fahrenheit - 32) * 5 / 9
	temperatura := math.Floor(rawValue*100) / 100

	fmt.Println("La temperatura en Celsius es de:", temperatura, "grados\n")
}
