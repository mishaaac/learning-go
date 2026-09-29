# Constantes tipadas y sin tipo en Go

Las constantes de Go pueden ser tipadas o `untyped`. Esta distinción controla si una constante puede adaptarse a un contexto compatible o si ya tiene un tipo fijado que sigue las reglas normales de asignación.

## 1. Qué debo entender

- `const count = 10` declara una **constante sin tipo**; `const count int = 10` declara una constante tipada.
- Una constante sin tipo tiene una clase, como entero sin tipo o string sin tipo, y también un tipo predeterminado que se usa cuando no se proporciona otro tipo.
- El contexto puede dar un tipo compatible a una constante sin tipo, pero su valor debe poder representarse mediante ese tipo.
- Si ningún contexto elige un tipo, Go usa el tipo predeterminado de la constante.
- Una constante tipada ya tiene un tipo fijado y no se adapta implícitamente a otro tipo numérico distinto.
- Las variables siempre están tipadas: inicializar una variable desde una constante sin tipo termina seleccionando un tipo concreto.
- Las constantes sin tipo suelen ser más reutilizables; las constantes tipadas son útiles cuando el tipo forma parte del significado buscado.

> ⚠️ **Corrección importante:** Decir que una constante tipada solo puede asignarse «al tipo correspondiente» es una simplificación. Sigue las reglas normales de asignabilidad de Go, pero una constante `int` tipada no se convierte implícitamente en `float64` o `byte`; para esos tipos numéricos distintos se requiere una conversión explícita.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| Constante sin tipo | Una constante sin un tipo concreto fijado | Puede adaptarse a un tipo contextual compatible |
| Clase de constante | La categoría sin tipo de una constante, como entero, punto flotante o string | Limita qué contextos son compatibles incluso antes de elegir un tipo concreto |
| Tipo predeterminado | El tipo concreto elegido cuando el contexto no proporciona otro | Convierte una constante sin tipo en un valor tipado cuando se necesita |
| Constante tipada | Una constante cuyo tipo ya está fijado | Sigue las reglas de asignación y operación de ese tipo |
| Representabilidad | Si un valor constante puede representarse mediante un tipo de destino | Una asignación fuera de rango o incompatible falla durante la compilación |
| Conversión explícita | Una forma como `float64(value)` | Crea intencionadamente una constante de otro tipo compatible |

Los tipos predeterminados son:

| Clase de constante sin tipo | Tipo predeterminado |
|---|---|
| Booleana | `bool` |
| `rune` | `rune` |
| Entera | `int` |
| Punto flotante | `float64` |
| Compleja | `complex128` |
| String | `string` |

La comparación central es:

| Característica | Constante sin tipo | Constante tipada |
|---|---|---|
| Tipo concreto fijado | No | Sí |
| Tiene tipo predeterminado | Sí | No es necesario; su tipo ya se conoce |
| Puede adaptarse a tipos contextuales compatibles | Sí, cuando es representable | No hay cambio implícito a otro tipo numérico distinto |
| Resulta útil para | Valores constantes matemáticos o generales reutilizables | Comunicar o imponer un tipo específico |

## 3. Cómo funciona en Go

Una constante sin tipo puede adaptarse a varios destinos compatibles:

```go
const count = 10

var integerCount int = count
var floatingCount float64 = count
var smallCount byte = count
```

El valor `10` puede representarse mediante los tres tipos de destino, por lo que el compilador usa el tipo exigido por cada declaración.

Cuando no se escribe un tipo de destino, se elige el tipo predeterminado:

```go
const count = 10
const ratio = 2.5
const letter = 'A'

var inferredCount = count // int
var inferredRatio = ratio // float64
var inferredLetter = letter // rune
```

Una constante tipada conserva su tipo declarado:

```go
const typedCount int = 10

var exactCount int = typedCount
var convertedCount float64 = float64(typedCount)
```

`exactCount` recibe directamente la constante `int` tipada. `convertedCount` requiere la conversión explícita porque `float64` es un tipo numérico distinto.

Una conversión también puede hacer que la propia expresión constante sea tipada:

```go
const converted = float64(10)
```

`converted` es una constante `float64` tipada aunque su declaración no contenga un tipo separado antes de `=`.

Cuando una constante tipada y otra sin tipo participan en una expresión compatible, el operando sin tipo puede adoptar el tipo del operando tipado:

```go
const base int = 10
const total = base + 2
```

`2` puede representarse como `int`, por lo que `total` también es una constante `int` tipada.

```text
constante sin tipo
       │
       ├── existe un contexto compatible
       │       ├── valor representable → usa el tipo contextual
       │       └── valor no representable → error de compilación
       │
       └── no hay tipo contextual → usa el tipo predeterminado
```

Las constantes tipadas son útiles cuando un tipo específico comunica significado, incluidos patrones posteriores que definen valores relacionados con `iota`. El uso detallado de `iota` queda fuera de este tema.

## 4. Ejemplos, diferencias y errores comunes

| Declaración | ¿Válida? | Resultado o motivo |
|---|---:|---|
| `const value = 10; var x byte = value` | ✅ | El valor sin tipo puede representarse mediante `byte` |
| `const value = 10; var x float64 = value` | ✅ | El valor sin tipo se adapta a `float64` |
| `const value = 10; var x = value` | ✅ | `x` recibe el tipo predeterminado `int` |
| `const value int = 10; var x int = value` | ✅ | El origen y el destino usan `int` |
| `const value int = 10; var x float64 = value` | ❌ | Un `int` tipado no se convierte implícitamente en `float64` |
| `const value int = 10; var x float64 = float64(value)` | ✅ | La conversión explícita produce una constante `float64` |
| `const value = 1000; var x byte = value` | ❌ | `1000` está fuera del rango de `byte` |
| `const value = 10.5; var x int = value` | ❌ | `10.5` no puede representarse mediante `int` |

> **No confundir:** `untyped` no significa compatibilidad universal. La clase de la constante y su valor exacto todavía deben ser válidos para el tipo de destino.

La misma constante sin tipo puede funcionar en un contexto y fallar en otro:

```go
const limit = 255

var byteLimit byte = limit
var integerLimit int = limit
```

Ambas declaraciones compilan. Si `limit` fuera `256`, la declaración de `byte` fallaría, mientras que la declaración de `int` seguiría siendo válida.

> **No confundir:** una constante sin tipo con nombre sigue siendo una constante; no se convierte en una variable porque distintos usos seleccionen tipos diferentes.

Una variable inferida desde una constante sin tipo deja de estar sin tipo:

```go
const source = 10
var value = source
```

`value` tiene el tipo concreto predeterminado `int`. Después no puede adaptarse como si todavía fuera una constante sin tipo.

Tanto las constantes tipadas como las constantes sin tipo son valores constantes inmutables. Su diferencia está relacionada con la selección del tipo y la compatibilidad, no con la posibilidad de reasignarlas.

## 5. Chuleta rápida

- `const count = 10` → constante entera sin tipo.
- `const count int = 10` → constante `int` tipada.
- Las constantes sin tipo tienen una clase y un tipo predeterminado.
- Tipos predeterminados → `bool`, `rune`, `int`, `float64`, `complex128`, `string`.
- El contexto puede elegir otro tipo compatible para una constante sin tipo.
- El valor de la constante debe poder representarse mediante el tipo elegido.
- `var count = untypedInteger` → tipo predeterminado `int`.
- Las constantes tipadas no cambian implícitamente a otros tipos numéricos distintos.
- `float64(typedCount)` → conversión explícita a una constante `float64` tipada.
- Un operando sin tipo puede adaptarse a un operando tipado compatible dentro de una expresión constante.
- Las variables siempre tienen tipos concretos y no conservan la flexibilidad `untyped`.
- Prefiere constantes sin tipo para ganar flexibilidad; usa constantes tipadas cuando el tipo comunique intención.

### Comprueba que realmente lo sabes

1. ¿Por qué la misma constante sin tipo `10` puede inicializar variables de tipo `int`, `float64` y `byte`, mientras que una constante `int` tipada no puede inicializar las tres directamente?
2. ¿Cuándo usa Go el tipo predeterminado de una constante sin tipo y qué tipo recibe cada clase de constante?
3. ¿Por qué una constante sin tipo con valor `255` puede inicializar un `byte`, mientras que la misma declaración falla después de cambiar el valor a `256`?
