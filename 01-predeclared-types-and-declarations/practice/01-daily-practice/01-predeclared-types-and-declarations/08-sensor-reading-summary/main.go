package main

import "fmt"

func main() {
	sampleCount := 3
	accumulatedTemperature := 73.5

	const warningThreshold = 24

	sensorLabel := "Room Sensor"
	unitSymbol := '℃'

	// Convertimos sampleCount antes de dividir para conservar la parte decimal.
	// Convertir el promedio después sería demasiado tarde si la división fuera entera.
	average := accumulatedTemperature / float64(sampleCount)

	temperatureWarning := average > warningThreshold

	fmt.Println("=== Sensor Reading Summary ===")
	fmt.Printf("Sensor: %s\n", sensorLabel)
	fmt.Printf("Samples: %d\n", sampleCount)
	fmt.Printf("Average temperature: %.1f%c\n", average, unitSymbol)
	fmt.Printf("Temperature warning: %t\n", temperatureWarning)
}
