# `append` en Go

`append` añade elementos al final de un slice y devuelve el slice resultante. Conservar ese valor de retorno es esencial porque el crecimiento puede cambiar la longitud, la capacidad y el backing array del slice.

## 1. Qué debo entender

- `append` recibe un slice de destino seguido de uno o más elementos.
- Funciona directamente con un slice `nil`.
- Devuelve el slice actualizado; el patrón habitual es `s = append(s, value)`.
- `source...` expande otro slice para poder añadir sus elementos.
- Go pasa el valor del slice a `append` por valor, por lo que debe devolver el descriptor resultante.

## 2. Conceptos clave

| Forma | Efecto |
|---|---|
| `append(s, value)` | Añade un elemento |
| `append(s, a, b)` | Añade varios elementos |
| `append(s, other...)` | Añade cada elemento de `other` |
| `append(bytes, text...)` | Añade los bytes de un string a un `[]byte` |
| `s = append(...)` | Conserva el slice devuelto |

## 3. Cómo funciona en Go

```go
var numbers []int

numbers = append(numbers, 10)
numbers = append(numbers, 20, 30)

more := []int{40, 50}
numbers = append(numbers, more...)

fmt.Println(numbers) // [10 20 30 40 50]
```

Conceptualmente:

```text
valor slice → append → valor slice actualizado
                     ├── mayor len
                     └── posible nuevo backing array
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `append(s, other)` intenta añadir `other` como un elemento; `append(s, other...)` añade sus elementos.

Ignorar el valor de retorno es inválido como llamada aislada y, más importante, perdería los cambios del descriptor del slice:

```go
numbers = append(numbers, 60) // correcto
```

Si hay capacidad suficiente, `append` puede reutilizar el backing array existente. En caso contrario, reserva un almacenamiento mayor y copia los elementos existentes; por eso otros slices que compartían el almacenamiento anterior podrían no observar cambios posteriores.

> ⚠️ **Corrección importante:** call-by-value forma parte de la explicación, pero no significa que los elementos del slice se copien al pasar el slice. El descriptor copiado sigue haciendo referencia al mismo backing array hasta que una operación como el crecimiento lo sustituye.

## 5. Chuleta rápida

- `append` añade elementos al final
- Funciona sobre un slice `nil`
- Conserva siempre el resultado
- Un valor → `s = append(s, value)`
- Varios valores → `s = append(s, a, b)`
- Otro slice → `s = append(s, other...)`
- `...` expande los elementos del slice fuente
- `append` aumenta `len`
- Puede reutilizar la capacidad existente
- Una capacidad insuficiente puede producir un nuevo backing array

### Comprueba que realmente lo sabes

1. ¿Por qué debe asignarse o usarse de otra forma el resultado de `append`?
2. ¿Qué cambia `...` al añadir otro slice?
3. ¿Cómo puede una asignación de memoria durante `append` cambiar la relación entre dos slices?
