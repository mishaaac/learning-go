# Conversión explícita de tipos en Go

Go normalmente exige que las conversiones entre valores tipados de distintos tipos numéricos sean visibles en el código. Esto hace explícitos el tipo de resultado elegido y cualquier posible cambio de valor, mientras que las condiciones booleanas deben expresarse mediante comparaciones reales en lugar de `truthiness`.

## 1. Qué debo entender

- Una conversión usa la forma `Tipo(valor)`, como `float64(x)` o `int(y)`.
- Las variables tipadas con tipos numéricos distintos no se promocionan automáticamente a un tipo común.
- La conversión elegida determina el tipo de la operación y puede cambiar el valor resultante.
- Las conversiones numéricas no son necesariamente exactas: pueden perderse fracciones, bits altos de un entero o precisión de punto flotante.
- Go no tiene `truthiness` para números, strings u otros valores no booleanos; una condición debe producir un valor booleano.
- Las comparaciones como `x == 0` y `text == ""` producen valores `bool` sin convertir sus operandos a `bool`.

> ⚠️ **Corrección importante:** La ausencia de promoción automática se aplica a valores tipados de tipos distintos. Una constante sin tipo representable puede adaptarse al tipo numérico requerido, y los alias como `byte` y `uint8` son el mismo tipo.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| Conversión explícita | `Tipo(valor)` produce un valor de otro tipo | El tipo deseado queda visible en el código |
| Valor tipado | Un valor cuyo tipo ya está fijado | Los valores numéricos tipados distintos suelen necesitar conversión antes de combinarse |
| Constante sin tipo | Una constante cuyo tipo concreto todavía no se ha seleccionado | Puede usarse como un tipo compatible cuando su valor es representable |
| Conversión reductora | Conversión a un tipo con menor rango o precisión | El valor convertido puede cambiar |
| `Truthiness` | Tratar un valor no booleano como verdadero o falso | Go no usa esta regla para números ni strings |
| Comparación | Una expresión como `x != 0` | Produce explícitamente la condición booleana |

Las operaciones habituales de la fuente son:

| Objetivo | Expresión en Go | Efecto |
|---|---|---|
| Convertir `int` a `float64` | `float64(x)` | Produce un valor de punto flotante |
| Convertir `float64` a `int` | `int(y)` | Descarta la parte fraccionaria hacia cero |
| Convertir `byte` a `int` | `int(data)` | Produce un `int` con el mismo valor |
| Convertir `int` a `byte` | `byte(x)` | Produce un `uint8`; pueden descartarse bits altos |
| Comprobar si un número es cero | `x == 0` | Produce un `bool` |
| Comprobar si un string está vacío | `text == ""` | Produce un `bool` |

## 3. Cómo funciona en Go

Dados un `int` y un `float64`, debe convertirse uno de los operandos antes de sumarlos:

```go
var x int = 10
var y float64 = 30.2

var sumFloat float64 = float64(x) + y
var sumInt int = x + int(y)
```

`sumFloat` vale `40.2`. `sumInt` vale `40` porque convertir el valor no constante `y` a `int` descarta su parte fraccionaria antes de la suma.

La conversión de punto flotante a entero trunca hacia cero en lugar de redondear:

```go
var positiveFloat float64 = 30.9
var negativeFloat float64 = -30.9

var positiveInt = int(positiveFloat)
var negativeInt = int(negativeFloat)
```

Los resultados son `30` y `-30`.

Los tipos enteros distintos siguen la misma regla explícita:

```go
var count int = 10
var data byte = 100

var sumAsInt int = count + int(data)
var sumAsByte byte = byte(count) + data
```

La primera suma se realiza como `int`; la segunda se realiza como `byte` (`uint8`). El tipo de destino es una decisión de diseño, no solo sintaxis.

Convertir a un tipo entero más pequeño puede cambiar el valor:

```go
var large int = 300
var reduced byte = byte(large)
```

`byte` tiene 8 bits, por lo que esta conversión conserva los 8 bits bajos y `reduced` se convierte en `44`. Una conversión no valida que el valor original quepa.

Las constantes sin tipo representables son una excepción importante a la necesidad habitual de conversión:

```go
var measurement float64 = 10
var total = measurement + 2
```

Las constantes `10` y `2` no tienen tipo y pueden representarse como `float64`, por lo que no hace falta una conversión explícita. Una variable tipada como `int` no podría reemplazar a `2` en la suma sin conversión.

Go obtiene valores booleanos expresando condiciones explícitas:

```go
var number = 10
var text = "Go"

var isZero = number == 0
var isNonZero = number != 0
var isEmpty = text == ""
```

`isZero`, `isNonZero` e `isEmpty` tienen el tipo `bool`. Los valores numérico y string se compararon; no se convirtieron a `bool`.

## 4. Ejemplos, diferencias y errores comunes

La dirección de la conversión afecta tanto al tipo como al valor:

| Expresión | Tipo de operación | Resultado del ejemplo |
|---|---|---:|
| `float64(x) + y` | `float64` | `40.2` |
| `x + int(y)` | `int` | `40` |
| `count + int(data)` | `int` | `110` |
| `byte(count) + data` | `byte` | `110` |

> **No confundir:** una conversión cambia o reexpresa un valor según las reglas de conversión; una comparación formula una pregunta y devuelve `true` o `false`.

Go rechaza una operación aritmética directa como `x + y` cuando `x` es un `int` tipado e `y` es un `float64` tipado. Convierte el operando que corresponda al tipo en el que deba realizarse el cálculo.

Una conversión numérica puede perder información:

- Convertir un float a un entero descarta la fracción hacia cero.
- Convertir un entero a un tipo entero más pequeño descarta bits altos después de extenderlo a precisión infinita.
- Convertir un entero o float a un tipo de punto flotante redondea a la precisión del destino.
- Convertir un valor de punto flotante no constante fuera del rango del entero de destino tiene un resultado dependiente de la implementación; valida primero el rango.

Las conversiones de constantes se comprueban durante la compilación. Por ejemplo, `int(30.2)` es inválido porque la constante `30.2` no puede representarse como `int`, mientras que `int(y)` está permitido cuando `y` es una variable `float64` y sigue las reglas de conversión numérica en ejecución.

> **No confundir:** la ausencia de `truthiness` en Go significa que `if number` e `if text` son inválidos. Escribe la condición deseada, como `number != 0` o `text != ""`.

Los números y strings no pueden convertirse directamente a `bool`; expresiones como `bool(1)` y `bool("true")` son inválidas.

> ⚠️ **Corrección importante:** Un tipo definido cuyo tipo subyacente sea `bool` puede convertirse explícitamente a `bool`. Esta es una conversión entre tipos booleanos, no `truthiness`.

```go
type Enabled bool

var custom Enabled = true
var standard bool = bool(custom)
```

El resultado sigue basándose en un valor booleano existente. Go nunca interpreta automáticamente un valor numérico o string como verdadero o falso.

## 5. Chuleta rápida

- Sintaxis de conversión → `Tipo(valor)`.
- Los valores numéricos tipados de tipos distintos suelen requerir conversión explícita.
- `float64(x)` → convierte un valor entero a `float64`.
- `int(y)` → convierte un float no constante a `int` truncando hacia cero.
- `int(data)` y `byte(x)` → convierten entre tipos enteros.
- ⚠️ Una conversión puede perder rango o precisión; no es una validación.
- Las constantes sin tipo representables pueden adaptarse sin conversión explícita.
- `byte` y `uint8` son alias, por lo que no necesitan conversión entre sí.
- Go no tiene `truthiness` numérico ni de strings.
- `x != 0` y `text != ""` → condiciones booleanas explícitas.
- Número/string → `bool` es inválido; los tipos booleanos compatibles sí pueden convertirse entre sí.

### Comprueba que realmente lo sabes

1. ¿Por qué `float64(x) + y` y `x + int(y)` producen tipos y valores distintos cuando `x` vale `10` e `y` vale `30.2`?
2. ¿Por qué una constante sin tipo puede combinarse con una variable `float64` sin conversión explícita, pero una variable `int` tipada no puede hacerlo?
3. ¿Por qué `bool(number)` e `if text` son inválidos y cómo deben expresarse esas condiciones en Go?
