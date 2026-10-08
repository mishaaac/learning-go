package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	total := 0

outer:
	for row := 1; row <= 3; row++ {
		for _, value := range values {
			switch {
			case value == 2:
				continue
			case row == 2 && value == 3:
				continue outer
			case value == 4:
				break
			default:
				total += row * value
			}
		}
		fmt.Println("row", row, total)
	}

	fmt.Println("total", total)

	// Predicción completa:
	// row 1 4
	// row 3 18
	// total 18
	// continue avanza la iteración del for interno y continue outer avanza la iteración del for etiquetado.
	// El continue etiquetado evita el reporte de la fila 2.
	// El break sin etiqueta termina solamente el switch, no ninguno de los dos loops.
}
