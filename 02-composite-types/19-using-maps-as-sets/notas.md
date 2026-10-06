# Uso de maps como conjuntos en Go

Go no tiene un tipo de conjunto built-in, pero un map puede modelar pertenencia única almacenando los elementos del conjunto como keys. El value suele ser `bool` para una lectura cómoda o `struct{}` cuando solo importa la presencia.

## 1. Qué debo entender

- Los elementos del conjunto se convierten en keys del map, así que el tipo del elemento debe ser comparable.
- Las keys de un map son únicas, por lo que una inserción repetida no crea duplicados.
- `map[T]bool` usa `true` para pertenencia y el zero value `false` para ausencia.
- `map[T]struct{}` no almacena un value significativo y comprueba la pertenencia con `comma ok`.
- `len(set)` informa la cantidad de keys únicas.

## 2. Conceptos clave

| Representación | Inserción | Comprobación de pertenencia |
|---|---|---|
| `map[T]bool` | `set[value] = true` | `if set[value]` |
| `map[T]struct{}` | `set[value] = struct{}{}` | `_, ok := set[value]` |

Ambas formas admiten inserción, comprobación de pertenencia, eliminación con `delete` y conteo con `len`.

## 3. Cómo funciona en Go

```go
values := []int{5, 10, 2, 5, 2}
set := map[int]bool{}

for _, value := range values {
    set[value] = true
}

fmt.Println(len(values)) // 5
fmt.Println(len(set))    // 3
fmt.Println(set[5])      // true
fmt.Println(set[99])     // false
```

Con un empty struct:

```go
compact := map[int]struct{}{5: {}}
_, present := compact[5]
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** en `map[T]bool`, una key presente almacenada con `false` parece ausente en `if set[value]`. Un conjunto basado en `bool` debe almacenar siempre `true`.

`map[T]struct{}` hace explícita la intención de “solo keys”, pero necesita `comma ok` para comprobar pertenencia. `map[T]bool` suele ser más fácil de leer y puede representar estados activados/desactivados si eso es realmente lo que se busca.

> ⚠️ **Corrección importante:** `struct{}` tiene tamaño cero y `bool` tiene un tamaño distinto de cero, pero la memoria total por entrada del map es un detalle de implementación. No supongas un ahorro garantizado de un byte por entrada; elige por claridad salvo que una medición demuestre que la memoria importa.

La unión, intersección y diferencia deben implementarse con estas operaciones básicas o proporcionarse mediante otra abstracción.

## 5. Chuleta rápida

- Go no tiene un tipo de conjunto built-in
- Elemento del conjunto → key del map
- El tipo del elemento debe ser comparable
- Una inserción duplicada conserva una sola key
- `len(set)` → cantidad de elementos únicos
- Forma sencilla → `map[T]bool`
- Insertar en forma bool → `set[x] = true`
- Comprobar forma bool → `if set[x]`
- Forma de solo keys → `map[T]struct{}`
- Comprobar forma struct → `_, ok := set[x]`
- Eliminar miembro → `delete(set, x)`

### Comprueba que realmente lo sabes

1. ¿Por qué los valores duplicados del origen producen un solo miembro del set?
2. ¿Qué condición debe mantener un conjunto `map[T]bool` para que `if set[x]` signifique pertenencia?
3. ¿Qué diferencia práctica existe entre `map[T]bool` y `map[T]struct{}`?
