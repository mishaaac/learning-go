package main

import "fmt"

const maxRetries = 3

func main() {
	// El snippet original tenía dos errores de compilación relacionados con
	// declaraciones: el segundo count := 20 intentaba redeclarar count sin una
	// variable nueva en el mismo bloque, y la variable local status_message se
	// declaraba pero nunca se usaba. El const de paquete sin usar no era un error.
	count := 10

	// Una asignación actualiza count sin intentar declararlo nuevamente en el mismo bloque.
	count = 20
	statusMessage := "ready"

	fmt.Println("=== Declaration Cleanup ===")
	fmt.Printf("Maximum retries: %d\n", maxRetries)
	fmt.Printf("Current count: %d\n", count)
	fmt.Printf("Status: %s\n", statusMessage)
}
