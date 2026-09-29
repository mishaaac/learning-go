# Booleanos en Go

El tipo predeclarado `bool` representa valores de verdad en Go. Sus dos valores posibles y su valor cero definido hacen que las variables booleanas sean predecibles incluso cuando no tienen un inicializador explícito.

## 1. Qué debo entender

- Una variable de tipo `bool` solo puede contener `true` o `false`.
- `true` y `false` son las constantes booleanas predeclaradas de Go.
- El **valor cero** de `bool` es `false`.
- `var flag bool` crea un `bool` cuyo valor inicial es automáticamente `false`.
- En `var isAwesome = true`, Go infiere el tipo de la variable como `bool` a partir de su inicializador.

## 2. Conceptos clave

| Concepto | Significado | Por qué importa |
|---|---|---|
| `bool` | El tipo booleano predeclarado de Go | Representa una condición con dos estados posibles |
| `true` | La constante booleana verdadera | Representa una condición que se cumple |
| `false` | La constante booleana falsa | Representa una condición que no se cumple |
| Valor cero | El valor usado cuando un `bool` no tiene un inicializador explícito | Un `bool` declarado siempre comienza con un valor definido |
| Inferencia de tipo | Go determina el tipo de una variable a partir de su inicializador | `var isAwesome = true` produce una variable de tipo `bool` |

Los dos valores son valores del mismo tipo; no crean tipos separados.

## 3. Cómo funciona en Go

Un tipo explícito sin un inicializador usa el valor cero:

```go
var flag bool
```

`flag` tiene el tipo `bool` y comienza como `false`. Escribir `= false` produciría el mismo valor inicial:

```go
var firstFlag bool
var secondFlag bool = false
```

Ambas variables son `false`. La primera forma, más corta, es suficiente cuando el valor cero expresa el estado inicial deseado.

Cuando existe un inicializador pero se omite el tipo, Go determina el tipo a partir de ese valor:

```go
var isAwesome = true
var isReady = false
```

Ambas variables tienen el tipo `bool`. Las constantes `true` y `false` tienen el tipo predeterminado `bool` cuando se necesita un tipo concreto y el contexto no proporciona otro tipo booleano.

```text
var flag bool
     │    │
     │    └── tipo explícito
     └─────── recibe false como valor cero

var isAwesome = true
     │           │
     │           └── inicializador
     └────────────── tipo inferido como bool
```

## 4. Ejemplos, diferencias y errores comunes

| Declaración | Tipo | Valor inicial | Cómo se determina |
|---|---|---|---|
| `var flag bool` | `bool` | `false` | El tipo es explícito; el valor es el valor cero |
| `var isAwesome = true` | `bool` | `true` | El valor es explícito; el tipo se infiere |
| `var isReady bool = true` | `bool` | `true` | Tanto el tipo como el valor son explícitos |

> **No confundir:** `false` es un valor booleano; `"false"` es un `string` que contiene texto.

```go
var flag = false
var label = "false"
```

`flag` tiene el tipo `bool`, mientras que `label` tiene el tipo `string`.

Una declaración sin inicializador no es una variable sin inicializar:

```go
var enabled bool
```

`enabled` ya es válida y tiene el valor `false`. Añadir `= false` es opcional y solo debe hacerse cuando aclare la intención.

La diferencia entre `var` y el operador de declaración corta `:=` pertenece al tema posterior sobre declaraciones de variables; no es necesaria para entender el valor cero de `bool`.

## 5. Chuleta rápida

- `bool` → tipo booleano predeclarado de Go.
- Valores posibles → `true` y `false`.
- Valor cero de `bool` → `false`.
- `var flag bool` → tipo `bool`, valor inicial `false`.
- `var isAwesome = true` → tipo inferido como `bool`, valor `true`.
- `= false` es innecesario cuando el valor cero es el estado inicial deseado.
- ⚠️ `false` es un `bool`; `"false"` es un `string`.
- Los detalles de `var` frente a `:=` corresponden a otro tema sobre declaraciones.

### Comprueba que realmente lo sabes

1. ¿Qué valor asigna `var flag bool` a `flag` y de dónde procede ese valor?
2. ¿Cómo determina Go el tipo y el valor de `var isAwesome = true`?
3. ¿Por qué `false` y `"false"` no son intercambiables aunque se parezcan?
