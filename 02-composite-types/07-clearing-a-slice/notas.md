# Vaciar un slice con `clear`

Para un slice, `clear` sustituye cada elemento actual por el zero value del tipo del elemento. Restablece el contenido sin cambiar la longitud ni la capacidad del slice.

## 1. Qué debo entender

- `clear(slice)` actúa sobre todos los elementos desde el índice `0` hasta `len(slice)-1`.
- Cada elemento pasa a contener el zero value de su tipo.
- El slice conserva la misma longitud y capacidad.
- `clear` no elimina elementos ni convierte el slice en `nil`.
- Limpiar un slice `nil` es un no-op seguro.

## 2. Conceptos clave

| Antes | Operación | Después |
|---|---|---|
| `[]string{"first", "second"}` | `clear(values)` | `[]string{"", ""}` |
| `[]int{4, 5}` | `clear(values)` | `[]int{0, 0}` |
| `len == 2` | `clear(values)` | `len == 2` |

`clear` se añadió como built-in en Go 1.21.

## 3. Cómo funciona en Go

```go
labels := []string{"first", "second", "third"}

clear(labels)

fmt.Printf("%q\n", labels) // ["" "" ""]
fmt.Println(len(labels))    // 3
fmt.Println(cap(labels))    // sin cambios
```

La operación equivale conceptualmente a asignar el zero value a cada elemento actual.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** limpiar los valores no es lo mismo que cambiar la longitud del slice a cero.

Para reutilizar el almacenamiento mientras el slice queda lógicamente vacío, vuelve a cortarlo:

```go
labels = labels[:0]
```

Eso cambia la longitud, pero no pone primero los elementos anteriores en su zero value. En cambio, `clear(labels)` pone en cero los elementos actuales, pero conserva la longitud.

> ⚠️ **Corrección importante:** “vaciar un slice” es ambiguo. `clear(s)`, `s = s[:0]` y `s = nil` producen longitudes, estados `nil` y comportamientos de retención de datos diferentes.

## 5. Chuleta rápida

- `clear(s)` → asigna zero values a los elementos actuales
- Los elementos de `[]string` pasan a `""`
- Los elementos de `[]int` pasan a `0`
- `len(s)` no cambia
- `cap(s)` no cambia
- El slice no se convierte en `nil`
- `clear(nilSlice)` → no-op seguro
- `s = s[:0]` → la longitud pasa a `0`, sin limpiar antes los valores
- `s = nil` → el slice se convierte en `nil`

### Comprueba que realmente lo sabes

1. ¿Qué permanece sin cambios después de aplicar `clear` a un slice?
2. ¿En qué se diferencia `clear(s)` de `s = s[:0]`?
3. ¿Qué valores quedan al limpiar un slice de structs?
