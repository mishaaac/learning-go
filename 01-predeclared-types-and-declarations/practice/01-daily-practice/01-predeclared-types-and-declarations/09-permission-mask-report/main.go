package main

import "fmt"

func main() {
	const (
		readMask    byte = 0b001
		writeMask   byte = 0b010
		executeMask byte = 0b100
	)

	// Combinamos las máscaras para habilitar los permisos de lectura y ejecución.
	var permissions byte = readMask | executeMask

	// Cada operación AND permite comprobar un permiso sin alterar el valor combinado.
	canRead := permissions&readMask != 0
	canWrite := permissions&writeMask != 0
	canExecute := permissions&executeMask != 0

	// La notación hexadecimal representa el mismo patrón de bits.
	var hexPermissions byte = 0x05

	matchesHexValue := permissions == hexPermissions

	fmt.Println("=== Permission Mask Report ===")
	fmt.Printf("Permission value: %d\n", permissions)
	fmt.Printf("Binary value:     %03b\n", permissions)
	fmt.Printf("Read:             %t\n", canRead)
	fmt.Printf("Write:            %t\n", canWrite)
	fmt.Printf("Execute:          %t\n", canExecute)
	fmt.Printf("Hex value:        0x%02X\n", hexPermissions)
	fmt.Printf("Values are equal: %t\n", matchesHexValue)
}
