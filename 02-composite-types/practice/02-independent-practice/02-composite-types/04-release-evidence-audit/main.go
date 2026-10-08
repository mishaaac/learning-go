package main

import "fmt"

type Evidence struct {
	Identifier  string
	Stage       string
	Description string
	ByteCount   int
	RuneCount   int
}

func main() {
	requiredStages := [4]string{"plan", "build", "test", "deploy"}
	requiredSet := make(map[string]struct{}, len(requiredStages))
	for _, stage := range requiredStages {
		requiredSet[stage] = struct{}{}
	}

	input := []struct {
		Identifier  string
		Stage       string
		Description string
	}{
		{Identifier: "K1", Stage: "plan", Description: "Draft"},
		{Identifier: "K2", Stage: "build", Description: "Café"},
		{Identifier: "K2", Stage: "test", Description: "duplicate"},
		{Identifier: "K3", Stage: "review", Description: "skip"},
		{Identifier: "K4", Stage: "test", Description: "世界"},
		{Identifier: "K5", Stage: "deploy", Description: "🚀"},
	}

	accepted := make([]Evidence, 0, len(input))
	usedIdentifiers := make(map[string]struct{}, len(input))
	stageCounts := make(map[string]int, len(requiredStages))

	duplicateRejections := 0
	unknownStageRejections := 0
	totalBytes := 0
	totalRunes := 0

	for _, entry := range input {
		if _, required := requiredSet[entry.Stage]; !required {
			unknownStageRejections++
			continue
		}

		if _, used := usedIdentifiers[entry.Identifier]; used {
			duplicateRejections++
			continue
		}

		byteCount := len([]byte(entry.Description))
		runeCount := len([]rune(entry.Description))
		accepted = append(accepted, Evidence{
			Identifier:  entry.Identifier,
			Stage:       entry.Stage,
			Description: entry.Description,
			ByteCount:   byteCount,
			RuneCount:   runeCount,
		})

		// El identificador se registra como usado únicamente después de aceptar la evidencia.
		usedIdentifiers[entry.Identifier] = struct{}{}
		stageCounts[entry.Stage]++
		totalBytes += byteCount
		totalRunes += runeCount
	}

	fmt.Println("=== Release Evidence Audit ===")
	for index, evidence := range accepted {
		fmt.Printf("Evidence %d: id=%s stage=%s description=%q bytes=%d runes=%d\n",
			index+1,
			evidence.Identifier,
			evidence.Stage,
			evidence.Description,
			evidence.ByteCount,
			evidence.RuneCount,
		)
	}

	fmt.Printf("Totals: input=%d accepted=%d duplicate-id=%d unknown-stage=%d\n",
		len(input), len(accepted), duplicateRejections, unknownStageRejections)

	completeCoverage := true
	for _, stage := range requiredStages {
		count, exists := stageCounts[stage]
		fmt.Printf("Stage %s: count=%d present=%t\n", stage, count, exists)
		if !exists || count == 0 {
			completeCoverage = false
		}
	}
	fmt.Printf("Complete coverage: %t\n", completeCoverage)
	fmt.Printf("Description totals: bytes=%d runes=%d\n", totalBytes, totalRunes)

	archive := make([]Evidence, len(accepted))
	copy(archive, accepted)
	clear(accepted)

	fmt.Printf("Cleared working records: len=%d values=%+v\n", len(accepted), accepted)
	fmt.Printf("Independent archive: len=%d values=%+v\n", len(archive), archive)

	// El array representa el conjunto fijo de etapas y proporciona un orden determinista para el reporte.
	// El slice conserva las evidencias aceptadas en orden, mientras los conjuntos validan pertenencia y unicidad.
	// El mapa de conteos permite consultar directamente la cobertura de cada etapa requerida.
}
