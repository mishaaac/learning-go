package main

import "fmt"

func main() {
	values := [5]int{10, 20, 30, 40, 50}
	shared := values[1:3]
	guarded := values[1:3:3]
	clone := make([]int, len(shared))
	copy(clone, shared)

	shared[0] = 99
	shared = append(shared, 77)
	guarded = append(guarded, 88)
	clear(shared[:2])

	fmt.Println(values)
	// Predicción: [10 0 0 77 50]
	fmt.Println(shared)
	// Predicción: [0 0 77]
	fmt.Println(guarded)
	// Predicción: [99 30 88]
	fmt.Println(clone)
	// Predicción: [20 30]

	// shared mantiene capacidad disponible, por lo que su append escribe 77 en el array values.
	// guarded tiene la capacidad restringida y su append crea almacenamiento independiente después de copiar 99 y 30.
	// clone ya era una copia independiente creada antes de las mutaciones y conserva 20 y 30.
	// clear pone en cero los dos primeros elementos de shared, que todavía comparte almacenamiento con values.
}
