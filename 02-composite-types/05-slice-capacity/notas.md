# Capacidad de slices en Go

Un slice registra tanto su longitud actual como la capacidad disponible desde su posición inicial en el backing array. La capacidad explica cuándo `append` puede reutilizar el almacenamiento y cuándo necesita obtener uno nuevo.

## 1. Qué debo entender

- `len(s)` cuenta los elementos actuales; `cap(s)` mide hasta dónde puede extenderse el slice dentro de su backing array.
- La relación invariable es `0 <= len(s) <= cap(s)`.
- `append` reutiliza el backing array cuando existe capacidad suficiente.
- Si la capacidad es insuficiente, `append` reserva nuevo almacenamiento y copia los elementos existentes.
- Reservar una capacidad razonable puede evitar asignaciones y copias repetidas.

## 2. Conceptos clave

| Concepto | Significado |
|---|---|
| Longitud | Elementos que actualmente pertenecen al slice |
| Capacidad | Elementos disponibles desde el inicio del slice hasta el final de su almacenamiento subyacente |
| Backing array | Almacenamiento que contiene los elementos del slice |
| Reasignación | Nuevo almacenamiento obtenido cuando el crecimiento no cabe |
| Preasignación | Reserva de capacidad antes de varios `append` |

## 3. Cómo funciona en Go

```go
values := make([]int, 0, 4)
fmt.Println(len(values), cap(values)) // 0 4

values = append(values, 10, 20, 30)
fmt.Println(len(values), cap(values)) // 3 4

values = append(values, 40)
fmt.Println(len(values), cap(values)) // 4 4
```

```text
backing array
[10][20][30][  ]
 ↑---- len ----↑
 ↑------ cap -------↑
```

El siguiente `append` necesita más capacidad, por lo que puede mover los elementos a un nuevo backing array.

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** la capacidad no es el número de índices válidos actualmente. Solo se puede acceder directamente a índices menores que `len(s)`.

Para un array, `cap(array) == len(array)`. Para un slice `nil`, ambos valores son `0`.

> ⚠️ **Corrección importante:** los factores exactos de crecimiento de la capacidad son detalles de implementación, no promesas del lenguaje Go. El comportamiento actual del runtime puede cambiar entre versiones de Go, tamaños de elementos y restricciones de asignación; el código no debe depender de capacidades como `1, 2, 4, 8` después de varios `append`.

Preasigna cuando conozcas un tamaño esperado, pero no trates la capacidad como un máximo fijo: `append` aún puede crecer más allá de ella.

## 5. Chuleta rápida

- `len(s)` → elementos actuales
- `cap(s)` → extensión disponible en el almacenamiento subyacente
- Siempre `0 <= len(s) <= cap(s)`
- Los índices válidos directos dependen de `len`
- La capacidad libre permite a `append` reutilizar almacenamiento
- Agotar la capacidad puede provocar una asignación y una copia
- La reasignación puede romper la memoria compartida con slices anteriores
- `make([]T, 0, n)` reserva capacidad para `append`
- `cap(nilSlice)` → `0`
- Los factores exactos de crecimiento no son garantías de la API

### Comprueba que realmente lo sabes

1. ¿Por qué un slice puede tener una capacidad mayor que su longitud sin permitir indexar toda esa capacidad?
2. ¿Qué trabajo puede ocurrir cuando `append` supera la capacidad actual?
3. ¿Por qué un programa no debe depender de una secuencia observada de valores de capacidad?
