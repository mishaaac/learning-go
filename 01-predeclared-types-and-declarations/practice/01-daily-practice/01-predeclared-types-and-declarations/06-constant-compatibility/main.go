package main

import "fmt"

func main() {
	const protocolLimit = 255

	var byteLimit byte = protocolLimit
	var integerLimit int = protocolLimit
	var floatingPointLimit float64 = protocolLimit

	const retryCount int = 255
	scaledRetryCount := float64(retryCount) * 4.5

	const calculatedLimit = 25 * 4.5

	// Si protocolLimit fuera 256, solo fallaría byteLimit porque byte admite valores de 0 a 255.
	// Los destinos int y float64 sí pueden representar ese valor.

	fmt.Println("=== Constant Compatibility ===")
	fmt.Printf("Byte limit: %d\n", byteLimit)
	fmt.Printf("Integer limit: %d\n", integerLimit)
	fmt.Printf("Floating-point limit: %.1f\n", floatingPointLimit)
	fmt.Printf("Scaled retry count: %.1f\n", scaledRetryCount)
	fmt.Printf("Constant expression result: %.1f\n", calculatedLimit)
}
