package main

import "fmt"

type Resource struct {
	Identifier string
	Label      string
	Priority   int
}

func main() {
	input := []Resource{
		{Identifier: "A1", Label: "Go", Priority: 3},
		{Identifier: "B2", Label: "世界", Priority: 0},
		{Identifier: "A1", Label: "duplicate", Priority: 9},
		{Identifier: "C3", Label: "🌞", Priority: 2},
	}

	catalog := make([]Resource, 0, len(input))
	resourcesByID := make(map[string]Resource, len(input))
	seenIDs := make(map[string]struct{}, len(input))
	duplicateCount := 0

	for _, resource := range input {
		if _, exists := seenIDs[resource.Identifier]; exists {
			duplicateCount++
			continue
		}

		seenIDs[resource.Identifier] = struct{}{}
		resourcesByID[resource.Identifier] = resource
		catalog = append(catalog, resource)
	}

	fmt.Println("=== Resource Registry ===")
	for index, resource := range catalog {
		fmt.Printf("Resource %d: id=%s label=%q priority=%d\n", index+1, resource.Identifier, resource.Label, resource.Priority)
	}
	fmt.Printf("Totals: input=%d accepted=%d duplicates=%d\n", len(input), len(catalog), duplicateCount)

	b2, b2Exists := resourcesByID["B2"]
	z9, z9Exists := resourcesByID["Z9"]
	fmt.Printf("Lookup B2: value=%+v present=%t\n", b2, b2Exists)
	fmt.Printf("Lookup Z9: value=%+v present=%t\n", z9, z9Exists)

	// El slice conserva los recursos aceptados en el orden de entrada.
	// El mapa resourcesByID permite recuperar un recurso directamente mediante su identificador.
	// El mapa seenIDs funciona como un conjunto para detectar identificadores duplicados.
	// comma-ok distingue la prioridad cero válida de B2 del valor cero producido por la ausencia de Z9.
}
