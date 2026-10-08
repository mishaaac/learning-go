package main

import "fmt"

func main() {
	original := []int{10, 20, 30, 40}
	shared := original[:2]
	limited := original[:2:2]

	shared = append(shared, 99)
	limited = append(limited, 77)
	limited[0] = 5

	fmt.Println(original)
	// Predicción: [10 20 99 40]
	fmt.Println(shared)
	// Predicción: [10 20 99]
	fmt.Println(limited)
	// Predicción: [5 20 77]

	// shared conserva capacidad disponible y su append reutiliza el array subyacente de original.
	// limited tiene su capacidad restringida a 2, por lo que append crea un almacenamiento independiente.
	// La actualización posterior mediante limited no afecta a original porque ya no comparten almacenamiento.
}
