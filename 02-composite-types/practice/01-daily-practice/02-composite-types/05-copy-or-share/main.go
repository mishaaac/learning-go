package main

import "fmt"

func main() {
	measurements := []int{10, 20, 30, 40}

	sliceSnapshot := make([]int, len(measurements), len(measurements)+2)
	copiedElements := copy(sliceSnapshot, measurements)
	arrayValue := [4]int(measurements)
	arrayPointer := (*[4]int)(measurements)

	measurements[0] = 11
	arrayPointer[1] = 22
	sliceSnapshot[2] = 300
	arrayValue[3] = 400

	fmt.Println("=== Copy or Share ===")
	fmt.Printf("Copied elements: %d\n", copiedElements)
	fmt.Printf("Measurements:    %v\n", measurements)
	fmt.Printf("Slice snapshot:  %v\n", sliceSnapshot)
	fmt.Printf("Array value:     %v\n", arrayValue)
	fmt.Printf("Array pointer:   %v\n", *arrayPointer)

	// copy solo puede escribir tantos elementos como permita la longitud del destino; su capacidad adicional no crea índices.
	// La conversión a un valor array copia los datos, mientras que el puntero al array comparte el almacenamiento del slice.
}
