package main

import "fmt"

func main() {
	grid := []string{"S.界", ".#.", "..X", "X.."}
	preTargetCount := 0

gridLoop:
	for rowNumber, row := range grid {
		for byteOffset, cell := range row {
			switch cell {
			case '#':
				continue
			case 'X':
				fmt.Printf("Target: row=%d byte-offset=%d\n", rowNumber, byteOffset)
				break gridLoop
			default:
				preTargetCount++
			}
		}
	}

	fmt.Printf("Pre-target cells: %d\n", preTargetCount)

	// El índice del range interno indica el byte inicial de cada runa dentro de su fila.
}
