package main

import "fmt"

func main() {
	readings := [...]int{2: 7, 4: 9}
	window := readings[1:4]

	window[1] = 8

	fmt.Println(len(readings), len(window), cap(window))
	// Predicción: 5 3 4
	fmt.Println(readings)
	// Predicción: [0 0 8 0 9]
	fmt.Println(window)
	// Predicción: [0 8 0]

	// El índice explícito más alto del literal es 4, por lo que Go infiere una longitud de 5.
	// El slice comparte el almacenamiento del array; window[1] corresponde a readings[2].
}
