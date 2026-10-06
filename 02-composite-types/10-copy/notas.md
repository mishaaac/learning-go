# `copy` en Go

El built-in `copy` transfiere elementos a posiciones que ya existen en un slice de destino. Es útil para crear copias independientes, copiar rangos concretos y copiar con seguridad entre regiones solapadas.

## 1. Qué debo entender

- La forma es `copy(destination, source)`.
- Copia `min(len(destination), len(source))` elementos.
- La capacidad no aumenta la cantidad copiada; solo importan las longitudes actuales.
- `copy` no hace crecer el destino.
- Se permiten regiones de origen y destino solapadas.

## 2. Conceptos clave

| Operación | Resultado |
|---|---|
| `n := copy(dst, src)` | Copia elementos y devuelve su cantidad |
| `copy(dst, src[2:])` | Copia un rango seleccionado del origen |
| `copy(s[:3], s[1:])` | Copia con seguridad regiones solapadas |
| `copy(dst, array[:])` | Usa un array mediante una vista slice |

## 3. Cómo funciona en Go

```go
source := []int{1, 2, 3, 4}
destination := make([]int, len(source))

count := copy(destination, source)
destination[0] = 99

fmt.Println(count)       // 4
fmt.Println(source)      // [1 2 3 4]
fmt.Println(destination) // [99 2 3 4]
```

Para un desplazamiento con solapamiento:

```go
values := []int{1, 2, 3, 4}
copy(values[:3], values[1:])
fmt.Println(values) // [2 3 4 4]
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `copy` sustituye elementos existentes del destino; `append` añade elementos y aumenta la longitud.

Este destino tiene capacidad pero no longitud, así que no se copia nada:

```go
destination := make([]int, 0, 4)
count := copy(destination, []int{1, 2}) // count == 0
```

Crea suficiente longitud en el destino cuando quieras una copia independiente. El número devuelto puede ignorarse cuando las longitudes ya hacen evidente el resultado.

Aplicar slicing a un array mediante `array[:]` permite usarlo como origen o destino porque `copy` trabaja con slices.

## 5. Chuleta rápida

- Sintaxis → `copy(dst, src)`
- Primer argumento → destino
- Segundo argumento → origen
- Valor devuelto → elementos copiados
- Cantidad → `min(len(dst), len(src))`
- La capacidad no determina la cantidad
- `copy` no hace crecer `dst`
- El slicing selecciona el rango copiado
- Se admite el solapamiento
- `array[:]` expone un array como slice

### Comprueba que realmente lo sabes

1. ¿Por qué copiar hacia `make([]int, 0, 10)` copia cero elementos?
2. ¿Cómo crearías una copia completamente independiente de un slice?
3. ¿Por qué el `copy` solapado es útil para desplazar elementos?
