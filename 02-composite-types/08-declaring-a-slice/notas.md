# Cómo declarar un slice en Go

Go ofrece varias declaraciones válidas de slices y cada una representa un estado inicial diferente. La elección entre `var`, un literal y `make` depende de si los valores ya existen y de cómo se rellenará el slice.

## 1. Qué debo entender

- Usa `var s []T` cuando el slice pueda quedar sin utilizar; comienza como `nil`.
- Usa un slice literal cuando ya conozcas los valores iniciales.
- Usa `make([]T, n)` cuando deban existir inmediatamente `n` elementos.
- Usa `make([]T, 0, n)` al comenzar vacío y reunir valores mediante `append`.
- Una estimación de capacidad reduce el crecimiento, pero no limita el tamaño final.

## 2. Conceptos clave

| Situación | Declaración | Estado inicial |
|---|---|---|
| Puede quedar vacío | `var data []int` | `nil`, longitud `0` |
| Debe estar vacío pero ser non-`nil` | `data := []int{}` | non-`nil`, longitud `0` |
| Valores conocidos | `data := []int{2, 4}` | longitud `2` |
| Longitud exacta e indexable | `make([]int, n)` | `n` elementos con zero values |
| `append` con tamaño estimado | `make([]int, 0, n)` | vacío con capacidad reservada |

## 3. Cómo funciona en Go

```go
var optional []int
fixedValues := []int{2, 4, 6}

transformed := make([]int, len(fixedValues))
for i, value := range fixedValues {
    transformed[i] = value * 2
}

collected := make([]int, 0, 10)
collected = append(collected, 7)
```

El primer `make` permite escrituras por índice; el segundo permite `append` progresivos sin crear elementos de relleno.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `[]T{}` y `var s []T` tienen longitud cero, pero solo el segundo es `nil`.

Esta declaración ya contiene cinco elementos:

```go
data := make([]int, 5)
```

`append` no rellena esos elementos; añade después de ellos. Si la longitud final es solo una estimación, prefiere `make([]T, 0, estimate)` y `append` para que no queden zero values sin utilizar en el resultado.

La distinción entre `nil` y vacío puede importar en límites de serialización o APIs, pero las operaciones normales de slices suelen funcionar con ambos.

## 5. Chuleta rápida

- Posiblemente sin uso → `var s []T`
- Valores conocidos → `s := []T{...}`
- Se necesita vacío non-`nil` → `s := []T{}`
- Longitud exacta y escrituras por índice → `make([]T, n)`
- `append` progresivos → `make([]T, 0, n)`
- `make([]T, n)` crea `n` elementos reales
- La capacidad es una estimación, no un máximo
- `append` siempre crece desde la longitud actual
- Prefiere la declaración que coincida con la forma de escribir los valores

### Comprueba que realmente lo sabes

1. ¿Por qué `make([]T, n)` es apropiado para la salida de una transformación por índices?
2. ¿Por qué `make([]T, 0, n)` es más seguro cuando la cantidad final solo es estimada?
3. ¿En qué situación puede importar la diferencia entre `nil` y vacío non-`nil`?
