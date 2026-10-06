# Slicing de slices en Go

Una slice expression crea una nueva vista sobre una parte de un almacenamiento existente. Su sintaxis compacta es útil, pero los elementos y la capacidad compartidos obligan a razonar sobre las mutaciones posteriores, especialmente `append`.

## 1. Qué debo entender

- `s[low:high]` incluye `low` y excluye `high`.
- Omitir `low` significa `0`; omitir `high` significa `len(s)`.
- El slicing no copia elementos; el resultado normalmente comparte el backing array.
- Un subslice puede tener capacidad más allá de los elementos que muestra actualmente.
- `s[low:high:max]` limita la capacidad resultante a `max-low`.

## 2. Conceptos clave

| Expresión | Longitud | Comportamiento de la capacidad |
|---|---:|---|
| `s[:2]` | `2` | Puede extenderse más allá del índice `2` |
| `s[1:]` | `len(s)-1` | Se extiende por la capacidad restante |
| `s[:]` | `len(s)` | Comparte el mismo rango visible |
| `s[low:high:max]` | `high-low` | Limitada a `max-low` |

## 3. Cómo funciona en Go

```go
letters := []string{"a", "b", "c", "d"}
left := letters[:2]
right := letters[1:3]

left[1] = "B"
fmt.Println(letters) // [a B c d]
fmt.Println(right)   // [B c]
```

```text
backing array: [a][B][c][d]
                └ left ┘
                   └ right ┘
```

Una vista con capacidad limitada se escribe así:

```go
left = letters[:2:2] // len 2, cap 2
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `sub := original[:2]` crea una vista; no copia los primeros dos elementos.

Con `sub := original[:2]`, `append(sub, value)` puede sobrescribir `original[2]` cuando existe capacidad compartida libre. Limitar la capacidad con `sub := original[:2:2]` obliga a que un `append` más allá de la longitud dos reserve almacenamiento separado.

La expresión completa controla el crecimiento futuro, no la memoria compartida de los elementos del rango actual. Incluso después de `s[:2:2]`, asignar `sub[0]` todavía modifica el elemento compartido.

Los límites inválidos se rechazan durante la compilación cuando son constantes demostrables o causan un `panic` en runtime.

## 5. Chuleta rápida

- `s[low:high]` → intervalo `[low, high)`
- `s[:high]` → comienza en `0`
- `s[low:]` → termina en `len(s)`
- `s[:]` → todo el slice visible
- El slicing normalmente comparte elementos
- Mutar elementos compartidos es visible mediante otras vistas
- La capacidad de un subslice puede superar su longitud
- `append` puede sobrescribir elementos posteriores compartidos
- `s[low:high:max]` → capacidad `max-low`
- Limitar la capacidad afecta a `append`, no a los elementos compartidos actuales

### Comprueba que realmente lo sabes

1. ¿Por qué cambiar un subslice puede cambiar el slice original?
2. ¿En qué condición añadir a un subslice puede sobrescribir un elemento posterior del original?
3. ¿Qué controla el tercer índice de `s[low:high:max]`?
