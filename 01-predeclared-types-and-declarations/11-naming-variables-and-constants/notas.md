# Nombres de variables y constantes en Go

Go separa las reglas del lenguaje que hacen válido un identificador de las convenciones que vuelven un nombre idiomático y legible. Los buenos nombres comunican el propósito de un valor, encajan con su alcance y usan las mayúsculas de forma coherente con las reglas de exportación de Go.

## 1. Qué debo entender

- Un identificador comienza con una letra Unicode o `_`; los caracteres posteriores también pueden ser dígitos Unicode.
- Las palabras reservadas como `var`, `func` y `type` no pueden utilizarse como identificadores.
- Los identificadores Unicode son legales, pero los puntos de código visualmente similares pueden crear nombres diferentes difíciles de distinguir.
- `_` por sí solo es el **identificador vacío** especial; nombres como `_0` y `__` son identificadores normales.
- Los nombres idiomáticos de varias palabras usan `mixedCaps` o `MixedCaps`, no `snake_case`; las constantes siguen la misma convención.
- La mayúscula inicial comunica visibilidad de exportación en los contextos definidos por Go, no que un nombre represente una constante.
- Los nombres cortos encajan en alcances pequeños y evidentes; los alcances mayores suelen necesitar nombres más descriptivos sin repetir innecesariamente el tipo del valor.

> ⚠️ **Corrección importante:** Una primera letra mayúscula no convierte cualquier identificador en exportado. El nombre también debe estar declarado en el bloque del paquete o ser el nombre de un campo o método; un identificador local sigue siendo local aunque comience con mayúscula.

## 2. Conceptos clave

| Concepto | Regla del lenguaje o convención | Por qué importa |
|---|---|---|
| Sintaxis de identificadores | Primer carácter: letra Unicode o `_`; caracteres restantes: letras Unicode, dígitos Unicode o `_` | Determina si el compilador acepta el token como identificador |
| Palabra reservada | Una palabra como `const`, `func` o `range` que no puede ser un identificador | Tener caracteres válidos no permite reutilizar una palabra reservada como nombre |
| Identificador vacío | `_` descarta un valor y no crea una vinculación | Es diferente de los nombres normales que simplemente contienen `_` |
| Identificador exportado | Un nombre que cumple los requisitos y comienza con una letra Unicode mayúscula | La capitalización controla el acceso desde otros paquetes en los contextos definidos |
| Nombre idiomático | Un nombre claro que sigue las convenciones de Go | El código válido todavía puede resultar innecesariamente confuso |
| Nombre adaptado al alcance | Un nombre cuyo detalle corresponde a la distancia durante la que debe entenderse | Los alcances pequeños necesitan menos contexto que las declaraciones a nivel de paquete |

Las elecciones habituales son:

| Situación | Estilo de nombre habitual |
|---|---|
| Alcance local muy pequeño | Nombre corto y claro |
| Índice de bucle | `i`, `j` |
| Clave y valor en `range` | `k`, `v` cuando su significado es evidente |
| Nombre no exportado de varias palabras | `indexCounter` |
| Nombre exportado de varias palabras | `IndexCounter` |
| Variable o constante a nivel de paquete | Nombre más descriptivo cuando su alcance mayor lo requiere |
| Constante | La misma convención `mixedCaps` o `MixedCaps`; no automáticamente `UPPER_SNAKE_CASE` |

El nombre debería explicar qué representa el valor. En la mayoría de los contextos, el sistema de tipos de Go ya comunica si ese valor es un `int`, un `string` u otro tipo.

## 3. Cómo funciona en Go

Los identificadores pueden usar letras Unicode, dígitos después del primer carácter y guiones bajos:

```go
func identifierValues() (int, int, string, string) {
    _0 := 0
    π := 3
    ａ := "fullwidth"
    a := "ASCII"

    return _0, π, ａ, a
}
```

Los cuatro nombres son válidos. La `ａ` de ancho completo y la `a` ASCII son puntos de código Unicode diferentes y, por tanto, identificadores distintos, aunque tengan una apariencia similar.

El identificador `_` se comporta de forma diferente a `_0` o `__`:

```go
func firstValue() int {
    first, _ := twoValues()
    return first
}

func twoValues() (int, int) {
    return 10, 20
}
```

`_` descarta el segundo resultado y no introduce una variable. En cambio, un nombre como `_0` crearía una variable normal.

Los identificadores de varias palabras usan mayúsculas en lugar de separadores:

```go
const maxRetries = 3
var packageRequestCount int
```

Los nombres no están exportados porque comienzan con letras minúsculas. Si una declaración a nivel de paquete forma intencionadamente parte de la API pública del paquete, su nombre comienza con una letra mayúscula:

```go
const DefaultRetries = 3
```

La capitalización describe la visibilidad de exportación; no distingue las constantes de las variables.

Los nombres locales cortos funcionan cuando su significado sigue siendo evidente:

```go
func sum(values []int) int {
    total := 0
    for _, v := range values {
        total += v
    }
    return total
}
```

`v` es legible porque su alcance se limita al pequeño cuerpo del bucle. `total` es un poco más descriptivo porque se utiliza en toda la función.

```text
alcance pequeño y evidente
        └── un nombre corto puede bastar

alcance mayor o menos evidente
        └── añade el contexto necesario para explicar el propósito del valor
```

Los nombres de receivers también suelen ser cortos y estar relacionados con el tipo del receiver. Su uso detallado pertenece al tema posterior sobre métodos.

## 4. Ejemplos, diferencias y errores comunes

| Nombre | ¿Válido? | ¿Suele ser idiomático aquí? | Motivo |
|---|---:|---:|---|
| `_0` | ✅ | ❌ | Es legal, pero comunica poco y se parece al identificador vacío |
| `π` | ✅ | ❌ | Es Unicode válido, pero puede ser más difícil de escribir o buscar de forma coherente |
| `ａ` | ✅ | ❌ | Puede confundirse visualmente con la `a` ASCII |
| `index_counter` | ✅ | ❌ | Go normalmente prefiere `indexCounter` |
| `indexCounter` | ✅ | ✅ | Nombre idiomático no exportado de varias palabras |
| `INDEX_COUNTER` | ✅ | ❌ | Las constantes no requieren `UPPER_SNAKE_CASE` en Go |
| `maxRetries` | ✅ | ✅ | Nombre idiomático no exportado para una constante o variable |
| `MaxRetries` | ✅ | ✅ cuando se desea exportar | La mayúscula inicial comunica visibilidad de exportación en una declaración que cumple los requisitos |
| `i` | ✅ | ✅ para un índice evidente | Su alcance pequeño proporciona el contexto que falta |
| `k`, `v` | ✅ | ✅ para roles evidentes de clave y valor | Su significado convencional está claro en un bucle `range` pequeño |

> **No confundir:** `_` es el identificador vacío; `_0` y `__` son identificadores normales que introducen vinculaciones.

> **No confundir:** la capitalización no marca constantes. Tanto `var` como `const` pueden usar nombres con minúscula o mayúscula; en las declaraciones que cumplen los requisitos, la capitalización determina en cambio la visibilidad de exportación.

Las palabras reservadas no son identificadores aunque estén formadas por letras válidas. Por ejemplo, `var`, `range` y `type` no pueden elegirse como nombres de variables o constantes.

Evita codificar el tipo en nombres como `userString` o `countInt` cuando el tipo no aporte un significado útil. Prefiere un nombre que describa el rol, como `user` o `count`; incluye palabras relacionadas con el tipo únicamente cuando distingan el concepto en sí.

Los nombres cortos no son mejores automáticamente. Si varios nombres breves resultan difíciles de seguir, el bloque podría necesitar nombres más claros o estar realizando demasiado trabajo. Por el contrario, un nombre muy largo dentro de un bucle de dos líneas puede repetir contexto que el lector ya tiene.

El Unicode poco habitual puede ser apropiado en un dominio especializado, pero los caracteres visualmente confundibles son riesgosos porque dos identificadores distintos pueden parecer casi iguales durante una revisión.

## 5. Chuleta rápida

- Inicio de identificador → letra Unicode o `_`.
- Caracteres posteriores → letras Unicode, dígitos Unicode o `_`.
- Las palabras reservadas no pueden ser identificadores.
- `_` solo → identificador vacío; `_0` y `__` son nombres normales.
- Los puntos de código Unicode visualmente similares pueden crear identificadores distintos.
- Nombres de varias palabras → `mixedCaps` o `MixedCaps`, no `snake_case`.
- Las constantes no requieren `UPPER_SNAKE_CASE`.
- Mayúscula inicial → significado de exportación solo en declaraciones de paquete, campo o método que cumplan los requisitos.
- Alcance pequeño y evidente → un nombre corto puede ser claro.
- Nombres habituales de bucles → `i`, `j`; roles evidentes en `range` → `k`, `v`.
- Alcance mayor → añade suficiente descripción para comunicar el propósito.
- Normalmente describe el rol del valor, no su tipo de Go.

### Comprueba que realmente lo sabes

1. ¿Por qué `ａ` y `a` pueden nombrar dos variables diferentes aunque se parezcan, y qué riesgo crea esto?
2. ¿Por qué `maxRetries` y `MaxRetries` comunican información sobre visibilidad en lugar de indicar si alguno de los nombres pertenece a una constante?
3. ¿Cómo deberían influir el alcance y el rol de una variable al elegir entre un nombre de una letra, una palabra corta o un nombre descriptivo de varias palabras?
