# Vaciar un map con `clear`

Para un map, `clear` elimina todas las entradas en una operación. El map sigue siendo utilizable si estaba inicializado, pero su longitud pasa a cero.

## 1. Qué debo entender

- `clear(m)` elimina todas las entradas key-value.
- Después, `len(m) == 0`.
- Un map inicializado permanece inicializado y admite escritura.
- Llamar a `clear` sobre un map `nil` es un no-op seguro.
- `clear` se comporta de forma diferente con maps y slices.

## 2. Conceptos clave

| Necesidad | Operación | Resultado |
|---|---|---|
| Eliminar una entrada del map | `delete(m, key)` | Solo se elimina esa key |
| Eliminar todas las entradas del map | `clear(m)` | `len(m) == 0` |
| Restablecer elementos de un slice | `clear(s)` | Los elementos pasan a zero values; la longitud se conserva |

`clear` se convirtió en built-in en Go 1.21.

## 3. Cómo funciona en Go

```go
scores := map[string]int{
    "hello": 5,
    "world": 10,
}

clear(scores)
fmt.Println(len(scores)) // 0

scores["new"] = 1       // todavía admite escritura
fmt.Println(scores)      // map[new:1]
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `clear(map)` elimina entradas, mientras que `clear(slice)` conserva los elementos y sustituye sus valores por zero values.

Para un map `nil`:

```go
var scores map[string]int
clear(scores) // seguro; scores sigue siendo nil
```

Limpiar un map inicializado no lo convierte en `nil`. Si la distinción importa, `m = nil` es una operación diferente; las escrituras futuras exigirían volver a inicializarlo.

Usa `delete` al seleccionar una key y `clear` cuando deban eliminarse todas las entradas actuales.

## 5. Chuleta rápida

- Eliminar todas las entradas → `clear(m)`
- Resultado → `len(m) == 0`
- El map inicializado sigue admitiendo escritura
- El map inicializado no se convierte en `nil`
- `clear(nilMap)` → no-op
- Una key → `delete(m, key)`
- Todas las keys → `clear(m)`
- `clear(map)` elimina entradas
- `clear(slice)` pone elementos en cero sin cambiar la longitud

### Comprueba que realmente lo sabes

1. ¿Se puede escribir inmediatamente en un map inicializado después de `clear`? ¿Por qué?
2. ¿En qué se diferencia `clear` para un map y para un slice?
3. ¿En qué se diferencia `clear(m)` de asignar `m = nil`?
