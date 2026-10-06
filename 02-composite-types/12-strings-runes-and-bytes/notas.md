# Strings, runes y bytes en Go

Un string de Go es una secuencia inmutable de bytes, mientras que un `rune` representa un Unicode code point. Esta distinción explica el comportamiento de la indexación, `len`, el slicing y las conversiones con texto UTF-8.

## 1. Qué debo entender

- Un string almacena bytes y no está obligado a contener UTF-8 válido.
- El código fuente de Go usa UTF-8, por lo que los string literals normales suelen contener texto codificado en UTF-8.
- `s[i]`, `s[a:b]` y `len(s)` operan sobre bytes.
- Un `rune` es un alias de `int32` y representa un Unicode code point.
- UTF-8 usa entre uno y cuatro bytes por code point, así que los límites de bytes no siempre son límites de texto.

## 2. Conceptos clave

| Forma | Representa |
|---|---|
| `byte` | Alias de `uint8`; un byte |
| `rune` | Alias de `int32`; un Unicode code point |
| `s[i]` | Byte en el índice `i` |
| `len(s)` | Número de bytes |
| `[]byte(s)` | Copia de los bytes del string |
| `[]rune(s)` | Unicode code points decodificados |

## 3. Cómo funciona en Go

```go
text := "Hello 🌞"

fmt.Println(len(text))   // 10 bytes
fmt.Println(text[0])     // 72, el byte de 'H'
fmt.Println([]byte(text))
fmt.Println([]rune(text))
```

```text
"Hello " → 6 bytes
"🌞"      → 4 bytes
total     → 10 bytes, 7 code points
```

Usa `range` para decodificar code points UTF-8:

```go
for _, r := range text {
    fmt.Printf("%c ", r)
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un índice de byte no es necesariamente un índice de carácter o code point.

Hacer slicing en medio de una codificación UTF-8 multibyte puede producir un string con UTF-8 inválido. Los strings siguen siendo inmutables, pero el slicing basado en bytes necesita límites válidos cuando el resultado debe representar texto.

> ⚠️ **Corrección importante:** `string(65)` convierte el entero como Unicode code point U+0041 y produce `"A"`; no formatea los dígitos decimales como `"65"`. El formateo numérico requiere una función de formato o conversión apropiada.

Convertir un `string` a `[]byte` conserva los bytes sin procesar; convertirlo a `[]rune` decodifica UTF-8 y sustituye las codificaciones inválidas por `utf8.RuneError`.

## 5. Chuleta rápida

- String → secuencia inmutable de bytes
- Un string puede contener UTF-8 inválido
- `byte` → alias de `uint8`
- `rune` → alias de `int32`, Unicode code point
- `s[i]` → un byte
- `s[a:b]` → rango de bytes
- `len(s)` → número de bytes
- UTF-8 → entre uno y cuatro bytes por code point
- `[]byte(s)` → bytes sin procesar
- `[]rune(s)` → code points decodificados
- `string(65)` → `"A"`, no `"65"`

### Comprueba que realmente lo sabes

1. ¿Por qué hacer slicing sobre un string UTF-8 válido puede producir UTF-8 inválido?
2. ¿Qué información diferente muestran `[]byte(text)` y `[]rune(text)`?
3. ¿Por qué `string(65)` no produce los dígitos `"65"`?
