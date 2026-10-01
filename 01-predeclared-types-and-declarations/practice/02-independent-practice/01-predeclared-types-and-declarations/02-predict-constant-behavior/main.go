package main

import "fmt"

const flexible = 255
const fixed int = 3

func main() {
	var octet byte = flexible
	var measurement float64 = flexible
	quotient := 7 / 2
	precise := float64(7) / 2
	converted := float64(fixed)

	fmt.Printf("%T %T %T %T %T\n", octet, measurement, quotient, precise, converted)
	// Predicción escrita antes de ejecutar: uint8 float64 int float64 float64

	fmt.Println(octet, measurement, quotient, precise, converted)
	// Predicción escrita antes de ejecutar: 255 255 3 3.5 3
	// La ejecución confirmó los tipos y valores de los cinco campos.

	// flexible es una constante entera sin tipo. El contexto declara octet como byte
	// y measurement como float64; el valor 255 puede representarse en ambos tipos.
	// La expresión constante 7 / 2 usa división entera y produce 3; al no existir
	// otro contexto, quotient recibe el tipo predeterminado int.
	// En precise, convertir 7 a float64 hace que 2 adopte ese tipo y conserva la fracción.
	// fixed ya es una constante de tipo int, por lo que su conversión explícita produce
	// el valor 3 con tipo float64 para inicializar converted.
}
