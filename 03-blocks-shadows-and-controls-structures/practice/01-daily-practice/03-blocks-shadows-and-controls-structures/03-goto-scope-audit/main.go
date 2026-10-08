package main

import "fmt"

func main() {
	value := -2
	if value < 0 {
		goto done
	}

	fmt.Println("processed", value)

done:
	fmt.Println("finished")

	// Predicción del programa A: compila e imprime "finished".
	// El salto permanece dentro del bloque de main y no omite ninguna declaración que esté en alcance en done.
	// Un if con ramas directas expresaría esta decisión con mayor claridad que goto.
	// Predicción del programa B: no compila porque el salto omite la declaración de message,
	// aunque message estaría en alcance al llegar a la etiqueta done.
	// Predicción del programa C: no compila porque goto intenta entrar desde fuera al bloque interno del if.
}
