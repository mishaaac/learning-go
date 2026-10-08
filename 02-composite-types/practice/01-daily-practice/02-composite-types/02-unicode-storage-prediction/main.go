package main

import "fmt"

func main() {
	text := "Go, 世界"
	bytes := []byte(text)
	runes := []rune(text)
	window := text[4:7]

	fmt.Println(len(text), len(bytes), len(runes))
	// Predicción: 10 10 6
	fmt.Printf("%q %d %c\n", window, text[4], runes[4])
	// Predicción: "世" 228 世

	// Los caracteres 世 y 界 ocupan tres bytes cada uno en UTF-8, pero cada uno cuenta como una runa.
	// text[4] obtiene el primer byte de 世, mientras que runes[4] obtiene el punto de código completo.
}
