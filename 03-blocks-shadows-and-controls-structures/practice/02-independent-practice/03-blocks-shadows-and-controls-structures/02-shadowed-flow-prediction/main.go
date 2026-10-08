package main

import "fmt"

func main() {
	status := "ready"
	total := 0

	for index := 1; index <= 4; index++ {
		if status := index % 2; status == 0 {
			total += index
			continue
		}

		switch index {
		case 3:
			status := "paused"
			fmt.Println(status, total)
			break
		default:
			total += index
		}

		fmt.Println(status, total)
	}

	fmt.Println(status, total)

	// Predicción completa:
	// ready 1
	// paused 3
	// ready 3
	// ready 7
	// El status del if es un int limitado a ese if; el status del case 3 es un string limitado a ese case.
	// Las demás impresiones observan el status externo, que conserva el valor ready.
	// continue avanza el for y omite las impresiones de los índices 2 y 4.
	// break termina solamente el switch, por eso la impresión posterior del índice 3 todavía se ejecuta.
}
