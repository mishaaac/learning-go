package main

import "fmt"

func main() {
	label := "Go→Lima🌞"
	runeCount := 0

	fmt.Println("=== Unicode Offset Report ===")
	for byteOffset, value := range label {
		fmt.Printf("Byte offset %d: rune=%U character=%c\n", byteOffset, value, value)
		runeCount++
	}

	fmt.Printf("Totals: bytes=%d runes=%d\n", len(label), runeCount)

	// El índice producido por range es un desplazamiento de bytes, no un contador de runas.
}
