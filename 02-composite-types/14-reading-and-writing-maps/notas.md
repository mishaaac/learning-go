# Lectura y escritura de maps en Go

La expresión `m[key]` accede a una entrada de un map, mientras que `m[key] = value` crea o reemplaza una. Una key ausente se lee como el zero value del tipo del value, lo que resulta práctico pero a veces ambiguo.

## 1. Qué debo entender

- Usa `m[key]` para leer un valor.
- Usa `m[key] = value` para insertar o reemplazar una entrada.
- Una key ausente devuelve el zero value del tipo del value.
- `:=` no puede asignar directamente a un índice de map porque la expresión de índice no es una variable nueva.
- Los values numéricos de un map pueden usar operaciones como `m[key]++`.

## 2. Conceptos clave

| Operación | Significado |
|---|---|
| `value := m[key]` | Lee una entrada |
| `m[key] = value` | Inserta o reemplaza una entrada |
| `m[key]++` | Incrementa un valor numérico, comenzando en cero si no existe |
| `len(m)` | Cuenta las entradas actuales |
| Key ausente | Devuelve el zero value del tipo del value |

## 3. Cómo funciona en Go

```go
totalWins := map[string]int{}

totalWins["Orcas"] = 1
totalWins["Lions"] = 2

fmt.Println(totalWins["Orcas"])   // 1
fmt.Println(totalWins["Kittens"]) // 0

totalWins["Kittens"]++
totalWins["Lions"] = 3

fmt.Println(totalWins["Kittens"]) // 1
fmt.Println(totalWins["Lions"])   // 3
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** leer una key ausente no inserta esa key. Solo produce el zero value.

Esto es inválido:

```go
totalWins["Orcas"] := 1 // error de compilación
```

Usa `=` porque un índice de map es un destino de asignación, no un identificador que se está declarando.

El comportamiento del zero value hace concisos los maps usados como contadores: un `int` ausente se comporta como `0`, por lo que `counts[word]++` funciona sin una rama de inicialización separada. Si un cero almacenado y una key ausente significan cosas distintas, usa la forma `comma ok` en vez de una lectura de un solo valor.

## 5. Chuleta rápida

- Leer → `m[key]`
- Insertar → `m[key] = value`
- Reemplazar → la misma sintaxis de asignación
- Key ausente → zero value
- Una lectura ausente no añade una entrada
- `m[key] := value` → inválido
- Contador numérico → `m[key]++`
- Una key numérica ausente comienza conceptualmente en `0`
- Necesitas información de presencia → usa `value, ok := m[key]`

### Comprueba que realmente lo sabes

1. ¿Por qué `counts[word]++` funciona cuando `word` todavía no es una key?
2. ¿Leer una key ausente cambia `len(m)`? ¿Por qué?
3. ¿Por qué `:=` es inválido en el lado izquierdo de una asignación a un índice de map?
