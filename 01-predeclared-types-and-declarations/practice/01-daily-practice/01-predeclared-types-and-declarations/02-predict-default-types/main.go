package main

import "fmt"

func main() {
	var count = 42
	var ratio = 3.5
	var letter = 'G'
	var message = "Go"
	var small int8 = 42

	fmt.Printf("%T %T %T %T %T\n", count, ratio, letter, message, small)
	// Predicción escrita antes de ejecutar: int float64 int32 string int8.

	fmt.Println(count, ratio, letter, message, small)
	// Predicción escrita antes de ejecutar: 42 3.5 71 Go 42.
	// La ejecución confirmó ambas predicciones.

	// Las primeras cuatro declaraciones usan los tipos predeterminados de sus literales.
	// En cambio, el tipo explícito de small aporta el contexto para representar 42 como int8.
}
