# Debugging Analysis

## Problem I Found

El programa no compila porque `goto summary` intenta saltar por encima de la declaración `result := "processed"`.

## Why It Happens

Go no permite que un `goto` omita una declaración cuando la variable declarada estaría en alcance en la etiqueta de destino. Si el salto fuera válido, `result` se usaría en `summary` sin haber ejecutado su inicialización.

## What I Changed

Inicialicé `result` con `"processed"` antes del `if`. Cuando `jobStatus` es `"pending"`, el bloque asigna `"skipped"`. Eliminé el salto porque una condición directa comunica las dos posibilidades sin alterar el flujo normal.

## Why the Fix Works

`result` siempre se inicializa antes del resumen único. El `if` modifica la misma variable cuando corresponde, por lo que `pending` produce `skipped` y cualquier otro estado conserva `processed`. Ya no existe un salto que cruce una declaración.

## Verification

Verifiqué las dos ramas requeridas:

```text
pending skipped
```

Después cambié temporalmente `jobStatus` a `"active"` y comprobé:

```text
active processed
```
