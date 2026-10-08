# `for` con solo condición en Go

Un `for` con solo una condición es el equivalente en Go de un bucle `while` convencional. La inicialización y las actualizaciones viven fuera o dentro del cuerpo, mientras que la cabecera solo expresa si la iteración debe continuar.

## 1. Qué debo entender

- La sintaxis es `for condition { ... }`.
- La condición se comprueba antes de cada iteración.
- La condición debe producir un `bool`.
- No se escriben puntos y coma cuando solo queda la condición.
- Normalmente, el cuerpo debe cambiar el estado para que la condición pueda volverse falsa.

## 2. Conceptos clave

| Bucle de tres partes | Bucle con solo condición |
|---|---|
| `for init; condition; post` | `for condition` |
| Init suele estar en la cabecera | La inicialización ocurre antes |
| La actualización suele estar en la cabecera | La actualización ocurre en el cuerpo |
| Usa puntos y coma | No usa puntos y coma |

## 3. Cómo funciona en Go

```go
value := 1

for value < 100 {
    fmt.Println(value)
    value *= 2
}
```

Los valores impresos son `1`, `2`, `4`, `8`, `16`, `32` y `64`. Después de que la actualización produzca `128`, la siguiente comprobación de la condición es falsa.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `for condition` comprueba antes del cuerpo, por lo que este puede ejecutarse cero veces.

Esta forma es más clara que una cláusula de tres partes vacía:

```go
for value < limit {
    value = nextValue(value)
}
```

Si el cuerpo nunca cambia algo que influya en la condición, el bucle puede ser infinito. Las actualizaciones ocultas en varias ramas también dificultan comprobar la terminación.

Usa un `for` de tres partes cuando la inicialización y un paso post regular formen parte natural de la cabecera del bucle.

## 5. Chuleta rápida

- Sintaxis → `for condition { ... }`
- Cumple el papel de `while`
- La condición se comprueba antes del cuerpo
- El cuerpo puede ejecutarse cero veces
- La condición debe ser booleana
- Sin puntos y coma en la cabecera
- Inicializa antes del bucle
- La actualización ocurre normalmente en el cuerpo
- El estado debe progresar hacia la terminación
- Prefiere el `for` de tres partes para contadores regulares

### Comprueba que realmente lo sabes

1. ¿Por qué un bucle con solo condición puede ejecutarse cero veces?
2. ¿Qué responsabilidad pasa al cuerpo al eliminar `init` y `post`?
3. ¿Cuándo resulta más claro un `for` de tres partes que uno con solo condición?
