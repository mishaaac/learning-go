package main

import "fmt"

func main() {
	routes := []struct {
		Name  string
		Codes []int
	}{
		{Name: "North", Codes: []int{2, 0, 8, -1, 99}},
		{Name: "East", Codes: []int{5, 11, 4}},
		{Name: "South", Codes: []int{7, -9, 3}},
		{Name: "West", Codes: []int{1}},
	}

	completedRoutes := 0
	lightCount := 0
	standardCount := 0
	heavyCount := 0

dispatchLoop:
	for _, route := range routes {
		routeCodeCount := 0

		for _, code := range route.Codes {
			switch code {
			case 0:
				continue
			case -1:
				completedRoutes++
				fmt.Printf("Route %s: closed after %d positive codes\n", route.Name, routeCodeCount)
				continue dispatchLoop
			case -9:
				fmt.Printf("System shutdown: route=%s code=%d\n", route.Name, code)
				break dispatchLoop
			}

			classification := ""
			switch {
			case code < 5:
				classification = "light"
				lightCount++
			case code < 10:
				classification = "standard"
				standardCount++
			default:
				classification = "heavy"
				heavyCount++
			}

			routeCodeCount++
			fmt.Printf("Route %s: code=%d classification=%s\n", route.Name, code, classification)
		}

		completedRoutes++
		fmt.Printf("Route %s: completed after %d positive codes\n", route.Name, routeCodeCount)
	}

	fmt.Printf("Totals: completed=%d light=%d standard=%d heavy=%d\n",
		completedRoutes, lightCount, standardCount, heavyCount)

	// Los totales globales se declaran fuera de los loops para que sobrevivan a cualquier salida etiquetada.
	// routeCodeCount pertenece solo a una ruta y deja de existir al terminar su iteración externa.
	// Los for-range conservan el orden; el switch con expresión maneja controles y el switch sin expresión clasifica rangos.
}
