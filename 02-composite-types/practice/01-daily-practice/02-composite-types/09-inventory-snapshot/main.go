package main

import "fmt"

type Product struct {
	SKU   string
	Name  string
	Stock int
}

func productsEqual(left, right []Product) bool {
	if len(left) != len(right) {
		return false
	}

	for index := range left {
		if left[index] != right[index] {
			return false
		}
	}

	return true
}

func main() {
	workingSet := []Product{
		{SKU: "A100", Name: "Keyboard", Stock: 12},
		{SKU: "B200", Name: "Mouse", Stock: 0},
		{SKU: "C300", Name: "Monitor", Stock: 5},
	}

	snapshot := make([]Product, len(workingSet))
	copy(snapshot, workingSet)

	indexBySKU := make(map[string]int, len(workingSet))
	for index, product := range workingSet {
		indexBySKU[product.SKU] = index
	}

	adjustments := []struct {
		SKU   string
		Delta int
	}{
		{SKU: "A100", Delta: 3},
		{SKU: "B200", Delta: 5},
		{SKU: "X999", Delta: 8},
	}

	fmt.Println("=== Inventory Adjustments ===")
	for _, adjustment := range adjustments {
		index, exists := indexBySKU[adjustment.SKU]
		if !exists {
			fmt.Printf("SKU %s: delta=%+d accepted=false\n", adjustment.SKU, adjustment.Delta)
			continue
		}

		workingSet[index].Stock += adjustment.Delta
		fmt.Printf("SKU %s: delta=%+d accepted=true stock=%d\n", adjustment.SKU, adjustment.Delta, workingSet[index].Stock)
	}

	fmt.Println("Snapshot:")
	for _, product := range snapshot {
		fmt.Printf("  %s | %s | stock=%d\n", product.SKU, product.Name, product.Stock)
	}

	fmt.Println("Working set:")
	for _, product := range workingSet {
		fmt.Printf("  %s | %s | stock=%d\n", product.SKU, product.Name, product.Stock)
	}

	fmt.Printf("Contents are equal: %t\n", productsEqual(snapshot, workingSet))

	// La copia del slice es suficiente porque Product solo contiene strings e ints, no datos mutables referenciados.
	// comma-ok permite distinguir un SKU ausente de un índice cero válido en el mapa de búsqueda.
}
