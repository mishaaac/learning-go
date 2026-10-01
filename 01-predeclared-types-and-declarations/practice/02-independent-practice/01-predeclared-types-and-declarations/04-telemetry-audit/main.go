package main

import "fmt"

// Agrupamos las configuraciones fijas de la sesión para expresar que no cambian.
// totalFrames conserva un tipo int concreto para hacer visible su límite con int32,
// y las máscaras usan byte porque representan bits. Las demás constantes permanecen
// sin tipo para adaptarse al contexto compatible donde se utilizan.
const (
	sessionName              = "Vega"
	totalFrames         int  = 20
	maximumRetries           = 4
	completionThreshold      = 75
	connectedMask       byte = 0b001
	encryptedMask       byte = 0b010
	cachedMask          byte = 0b100
)

func main() {
	// int32 proporciona el entero de ancho fijo requerido para los frames recibidos.
	var receivedFrames int32 = 13

	// byte comunica que el valor almacena un conjunto compacto de bits.
	var currentFlags byte = connectedMask | cachedMask

	// rune expresa que el marcador representa un único punto de código Unicode.
	var statusMarker rune = '✓'

	// La declaración sin inicializador hace visible la ausencia de una nota opcional.
	var optionalNote string

	// Las declaraciones cortas asignan a retriesUsed el tipo predeterminado int y a
	// signalSample el tipo predeterminado complex128 de una constante compleja.
	retriesUsed := 2
	signalSample := complex(3.0, -4.5)

	// Las conversiones a float64 compatibilizan los tipos, evitan la división entera
	// y hacen que completionPercentage tenga tipo float64.
	completionPercentage := float64(receivedFrames) / float64(totalFrames) * 100

	// maximumRetries adopta int al operar con retriesUsed, mientras que el umbral
	// adopta float64 al compararse con el porcentaje; la comparación produce un bool.
	remainingRetries := maximumRetries - retriesUsed
	hasReachedThreshold := completionPercentage >= completionThreshold

	// Las comparaciones de cada máscara con cero producen los tres valores bool.
	isConnected := currentFlags&connectedMask != 0
	isEncrypted := currentFlags&encryptedMask != 0
	isCached := currentFlags&cachedMask != 0

	// real e imag extraen componentes float64 del valor complex128.
	realComponent := real(signalSample)
	imaginaryComponent := imag(signalSample)

	fmt.Println("=== Telemetry Audit ===")
	fmt.Printf("Session: %s\n", sessionName)
	fmt.Printf("Status marker: %c\n", statusMarker)
	fmt.Printf("Received frames: %d\n", receivedFrames)
	fmt.Printf("Total frames: %d\n", totalFrames)
	fmt.Printf("Completion: %.0f%%\n", completionPercentage)
	fmt.Printf("Retries used: %d\n", retriesUsed)
	fmt.Printf("Retries remaining: %d\n", remainingRetries)
	fmt.Printf("Reached %d%% threshold: %t\n", completionThreshold, hasReachedThreshold)

	fmt.Println("Flags:")
	fmt.Printf("Combined value: %03b\n", currentFlags)
	fmt.Printf("Connected: %t\n", isConnected)
	fmt.Printf("Encrypted: %t\n", isEncrypted)
	fmt.Printf("Cached: %t\n", isCached)

	fmt.Println("Signal:")
	fmt.Printf("Sample: %v\n", signalSample)
	fmt.Printf("Real component: %.1f\n", realComponent)
	fmt.Printf("Imaginary component: %.1f\n", imaginaryComponent)

	fmt.Printf("Optional note: %q\n", optionalNote)
}
