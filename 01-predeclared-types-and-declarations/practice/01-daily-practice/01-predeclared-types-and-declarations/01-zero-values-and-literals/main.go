package main

import "fmt"

func main() {
	// Estas declaraciones muestran los valores cero que Go asigna automáticamente.
	var isReady bool
	var itemCount int
	var statusMessage string

	decimalValue := 42
	binaryValue := 0b_0010_1010
	octalValue := 0o_52
	hexadecimalValue := 0x_2A

	var letter rune = 'A'
	text := "A"

	fmt.Println("=== Zero Values and Literals ===")
	fmt.Printf("Boolean zero value: %t\n", isReady)
	fmt.Printf("Numeric zero value: %d\n", itemCount)
	fmt.Printf("String zero value: %q\n", statusMessage)

	fmt.Println("Integer representations:")
	fmt.Printf("Decimal literal value: %d\n", decimalValue)
	fmt.Printf("Binary literal value: %d\n", binaryValue)
	fmt.Printf("Octal literal value: %d\n", octalValue)
	fmt.Printf("Hexadecimal literal value: %d\n", hexadecimalValue)

	fmt.Println("Text representations:")
	fmt.Printf("Rune: %c (code point %d)\n", letter, letter)
	fmt.Printf("String: %q\n", text)
}
