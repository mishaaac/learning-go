# Strings y runes en Go

Go usa `string` para datos de texto inmutables y `rune` para expresar la intención de trabajar con un code point Unicode. Sus literales, valores cero, alias y comportamiento al nivel de bytes son importantes cuando el código contiene algo más que texto ASCII sencillo.

## 1. Qué debo entender

- `string` es un tipo predeclarado cuyo valor cero es el string vacío `""`.
- Un string es una **secuencia inmutable de bytes**; una variable puede recibir otro string, pero el contenido de un string existente no puede cambiarse.
- Los strings admiten igualdad, comparaciones de orden y concatenación con `+`.
- `rune` es un alias de `int32` y es el nombre convencional para un code point Unicode.
- `byte` es un alias de `uint8`.
- Un literal `rune` usa comillas simples y su tipo predeterminado es `rune`; un literal `string` usa comillas dobles o invertidas y su tipo predeterminado es `string`.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| `string` | Una secuencia inmutable de bytes | Almacena texto y otros datos en bytes, pero no puede editarse en el mismo lugar |
| String vacío | `""`, el valor cero de `string` | Una variable `string` sin inicializador tiene un valor válido y definido |
| `rune` | Alias de `int32`, usado convencionalmente para un code point Unicode | El nombre comunica la intención de procesar caracteres |
| `byte` | Alias de `uint8` | Comunica que un valor numérico representa un byte |
| Literal `string` | Texto escrito con `"..."` o comillas invertidas | Su tipo predeterminado es `string` |
| Literal `rune` | Un code point o escape escrito con comillas simples | Su tipo predeterminado es `rune` |

Los principales operadores de strings son:

| Propósito | Operadores | Comportamiento |
|---|---|---|
| Igualdad | `==`, `!=` | Comparan los valores completos de los strings |
| Orden | `<`, `<=`, `>`, `>=` | Comparan los strings lexicográficamente, byte por byte |
| Concatenación | `+` | Producen un string nuevo a partir de dos strings |

Los operadores de orden no realizan una ordenación dependiente de la configuración regional ni del lenguaje humano.

## 3. Cómo funciona en Go

Las variables declaradas sin inicializadores reciben sus valores cero:

```go
var text string
var symbol rune
```

`text` comienza como `""`, mientras que `symbol` comienza como `0`, el valor cero numérico de `int32` y, por tanto, de su alias `rune`.

El código fuente de Go puede contener directamente texto y code points Unicode:

```go
var greeting = "Hello"
var world = "世界"
var initial rune = 'J'
```

Cuando no se proporciona otro tipo, `"世界"` tiene el tipo predeterminado `string` y `'J'` tiene el tipo predeterminado `rune`.

Los strings pueden compararse y concatenarse:

```go
const fullMessage = "Hello, " + "Go"
const same = "Go" == "Go"
const comesFirst = "apple" < "banana"
```

`fullMessage` es un string recién formado. La concatenación no modifica ninguno de los operandos.

Las variables `string` pueden reasignarse:

```go
func replaceMessage() string {
    message := "first value"
    message = "another value"
    return message
}
```

La variable cambia el valor `string` que contiene. El valor original no se modifica.

Como `rune` y `byte` son alias, sus valores pueden usarse con los tipos subyacentes sin conversión:

```go
var initial rune = 'J'
var codePoint int32 = initial

var data byte = 65
var number uint8 = data
```

Los nombres de los alias comunican la intención: `rune` sugiere un code point, mientras que `byte` sugiere datos en bytes.

El texto Unicode de un literal `string` en el código fuente se codifica como UTF-8. Aun así, indexar un string lee un byte, no un `rune`:

```go
var word = "世界"
var firstByte = word[0]
```

`firstByte` tiene el tipo `byte` (`uint8`) y es solo el primer byte de la codificación UTF-8. La conversión e iteración detalladas entre strings, bytes y runes pertenecen al tema posterior sobre procesamiento de texto.

## 4. Ejemplos, diferencias y errores comunes

| Aspecto | `string` | `rune` |
|---|---|---|
| Uso previsto | Texto o datos en bytes | Un code point Unicode |
| Ejemplo de literal | `"Go"` | `'G'` |
| Delimitadores del literal | Comillas dobles o invertidas | Comillas simples |
| Tipo predeterminado del literal | `string` | `rune` |
| Relación con otro tipo | Tipo definido predeclarado | Alias de `int32` |
| Valor cero | `""` | `0` |
| Mutabilidad | Inmutable | Un valor numérico, no un contenedor |

> **No confundir:** reasignar una variable `string` cambia el valor almacenado en la variable; no modifica el string anterior.

Esta reasignación es válida:

```go
func rename() string {
    name := "Gopher"
    name = "Go developer"
    return name
}
```

Una asignación como `name[0] = 'g'` es inválida porque los elementos de un string son de solo lectura. Crear un string diferente y asignarlo a `name` sí es válido.

> **No confundir:** `"G"` es un `string`; `'G'` es un `rune`. Las comillas dobles y simples no son intercambiables.

Usa el alias que comunique el propósito del valor:

```go
var firstInitial rune = 'J'
var lastInitial rune = 'B'
```

Usar `int32` compilaría, pero `rune` hace más clara la intención del código. Como `rune` es un alias y no un tipo de validación, un valor `int32` arbitrario no es necesariamente un code point Unicode válido.

> **No confundir:** indexar un string devuelve un byte. Un code point Unicode puede ocupar más de un byte UTF-8 y un carácter visible puede contener a veces más de un code point.

El orden de strings es útil para una comparación léxica determinista, pero no debe tratarse como un orden alfabético adaptado a un idioma.

## 5. Chuleta rápida

- `string` → secuencia inmutable de bytes.
- Valor cero de `string` → `""`.
- Comparaciones de strings → `==`, `!=`, `<`, `<=`, `>`, `>=`.
- Orden de strings → comparación lexicográfica por bytes, no ordenación regional.
- Concatenación de strings → `+`, que produce un string nuevo.
- Reasignar una variable está permitido; modificar un elemento del string no.
- `rune` = `int32`; úsalo para comunicar la intención de representar un code point Unicode.
- `byte` = `uint8`; úsalo para comunicar la intención de representar datos en bytes.
- Literal `rune` → `'G'`; tipo predeterminado `rune`.
- Literal `string` → `"Go"`; tipo predeterminado `string`.
- ⚠️ Indexar un string produce un byte, no necesariamente un `rune` completo ni un carácter visible.

### Comprueba que realmente lo sabes

1. ¿Por qué puede reasignarse una variable `string` aunque los strings sean inmutables?
2. ¿De qué tipos son alias `rune` y `byte`, y qué intención comunica cada nombre?
3. ¿Por qué `word[0]` podría no representar el primer carácter visible de un string Unicode?
