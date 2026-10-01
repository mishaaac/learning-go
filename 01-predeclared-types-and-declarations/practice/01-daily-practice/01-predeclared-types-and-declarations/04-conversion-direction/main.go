package main

import "fmt"

func main() {
	var itemCount int = 5
	var unitPrice float64 = 3.5

	floatingPointTotal := float64(itemCount) * unitPrice

	// Convertir el precio a int descarta su parte decimal antes de la multiplicación.
	integerTotal := itemCount * int(unitPrice)

	originalValue := 300

	// La conversión a byte conserva solo los ocho bits menos significativos de 300.
	narrowedValue := byte(originalValue)

	fmt.Println("=== Conversion Direction ===")
	fmt.Printf("Total with decimal price: %.1f\n", floatingPointTotal)
	fmt.Printf("Total after price truncation: %d\n", integerTotal)
	fmt.Printf("Original integer: %d\n", originalValue)
	fmt.Printf("Value after byte conversion: %d\n", narrowedValue)
}
