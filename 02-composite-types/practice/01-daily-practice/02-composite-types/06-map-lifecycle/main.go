package main

import "fmt"

func main() {
	inventory := map[string]int{
		"keyboard": 12,
		"mouse":    0,
		"legacy":   4,
	}

	mouseStock, mouseExists := inventory["mouse"]
	monitorStock, monitorExists := inventory["monitor"]

	fmt.Println("=== Map Lifecycle ===")
	fmt.Printf("Mouse: stock=%d present=%t\n", mouseStock, mouseExists)
	fmt.Printf("Monitor: stock=%d present=%t\n", monitorStock, monitorExists)

	inventory["keyboard"] = 15
	inventory["monitor"] = 6
	delete(inventory, "legacy")

	_, legacyExists := inventory["legacy"]
	fmt.Printf("After updates: keyboard=%d mouse=%d monitor=%d\n", inventory["keyboard"], inventory["mouse"], inventory["monitor"])
	fmt.Printf("After delete: legacy present=%t entries=%d\n", legacyExists, len(inventory))

	clear(inventory)
	fmt.Printf("After clear: entries=%d\n", len(inventory))

	inventory["cable"] = 9
	cableStock, cableExists := inventory["cable"]
	fmt.Printf("Next cycle: cable=%d present=%t entries=%d\n", cableStock, cableExists, len(inventory))

	// El resultado adicional de comma-ok distingue un valor cero almacenado de una clave ausente.
	// clear elimina las entradas, pero conserva el mapa inicializado y listo para nuevas escrituras.
}
