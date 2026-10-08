package main

import "fmt"

func main() {
	routes := []string{"north", "api!tail", "世界", "runXstop", "later"}
	completedRoutes := 0
	rejectedRoutes := 0
	ordinaryRunes := 0

auditLoop:
	for _, route := range routes {
		for _, value := range route {
			switch value {
			case '!':
				rejectedRoutes++
				fmt.Printf("Route %q: rejected\n", route)
				continue auditLoop
			case 'X':
				fmt.Printf("Route %q: shutdown\n", route)
				break auditLoop
			default:
				ordinaryRunes++
			}
		}

		completedRoutes++
		fmt.Printf("Route %q: completed\n", route)
	}

	fmt.Printf("Totals: completed=%d rejected=%d ordinary-runes=%d\n",
		completedRoutes, rejectedRoutes, ordinaryRunes)

	// continue auditLoop expresa que ! rechaza únicamente la ruta actual sin inspeccionar su texto restante.
	// break auditLoop expresa que X detiene ambos loops y deja intactas todas las rutas posteriores.
	// Un control sin etiqueta afectaría solo al switch o al loop interno, mientras goto ocultaría la estructura del recorrido.
}
