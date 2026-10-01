package main

import "fmt"

func main() {
	const weightPerUnit float64 = 2.75

	var facilityName string = "North Dock"
	var shipmentMarker rune = 'A'
	var deliveredUnits int = 1250
	var damagedUnits int = 18

	// Los valores cero comunican que la inspección no comenzó y que no existe una nota.
	var inspectionStarted bool
	var optionalNote string

	// Los resultados derivados permanecen vinculados con los datos originales del envío.
	acceptedUnits := deliveredUnits - damagedUnits
	totalAcceptedWeight := float64(acceptedUnits) * weightPerUnit

	fmt.Println("=== Inventory State Report ===")
	fmt.Printf("Facility: %s\n", facilityName)
	fmt.Printf("Shipment marker: %c\n", shipmentMarker)
	fmt.Printf("Delivered units: %d\n", deliveredUnits)
	fmt.Printf("Damaged units: %d\n", damagedUnits)
	fmt.Printf("Accepted units: %d\n", acceptedUnits)
	fmt.Printf("Weight per accepted unit: %.2f kg\n", weightPerUnit)
	fmt.Printf("Total accepted weight: %.2f kg\n", totalAcceptedWeight)
	fmt.Printf("Inspection started: %t\n", inspectionStarted)
	fmt.Printf("Optional note: %q\n", optionalNote)
}
