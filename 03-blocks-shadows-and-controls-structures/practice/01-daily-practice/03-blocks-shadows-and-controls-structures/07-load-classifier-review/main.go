package main

import "fmt"

func main() {
	readings := []int{0, 3, 7, 10, 14}

	for _, reading := range readings {
		classification := ""

		switch {
		case reading == 0:
			classification = "idle"
		case reading < 7:
			classification = "normal"
		case reading < 10:
			classification = "high"
		default:
			classification = "critical"
		}

		fmt.Printf("Load %d: %s\n", reading, classification)
	}

	// for-range comunica que se procesa cada lectura una vez y conserva el orden del slice.
	// El switch sin expresión presenta los umbrales como una sola decisión ordenada y mutuamente excluyente.
	// Una cadena if-else también sería válida, pero no ofrecería una ventaja de claridad para estas categorías paralelas.
}
