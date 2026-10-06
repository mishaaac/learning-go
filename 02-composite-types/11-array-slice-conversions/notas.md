# Conversión entre arrays y slices

Los arrays y slices pueden exponer o reproducir la misma secuencia de formas diferentes. La pregunta importante es si una operación comparte el almacenamiento original o copia sus elementos.

## 1. Qué debo entender

- `array[:]` crea una vista slice y comparte el almacenamiento del array.
- `[N]T(slice)` crea un valor array independiente copiando elementos.
- `(*[N]T)(slice)` crea un pointer a array que comparte almacenamiento con el slice.
- En ambas formas de slice a array, `N` no puede superar `len(slice)`.
- La longitud del array de destino debe ser explícita; `[...]T(slice)` no es una conversión válida.

## 2. Conceptos clave

| Operación | ¿Copia? | ¿Comparte almacenamiento? |
|---|---:|---:|
| `array[:]` | No | Sí |
| `array[1:3]` | No | Sí |
| `[N]T(slice)` | Sí | No |
| `(*[N]T)(slice)` | No | Sí |

## 3. Cómo funciona en Go

```go
array := [4]int{1, 2, 3, 4}
view := array[:]
view[0] = 10
fmt.Println(array) // [10 2 3 4]

slice := []int{5, 6, 7, 8}
copyArray := [4]int(slice)
sharedArray := (*[4]int)(slice)

slice[0] = 50
fmt.Println(copyArray)    // [5 6 7 8]
fmt.Println(sharedArray)  // &[50 6 7 8]
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `[4]int(slice)` copia, mientras que `(*[4]int)(slice)` comparte el almacenamiento subyacente del slice.

Una conversión a un array más corto solo copia el prefijo necesario:

```go
firstTwo := [2]int(slice)
```

Si `len(slice) < N`, convertir a `[N]T` o `*[N]T` causa un `panic` en runtime. La capacidad adicional no ayuda; la regla usa la longitud.

Si una API convierte repetidamente arrays de distintas longitudes solo para aceptarlos, un parámetro slice suele ser un diseño más claro porque la longitud de un array forma parte de su tipo.

## 5. Chuleta rápida

- Array a slice → `array[:]`
- El slicing de un array comparte almacenamiento
- Slice a valor array → `[N]T(slice)`
- La conversión a valor array copia
- Slice a pointer a array → `(*[N]T)(slice)`
- La conversión a pointer a array comparte almacenamiento
- Regla obligatoria → `N <= len(slice)`
- `cap(slice)` no vuelve válida una conversión demasiado grande
- Los arrays de destino más pequeños copian el prefijo
- `[...]T(slice)` no es una conversión válida

### Comprueba que realmente lo sabes

1. ¿Qué conversión debes usar cuando el resultado tiene que ser independiente del slice?
2. ¿Por qué un slice con `len == 2` y `cap == 8` no puede convertirse a `[4]int`?
3. ¿Cómo puede la modificación de `array[:]` afectar al array original?
