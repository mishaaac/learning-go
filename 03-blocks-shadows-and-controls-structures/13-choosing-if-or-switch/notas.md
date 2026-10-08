# Elegir entre `if` y `switch` en Go

Tanto una cadena `if` como un switch sin expresión pueden evaluar condiciones booleanas. La elección es principalmente comunicativa: un switch presenta alternativas relacionadas como una sola decisión, mientras que `if` suele ser más claro para comprobaciones separadas o un flujo asimétrico.

## 1. Qué debo entender

- `if`/`else if` y `switch {}` pueden expresar secuencias de condiciones equivalentes.
- Un switch indica que sus cases son alternativas relacionadas.
- Un `if` es natural para una condición, guard clauses o comprobaciones menos relacionadas.
- Ambas formas eligen la primera rama exitosa cuando se escriben como una sola cadena.
- Refactoriza cuando ninguna forma pueda presentar la lógica de manera coherente.

## 2. Conceptos clave

| Situación | Normalmente más claro |
|---|---|
| Una condición principal | `if` |
| Salida temprana o guard | `if` |
| Varias alternativas relacionadas | `switch` |
| Caso restante claro | `default` en `switch` |
| Responsabilidades sin relación | `if` separados o refactorización |

## 3. Cómo funciona en Go

```go
switch {
case value%3 == 0 && value%5 == 0:
    fmt.Println("FizzBuzz")
case value%3 == 0:
    fmt.Println("Fizz")
case value%5 == 0:
    fmt.Println("Buzz")
default:
    fmt.Println(value)
}
```

Estos cases responden una pregunta relacionada sobre el mismo valor, por lo que el switch hace visibles las alternativas sin repetir sentencias `continue`.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** un switch no es automáticamente mejor solo porque haya varias condiciones; los cases deben pertenecer a una misma decisión conceptual.

Un guard con `if` suele ser más directo:

```go
if err != nil {
    return err
}
```

El orden de los cases importa cuando los predicados se superponen. En FizzBuzz, la divisibilidad entre ambos valores debe comprobarse antes que la divisibilidad entre cada uno por separado.

Varios `if` independientes pueden ejecutarse todos; una cadena `if`/`else if` o un switch seleccionan solo una rama. Elige según el comportamiento necesario, no solo por el estilo visual.

## 5. Chuleta rápida

- Una condición → normalmente `if`
- Guard clause → `if`
- Alternativas relacionadas → `switch`
- Un switch sin expresión maneja predicados booleanos
- Se ejecuta el primer case coincidente
- `default` hace explícito el resto
- Los cases superpuestos exigen ordenar con cuidado
- Varios `if` separados pueden ejecutarse todos
- `if`/`else if` selecciona una sola rama
- Cases sin relación sugieren `if` o refactorización

### Comprueba que realmente lo sabes

1. ¿Por qué FizzBuzz se comunica bien mediante un switch sin expresión?
2. ¿En qué se diferencian varios `if` separados de una cadena `if`/`else if`?
3. ¿Qué indica que un switch propuesto debería refactorizarse?
