# `switch` sin expresión en Go

Un switch sin expresión coloca una condición booleana en cada `case` en vez de comparar un valor con varias alternativas. Se comporta como un switch sobre `true` y resulta útil para rangos o predicados relacionados.

## 1. Qué debo entender

- La sintaxis es `switch { case condition: ... }`.
- La expresión de cada case debe ser booleana.
- Los cases se comprueban en orden y solo se ejecuta el primero que sea verdadero.
- Una sentencia inicial puede calcular un valor usado por todos los cases.
- Usa un expression switch cuando todos los cases sean comparaciones de igualdad contra el mismo valor.

## 2. Conceptos clave

| Pregunta expresada | Forma más adecuada |
|---|---|
| “¿Qué valor es igual a `x`?” | `switch x` |
| “¿Qué condición relacionada es verdadera?” | `switch {}` |
| “Ninguna condición anterior coincidió” | `default` |

Un switch sin expresión equivale conceptualmente a `switch true`.

## 3. Cómo funciona en Go

```go
switch length := len(word); {
case length < 5:
    fmt.Println("short")
case length > 10:
    fmt.Println("long")
default:
    fmt.Println("medium")
}
```

`length` se calcula una vez y permanece disponible en todas las cláusulas de ese switch.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `switch { case x == 2: ... }` funciona, pero `switch x { case 2: ... }` expresa más directamente la igualdad repetida.

El orden importa cuando las condiciones se superponen. Coloca una condición más específica antes de otra más amplia que también coincidiría:

```go
switch {
case value%3 == 0 && value%5 == 0:
    fmt.Println("FizzBuzz")
case value%3 == 0:
    fmt.Println("Fizz")
}
```

Un switch sin expresión resulta más claro cuando sus condiciones son predicados diferentes sobre una misma decisión, no comprobaciones sin relación reunidas únicamente porque la sintaxis lo permite.

## 5. Chuleta rápida

- Forma → `switch { ... }`
- Conceptualmente → `switch true`
- Los cases contienen condiciones booleanas
- Se ejecuta el primer case verdadero
- El orden importa para condiciones superpuestas
- `default` maneja el resto
- Se permite una sentencia inicial
- Predicados relacionados diferentes → switch sin expresión
- Misma expresión comparada con valores → expression switch
- Evita agrupar condiciones sin relación

### Comprueba que realmente lo sabes

1. ¿Por qué la condición superpuesta más específica debe aparecer primero?
2. ¿Cuándo comunica mejor la intención `switch x` que `switch {}`?
3. ¿Qué relación deberían tener las condiciones de un switch sin expresión?
