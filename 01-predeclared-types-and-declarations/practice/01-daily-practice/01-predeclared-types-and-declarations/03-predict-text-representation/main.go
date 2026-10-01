package main

import "fmt"

func main() {
	text := "世界"
	var symbol rune = '世'
	interpreted := "line1\nline2"
	raw := `line1\nline2`

	fmt.Printf("%T %T\n", text[0], symbol)
	// Predicción escrita antes de ejecutar: uint8 int32.
	// Al indexar un string se obtiene uno de sus bytes UTF-8, por eso text[0]
	// tiene el tipo byte (alias de uint8). symbol fue declarado como rune,
	// alias de int32, para representar el punto de código Unicode '世'.

	fmt.Println(len(text) > 2)
	// Predicción escrita antes de ejecutar: true.

	fmt.Println(interpreted == raw)
	// Predicción escrita antes de ejecutar: false.
	// La ejecución confirmó las tres predicciones.

	// La cadena interpretada convierte \n en un salto de línea; la cadena cruda conserva
	// esos dos caracteres literalmente, por lo que sus valores no son iguales.
}
