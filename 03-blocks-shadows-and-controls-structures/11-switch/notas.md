# `switch` en Go

Un expression `switch` selecciona el primer `case` coincidente y se detiene automáticamente después de esa cláusula. Go admite varias expresiones por case, inicialización local y `fallthrough` explícito, pero los cases normales no necesitan `break`.

## 1. Qué debo entender

- `switch expression` compara una expresión con las expresiones de los cases.
- Los cases se consideran en orden de aparición y se ejecuta la primera coincidencia.
- Varios valores pueden compartir un case separándolos mediante comas.
- Los cases no hacen fall-through automáticamente.
- Una sentencia inicial opcional puede declarar un valor cuyo scope cubre todo el switch.

## 2. Conceptos clave

| Elemento | Comportamiento |
|---|---|
| `case 1, 2, 3:` | Cualquier valor listado selecciona la cláusula |
| `default:` | Se ejecuta cuando ningún case coincide |
| Case vacío | No realiza ninguna acción |
| `fallthrough` | Continúa incondicionalmente en el cuerpo de la siguiente cláusula |
| `break` | Termina el `switch`, `for` o `select` contenedor más cercano |

## 3. Cómo funciona en Go

```go
switch size := len(word); size {
case 1, 2, 3, 4:
    fmt.Println("short")
case 5:
    description := "exact"
    fmt.Println(description)
case 6, 7, 8, 9:
    // Intencionalmente no se hace nada.
default:
    fmt.Println("long")
}
```

`size` está disponible en la expresión del switch y en todas las cláusulas, mientras que `description` pertenece solo al bloque de su case.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un case vacío no hace nada; no continúa en el siguiente case.

`fallthrough` debe ser la última sentencia no vacía de un case que no sea el último. Transfiere el control al cuerpo de la siguiente cláusula sin comprobar si coincide su expresión, por lo que debe ser poco frecuente y deliberado.

> ⚠️ **Corrección importante:** un expression switch no está limitado a enteros, pero su expresión y los valores de los cases deben admitir una comparación de igualdad válida. Los channels son comparables; los slices, maps y funciones no.

Dentro de un `switch` anidado en un bucle, un `break` sin label sale del switch. Usa `break loopLabel` para salir de un bucle exterior etiquetado explícitamente.

## 5. Chuleta rápida

- Forma → `switch expression { ... }`
- No se necesitan paréntesis en la expresión
- Se ejecuta el primer case coincidente
- Varios valores → lista de case separada por comas
- `default` maneja la falta de coincidencia
- Los cases no hacen fall-through automáticamente
- Case vacío → ninguna acción
- `fallthrough` entra explícitamente al siguiente cuerpo
- Sentencia inicial → `switch statement; expression`
- Las declaraciones de un case permanecen en su bloque
- `break` sin label sale de la estructura más cercana

### Comprueba que realmente lo sabes

1. ¿Por qué normalmente no hace falta `break` al final de un case de Go?
2. ¿En qué se diferencia `fallthrough` de coincidir normalmente con el siguiente case?
3. ¿Por qué un `break` dentro de un switch anidado en un bucle no sale necesariamente del bucle?
