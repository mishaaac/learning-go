# Uso de `make` para slices

`make` crea un slice inicializado con una longitud elegida y una capacidad opcional. La distinción crucial es que la longitud crea elementos utilizables, mientras que la capacidad adicional solo reserva espacio para crecer posteriormente.

## 1. Qué debo entender

- `make([]T, length)` crea `length` elementos inicializados con el zero value.
- Sin un tercer argumento, la capacidad es igual a la longitud.
- `make([]T, length, capacity)` separa los elementos existentes del espacio reservado.
- Los índices solo son válidos por debajo de `len`, no por debajo de `cap`.
- Para construir mediante `append`, `make([]T, 0, capacity)` suele ser la forma buscada.

## 2. Conceptos clave

| Expresión | `len` | `cap` | Elementos iniciales |
|---|---:|---:|---|
| `make([]int, 5)` | 5 | 5 | Cinco zero values |
| `make([]int, 5, 10)` | 5 | 10 | Cinco zero values |
| `make([]int, 0, 10)` | 0 | 10 | Ninguno |

La relación obligatoria es `length <= capacity`.

## 3. Cómo funciona en Go

```go
indexed := make([]int, 3)
indexed[0] = 10

collected := make([]int, 0, 3)
collected = append(collected, 10, 20)

fmt.Println(indexed)   // [10 0 0]
fmt.Println(collected) // [10 20]
```

```text
make([]T, len, cap)
          │    └── extensión reservada
          └─────── elementos existentes
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `make([]int, 5)` no es un slice vacío con espacio para cinco valores; ya contiene cinco valores.

Este error común añade un sexto elemento:

```go
values := make([]int, 5)
values = append(values, 10) // [0 0 0 0 0 10]
```

Usa indexación directa cuando conozcas la longitud final y vayas a rellenar todas las posiciones. Usa longitud `0` más capacidad cuando los valores lleguen progresivamente mediante `append`.

Si argumentos constantes especifican `length > capacity`, la compilación falla. Si valores de runtime producen esa relación, la llamada causa un `panic`.

## 5. Chuleta rápida

- `make([]T, n)` → `len == cap == n`
- Los elementos creados contienen zero values
- `make([]T, n, c)` exige `n <= c`
- `len` → elementos existentes e indexables
- `cap` → espacio reservado para crecer
- `make([]T, 0, c)` → slice vacío, non-`nil` y preasignado
- La capacidad por sí sola no vuelve válido un índice
- `append` añade después de la longitud actual
- Longitud final conocida → asigna la longitud e indexa
- Colección progresiva → longitud `0`, capacidad estimada y después `append`

### Comprueba que realmente lo sabes

1. ¿Por qué añadir a `make([]int, 5)` produce un sexto elemento?
2. ¿Cuándo debería ser cero el segundo argumento de `make`?
3. ¿Qué determina si un índice es válido actualmente: la longitud o la capacidad?
