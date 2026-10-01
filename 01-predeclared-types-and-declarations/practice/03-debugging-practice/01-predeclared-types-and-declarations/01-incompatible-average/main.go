package main

import "fmt"

func main() {
	var sampleCount int = 4
	var accumulatedTemperature float64 = 97.0

	// Convertimos el conteo antes de dividir para compatibilizar los tipos y conservar la fracción.
	averageTemperature := accumulatedTemperature / float64(sampleCount)

	fmt.Println("=== Temperature Average ===")
	fmt.Printf("Sample count: %d\n", sampleCount)
	fmt.Printf("Accumulated temperature: %.2f\n", accumulatedTemperature)
	fmt.Printf("Average temperature: %.2f\n", averageTemperature)
}
