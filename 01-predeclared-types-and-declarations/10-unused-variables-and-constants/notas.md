# Variables y constantes no utilizadas en Go

Go rechaza las variables declaradas dentro del cuerpo de una función cuando nunca se utilizan, lo que evita que restos accidentales permanezcan en código compilable. Esta comprobación no demuestra que cada valor asignado sea útil, y los parámetros, las variables de paquete y las constantes sin uso siguen reglas diferentes.

## 1. Qué debo entender

- Una variable declarada dentro del cuerpo de una función debe utilizarse; de lo contrario, la compilación falla con el error `declared and not used`.
- Limitarse a asignar otro valor a esa variable no cuenta como utilizar su valor.
- Una vez que la variable se utiliza en algún lugar, el compilador no realiza un análisis completo del flujo de datos para comprobar que cada valor asignado se lea posteriormente.
- Los parámetros y receivers de funciones pueden permanecer sin utilizar aunque las variables declaradas en el cuerpo de la función no puedan hacerlo.
- Las variables a nivel de paquete pueden permanecer sin utilizar, y tanto las constantes locales como las de paquete pueden quedar sin uso.
- El identificador vacío `_` puede descartar deliberadamente un valor, pero no debería utilizarse para ocultar código que tendría que eliminarse.
- En el ejemplo estudiado, `go vet` no informa las asignaciones sobrescritas o finales que nunca se leen; analizadores especializados pueden realizar comprobaciones más profundas.

> ⚠️ **Corrección importante:** La regla estricta sobre variables sin uso se aplica a las variables declaradas dentro del cuerpo de una función. Los parámetros y receivers sin utilizar están permitidos, por lo que «toda variable local debe utilizarse» resulta demasiado amplio sin esta distinción.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| Variable local sin utilizar | Una variable declarada en el cuerpo de una función que nunca se utiliza | Produce un error de compilación |
| Uso de una variable | Una operación que lee o utiliza de otra forma el valor de la variable declarada | Satisface la comprobación del compilador sobre variables sin uso |
| Asignación no leída | Un valor asignado a una variable utilizada, pero sobrescrito o abandonado antes de leerse | Puede compilar aunque la asignación sea innecesaria |
| Variable de paquete | Una variable declarada fuera de los cuerpos de funciones | El compilador permite que permanezca sin uso |
| Constante sin utilizar | Una constante declarada a la que nunca se hace referencia | Está permitida porque la evaluación de constantes no tiene efectos secundarios en tiempo de ejecución |
| Identificador vacío | `_`, un marcador que descarta un valor sin crear una vinculación | Resulta útil cuando un valor producido no se necesita intencionadamente |

La comparación clave es:

| Situación | Resultado del compilador |
|---|---|
| Variable declarada en el cuerpo de una función y nunca utilizada | ❌ Error |
| Parámetro o receiver de función nunca utilizado | ✅ Permitido |
| Variable a nivel de paquete nunca utilizada | ✅ Permitido |
| Constante local o de paquete nunca utilizada | ✅ Permitido |
| Asignación concreta sobrescrita antes de una lectura posterior | ✅ Puede estar permitida |
| Asignación final concreta que nunca se lee | ✅ Puede estar permitida |

La regla del compilador trata principalmente de si se utiliza la vinculación de la variable, no de si cada valor que pasa por ella contribuye al resultado del programa.

## 3. Cómo funciona en Go

Esta declaración fallaría si se quitara el comentario porque `value` nunca se utiliza:

```go
func unusedLocal() {
    // value := 10 // error de compilación: declared and not used
}
```

Volver a asignar un valor a la variable no corregiría el problema si su valor siguiera sin leerse:

```go
func stillUnused() {
    // value := 10
    // value = 20 // la asignación por sí sola no es un uso significativo
}
```

El ejemplo de la fuente compila porque `fmt.Println(x)` utiliza `x`:

```go
import "fmt"

func demonstrateAssignments() {
    x := 10
    x = 20
    fmt.Println(x)
    x = 30
}
```

Los valores individuales tienen resultados diferentes:

```text
x := 10       → se sobrescribe antes de leerse
x = 20        → fmt.Println lo lee
x = 30        → no se lee posteriormente
```

El compilador observa un uso válido de la variable `x`; no rechaza la primera y la última asignación como asignaciones muertas.

Los parámetros, las variables de paquete y las constantes pueden permanecer sin utilizar:

```go
var packageValue = 10

const packageLimit = 20

func process(unusedParameter int) {
    const localLimit = 30
}
```

Las tres declaraciones sin uso están permitidas. Esto no significa que conservar nombres innecesarios sea un buen diseño; únicamente describe lo que acepta el compilador.

El identificador vacío puede descartar un valor explícitamente:

```go
func discardValue() {
    value := 10
    _ = value
}
```

`_ = value` evalúa `value` y evita el error de variable sin utilizar. Cuando una declaración es realmente innecesaria, eliminarla suele ser más claro que añadir una asignación vacía únicamente para silenciar al compilador.

```text
variable declarada en el cuerpo de una función
               │
               ├── nunca se utiliza → error de compilación
               │
               └── se utiliza al menos una vez → compila
                                                │
                                                └── pueden quedar asignaciones individuales sin leer
```

## 4. Ejemplos, diferencias y errores comunes

| Ejemplo | ¿Válido? | Motivo |
|---|---:|---|
| `func run() { value := 10 }` | ❌ | Una variable declarada en el cuerpo de la función nunca se utiliza |
| `func run(value int) {}` | ✅ | Un parámetro de función sin utilizar está permitido |
| `var packageValue = 10` | ✅ | Una variable a nivel de paquete sin utilizar está permitida |
| `func run() { const limit = 10 }` | ✅ | Una constante local sin utilizar está permitida |
| `value := 10; _ = value` dentro de una función | ✅ | La asignación vacía descarta explícitamente el valor |
| Declarar `x`, imprimirla una vez y asignar después un valor final que no se lee | ✅ | La vinculación se utiliza aunque la asignación final no se lea |

> **No confundir:** una variable sin utilizar no es lo mismo que una asignación no leída. La primera se rechaza dentro del cuerpo de una función; la segunda puede permanecer después de que la variable se haya utilizado de otra forma.

Para la secuencia de asignaciones estudiada, tanto el compilador como `go vet` estándar aceptan el código. Esto no debe generalizarse como si `go vet` no tuviera ninguna comprobación de asignaciones; simplemente no proporciona una detección general de asignaciones muertas para este caso. Herramientas de análisis estático más especializadas pueden informar valores asignados que nunca se leen.

> ⚠️ **Corrección importante:** Go no garantiza el contenido exacto del binario generado. Una constante sin uso no tiene evaluación ni efectos secundarios en tiempo de ejecución y puede descartarse, pero afirmar que estará ausente de todos los binarios compilados excede las reglas del lenguaje.

Las variables de paquete sin uso son legales, pero aun así pueden dificultar la comprensión del programa al sugerir un estado que no tiene ningún propósito. El permiso del compilador no es una recomendación para conservarlas.

> **No confundir:** `_ = value` evita intencionadamente el error de variable sin utilizar; no demuestra que la declaración original contribuya con un comportamiento útil.

## 5. Chuleta rápida

- Variable declarada en el cuerpo de una función y nunca utilizada → error de compilación.
- Una asignación posterior por sí sola no cuenta como lectura del valor de la variable.
- Un uso válido de una variable no vuelve útil cada asignación.
- Las asignaciones sobrescritas o finales sin lectura todavía pueden compilar.
- Los parámetros y receivers sin utilizar están permitidos.
- Las variables a nivel de paquete sin utilizar están permitidas.
- Las constantes locales y de paquete sin utilizar están permitidas.
- `_` descarta un valor sin introducir una vinculación.
- Prefiere eliminar código muerto antes que añadir `_ = value` solo para silenciar al compilador.
- Las asignaciones muertas estudiadas no son informadas por `go vet` estándar.
- Los analizadores estáticos especializados pueden detectar más asignaciones no leídas.
- ⚠️ El contenido del binario es un detalle de implementación; aquí solo puede darse por segura la ausencia de evaluación y efectos secundarios de constantes en tiempo de ejecución.

### Comprueba que realmente lo sabes

1. ¿Por qué `x := 10; x = 20` sigue fallando cuando `x` nunca se lee, mientras que el ejemplo más largo que contiene `fmt.Println(x)` sí compila?
2. En la secuencia `x := 10; x = 20; fmt.Println(x); x = 30`, ¿qué valores asignados se leen y qué comprueba realmente el compilador?
3. ¿Por qué pueden compilar un parámetro, una variable a nivel de paquete o una constante sin utilizar, aunque una variable sin uso declarada dentro del cuerpo de una función no pueda hacerlo?
