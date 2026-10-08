package main

import "fmt"

func main() {
	readings := []int{-2, 0, 4, 9, 15, 6}
	rejectedCount := 0
	standardCount := 0
	elevatedCount := 0
	criticalCount := 0

intakeLoop:
	for _, reading := range readings {
		switch {
		case reading < 0:
			rejectedCount++
			fmt.Printf("Reading %d: rejected\n", reading)
		case reading == 0:
			fmt.Printf("Reading %d: ignored\n", reading)
		case reading < 8:
			standardCount++
			fmt.Printf("Reading %d: standard\n", reading)
		case reading < 15:
			elevatedCount++
			fmt.Printf("Reading %d: elevated\n", reading)
		default:
			criticalCount++
			fmt.Printf("Reading %d: critical\n", reading)
			break intakeLoop
		}
	}

	fmt.Printf("Totals: rejected=%d standard=%d elevated=%d critical=%d\n",
		rejectedCount, standardCount, elevatedCount, criticalCount)

	// for-range procesa el stream en orden y el break etiquetado impide alcanzar datos posteriores al evento crítico.
	// El switch sin expresión modela una sola decisión ordenada, por lo que cada lectura tiene como máximo un resultado.
}
