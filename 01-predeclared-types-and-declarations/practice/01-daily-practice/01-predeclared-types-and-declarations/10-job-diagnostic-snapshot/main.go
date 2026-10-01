package main

import "fmt"

const (
	jobName                 = "Atlas"
	maximumAttempts     int = 3
	completionThreshold     = 75
)

func main() {
	// int32 expresa de forma explícita el tipo requerido para las unidades completadas.
	var completedUnits int32 = 7
	var totalUnits int = 10

	attemptsUsed := 1

	// El porcentaje es una variable porque depende de valores calculados durante la ejecución.
	completionPercentage := float64(completedUnits) / float64(totalUnits) * 100

	remainingAttempts := maximumAttempts - attemptsUsed
	hasReachedThreshold := completionPercentage >= completionThreshold

	statusMarker := '✓'

	// La ausencia de un inicializador permite mostrar el valor cero de int.
	var errorCount int

	heading := "=== Diagnostic Snapshot - " + jobName + " ==="

	fmt.Println(heading)
	fmt.Printf("Completed units: %d\n", completedUnits)
	fmt.Printf("Total units: %d\n", totalUnits)
	fmt.Printf("Completion: %.0f%%\n", completionPercentage)
	fmt.Printf("Attempts used: %d\n", attemptsUsed)
	fmt.Printf("Remaining attempts: %d\n", remainingAttempts)
	fmt.Printf("Reached %d%% threshold: %t\n", completionThreshold, hasReachedThreshold)
	fmt.Printf("Status marker: %c\n", statusMarker)
	fmt.Printf("Error count: %d\n", errorCount)
}
