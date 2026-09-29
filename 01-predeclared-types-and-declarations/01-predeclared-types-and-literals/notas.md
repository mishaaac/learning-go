# Tipos predeclarados y literales en Go

Go proporciona tipos básicos con valores iniciales definidos y varias formas literales para escribir constantes directamente en el código fuente. Entender su sintaxis y comportamiento predeterminado evita errores de tipos, comillas y representación numérica.

## 1. Qué debo entender

- Los **tipos predeclarados** están disponibles sin importar un paquete. Incluyen `bool`, `string`, los tipos enteros, los tipos de punto flotante y alias como `byte` y `rune`.
- Una variable declarada sin un valor inicial explícito recibe el **valor cero** de su tipo: `false` para `bool`, `0` para los tipos numéricos y `""` para `string`.
- Las formas literales estudiadas aquí escriben constantes enteras, de punto flotante, `rune` y `string` directamente en el código fuente.
- Los literales enteros pueden usar notación decimal, binaria, octal o hexadecimal. `_` puede agrupar dígitos sin cambiar el valor.
- Un literal `rune` representa un punto de código Unicode y usa comillas simples. Un literal `string` usa comillas dobles o comillas invertidas.
- Las constantes literales estudiadas aquí comienzan como **constantes sin tipo**. El contexto puede darles un tipo; de lo contrario, Go usa su tipo predeterminado.

> ⚠️ **Corrección importante:** No toda construcción llamada literal en Go carece de tipo. Los literales enteros, de punto flotante, `rune` y `string` estudiados aquí representan constantes sin tipo, mientras que un literal compuesto como `[]int{1, 2}` tiene un tipo.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| Tipo predeclarado | Un tipo ya declarado por el lenguaje | Puede usarse sin una importación ni una declaración del programador |
| Valor cero | El valor asignado cuando una variable no tiene un inicializador explícito | Toda variable declarada comienza con un valor definido |
| Literal | Notación en el código fuente para un valor constante | El delimitador o prefijo determina cómo lo interpreta Go |
| Constante sin tipo | Una constante que todavía no tiene un tipo fijo | Puede adaptarse a un tipo compatible exigido por el contexto |
| Tipo predeterminado | El tipo elegido cuando el contexto no aporta otro | Entero → `int`, punto flotante → `float64`, `rune` → `rune` y string → `string` |

Las bases de los enteros usan estas formas:

| Base | Prefijo | Ejemplo | Uso habitual |
|---|---|---|---|
| Decimal | Ninguno | `1234` | Valores numéricos normales |
| Binaria | `0b` o `0B` | `0b1010` | Patrones de bits visibles |
| Octal | `0o` o `0O` | `0o777` | Valores como permisos POSIX |
| Hexadecimal | `0x` o `0X` | `0xFF` | Patrones compactos de bits y bytes |

La forma octal antigua con un cero inicial, como `0777`, es válida, pero `0o777` comunica la base con mayor claridad.

## 3. Cómo funciona en Go

Declarar variables básicas sin inicializadores aplica automáticamente sus valores cero:

```go
var enabled bool
var count int
var price float64
var name string
```

Sus valores son `false`, `0`, `0` y `""`, respectivamente. Esto evita variables sin inicializar cuyo contenido sea indeterminado.

Los guiones bajos pueden mejorar la legibilidad de los literales numéricos:

```go
const population = 1_234_567
const mask = 0b1111_0000
const color = 0xFF_A0_00
```

Un guion bajo puede separar dígitos y puede aparecer inmediatamente después de un prefijo de base, como en `0x_FF`. No puede terminar un literal, aparecer dos veces seguidas ni separar un dígito de un punto decimal o de un marcador de exponente.

Los literales de punto flotante pueden usar notación decimal y un exponente en base 10 con `e` o `E`:

```go
const avogadro = 6.03e23
```

Un literal de punto flotante hexadecimal empieza con `0x` o `0X` y requiere un exponente `p` o `P`, que representa una potencia de dos:

```go
const hexadecimalFloat = 0x12.34p5
```

Aquí, `0x12.34p5` equivale a `582.5` en decimal.

Un literal `rune` contiene un punto de código Unicode o un escape válido:

```go
const letter = 'a'
const newLine = '\n'
const tab = '\t'
const quote = '\''
const backslash = '\\'
```

El mismo valor `rune` puede escribirse con escapes numéricos:

```go
const letterDirect = 'a'
const letterOctal = '\141'
const letterHex = '\x61'
const letterUnicode16 = '\u0061'
const letterUnicode32 = '\U00000061'
```

Los `string` entre comillas dobles interpretan los escapes. Los `string` entre comillas invertidas permiten saltos de línea directos y no dan un significado especial a las barras inversas:

```go
const interpreted = "Greetings and\n\"Salutations\""

const raw = `Greetings and
"Salutations"`
```

Un `string` sin interpretar no puede contener una comilla invertida. Un `string` interpretado no puede contener un salto de línea literal ni una comilla doble sin escapar.

Las constantes literales sin tipo reciben un tipo predeterminado solo cuando se necesita un tipo y el contexto no selecciona otro:

```go
var count = 42
var ratio = 3.5
var letter = 'a'
var message = "hello"

var small int8 = 42
```

Las primeras cuatro variables tienen los tipos predeterminados `int`, `float64`, `rune` (un alias de `int32`) y `string`. El último literal toma en cambio el tipo contextual `int8` porque `42` puede representarse con ese tipo.

```text
literal en el código fuente
            │
            └── constante sin tipo
                    ├── tipo compatible aportado por el contexto
                    └── tipo predeterminado si el contexto no aporta ninguno
```

## 4. Ejemplos, diferencias y errores comunes

| Característica | `string` interpretado | `string` sin interpretar |
|---|---|---|
| Delimitador | `"..."` | Comillas invertidas: `` `...` `` |
| Secuencias de escape | Se interpretan | No se interpretan |
| Salto de línea directo | No permitido | Permitido |
| Comilla doble dentro | Debe escaparse | Permitida directamente |
| Barra inversa dentro | Inicia un escape | Se conserva como barra inversa |
| Comilla invertida dentro | No actúa como delimitador | No permitida |

Estos dos literales producen el mismo texto:

```go
const interpretedText = "Greetings and\n\"Salutations\""
```

```go
const rawText = `Greetings and
"Salutations"`
```

> **No confundir:** `'a'` es un literal `rune` para un punto de código; `"a"` es un literal `string`.

En un `string` entre comillas dobles, una comilla doble necesita `\"`, mientras que una comilla simple puede aparecer directamente. En un literal `rune`, una comilla simple necesita `\'`.

```go
const apostropheText = "It's valid"
const singleQuote = '\''
```

Entre los errores numéricos habituales están la posición inválida de los guiones bajos y una notación octal poco clara:

- `123_` es inválido porque el guion bajo termina el literal.
- `1__234` es inválido porque los guiones bajos no pueden ser consecutivos.
- `1_.5` es inválido porque un guion bajo no puede tocar el punto decimal.
- `0777` es un octal válido, pero `0o777` es más claro.

Usa la base decimal para valores enteros y de punto flotante ordinarios. Las bases binaria y hexadecimal son útiles cuando importa la agrupación de bits, mientras que la octal sirve para valores escritos convencionalmente en base 8, como los permisos POSIX.

> **No confundir:** `e` en un literal de punto flotante decimal escala por una potencia de 10; `p` en un literal de punto flotante hexadecimal escala por una potencia de 2.

Los escapes numéricos de `rune` son válidos, pero el carácter visible suele ser más claro, salvo que el escape comunique mejor la intención.

## 5. Chuleta rápida

- Los tipos básicos predeclarados pueden usarse sin importaciones.
- Valores cero: `bool` → `false`, tipos numéricos → `0`, `string` → `""`.
- Prefijos enteros: `0b` binario, `0o` octal, `0x` hexadecimal; los dígitos ordinarios sin prefijo son decimales, salvo la forma octal antigua con cero inicial.
- `_` mejora la legibilidad numérica sin cambiar el valor; su posición está restringida.
- Los valores de punto flotante decimales usan `e`; los hexadecimales usan `0x` con `p`.
- Un literal `rune` usa comillas simples: `'a'`.
- Un `string` interpretado usa comillas dobles y procesa escapes.
- Un `string` sin interpretar usa comillas invertidas, permite saltos de línea directos y no procesa escapes.
- ⚠️ Prefiere `0o777` frente a la forma menos clara con cero inicial `0777`.
- Los literales numéricos, `rune` y `string` estudiados aquí son constantes sin tipo y usan el contexto o un tipo predeterminado.

### Comprueba que realmente lo sabes

1. ¿Qué valores reciben las variables `bool`, numéricas y `string` sin inicializar, y por qué es útil este comportamiento?
2. ¿Cuándo elegirías un `rune`, un `string` interpretado o un `string` sin interpretar para el mismo carácter o texto visible?
3. ¿Por qué el literal `42` puede convertirse en `int` en `var count = 42`, pero en `int8` en `var small int8 = 42`?
