package main

import "fmt"

func main() {
	commands := []string{"ok", "retry", "ignored", "retry", "retry", "ok", "fatal"}
	okCount := 0
	retryCount := 0
	stopReason := "input completed"

commandLoop:
	for _, command := range commands {
		switch command {
		case "ok":
			okCount++
		case "retry":
			retryCount++
			if retryCount == 3 {
				stopReason = "retry budget exhausted"
				break commandLoop
			}
		case "ignored":
			continue
		case "fatal":
			stopReason = "fatal command"
			break commandLoop
		}
	}

	fmt.Printf("Stop reason: %s\n", stopReason)
	fmt.Printf("Summary: ok=%d retry=%d\n", okCount, retryCount)

	// El break etiquetado termina el loop completo; un break sin etiqueta solo terminaría el switch.
}
