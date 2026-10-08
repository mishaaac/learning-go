package main

import "fmt"

type InternalJob struct {
	Name     string
	Attempts int
	Active   bool
}

type ExportJob struct {
	Name     string
	Attempts int
	Active   bool
}

func main() {
	internal := InternalJob{
		Name:     "Atlas",
		Attempts: 2,
		Active:   true,
	}

	exported := ExportJob(internal)

	// La estructura anónima es apropiada porque este snapshot solo se utiliza localmente una vez.
	var localSnapshot struct {
		Name     string
		Attempts int
		Active   bool
	} = internal

	convertedBack := InternalJob(exported)

	fmt.Println("=== Struct Type Boundaries ===")
	fmt.Printf("Internal job: %+v\n", internal)
	fmt.Printf("Export job: %+v\n", exported)
	fmt.Printf("Local snapshot: %+v\n", localSnapshot)
	fmt.Printf("Internal equals converted export: %t\n", internal == convertedBack)
	fmt.Printf("Internal equals local snapshot: %t\n", internal == localSnapshot)

	// Los tipos con nombre representan roles reutilizables distintos y evitan mezclarlos accidentalmente.
	// Aunque sus campos coinciden, ambos tipos tienen identidades diferentes y la asignación directa no está permitida.
	// La conversión explícita hace visible la decisión de cruzar ese límite entre dominios.
}
