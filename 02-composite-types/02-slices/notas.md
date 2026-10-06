# Slices en Go

Un slice es una vista flexible sobre una parte contigua de un array subyacente. Es la opción habitual en Go para secuencias de longitud variable porque la longitud no forma parte del tipo del slice.

## 1. Qué debo entender

- `[]T` es un tipo slice; su longitud actual no forma parte de ese tipo.
- Un slice describe un almacenamiento subyacente mediante una longitud y una capacidad.
- El zero value de un slice es `nil`, con longitud y capacidad `0`.
- La indexación está limitada por `len`, incluso cuando existe más capacidad.
- Los slices no se pueden comparar entre sí con `==`; solo pueden compararse directamente con `nil`.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| `[]int{1, 2, 3}` | Slice literal | Crea un slice listo para usar con tres elementos |
| `var values []int` | Slice `nil` | Se puede usar con `len`, `cap`, `append` y `range` |
| `[][]int` | Slice de slices | Los slices internos pueden tener longitudes diferentes |
| `slices.Equal` | Igualdad elemento por elemento | Se usa en vez de `==` con elementos comparables |
| `slices.EqualFunc` | Igualdad mediante una función | Permite una lógica de comparación personalizada |

## 3. Cómo funciona en Go

```go
values := []int{10, 20, 30}
values[1] = 25

var empty []int
fmt.Println(values, len(values)) // [10 25 30] 3
fmt.Println(empty == nil)        // true
```

Los elementos indexados de un literal pueden dejar espacios con el zero value:

```go
sparse := []int{0: 1, 3: 8} // [1 0 0 8]
```

Para comparar el contenido:

```go
fmt.Println(slices.Equal([]int{1, 2}, []int{1, 2})) // true
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un slice puede crecer mediante `append`, pero escribir directamente en `values[len(values)]` sigue estando fuera de rango.

`nil` y `[]int{}` tienen longitud `0`, pero solo el primero es igual a `nil`. La mayoría de las operaciones los tratan de forma similar, así que distínguelos únicamente cuando la representación importe.

> ⚠️ **Corrección importante:** los slices no son directamente comparables con otros slices. `slices.Equal` requiere valores de elementos comparables; usa `slices.EqualFunc` cuando la igualdad necesite lógica personalizada. Comparar slices con tipos de elemento incompatibles también causa un error de compilación, salvo que una llamada adecuada a `EqualFunc` admita ambos tipos.

`reflect.DeepEqual` tiene una semántica más amplia, que incluye considerar diferentes un slice `nil` y uno vacío non-`nil`. Prefiere las funciones específicas de `slices` cuando su semántica coincida con la tarea.

## 5. Chuleta rápida

- `[]T` → slice de `T`
- La longitud no forma parte del tipo del slice
- `var s []T` → `nil`, `len == 0`, `cap == 0`
- `[]T{}` → vacío pero non-`nil`
- Los índices válidos dependen de `len`, no de `cap`
- `[][]T` → slice de slices
- `slice == slice` → error de compilación
- `slice == nil` → válido
- `slices.Equal` → igualdad normal de elementos
- `slices.EqualFunc` → igualdad personalizada

### Comprueba que realmente lo sabes

1. ¿Por qué dos slices de longitudes diferentes pueden tener el mismo tipo?
2. ¿Por qué una capacidad adicional no convierte `s[len(s)]` en un índice válido?
3. ¿Cuándo sería más apropiado `slices.EqualFunc` que `slices.Equal`?
