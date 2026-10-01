# Answers

## Question 1

Cuando una variable se declara sin inicializador, Go le asigna el valor cero de su tipo:

- `bool`: `false`.
- Tipos enteros (`int`, `int8`, `uint`, `byte`, `rune`, etc.): `0`.
- Tipos de punto flotante (`float32`, `float64`): `0`.
- Tipos complejos (`complex64`, `complex128`): `0 + 0i`.
- `string`: `""`.

Esto garantiza que toda variable declarada tiene desde el principio un valor válido y definido; nunca contiene datos indeterminados por no haber sido inicializada explícitamente.

## Question 2

Los literales estudiados se escriben de estas formas:

- **Enteros:** decimal (`42`), binario (`0b101010`), octal (`0o52`, o la forma antigua `052`) y hexadecimal (`0x2A`). Se puede usar `_` para separar dígitos, por ejemplo `1_000` o `0b0010_1010`. El guion bajo puede ir entre dígitos o justo después del prefijo de base, pero no al final, repetido consecutivamente ni junto al punto decimal o al marcador/signo del exponente.
- **Punto flotante decimal:** puede contener punto decimal y un exponente opcional `e` o `E`, que representa una potencia de diez; por ejemplo, `3.14`, `.5`, `1e6` o `6.02e23`.
- **Punto flotante hexadecimal:** empieza con `0x` o `0X` y debe incluir un exponente `p` o `P`, que representa una potencia de dos; por ejemplo, `0x1.8p1`.
- **Rune:** usa comillas simples y representa un solo code point Unicode o un escape válido, por ejemplo `'A'`, `'界'`, `'\n'`, `'\x41'`, `'\u0041'` o `'\U00000041'`.
- **String interpretado:** usa comillas dobles, interpreta escapes y no permite saltos de línea literales; por ejemplo, `"hola\nGo"`. Una comilla doble interior debe escribirse como `\"`.
- **String raw (sin interpretar):** usa comillas invertidas, conserva las barras inversas y permite saltos de línea y comillas dobles. No puede contener una comilla invertida.

Estos literales básicos representan inicialmente constantes sin tipo. Sus delimitadores importan: `'A'` es un `rune`, mientras que `"A"` es un `string`.

## Question 3

Las familias numéricas se distinguen por lo que representan y por sus operaciones:

- Los **enteros** con signo son `int`, `int8`, `int16`, `int32` e `int64`; los enteros sin signo son `uint`, `uint8`, `uint16`, `uint32`, `uint64` y `uintptr`. Representan números enteros exactos. Los tipos con ancho fijo tienen un rango definido; `int` y `uint` tienen 32 o 64 bits según la plataforma. Admiten aritmética entera, módulo, operaciones bitwise y desplazamientos. La división entera trunca hacia cero.
- Los **puntos flotantes**, `float32` y `float64`, representan números reales mediante aproximaciones IEEE 754. `float64` es la opción habitual porque ofrece más precisión. Admiten aritmética y comparaciones, pero no el operador `%`.
- Los **complejos**, `complex64` y `complex128`, contienen una parte real y otra imaginaria. Sus componentes son, respectivamente, `float32` y `float64`. Admiten aritmética, igualdad y desigualdad, pero no comparaciones de orden.

`byte` es un alias de `uint8` y comunica que el valor representa un byte. `rune` es un alias de `int32` y comunica que el valor representa un code point Unicode.

Un `string` es una secuencia inmutable de bytes. El texto fuente suele almacenarse en UTF-8, pero `text[0]` devuelve un `byte`, no necesariamente el primer `rune` ni el primer carácter visible. Un code point Unicode puede ocupar varios bytes; para trabajar por code points se debe decodificar o recorrer el string como runes.

## Question 4

Una constante puede ser **tipada** o **sin tipo**:

```go
const untypedLimit = 10
const typedLimit int = 10
```

Una constante sin tipo conserva una clase o *kind*: booleana, `rune`, entera, de punto flotante, compleja o string. Mientras siga sin tipo, puede adoptar el tipo que exija un contexto compatible. Por ejemplo, `const limit = 10` puede inicializar un `int`, un `float64` o un `byte`, porque su clase es compatible y el valor `10` es representable en esos tipos.

Si no existe un tipo contextual, Go usa el tipo predeterminado de la clase:

- booleana -> `bool`;
- `rune` -> `rune` (`int32`);
- entera -> `int`;
- punto flotante -> `float64`;
- compleja -> `complex128`;
- string -> `string`.

El contexto puede proceder de una declaración, una asignación o de otro operando tipado. No basta con que la clase sea compatible: el valor también debe ser representable. Por eso `var mask byte = 255` es válido, `var mask byte = 256` falla por rango y `var count int = 10.5` falla porque el valor no es un entero exacto.

Una constante tipada ya tiene un tipo fijo y sigue las reglas normales de asignabilidad y operaciones. Por ejemplo, una constante `int` no se convierte implícitamente en `float64`; hace falta `float64(typedLimit)`. Al inicializar una variable desde una constante sin tipo siempre se termina seleccionando un tipo concreto para la variable.

## Question 5

Una variable `int` y una variable `float64` ya son valores tipados de tipos diferentes. Go no aplica promociones numéricas implícitas para encontrar un tipo común, así que no se pueden sumar directamente:

```go
var count int = 10
var measurement float64 = 30.2

var asFloat = float64(count) + measurement // 40.2, tipo float64
var asInt = count + int(measurement)       // 40, tipo int
```

En cambio, una constante numérica sin tipo puede adoptar el tipo exigido por el otro operando si su valor es compatible y representable:

```go
var total = measurement + 2 // 2 se usa como float64
```

La dirección de la conversión cambia el tipo en el que se realiza la operación y puede cambiar el valor. Convertir el `int` a `float64` conserva aquí la fracción de `measurement`; convertir `measurement` a `int` primero descarta su parte fraccionaria hacia cero. Otras conversiones pueden perder precisión o bits, por lo que la dirección debe elegirse según el significado deseado del cálculo.

## Question 6

Cada bloque tiene su propio scope. Una declaración corta dentro de un bloque anidado no reutiliza una variable que solo existe en un bloque exterior: declara una variable nueva con el mismo nombre, la cual oculta temporalmente a la exterior.

```go
count := 10

if count > 0 {
    count, label := 20, "inner"
    _, _ = count, label
}

// count todavía vale 10
```

El `count` del `if` es otro binding y deja de existir al terminar el bloque; por eso asignarle `20` no modifica el `count` exterior.

En una declaración corta, un nombre solo puede reutilizar una variable si ya fue declarado en el mismo bloque (o es un parámetro cuando la declaración aparece en el cuerpo de esa función), conserva el mismo tipo y la declaración introduce al menos una variable nueva distinta de `_`. Un nombre encontrado únicamente en un bloque exterior cuenta como una declaración nueva en el bloque interior y produce shadowing. Para actualizar intencionalmente la variable exterior se debe usar asignación con `=`, no `:=`.

## Question 7

`var` y `:=` tienen estas diferencias:

| Característica | `var` | `:=` |
|---|---|---|
| Scope permitido | Nivel de paquete y dentro de funciones | Solo dentro de funciones |
| Inicializador | Puede omitirse si se indica el tipo | Siempre es obligatorio |
| Tipo explícito | Permitido: `var count int` | No forma parte de su sintaxis |
| Valor cero | Puede declarar directamente una variable con su valor cero | No puede hacerlo sin proporcionar un inicializador |
| Inferencia | Ocurre si se omite el tipo: `var count = 10` | Ocurre siempre: `count := 10` |
| Varias variables | Sí | Sí |
| Bloques agrupados | Sí: `var (...)` | No |
| Reutilización en el mismo bloque | `var` no puede redeclarar el nombre | Puede reutilizarlo si también declara al menos un nombre nuevo no `_` y el reutilizado conserva su tipo |

`var` resulta apropiado cuando importa un tipo explícito, se desea comenzar con el valor cero, se declara a nivel de paquete o se agrupan declaraciones. `:=` es conciso para variables locales nuevas cuyo tipo correcto resulta evidente a partir del inicializador.

`:=` declara variables; `=` solo asigna valores a variables existentes. Además, una declaración corta dentro de un bloque interior puede crear shadowing aunque exista un nombre idéntico fuera de ese bloque.

## Question 8

Elegiría nombres idiomáticos en `mixedCaps` y tipos que expresen la función de cada dato. Dentro del cuerpo de una función, las declaraciones podrían ser:

```go
serviceName := "payments"                  // string inferido
var statusMarker rune = '✓'                // un code point Unicode
enabled := true                            // bool inferido
var requestCount int                       // contador general; empieza en 0
var processedBytes uint64 = 1_024          // ancho fijo para un formato externo
var latencyMilliseconds float64 = 12.5     // medición decimal
var permissionMask byte = 0b1110_0100      // byte/uint8, patrón de bits visible
const warningThreshold = 10.0              // constante numérica sin tipo
var message string                         // mensaje opcional; empieza como ""

_ = serviceName
_ = statusMarker
_ = enabled
_ = requestCount
_ = processedBytes
_ = latencyMilliseconds
_ = permissionMask
_ = warningThreshold
_ = message
```

En código real usaría esos valores para construir, imprimir o devolver el diagnóstico; las asignaciones a `_` solo muestran que una variable local no puede quedar sin uso. Si un dato no tiene ningún propósito, es mejor eliminar su declaración que silenciar el compilador.

Las decisiones son las siguientes:

- `string` representa el nombre y el mensaje. `var message string` aprovecha que su valor cero es `""` para expresar «sin mensaje».
- `rune` deja claro que el marcador es un code point, y el literal `'✓'` usa las comillas simples correctas.
- `bool` representa `enabled`; una condición sería `if enabled`, no una prueba de *truthiness* sobre números o strings. Para estos escribiría condiciones explícitas, como `requestCount != 0` o `message != ""`.
- `int` es adecuado para un contador de propósito general. `uint64` se justifica si un protocolo, archivo o API exige exactamente ese ancho y valores no negativos.
- `float64` es la elección general para una medición decimal, recordando que su representación es aproximada.
- `byte`, alias de `uint8`, expresa que la máscara ocupa ocho bits; el literal binario hace visible el patrón de permisos y `0b1110_0100` cabe en su rango.
- `warningThreshold` permanece sin tipo para poder usarse en contextos numéricos compatibles. Si el tipo fuera parte de su contrato, usaría una constante tipada, por ejemplo `const warningThreshold float64 = 10.0`.

No mezclaría directamente `requestCount`, `processedBytes` y `latencyMilliseconds`, porque son valores tipados diferentes. Elegiría primero el tipo correcto para la operación y convertiría de forma explícita, por ejemplo `float64(requestCount) + latencyMilliseconds`, comprobando el rango antes de realizar una conversión reductora.

Por último, usaría `=` para actualizar variables ya existentes. Reservaría `:=` para bindings locales realmente nuevos y evitaría repetir esos nombres en bloques interiores, porque hacerlo podría ocultar accidentalmente el valor exterior.
