package main

import "fmt"

type Report struct {
	Name   string
	Values []int
}

func main() {
	// Diagnóstico previo a la reparación:
	// 1. Escribir en totals sin inicializarlo provoca un fallo en ejecución porque el mapa es nil.
	// 2. Acceder a values[0] con longitud cero provoca un fallo en ejecución; la capacidad no crea índices válidos.
	// 3. copy no copia elementos hacia snapshot con longitud cero, lo que produce un resultado incorrecto.
	// 4. Comparar dos Report con == causa un fallo de compilación porque el campo slice no es comparable.

	totals := make(map[string]int)
	totals["ready"] = 1

	values := make([]int, 1, 3)
	values[0] = 10
	values = append(values, 20, 30)

	snapshot := make([]int, len(values))
	copied := copy(snapshot, values)

	current := Report{Name: "batch", Values: values}
	archived := Report{Name: "batch", Values: snapshot}
	equalBeforeMutation := current.Name == archived.Name &&
		[3]int(current.Values) == [3]int(archived.Values)

	fmt.Println(equalBeforeMutation)
	values[0] = 99
	fmt.Println(totals["ready"], copied)
	fmt.Println(current.Values)
	fmt.Println(archived.Values)

	// Los arrays de tres enteros sí son comparables, por eso las conversiones permiten revisar todas las lecturas.
	// La comparación también revisa el nombre para cubrir todo el contenido de cada reporte.
	// El snapshot tiene longitud suficiente para copy y almacenamiento propio, por eso la mutación no cambia archived.
}
