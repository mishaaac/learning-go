# `len` en Go

`len` es una función built-in que informa la longitud de varios tipos de Go. Su significado exacto depende del operando, por lo que entender qué cuenta evita errores de indexación y procesamiento de texto.

## 1. Qué debo entender

- Para arrays y slices, `len` es el número de elementos.
- Para strings, `len` cuenta bytes, no Unicode code points.
- Para maps, `len` cuenta entradas.
- Para channels, `len` cuenta los elementos que están actualmente en el buffer.
- Un slice o map `nil` tiene longitud `0`.

## 2. Conceptos clave

| Operando | Significado de `len(value)` |
|---|---|
| Array | Número de elementos del array |
| Slice | Número actual de elementos del slice |
| String | Número de bytes |
| Map | Número de entradas key-value |
| Channel | Número de elementos actualmente en el buffer |

`len` tiene reglas de tipos definidas por el lenguaje. Una función normal no puede reproducir exactamente la misma operación única sobre todas estas categorías de tipos no relacionados.

## 3. Cómo funciona en Go

```go
array := [3]int{10, 20, 30}
slice := []int{10, 20}
var nilSlice []int
counts := map[string]int{"go": 2}

fmt.Println(len(array))    // 3
fmt.Println(len(slice))    // 2
fmt.Println(len(nilSlice)) // 0
fmt.Println(len(counts))   // 1
fmt.Println(len("🌞"))      // 4 bytes UTF-8
```

Para un slice no vacío, los índices válidos van desde `0` hasta `len(slice)-1`.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `len(slice) == 0` no demuestra que el slice sea `nil`; un slice vacío non-`nil` también tiene longitud `0`.

Llamar a `len` sobre un valor `nil` admitido es seguro:

```go
var values []int
if len(values) == 0 {
    fmt.Println("no values")
}
```

No uses `len(text)` como contador de caracteres para texto UTF-8 arbitrario. Además, la longitud de un channel solo es una observación momentánea del buffer y normalmente no es una base fiable para coordinar trabajo concurrente.

## 5. Chuleta rápida

- `len(array)` → número de elementos
- `len(slice)` → número actual de elementos
- `len(string)` → bytes
- `len(map)` → entradas
- `len(channel)` → elementos actualmente en el buffer
- `len(nilSlice)` → `0`
- `len(nilMap)` → `0`
- Último índice válido de un slice → `len(s)-1`, solo si no está vacío
- `len(s) == 0` no implica `s == nil`
- Tipo de operando no admitido → error de compilación

### Comprueba que realmente lo sabes

1. ¿Por qué `len("🌞")` puede ser diferente del número de símbolos visibles?
2. ¿Qué se puede y qué no se puede concluir de `len(s) == 0`?
3. ¿Por qué `len(channel)` es un mecanismo deficiente de sincronización?
