package main

import "fmt"

func main() {
	acceptedCount := 0
	acceptedTotal := 0

	for ticket := 1; ticket <= 12; ticket++ {
		if ticket%2 == 0 {
			continue
		}

		if ticket > 9 {
			break
		}

		fmt.Printf("Accepted ticket: %d\n", ticket)
		acceptedCount++
		acceptedTotal += ticket
	}

	fmt.Printf("Summary: count=%d total=%d\n", acceptedCount, acceptedTotal)

	// La actualización permanece en la cabecera para que continue no impida el avance del ticket.
}
