package main

import "fmt"

func main() {
	count := 10

	// La declaración corta crea un count nuevo porque se encuentra en otro bloque.
	{
		count, label := 20, "inner"
		fmt.Println(count, label)
		// Predicción escrita antes de ejecutar: 20 inner.
	}

	fmt.Println(count)
	// Predicción escrita antes de ejecutar: 10.
	// La ejecución confirmó ambas predicciones de la versión original.

	// En la segunda versión, la asignación modifica la variable declarada fuera del bloque.
	updatedCount := 10

	{
		updatedCount = 20
		label := "inner"
		fmt.Println(updatedCount, label)
	}

	fmt.Println(updatedCount)
}
