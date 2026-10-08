package main

import "fmt"

func main() {
	status := "queued"
	count := 1

	if count > 0 {
		status, count := "running", count+1
		fmt.Println(status, count)

		if count == 2 {
			status = "checked"
			note := status
			fmt.Println(note, count)
		}

		fmt.Println(status, count)
	}

	fmt.Println(status, count)

	// Predicción completa:
	// running 2
	// checked 2
	// checked 2
	// queued 1
	// La declaración corta crea status y count nuevos dentro del primer if porque los anteriores están en otro bloque.
	// La asignación del segundo if modifica el status del primer if; note solo existe dentro del bloque más interno.
	// La última línea observa las variables de main, que nunca fueron modificadas por las variables internas.
}
