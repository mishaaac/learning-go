# `for` infinito en Go

`for { ... }` es la forma explícita de bucle infinito en Go. Al no existir una condición que pueda volverse falsa, la ejecución se repite hasta que algo dentro del cuerpo sale o la ejecución se detiene externamente.

## 1. Qué debo entender

- La sintaxis es `for { ... }`.
- Omitir la condición equivale a una condición siempre verdadera.
- Esta forma no tiene init, condition, post ni puntos y coma.
- Un bucle infinito práctico suele contener una salida como `break` o `return`.
- Sin una salida interna, el bucle solo termina por un evento externo o por la terminación del programa.

## 2. Conceptos clave

| Forma | Control en la cabecera |
|---|---|
| `for init; condition; post` | Inicialización, prueba y actualización |
| `for condition` | Solo prueba |
| `for {}` | Sin prueba en la cabecera; se repite indefinidamente |

Un timeout externo o `Ctrl-C` detiene el programa, no la lógica propia del bucle.

## 3. Cómo funciona en Go

```go
for {
    fmt.Println("Hello")
}
```

Una versión más práctica decide cuándo salir:

```go
for {
    value := readValue()
    if value == "stop" {
        break
    }
    process(value)
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un timeout de la plataforma no vuelve finito a `for {}`; interrumpe externamente todo el programa.

Un bucle infinito puede ser intencional cuando un servidor, worker o procesador de eventos espera repetidamente, pero el diseño que lo rodea todavía necesita un camino de cierre o cancelación.

Un bucle vacío también puede consumir CPU continuamente. Si el cuerpo espera input u otra operación bloqueante, la ejecución se comporta de otra forma aunque el bucle siga siendo lógicamente infinito.

Usa `for condition` cuando la regla de continuación pertenezca naturalmente a la cabecera; usa `for {}` cuando las decisiones de salida ocurran dentro del cuerpo.

## 5. Chuleta rápida

- Forma infinita → `for { ... }`
- Sin init
- Sin condition
- Sin post
- Sin puntos y coma
- La condición ausente actúa como `true`
- El cuerpo se repite indefinidamente
- `break` sale del bucle
- `return` sale de la función
- Un timeout externo no es lógica de terminación del bucle
- El código real debe tener un camino deliberado de cierre

### Comprueba que realmente lo sabes

1. ¿Por qué `for {}` es infinito aunque un ejecutor online lo detenga después de unos segundos?
2. ¿Cuándo resulta más claro `for {}` que `for condition`?
3. ¿Qué diferencia a un bucle infinito activo de uno bloqueado esperando input?
