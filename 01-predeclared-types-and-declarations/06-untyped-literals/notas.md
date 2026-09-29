# Literales sin tipo en Go

Los literales básicos estudiados aquí representan constantes que inicialmente no tienen un tipo fijo de Go. Esto permite que un literal se adapte a un contexto compatible, pero su valor todavía debe poder representarse mediante el tipo de destino.

## 1. Qué debo entender

- Un literal entero, de punto flotante, `rune` o `string` representa inicialmente una **constante sin tipo**.
- El contexto puede dar un tipo concreto a una constante sin tipo sin necesidad de una conversión explícita.
- Si ningún contexto proporciona un tipo, Go usa el tipo predeterminado de esa clase de literal, como `int` para un literal entero o `float64` para uno de punto flotante.
- Un literal solo puede adaptarse cuando su clase es compatible con el tipo requerido y su valor puede representarse mediante ese tipo.
- Una variable ya tiene un tipo concreto, por lo que dos variables numéricas tipadas con tipos distintos normalmente requieren una conversión explícita antes de combinarse.
- Un literal compatible también puede inicializar un tipo definido cuyo tipo subyacente acepte ese valor.

> ⚠️ **Corrección importante:** La afirmación «los literales son `untyped`» se aplica a las formas de literales básicos estudiadas aquí. Un literal compuesto como `[]int{1, 2}` incluye su tipo y no es una constante sin tipo.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| Constante sin tipo | Una constante cuyo tipo concreto de Go todavía no se ha elegido | Puede adaptarse al tipo compatible requerido por el contexto |
| Tipo contextual | Un tipo requerido por una asignación, declaración o expresión | Puede determinar el tipo usado para una constante sin tipo |
| Tipo predeterminado | El tipo elegido cuando ningún otro contexto proporciona uno | Un literal entero usa `int`; uno de punto flotante usa `float64` |
| Representabilidad | Si un valor constante puede almacenarse como el tipo requerido | Un valor fuera de rango o incompatible causa un error de compilación |
| Valor tipado | Un valor cuyo tipo ya está fijado | No recibe la misma flexibilidad que una constante sin tipo |
| Tipo definido | Un nuevo tipo declarado a partir de un tipo subyacente | Un literal compatible y representable puede inicializarlo directamente |

La comparación central es:

| Situación | Resultado |
|---|---|
| Constante sin tipo + valor tipado compatible | La constante puede adoptar el tipo del valor tipado |
| Dos valores tipados de tipos numéricos distintos | Normalmente se requiere una conversión explícita |
| Literal incompatible con el tipo requerido | Error de compilación |
| Literal numérico compatible fuera del rango de destino | Error de compilación |

Para destinos enteros, la constante debe ser un entero exacto dentro del rango. Para destinos de punto flotante, la constante no debe causar overflow; su valor exacto puede redondearse a la precisión del tipo de destino.

## 3. Cómo funciona en Go

Un tipo de destino explícito proporciona el contexto para un literal:

```go
var distance float64 = 10
```

El literal `10` permanece sin tipo hasta que la declaración requiere `float64`. Como su valor puede representarse mediante `float64`, no se necesita la conversión `float64(10)`.

Una constante sin tipo también puede adaptarse a un operando tipado dentro de una expresión:

```go
var unitPrice float64 = 200.3
var total = unitPrice * 5
```

`unitPrice` fija el tipo de la operación como `float64`, y la constante sin tipo `5` puede representarse mediante ese tipo. Por tanto, `total` tiene el tipo `float64`.

Si toda la expresión es constante, puede permanecer sin tipo hasta que la asignación proporcione uno:

```go
var calculated float64 = 200.3 * 5
```

Sin un tipo de destino, Go elige los tipos predeterminados:

```go
var count = 10       // int
var ratio = 200.3    // float64
```

Los literales también pueden inicializar tipos definidos compatibles:

```go
type Score int16

var score Score = 100
```

El compilador aplica este proceso de decisión:

```text
literal básico
     │
     └── constante sin tipo
             │
             ├── existe un contexto compatible
             │       ├── valor representable → usa el tipo contextual
             │       └── valor no representable → error de compilación
             │
             └── no hay tipo contextual → usa el tipo predeterminado
```

Las constantes numéricas tienen valores exactos mientras siguen siendo constantes. Cuando un valor se convierte en un valor concreto de punto flotante, se representa con la precisión limitada de ese tipo.

## 4. Ejemplos, diferencias y errores comunes

| Declaración o expresión | ¿Válida? | Motivo |
|---|---:|---|
| `var value float64 = 10` | ✅ | `10` puede representarse mediante `float64` |
| `var value = 10` | ✅ | Sin otro contexto, `10` recibe el tipo predeterminado `int` |
| `var small byte = 255` | ✅ | `255` se encuentra dentro del rango de `byte` |
| `var small byte = 1000` | ❌ | `1000` se encuentra fuera del rango de `byte` |
| `var whole int = 10.0` | ✅ | El valor exacto de la constante es el entero `10` |
| `var whole int = 10.5` | ❌ | `10.5` no puede representarse mediante `int` |
| `var label string = 10` | ❌ | Una constante numérica no es compatible con `string` |

> ⚠️ **Corrección importante:** Un literal de punto flotante no es rechazado por un destino entero solamente porque contiene un punto decimal. `10.0` tiene un valor integral exacto y puede inicializar un `int`; `10.5` no puede hacerlo.

> **No confundir:** una constante sin tipo es flexible en tiempo de compilación; no es un valor con tipado dinámico que cambie de tipo durante la ejecución.

Una variable tipada no se adapta de la misma manera que un literal:

```go
var count int = 10
var measurement float64 = 30.2

var valid = measurement + 10
// var invalid = measurement + count // tipos distintos: float64 e int
```

El `10` sin tipo puede usarse como `float64`, pero la variable tipada `count` sigue siendo un `int`. Para combinar `measurement` y `count` es necesario elegir y escribir una conversión explícita.

> **No confundir:** la compatibilidad y el rango son comprobaciones distintas. `1000` es numérico y, por tanto, compatible en clase con `byte`, pero su valor no puede representarse mediante `byte`.

Las asignaciones desde constantes sin tipo se comprueban en tiempo de compilación. Esto impide que una constante incompatible o fuera de rango se convierta silenciosamente en un valor concreto.

## 5. Chuleta rápida

- Los literales básicos enteros, de punto flotante, `rune` y `string` representan constantes sin tipo.
- El contexto puede proporcionar un tipo concreto compatible.
- Sin contexto → usa el tipo predeterminado; entero → `int`, punto flotante → `float64`.
- El valor de un literal debe poder representarse mediante el tipo de destino.
- `var value float64 = 10` → válido sin una conversión explícita.
- `measurement * 5` → `5` puede adaptarse al tipo de `measurement`.
- Los valores tipados de tipos numéricos distintos normalmente requieren una conversión explícita.
- Los tipos definidos pueden inicializarse mediante literales compatibles y representables.
- ⚠️ `var whole int = 10.0` es válido; `var whole int = 10.5` no lo es.
- ⚠️ `var small byte = 1000` falla porque el valor está fuera de rango.
- Las constantes numéricas y `string` no son intercambiables.
- `untyped` significa flexibilidad en tiempo de compilación, no ausencia de reglas del sistema de tipos.

### Comprueba que realmente lo sabes

1. ¿Por qué `measurement + 5` compila cuando `measurement` es un `float64`, mientras que sumar una variable de tipo `int` normalmente no compila?
2. ¿Por qué `10.0` puede inicializar un `int`, pero `10.5` no puede hacerlo?
3. ¿Qué sucede cuando un literal sin tipo no tiene un tipo contextual y en qué se diferencia esto de asignarlo a un tipo definido como `Score`?
