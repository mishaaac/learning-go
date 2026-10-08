package main

import "fmt"

func main() {
	batchSlots := [4]string{0: "queued", 2: "complete"}

	var optionalResults []int
	emptyResults := []int{}
	labels := []string{"build", "test", "deploy"}
	outputSpace := make([]int, 3)

	results := make([]int, 0, 5)
	results = append(results, 21, 34)

	fmt.Println("=== Sequence Starting States ===")
	fmt.Printf("Fixed batch slots: %q\n", batchSlots)
	fmt.Printf("Nil slice: values=%v len=%d cap=%d nil=%t\n",
		optionalResults, len(optionalResults), cap(optionalResults), optionalResults == nil)
	fmt.Printf("Empty non-nil slice: values=%v len=%d cap=%d nil=%t\n",
		emptyResults, len(emptyResults), cap(emptyResults), emptyResults == nil)
	fmt.Printf("Known labels: %v len=%d cap=%d\n", labels, len(labels), cap(labels))
	fmt.Printf("Output space before writes: %v len=%d cap=%d\n",
		outputSpace, len(outputSpace), cap(outputSpace))

	outputSpace[0] = 7
	outputSpace[1] = 14
	outputSpace[2] = 21
	fmt.Printf("Output space after writes: %v\n", outputSpace)
	fmt.Printf("Reserved results: %v len=%d cap=%d\n", results, len(results), cap(results))

	// La capacidad reservada prepara almacenamiento, pero no crea índices válidos más allá de la longitud.
}
