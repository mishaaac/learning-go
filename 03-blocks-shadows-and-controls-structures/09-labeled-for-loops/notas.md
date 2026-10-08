# Labels en bucles `for` de Go

Los labels permiten que `break` o `continue` actúen sobre una estructura de control contenedora en vez de sobre la más cercana. Son especialmente útiles en bucles anidados cuando una decisión interior determina qué debe ocurrir con el bucle exterior.

## 1. Qué debo entender

- Un label es un identificador seguido de `:` antes de una sentencia.
- `break` y `continue` sin label actúan sobre la estructura aplicable más cercana.
- `continue label` inicia la siguiente iteración del `for` etiquetado.
- `break label` termina el `for`, `switch` o `select` etiquetado cuando corresponda.
- Un label de `continue` debe identificar un bucle `for` contenedor.

## 2. Conceptos clave

| Sentencia | Efecto dentro de bucles anidados |
|---|---|
| `continue` | Siguiente iteración del bucle interior |
| `continue outer` | Siguiente iteración del bucle exterior etiquetado |
| `break` | Sale de la estructura aplicable más cercana |
| `break outer` | Sale de la estructura exterior etiquetada |

## 3. Cómo funciona en Go

```go
samples := []string{"hello", "apple"}

outer:
for _, sample := range samples {
    for _, r := range sample {
        if r == 'l' {
            continue outer
        }
        fmt.Printf("%c", r)
    }
    fmt.Println(" accepted")
}
```

Cuando un rune interior es `'l'`, la ejecución omite el resto del trabajo interior y de la iteración exterior actual.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `continue` sin label afecta al bucle interior; `continue outer` afecta al bucle etiquetado explícitamente.

Un patrón habitual rechaza un elemento exterior en cuanto un valor interior es inválido:

```go
outer:
for _, item := range items {
    for _, part := range item.Parts {
        if !valid(part) {
            continue outer
        }
    }
    accept(item)
}
```

Los labels deben aclarar qué bucle se selecciona. Si el flujo anidado sigue siendo difícil de seguir, extraer la lógica a una función puede ser más claro que añadir más labels.

## 5. Chuleta rápida

- Sintaxis de label → `outer:`
- Coloca el label inmediatamente antes de la sentencia objetivo
- `continue` simple → `for` más cercano
- `continue outer` → siguiente iteración exterior
- `break` simple → estructura aplicable más cercana
- `break outer` → salir de la estructura etiquetada
- El label de `continue` debe nombrar un `for` contenedor
- El control etiquetado se usa principalmente en bucles anidados
- También omite el código restante del cuerpo exterior
- Prefiere labels descriptivos

### Comprueba que realmente lo sabes

1. ¿Qué código se omite cuando un bucle interior ejecuta `continue outer`?
2. ¿Por qué `continue` solo puede seleccionar un `for` etiquetado?
3. ¿Cuándo puede ser más claro extraer una función que usar un label?
