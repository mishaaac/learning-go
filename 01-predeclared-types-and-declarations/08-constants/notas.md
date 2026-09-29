# Constantes con `const` en Go

`const` da un nombre a un valor definido por las reglas de constantes en tiempo de compilación de Go. Impide la reasignación, pero no convierte una variable arbitraria calculada en tiempo de ejecución ni una estructura de datos en un valor inmutable.

## 1. Qué debo entender

- Una declaración `const` asocia un identificador con una **expresión constante** evaluada en tiempo de compilación.
- Las constantes pueden declararse a nivel de paquete o dentro de una función, y las constantes relacionadas pueden agruparse con `const (...)`.
- Una constante puede tener un tipo explícito o permanecer sin tipo; la diferencia detallada se estudia por separado.
- Los valores constantes de Go están limitados a booleanos, valores numéricos, `rune` y `string`.
- Las expresiones constantes pueden combinar operandos constantes, operadores permitidos, conversiones y ciertas llamadas a funciones predeclaradas.
- Una variable o el resultado de una función ordinaria no es una constante, aunque su valor en tiempo de ejecución parezca predecible.
- `const` no es un mecanismo general para hacer inmutables arrays, slices, maps, structs, campos o variables calculadas en tiempo de ejecución.

> ⚠️ **Corrección importante:** «El compilador puede conocer el valor» es una idea útil, pero incompleta. El valor también debe pertenecer a una de las clases de constantes permitidas por Go y proceder de una expresión constante válida; un valor completo de array, slice, map o struct sigue sin poder ser una constante.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| Declaración de constante | Una declaración que comienza con `const` | Da un nombre a un valor que cumple las reglas de constantes de Go |
| Expresión constante | Una expresión evaluada en tiempo de compilación mediante operaciones constantes permitidas | Puede contener más de un literal, como `20 * 10` |
| Constante tipada | Una constante declarada o convertida con un tipo constante explícito | Su valor debe poder representarse mediante ese tipo |
| Constante sin tipo | Una constante sin un tipo concreto fijado todavía | Puede adaptarse a un contexto compatible |
| Valor de tiempo de ejecución | Un valor producido a partir de variables o de la ejecución de funciones ordinarias | No puede inicializar una constante |
| Inmutabilidad | La imposibilidad de modificar un valor después de crearlo | Las constantes de Go no son una característica general de variables inmutables |

Las clases de constantes permitidas son:

| Clase | Ejemplo |
|---|---|
| Booleana | `true` |
| Entera | `10` |
| Punto flotante | `3.14` |
| Compleja | `2 + 3i` |
| `rune` | `'A'` |
| `string` | `"hello"` |

La comparación central es:

| Característica | `const` | `var` |
|---|---|---|
| Scope de paquete o función | Sí | Sí |
| Reasignación | No | Sí |
| Inicializador calculado en tiempo de ejecución | No | Sí |
| Valores compuestos como slices o structs | No | Sí |
| Declaración agrupada con `(...)` | Sí | Sí |

## 3. Cómo funciona en Go

Una constante puede tener un tipo explícito o no tenerlo:

```go
const maxCount int64 = 10
const greeting = "hello"
const total = 20 * 10
```

`total` es válida porque ambos operandos y la multiplicación forman una expresión constante, por lo que su valor se calcula como `200` durante la compilación.

Las constantes están permitidas tanto a nivel de paquete como dentro de funciones:

```go
const packageLabel = "learning-go"

func localLabel() string {
    const prefix = "item"
    return prefix
}
```

Las constantes relacionadas pueden agruparse:

```go
const (
    idKey   = "id"
    nameKey = "name"
)
```

Ciertas funciones predeclaradas pueden participar en expresiones constantes bajo las condiciones definidas por el lenguaje:

```go
const textLength = len("Go")
const arrayCapacity = cap([4]int{})
const complexValue = complex(2, 3)
const realPart = real(complexValue)
const imaginaryPart = imag(complexValue)
```

`textLength` vale `2`, `arrayCapacity` vale `4` y las operaciones con números complejos siguen siendo constantes porque sus operandos cumplen las reglas de las expresiones constantes. `len` y `cap` no producen constantes para todos los argumentos posibles.

Las variables interrumpen las expresiones constantes:

```go
func calculateTotal() int {
    x := 5
    y := 10
    return x + y
}
```

El resultado es un valor `int` calculado cuando se ejecuta la función. Ni `x + y` ni una llamada a `calculateTotal()` pueden inicializar una `const`.

```text
operandos constantes permitidos
            │
            └── operaciones constantes permitidas
                         │
                         └── expresión constante → puede inicializar const

variable o llamada a función ordinaria
            │
            └── valor de tiempo de ejecución → no puede inicializar const
```

`iota` es una constante entera predeclarada disponible dentro de declaraciones de constantes. Su uso detallado se pospone hasta el tema posterior sobre la definición de valores constantes relacionados.

## 4. Ejemplos, diferencias y errores comunes

| Declaración | ¿Válida? | Motivo |
|---|---:|---|
| `const answer = 6 * 7` | ✅ | Solo utiliza operandos constantes y una operación permitida |
| `const title = "Go"` | ✅ | Un string puede ser una constante |
| `const length = len("Go")` | ✅ | La longitud de un string constante es constante |
| `const size = len([4]int{})` | ✅ | La longitud de este array cumple las reglas de `len` constante |
| `const total = x + y` cuando `x` e `y` son variables | ❌ | Los valores de variables no son operandos constantes |
| `const total = calculateTotal()` | ❌ | Una llamada a una función ordinaria no es una expresión constante |
| `const items = []int{1, 2}` | ❌ | Un slice no es un valor constante permitido |
| `const point = struct{ X int }{X: 1}` | ❌ | Un struct no es un valor constante permitido |

> **No confundir:** una expresión constante está determinada por las reglas del lenguaje Go, no solamente por la posibilidad de que una persona pueda predecir su resultado.

No se puede asignar otro valor a una constante declarada:

```go
const maxAttempts = 3
// maxAttempts = 4 // no se puede asignar a una constante
```

La asignación inválida está comentada para que el ejemplo siga compilando. Las sentencias de incremento y decremento como `maxAttempts++` son inválidas por la misma razón.

> **No confundir:** una constante con nombre no es una variable marcada como `readonly`. Es un nombre para un valor constante y no proporciona almacenamiento inmutable en tiempo de ejecución.

Un resultado de tiempo de ejecución debe almacenarse en una variable:

```go
var currentTotal = calculateTotal()
```

El código puede decidir no reasignar `currentTotal`, pero Go no impone esa decisión mediante `const`. La misma limitación se aplica a arrays, slices, maps, structs y campos individuales de structs.

Las llamadas a `complex`, `real` e `imag` producen constantes únicamente cuando sus operandos cumplen las reglas de constantes. Del mismo modo, `len` y `cap` son constantes solo para argumentos específicos, como un string constante o una expresión de array que cumpla los requisitos; no son automáticamente constantes para slices u otros valores de tiempo de ejecución.

## 5. Chuleta rápida

- `const` → da nombre a un valor permitido por las reglas de constantes en tiempo de compilación de Go.
- Clases permitidas → booleana, entera, punto flotante, compleja, `rune` y `string`.
- `const maxCount int64 = 10` → constante tipada.
- `const total = 20 * 10` → expresión constante sin tipo con valor `200`.
- `const (...)` → agrupa constantes relacionadas.
- Las constantes pueden aparecer a nivel de paquete o dentro de funciones.
- Una constante no puede reasignarse, incrementarse ni decrementarse.
- Las variables y llamadas a funciones ordinarias no pueden ser operandos constantes.
- Algunos usos de `complex`, `real`, `imag`, `len` y `cap` producen constantes.
- Los arrays, slices, maps, structs y campos de structs no pueden declararse constantes.
- `iota` está disponible en declaraciones de constantes y se estudia por separado.
- ⚠️ `const` no es una característica general de variables inmutables o `readonly`.

### Comprueba que realmente lo sabes

1. ¿Por qué `const total = 20 * 10` es válida, mientras que `const total = x + y` no lo es cuando `x` e `y` son variables inicializadas con `20` y `10`?
2. ¿Por qué ni un literal de slice ni el resultado predecible de una función ordinaria pueden hacerse inmutables simplemente colocándolos en una declaración `const`?
3. ¿Bajo qué tipo de condiciones pueden las llamadas a `len`, `cap`, `complex`, `real` o `imag` participar en una expresión constante?
