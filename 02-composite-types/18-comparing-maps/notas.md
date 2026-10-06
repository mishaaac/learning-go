# Comparación de maps en Go

Los maps no pueden compararse entre sí con `==`, pero el paquete estándar `maps` permite comparar su contenido. Elige igualdad normal o una función personalizada según el tipo de value y la semántica buscada.

## 1. Qué debo entender

- `mapA == mapB` es inválido; un map solo puede compararse directamente con `nil`.
- `maps.Equal` comprueba que existan las mismas keys y values.
- El orden de iteración del map no afecta a la igualdad.
- `maps.Equal` exige values que admitan `==`.
- `maps.EqualFunc` delega la comparación de values a una función proporcionada.

## 2. Conceptos clave

| Función | Comparación de values | Caso de uso |
|---|---|---|
| `maps.Equal` | `==` | Values comparables con igualdad normal |
| `maps.EqualFunc` | Función proporcionada | Igualdad personalizada o values no comparables |
| `m == nil` | Comparación con `nil` | Comprobar si un map es `nil` |

Ambas funciones de comparación consideran las entradas, no el orden del literal o de iteración.

## 3. Cómo funciona en Go

```go
left := map[string]int{
    "hello": 5,
    "world": 10,
}
right := map[string]int{
    "world": 10,
    "hello": 5,
}

fmt.Println(maps.Equal(left, right)) // true
```

Comparación personalizada:

```go
equal := maps.EqualFunc(left, right, func(a, b int) bool {
    return a == b
})
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** el orden de un map no forma parte de su contenido, así que un orden diferente en el literal o la iteración no hace diferentes entradas iguales.

`maps.Equal` primero exige el mismo número de entradas y después values coincidentes para cada key. También considera iguales un map `nil` y uno vacío inicializado porque ninguno contiene entradas; eso no significa que su estado `nil` sea el mismo.

> ⚠️ **Corrección importante:** `maps.Equal` no está disponible para maps cuyo tipo de value no sea comparable, como `map[string][]int`. Usa `maps.EqualFunc` y define cómo deben compararse esos values.

## 5. Chuleta rápida

- `mapA == mapB` → inválido
- `m == nil` → válido
- Mismas entradas → `maps.Equal(a, b)`
- El orden no importa
- `maps.Equal` compara values con `==`
- Por tanto, los values deben ser comparables
- Igualdad personalizada → `maps.EqualFunc`
- Las keys siguen la igualdad normal de las keys de map
- Un map `nil` y uno vacío son iguales por contenido
- La igualdad de contenido es diferente del estado `nil`

### Comprueba que realmente lo sabes

1. ¿Por qué el orden de las entradas no afecta a `maps.Equal`?
2. ¿Por qué `maps.Equal` no puede comparar directamente values de `map[string][]int`?
3. ¿Cómo pueden dos maps ser iguales por contenido si solo uno es `nil`?
