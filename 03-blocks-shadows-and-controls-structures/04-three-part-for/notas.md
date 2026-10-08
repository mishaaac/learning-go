# El `for` de tres partes en Go

El `for` de tres partes de Go muestra en una cabecera las reglas de inicialización, continuación y actualización. Es útil cuando una iteración tiene límites numéricos o posicionales claros.

## 1. Qué debo entender

- La forma es `for init; condition; post { ... }`.
- `init` se ejecuta una vez antes de la primera comprobación de la condición.
- `condition` se comprueba antes de cada iteración y debe ser booleana.
- `post` se ejecuta después de cada iteración completada y antes de volver a comprobar la condición.
- Las variables declaradas en `init` tienen como scope todo el bucle, incluido su cuerpo.

## 2. Conceptos clave

| Parte | Ejemplo | Momento de ejecución |
|---|---|---|
| Init | `i := 0` | Una vez, antes del bucle |
| Condición | `i < 10` | Antes de cada iteración |
| Post | `i++` | Después de cada iteración completada |

Se pueden omitir una o más partes, pero los puntos y coma permanecen al usar una cláusula `for`.

## 3. Cómo funciona en Go

```go
for i := 0; i < 3; i++ {
    fmt.Println(i)
}
```

```text
init → condición ─false→ terminar
          │true
          ▼
         cuerpo
          ▼
         post ──────────┘
```

La inicialización puede estar fuera del bucle:

```go
i := 0
for ; i < 3; i++ {
    fmt.Println(i)
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `init` se ejecuta una vez; `post` se ejecuta repetidamente después de las iteraciones.

`init` y `post` son simple statements. Una declaración `var` no es válida directamente en `init`; usa una short declaration o declara antes la variable.

> ⚠️ **Corrección importante:** la sentencia `post` no puede ser una short variable declaration. `i++`, las asignaciones y las llamadas a funciones pueden ser válidas, pero `j := i + 1` no está permitido allí.

Si la actualización es compleja, omite `post` y actualiza dentro del cuerpo. Asegúrate de que todos los caminos sigan progresando o el bucle podría no terminar.

## 5. Chuleta rápida

- Forma → `for init; condition; post`
- Sin paréntesis alrededor
- Init se ejecuta una vez
- La condición se ejecuta antes de cada iteración
- Post se ejecuta después de cada iteración completada
- La condición debe producir `bool`
- La variable de init vive durante todo el bucle
- Se pueden omitir partes
- Los puntos y coma permanecen en una cláusula de tres partes
- `var` no es una simple statement válida en init
- Post no puede usar `:=`

### Comprueba que realmente lo sabes

1. ¿En qué orden se ejecutan la condición, el cuerpo y post?
2. ¿Por qué `j := i + 1` es inválido como sentencia post?
3. ¿Qué riesgo aparece al mover la actualización al cuerpo del bucle?
