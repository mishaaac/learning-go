package main

import "fmt"

func main() {
	candidates := []string{"go", "api!", "世界", "bad#tag", "loop"}
	acceptedCount := 0
	rejectedCount := 0

candidatesLoop:
	for _, candidate := range candidates {
		for _, value := range candidate {
			if value == '!' || value == '#' {
				rejectedCount++
				continue candidatesLoop
			}
		}

		fmt.Printf("Accepted: %s\n", candidate)
		acceptedCount++
	}

	fmt.Printf("Summary: accepted=%d rejected=%d\n", acceptedCount, rejectedCount)

	// El continue etiquetado abandona inmediatamente el candidato rechazado y continúa con el siguiente.
}
