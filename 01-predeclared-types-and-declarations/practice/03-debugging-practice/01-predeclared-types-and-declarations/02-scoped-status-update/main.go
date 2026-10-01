package main

import "fmt"

func main() {
	jobState := "queued"
	attemptCount := 0

	{
		// La asignación actualiza los bindings externos en lugar de crear variables internas.
		jobState = "running"
		attemptCount = attemptCount + 1

		fmt.Println("=== Job Status Update ===")
		fmt.Printf("During update: state=%s, attempts=%d\n", jobState, attemptCount)
	}

	fmt.Printf("Stored state: state=%s, attempts=%d\n", jobState, attemptCount)
}
