# Eliminar elementos de un map en Go

El built-in `delete` elimina la entrada asociada a una key de un map. Está diseñado para ser seguro cuando la key está ausente o el map es `nil`.

## 1. Qué debo entender

- La sintaxis es `delete(m, key)`.
- Si la key existe, se elimina toda su entrada key-value.
- Si la key está ausente, la operación no hace nada.
- Si el map es `nil`, la operación tampoco hace nada.
- `delete` no devuelve ningún valor.

## 2. Conceptos clave

| Situación | Resultado de `delete(m, key)` |
|---|---|
| La key existe | Se elimina la entrada |
| La key está ausente | No-op |
| El map es `nil` | No-op |
| Se necesita el valor anterior | Se lee antes de eliminar |

## 3. Cómo funciona en Go

```go
scores := map[string]int{
    "hello": 5,
    "world": 10,
}

delete(scores, "hello")
delete(scores, "missing")

fmt.Println(len(scores)) // 1

var nilScores map[string]int
delete(nilScores, "hello") // seguro
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** asignar el zero value mantiene la key presente; `delete` elimina la propia key.

```go
scores["world"] = 0       // la key todavía existe
delete(scores, "world")  // la key deja de existir
```

No es necesario comprobar la presencia antes de eliminar. Si importa el valor eliminado o su presencia anterior, léelo primero con `comma ok` y después llama a `delete`.

Intentar asignar el resultado de `delete` es inválido porque el built-in no tiene valor de retorno.

## 5. Chuleta rápida

- Eliminar una entrada → `delete(m, key)`
- Key existente → eliminada
- Key ausente → no-op
- Map `nil` → no-op
- No hace falta comprobar antes la existencia
- `delete` no devuelve nada
- Asignar cero no es eliminar
- Necesitas el valor anterior → léelo antes de eliminar
- `len(m)` solo disminuye al eliminar una key existente

### Comprueba que realmente lo sabes

1. ¿Por qué asignar `m[key] = 0` es diferente de eliminar la key?
2. ¿Qué ocurre cuando `delete` recibe un map `nil`?
3. ¿Cómo conservarías el valor y la información de presencia antes de eliminar?
