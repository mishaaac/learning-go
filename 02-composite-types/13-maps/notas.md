# Maps en Go

Un map almacena asociaciones entre keys únicas y valores. Es la opción natural cuando los datos se recuperan mediante una key significativa en vez de una posición entera secuencial.

## 1. Qué debo entender

- Un tipo map se escribe `map[K]V`.
- Las keys deben ser comparables; los values pueden tener cualquier tipo.
- El zero value es `nil`: las lecturas son seguras, pero las escrituras causan `panic`.
- Un literal vacío o `make` crea un map inicializado en el que se puede escribir.
- Los maps crecen al añadir entradas y no garantizan el orden de iteración.

## 2. Conceptos clave

| Forma | Estado | ¿Admite escritura? |
|---|---|---:|
| `var counts map[string]int` | `nil`, longitud `0` | No |
| `map[string]int{}` | Inicializado, longitud `0` | Sí |
| `map[string]int{"go": 1}` | Inicializado con datos | Sí |
| `make(map[string]int, 100)` | Vacío con una estimación inicial de tamaño | Sí |

La estimación pasada a `make` no es la longitud del map ni un tamaño máximo.

## 3. Cómo funciona en Go

```go
teams := map[string][]string{
    "Orcas": {"Fred", "Ralph"},
    "Lions": {"Sarah"},
}

teams["Kittens"] = []string{"Waldo"}
fmt.Println(len(teams)) // 3

var missing map[string]int
fmt.Println(missing["key"]) // 0
```

La búsqueda en un map utiliza una key en vez de una posición secuencial:

```text
key ──búsqueda──> value
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** leer `nilMap[key]` es válido, pero asignar `nilMap[key] = value` causa un `panic` en runtime.

Los slices y maps son flexibles y tienen zero values `nil`, pero modelan relaciones diferentes: un slice usa posiciones enteras ordenadas, mientras que un map usa keys comparables únicas y no tiene un orden de iteración definido.

> ⚠️ **Corrección importante:** no todo valor de Go puede ser una key de map. Los slices, maps y funciones no son comparables y, por tanto, no pueden ser tipos de key; los arrays y structs solo son válidos cuando todos sus componentes son comparables.

Dos maps no pueden compararse directamente con `==`, salvo que un map sí puede compararse con `nil`.

## 5. Chuleta rápida

- Tipo → `map[K]V`
- Tipo de key → debe ser comparable
- Tipo de value → puede ser cualquier tipo
- Zero value → `nil`
- Leer de un map `nil` → zero value del tipo del value
- Escribir en un map `nil` → `panic`
- Map vacío modificable → `map[K]V{}` o `make(map[K]V)`
- `make(map[K]V, n)` → estimación inicial, no longitud ni límite
- `len(m)` → número de entradas
- El orden de iteración no está especificado
- `map == map` es inválido; `map == nil` es válido

### Comprueba que realmente lo sabes

1. ¿Por qué un struct puede ser a veces una key de map mientras que un slice no puede serlo?
2. ¿Cuál es la diferencia crítica de comportamiento entre un map `nil` y uno vacío inicializado?
3. ¿Cuándo debería preferirse un map en lugar de un slice?
