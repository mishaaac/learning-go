package main

import "fmt"

// int conserva el tipo del programa original para el tamaño del lote y representa exactamente 10.
const batchSize int = 10

func main() {
	// Diagnóstico del programa original:
	// 1. completed / batchSize mezcla valores con tipos int32 e int, que no son compatibles.
	// 2. El valor entero 1 no puede inicializar un bool porque Go no utiliza truthiness.
	// 3. La constante 300 no es representable como byte, cuyo valor máximo es 255.
	// 4. label := "complete" no declara una variable nueva en el mismo bloque.
	// 5. unusedMessage se declara como variable local, pero el reporte original no la utiliza.

	// int32 conserva el tipo fijo del dato original y representa exactamente el conteo 7.
	var completedCount int32 = 7

	// Convertimos ambos operandos a float64 para compatibilizar sus tipos y conservar la fracción;
	// por ese contexto, progressRatio se infiere también como float64.
	progressRatio := float64(completedCount) / float64(batchSize)

	// La comparación produce el tipo bool requerido sin intentar convertir un número a bool.
	hasCompletedItems := completedCount != 0

	// int puede conservar exactamente el código 300, a diferencia de byte.
	var batchCode int = 300

	// Los literales de texto hacen que ambas declaraciones cortas infieran string,
	// el tipo apropiado para la etiqueta y el mensaje del reporte.
	statusLabel := "ready"
	statusLabel = "complete"
	statusMessage := "batch processed"

	fmt.Println("=== Batch Progress Report ===")
	fmt.Printf("Batch size: %d\n", batchSize)
	fmt.Printf("Completed count: %d\n", completedCount)
	fmt.Printf("Progress ratio: %.2f\n", progressRatio)
	fmt.Printf("Has completed items: %t\n", hasCompletedItems)
	fmt.Printf("Batch code: %d\n", batchCode)
	fmt.Printf("Status label: %s\n", statusLabel)
	fmt.Printf("Message: %s\n", statusMessage)
}
