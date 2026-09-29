# `var` vs. `:=` en Go

Go ofrece declaraciones con `var` y declaraciones cortas de variables con `:=`. Elegir entre ambas formas comunica si lo más importante es el scope, un tipo explícito, un zero value o una inferencia local y concisa del tipo.

## 1. Qué debo entender

- `var` funciona a nivel de package y dentro de funciones; `:=` funciona únicamente dentro de funciones.
- `var` puede indicar un tipo explícitamente, inferirlo desde un inicializador o inicializar un tipo declarado con su **zero value**.
- `:=` siempre requiere expresiones de inicialización e infiere a partir de ellas los tipos de las variables.
- `:=` es una declaración, no una forma abreviada de la asignación con `=`.
- Una declaración corta puede reutilizar variables del mismo bloque únicamente cuando también declara al menos una variable nueva distinta de `_` y conserva los tipos de las variables reutilizadas.
- El scope importa: usar `:=` en un bloque interior puede crear una variable nueva que haga shadowing a un nombre de un bloque exterior.
- Normalmente, varias variables deberían declararse juntas solo cuando sus valores estén relacionados de forma natural, como los valores devueltos por una misma operación.

> ⚠️ **Corrección importante:** «Al menos una variable nueva» no es la regla completa de redeclaración de `:=`. Las variables reutilizadas deben haberse declarado antes en el mismo bloque —o en la lista de parámetros de la función cuando la declaración corta está en su cuerpo—, deben conservar el mismo tipo y `_` no cuenta como una variable nueva.

## 2. Conceptos clave

| Característica | `var` | `:=` |
|---|---|---|
| Scope permitido | A nivel de package o dentro de funciones | Solo dentro de funciones |
| Inicializador | Opcional cuando se indica un tipo | Obligatorio |
| Tipo explícito | Permitido | No está permitido en la sintaxis de la declaración |
| Inferencia de tipo | Se usa cuando se omite el tipo | Se usa siempre |
| Declaración con zero value | `var count int` | No disponible |
| Múltiples variables | Permitidas | Permitidas |
| Declaración agrupada | `var (...)` | No disponible |
| Reutilización de nombres existentes | Una declaración `var` no puede redeclarar un nombre en el mismo bloque | Es posible bajo las reglas de redeclaración corta |

Las formas principales de declaración son:

| Forma | Significado |
|---|---|
| `var count int = 10` | Tipo y valor inicial explícitos |
| `var count = 10` | Tipo inferido desde el inicializador |
| `var count int` | Tipo explícito inicializado con su zero value |
| `count := 10` | Declaración corta local con un tipo inferido |
| `var (...)` | Un grupo de declaraciones `var` separadas |

`var count = 10` y `count := 10` producen el mismo tipo y valor cuando aparecen en una función y `count` es nueva. Sin embargo, su sintaxis y los contextos en los que están permitidas siguen siendo diferentes.

## 3. Cómo funciona en Go

Las tres formas habituales de `var` expresan intenciones diferentes:

```go
func declareValues() (int, int, bool) {
    var count int = 10
    var limit = 20
    var ready bool

    return count, limit, ready
}
```

`count` tiene tipo y valor explícitos, `limit` se infiere como `int` y `ready` recibe el zero value `false`.

Una declaración puede introducir varias variables relacionadas:

```go
func coordinates() (int, string) {
    var x, label = 10, "start"
    return x, label
}
```

Sus tipos se infieren por separado: `x` es `int` y `label` es `string`.

Una declaración `var` agrupada combina especificaciones de declaración separadas:

```go
var (
    maxRetries int = 3
    serviceName     = "learner"
    enabled    bool
)
```

Esta sintaxis es válida a nivel de package y dentro de funciones. A nivel de package se requiere `var` porque `:=` no está permitido allí.

Dentro de una función, `:=` es la forma concisa habitual cuando el tipo inferido es el tipo deseado:

```go
func message() string {
    text := "hello"
    return text
}
```

Una declaración corta puede combinar una variable existente del mismo bloque con una variable nueva:

```go
func values() (int, string) {
    count := 10
    count, label := 30, "thirty"
    return count, label
}
```

La segunda declaración corta asigna `30` a la variable existente `count` y declara `label`. Es válida porque `label` es nueva y `count` sigue siendo un `int`.

La declaración y la asignación son operaciones diferentes:

```go
func updateCount() int {
    count := 10
    count = 20
    return count
}
```

La primera sentencia declara `count`; la segunda asigna un valor nuevo a la variable que ya existe.

## 4. Ejemplos, diferencias y errores comunes

| Situación | ¿Válida? | Motivo |
|---|---:|---|
| Nivel de package: `var count = 10` | ✅ | `var` está permitido a nivel de package |
| Nivel de package: `count := 10` | ❌ | Las declaraciones cortas están restringidas a las funciones |
| Local: `var count int` | ✅ | `count` comienza con el zero value `0` |
| Local: `count := 10` | ✅ | El inicializador da a `count` el tipo inferido `int` |
| Después de `count := 10`: `count := 20` | ❌ | La declaración corta no introduce ninguna variable nueva |
| Después de `count := 10`: `count, label := 20, "twenty"` | ✅ | `label` es una variable nueva distinta de `_` |
| Después de `count := 10`: `count, _ := 20, "ignored"` | ❌ | `_` no crea un binding y no cuenta como variable nueva |

> **No confundir:** `:=` declara al menos una variable; `=` solamente asigna valores a variables que ya están declaradas.

Una variable exterior no se redeclara mediante `:=` en un bloque interior. Se crea una variable nueva que hace shadowing a la exterior:

```go
func shadowExample() int {
    count := 10

    if count > 0 {
        count, label := 20, "inner"
        _, _ = count, label
    }

    return count
}
```

La función devuelve `10`. La variable `count` interior es diferente y su scope termina con el bloque del `if`.

> **No confundir:** usar el mismo nombre no implica que sea la misma variable. El bloque en el que se declara un identificador determina su binding y su scope.

Estas dos declaraciones locales son válidas:

```go
var data byte = 20
otherData := byte(20)
```

La primera forma destaca directamente el tipo deseado; la segunda usa una conversión explícita antes de la inferencia. Preferir `var data byte = 20` cuando el tipo no predeterminado forma parte de la intención es una elección de estilo, no un requisito del lenguaje.

Las declaraciones agrupadas deben comunicar una relación en lugar de limitarse a ahorrar líneas. Son especialmente naturales al recibir varios resultados o usar el comma-ok idiom; esos mecanismos se estudian por separado.

Las variables de package son mutables salvo que otra regla impida la mutación. Mantenerlas escasas o sin cambios efectivos es una recomendación de diseño que facilita seguir el flujo de datos; `var` no impone inmutabilidad.

## 5. Chuleta rápida

- `var` → permitido a nivel de package y dentro de funciones.
- `:=` → permitido solo dentro de funciones.
- `var count int = 10` → tipo y valor explícitos.
- `var count = 10` → tipo inferido.
- `var count int` → zero value de `int`.
- `count := 10` → declaración local concisa con inferencia.
- `var (...)` → agrupa declaraciones separadas.
- `:=` puede reutilizar variables del mismo bloque solo si al menos una variable distinta de `_` es nueva y los tipos reutilizados no cambian.
- Un nombre exterior usado con `:=` en un bloque interior puede quedar oculto por una variable nueva.
- `_` nunca introduce un binding y no satisface la regla de «variable nueva».
- `:=` declara; `=` asigna.
- Declara varias variables juntas cuando estén relacionadas y evita abusar de variables mutables de package.

### Comprueba que realmente lo sabes

1. ¿Cuándo comunica `var` una intención que `:=` no puede expresar directamente?
2. ¿Por qué la variable exterior `count` permanece sin cambios cuando un bloque interior ejecuta `count, label := 20, "inner"`?
3. Después de `count := 10`, ¿por qué `count, label := 20, "twenty"` es válido, pero `count, _ := 20, "ignored"` no lo es?
