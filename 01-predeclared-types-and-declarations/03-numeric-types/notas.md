# Tipos numéricos en Go

Los tipos numéricos predeclarados de Go abarcan enteros, valores de punto flotante y números complejos. Elegir entre ellos requiere entender el tamaño, la precisión, los operadores, las conversiones y la diferencia entre representaciones exactas y aproximadas.

## 1. Qué debo entender

- Go tiene tres familias numéricas: **enteros**, **números de punto flotante** y **números complejos**.
- Todo tipo numérico tiene un valor cero de `0`; para los tipos complejos esto significa `0 + 0i`.
- Usa `int` para el trabajo general con enteros, salvo que un formato externo exija un tamaño o signo exacto.
- Los tipos enteros de ancho fijo siguen siendo distintos de `int`, aunque tengan el mismo tamaño en una plataforma.
- Prefiere `float64` para el trabajo general con punto flotante, pero recuerda que sus valores son aproximaciones.
- Usa `complex64` o `complex128` cuando un valor necesite componentes real e imaginario.
- Los literales numéricos comienzan como constantes sin tipo y reciben un tipo del contexto o un tipo predeterminado como `int`, `float64` o `complex128`.

## 2. Conceptos clave

| Familia | Tipos predeclarados | Representación |
|---|---|---|
| Enteros con signo | `int8`, `int16`, `int32`, `int64`, `int` | Números enteros, incluidos los negativos |
| Enteros sin signo | `uint8`, `uint16`, `uint32`, `uint64`, `uint`, `uintptr` | Cero y números enteros positivos |
| Punto flotante | `float32`, `float64` | Valores reales aproximados mediante IEEE 754 |
| Complejos | `complex64`, `complex128` | Una parte real y una parte imaginaria |

Los rangos de los enteros de ancho fijo son:

| Tipo | Rango |
|---|---|
| `int8` | −128 a 127 |
| `int16` | −32 768 a 32 767 |
| `int32` | −2 147 483 648 a 2 147 483 647 |
| `int64` | −9 223 372 036 854 775 808 a 9 223 372 036 854 775 807 |
| `uint8` | 0 a 255 |
| `uint16` | 0 a 65 535 |
| `uint32` | 0 a 4 294 967 295 |
| `uint64` | 0 a 18 446 744 073 709 551 615 |

Los nombres enteros especiales tienen funciones concretas:

| Tipo | Significado |
|---|---|
| `byte` | Alias de `uint8` |
| `rune` | Alias de `int32`; su uso detallado se estudia más adelante |
| `int` | Entero con signo cuyo tamaño es de 32 o 64 bits |
| `uint` | Entero sin signo con el mismo tamaño que `int` |
| `uintptr` | Entero sin signo suficientemente grande para contener los bits sin interpretar de un valor pointer; su uso detallado se estudia más adelante |

`int` y `uint` suelen tener 32 bits en sistemas de 32 bits y 64 bits en la mayoría de los sistemas de 64 bits, pero el lenguaje solo garantiza que cada uno tenga 32 o 64 bits.

Los tipos de punto flotante y complejos se relacionan por el tamaño de sus componentes:

| Tipo | Tamaño o componentes | Nota práctica |
|---|---|---|
| `float32` | 32 bits; alrededor de 6–7 dígitos decimales de precisión | Úsalo cuando lo exija un formato existente o una medición demuestre una necesidad real de memoria |
| `float64` | 64 bits; tipo predeterminado de los literales de punto flotante | Preferido para el trabajo general con punto flotante |
| `complex64` | Dos componentes `float32` | Valores complejos de menor precisión |
| `complex128` | Dos componentes `float64`; tipo complejo predeterminado | Preferido para el trabajo general con números complejos |

Algunos límites útiles del rango de punto flotante son:

| Tipo | Mayor valor absoluto finito, aproximadamente | Menor valor positivo distinto de cero, aproximadamente |
|---|---:|---:|
| `float32` | `3.4e38` | `1.4e-45` |
| `float64` | `1.8e308` | `4.9e-324` |

## 3. Cómo funciona en Go

Los tipos enteros admiten operaciones aritméticas, comparaciones, operaciones bitwise y desplazamientos:

| Categoría | Operadores |
|---|---|
| Aritméticos | `+`, `-`, `*`, `/`, `%` |
| Comparación | `==`, `!=`, `<`, `<=`, `>`, `>=` |
| Bitwise | `&`, \|, `^`, `&^` |
| Desplazamientos | `<<`, `>>` |
| Asignación aritmética | `+=`, `-=`, `*=`, `/=`, `%=` |
| Asignación bitwise | `&=`, \|=, `^=`, `&^=`, `<<=`, `>>=` |

La división de enteros produce un entero y trunca hacia cero:

```go
const positiveQuotient = 5 / 3
const negativeQuotient = -5 / 3
```

Los resultados son `1` y `-1`. Convierte los operandos antes de dividir cuando necesites un resultado de punto flotante:

```go
var numerator int = 5
var denominator int = 3
var ratio = float64(numerator) / float64(denominator)
```

`ratio` es aproximadamente `1.6666666666666667`.

Los tipos enteros con nombre no se mezclan implícitamente:

```go
var count int = 10
var count32 int32 = int32(count)
```

La conversión explícita es necesaria porque `int` e `int32` son tipos distintos. `byte` y `uint8` son una excepción a esta distinción porque `byte` es un alias de `uint8`.

Los operadores bitwise exponen directamente los patrones de bits de los enteros:

```go
const mask = 0b1111
const selected = mask & 0b0101
const shifted = selected << 1
```

Los tipos de punto flotante admiten `+`, `-`, `*`, `/` y comparaciones, pero no `%`. Su gran rango implica una precisión limitada, por lo que muchas fracciones decimales se almacenan como aproximaciones binarias cercanas.

Una comparación sencilla con tolerancia absoluta puede escribirse así:

```go
func almostEqual(left, right, tolerance float64) bool {
    difference := left - right
    if difference < 0 {
        difference = -difference
    }
    return difference <= tolerance
}
```

La tolerancia adecuada depende de la magnitud de los valores y de la precisión exigida por el problema.

La función built-in `complex` combina los componentes real e imaginario:

```go
var value = complex(20.3, 10.2)

var real32 float32 = 2.5
var imaginary32 float32 = 3.1
var value32 = complex(real32, imaginary32)

var realPart = real(value)
var imaginaryPart = imag(value)
```

`value` recibe el tipo predeterminado `complex128`; `value32` tiene el tipo `complex64`. `real` e `imag` extraen los dos componentes: devuelven `float32` para `complex64` y `float64` para `complex128`.

Los valores complejos admiten `+`, `-`, `*`, `/`, `==` y `!=`, pero no comparaciones de orden. Los literales imaginarios usan el sufijo `i`:

```go
var imaginaryValue = 2.5i
var combined = 3 + 2.5i
```

El paquete estándar `math/cmplx` proporciona operaciones adicionales para `complex128`, como el cálculo de la magnitud.

## 4. Ejemplos, diferencias y errores comunes

Usa el tipo entero que corresponda a la restricción:

| Situación | Elección habitual | Motivo |
|---|---|---|
| Trabajo general con enteros | `int` | Es el tipo entero convencional de propósito general en Go |
| Formato binario o protocolo | Tipo con signo o sin signo del ancho exacto | La representación externa determina el tamaño y el signo |
| Algoritmo para varios tipos enteros | Generics con una restricción de enteros adecuada | Una implementación puede aceptar los tipos enteros previstos |

El código antiguo puede contener implementaciones separadas para valores con signo y sin signo. La fuente menciona `strconv.FormatInt` y `strconv.FormatUint` como ejemplos conocidos de APIs separadas para `int64` y `uint64`.

> **No confundir:** tener el mismo tamaño de almacenamiento no vuelve intercambiables a `int`, `int32` e `int64`. Go sigue tratándolos como tipos con nombre distintos y normalmente exige una conversión explícita.

`byte` es un alias verdadero, por lo que puede asignarse a `uint8`, compararse con él y usarse en operaciones aritméticas con él sin conversión. El nombre `byte` resulta útil cuando el valor representa datos en bytes.

La elección para punto flotante sigue una regla parecida: usa normalmente `float64` y elige `float32` cuando lo justifiquen la interoperabilidad, el almacenamiento o requisitos de memoria medidos.

> ⚠️ **Corrección importante:** Dividir por un cero constante es un error de compilación tanto en expresiones constantes enteras como de punto flotante. En ejecución, un divisor entero igual a cero provoca un panic; un divisor de punto flotante igual a cero produce un infinito si el numerador no es cero y `NaN` para `0 / 0`.

```go
func divideInteger(numerator, divisor int) int {
    return numerator / divisor
}

func divideFloat(numerator, divisor float64) float64 {
    return numerator / divisor
}
```

Llamar a `divideInteger` con un divisor igual a cero provoca un panic. Llamar a `divideFloat` con un divisor igual a cero en tiempo de ejecución sigue el comportamiento de IEEE 754.

Los valores de punto flotante no deben representar dinero ni otra cantidad que exija una representación decimal exacta. Las comparaciones exactas `==` y `!=` son válidas en Go, pero no comprueban igualdad aproximada; usa una tolerancia adecuada para el dominio cuando busques una aproximación.

El resultado de `complex` depende de sus argumentos:

| Argumentos | Resultado |
|---|---|
| Dos constantes numéricas sin tipo | Constante compleja sin tipo; `complex128` predeterminado cuando se necesita un tipo concreto predeterminado |
| Dos valores `float32` | `complex64` |
| Un `float32` y una constante sin tipo representable | `complex64` |
| Dos valores `float64` | `complex128` |
| Un `float64` y una constante sin tipo representable | `complex128` |

> ⚠️ **Corrección importante:** Dos argumentos tipados con tamaños distintos, como `float32` y `float64`, no usan `complex128` como alternativa automática; sus tipos deben igualarse explícitamente antes de llamar a `complex`.

La igualdad de números complejos tiene el mismo problema de exactitud que sus componentes de punto flotante. Para una comparación aproximada, puede aplicarse una tolerancia a la magnitud de la diferencia, por ejemplo con `cmplx.Abs(left - right)`.

Go proporciona directamente tipos complejos, que pueden ser útiles en trabajos como conjuntos de Mandelbrot o ecuaciones cuadráticas. Las funciones numéricas más amplias, como matrices y álgebra lineal, quedan fuera de los tipos numéricos built-in del lenguaje; la fuente menciona Gonum como una opción de terceros para cálculo numérico.

## 5. Chuleta rápida

- Familias numéricas → enteros, punto flotante y números complejos.
- Todo valor cero numérico → `0`; valor cero complejo → `0 + 0i`.
- Entero general → `int`; formato externo fijo → tipo entero del ancho exacto.
- `byte` = `uint8`; `rune` = `int32`.
- `int` y `uint` tienen 32 o 64 bits; los tipos de ancho fijo siguen siendo distintos.
- La división entera trunca hacia cero; `%` solo admite enteros.
- Operadores bitwise para enteros → `&`, `|`, `^`, `&^`, `<<`, `>>`.
- Punto flotante general → `float64`; `float32` tiene alrededor de 6–7 dígitos decimales de precisión.
- ⚠️ Los floats son aproximados: evítalos para cantidades decimales exactas y usa una tolerancia adecuada para igualdad aproximada.
- La división constante por cero es inválida; la división entera por cero en ejecución provoca panic; la división de floats en ejecución sigue IEEE 754.
- `complex64` usa componentes `float32`; `complex128` usa componentes `float64`.
- `complex(realPart, imaginaryPart)`, `real(value)`, `imag(value)`; los literales imaginarios terminan en `i`.

### Comprueba que realmente lo sabes

1. ¿Cuándo deberías elegir `int` en lugar de un entero de ancho fijo y por qué un `int` no puede asignarse directamente a un `int32`?
2. ¿Cómo se diferencian la división entera y la de punto flotante en el resultado, la precisión y el comportamiento cuando un divisor en ejecución es cero?
3. ¿Qué tipo produce `complex` con dos valores `float32`, con dos constantes sin tipo y con valores tipados `float32` y `float64`?
