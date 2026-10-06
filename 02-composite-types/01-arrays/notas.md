# Arrays en Go

Un array es una secuencia de longitud fija cuyos valores tienen un solo tipo. Los arrays importan porque su longitud forma parte de su tipo y porque proporcionan el almacenamiento subyacente usado por los slices.

## 1. Qué debo entender

- Un tipo array tiene la forma `[N]T`; tanto `N` como `T` definen el tipo.
- Cada elemento comienza con el zero value de su tipo, salvo que un literal proporcione otro valor.
- Los índices de un array comienzan en `0` y terminan en `len(array)-1`.
- Los arrays solo se pueden comparar directamente cuando el tipo de sus elementos es comparable.
- Los arrays son apropiados cuando se conoce de antemano un tamaño exacto y significativo; los slices son más flexibles para la mayoría de las secuencias.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| `[3]int` | Array de tres elementos `int` | Es un tipo diferente de `[4]int` |
| `[...]int{1, 2}` | Longitud inferida por el compilador | El tipo resultante es `[2]int` |
| Literal con índices | Valores asignados en índices concretos | Los elementos omitidos conservan su zero value |
| `[2][3]int` | Array cuyos elementos son arrays `[3]int` | Go compone arrays en vez de tener un tipo matriz separado |
| `len(array)` | Longitud fija del array | En un array, `len` está determinada por su tipo |

## 3. Cómo funciona en Go

```go
var counts [3]int                 // [0 0 0]
scores := [...]int{10, 20, 30}    // tipo [3]int
sparse := [6]int{0: 1, 4: 9}      // [1 0 0 0 9 0]

counts[1] = 7
fmt.Println(counts[1], len(counts)) // 7 3
fmt.Println(scores == [3]int{10, 20, 30}) // true
```

Un array anidado es un array de arrays:

```go
var grid [2][3]int
grid[1][2] = 5
```

```text
[2][3]int
├── [3]int
└── [3]int
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `[...]int{1, 2, 3}` es un array, mientras que `[]int{1, 2, 3}` es un slice.

Un `[3]int` no se puede asignar a un `[4]int`, y una variable normal no puede determinar la longitud de un array en runtime. Un índice constante fuera de rango se rechaza durante la compilación; un índice inválido calculado dinámicamente causa un `panic` en runtime.

> ⚠️ **Corrección importante:** los arrays no siempre son comparables. `[3]int` admite `==`, pero un array como `[2][]int` no, porque los slices no son comparables.

Los arrays son útiles cuando el tamaño forma parte del significado del dato, como en un digest de tamaño fijo. Su función indirecta más habitual es servir como backing storage de los slices.

## 5. Chuleta rápida

- `[N]T` → array de `N` valores de tipo `T`
- La longitud forma parte del tipo del array
- `var a [3]int` → `[0 0 0]`
- `[...]T{...}` → el compilador infiere la longitud
- Los literals con índices dejan las posiciones omitidas en su zero value
- Índices válidos → desde `0` hasta `len(a)-1`
- `[2][3]int` → array de dos arrays `[3]int`
- `==` y `!=` requieren elementos comparables
- `[3]int` y `[4]int` son tipos diferentes
- Los arrays suelen proporcionar el backing storage de los slices

### Comprueba que realmente lo sabes

1. ¿Por qué una función que recibe `[3]int` no puede recibir también `[4]int`?
2. ¿Qué determina si dos arrays pueden compararse con `==`?
3. ¿Qué diferencia hay entre un índice constante inválido y un índice inválido calculado en runtime?
