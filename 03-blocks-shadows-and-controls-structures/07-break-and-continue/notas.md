# `break` y `continue` en Go

`break` termina un bucle, mientras que `continue` omite el resto de la iteración actual y comienza la siguiente. Usadas con cuidado, ambas sentencias hacen explícitas las condiciones de salida y reducen el anidamiento innecesario.

## 1. Qué debo entender

- Dentro de un bucle, `break` sale de la estructura contenedora aplicable más cercana.
- `continue` inicia la siguiente iteración del `for` contenedor más cercano.
- En un bucle de tres partes, `continue` todavía pasa por la sentencia post antes de volver a comprobar.
- `for { ...; if !condition { break } }` puede modelar un patrón `do/while`.
- Los `continue` tempranos pueden mantener casos independientes alineados y legibles.

## 2. Conceptos clave

| Sentencia | Efecto |
|---|---|
| `break` | Abandona por completo el bucle actual |
| `continue` | Omite el resto del cuerpo de esta iteración |
| `return` | Abandona toda la función |
| Forma etiquetada | Selecciona explícitamente un bucle contenedor |

## 3. Cómo funciona en Go

```go
for i := 1; i <= 10; i++ {
    if i == 8 {
        break
    }
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```

Esto imprime los valores impares menores que `8`. Una repetición que ejecuta primero el cuerpo puede escribirse así:

```go
for {
    doWork()
    if !shouldContinue() {
        break
    }
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `break` termina el bucle; `continue` termina solo la iteración actual.

Al traducir `do { ... } while (condition)`, la prueba de salida de Go se niega porque pregunta cuándo detenerse:

```go
if !condition {
    break
}
```

Varias comprobaciones tempranas con `continue` pueden ser más claras que ramas `if`/`else` profundamente anidadas, pero un exceso de saltos también puede fragmentar el flujo. Cada sentencia debe dejar evidente el siguiente punto de ejecución.

Un `break` sin label dentro de un `switch` anidado en un bucle sale del `switch`, no del bucle; los labels se estudian por separado.

## 5. Chuleta rápida

- `break` → salir del bucle
- `continue` → siguiente iteración
- `return` → salir de la función
- `continue` omite el resto del cuerpo
- En un bucle de tres partes, `continue` pasa a `post`
- Patrón con cuerpo primero → `for { ...; if !condition { break } }`
- La negación expresa cuándo detenerse
- Un `continue` temprano puede reducir el anidamiento
- El control sin label actúa sobre la estructura aplicable más cercana
- Usa labels cuando debas seleccionar un bucle exterior

### Comprueba que realmente lo sabes

1. En un bucle de tres partes, ¿qué ocurre después de `continue` y antes de la siguiente comprobación?
2. ¿Por qué se niega la condición al modelar `do/while` mediante `break`?
3. ¿Cómo puede `continue` reducir el anidamiento sin cambiar el resultado?
