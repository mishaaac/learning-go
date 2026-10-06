# El idioma `comma ok` en Go

Una lectura de map con un solo valor no puede distinguir una key ausente de un zero value almacenado. La forma `comma ok` devuelve tanto el valor como un resultado explícito de presencia.

## 1. Qué debo entender

- `value, ok := m[key]` realiza una lectura de map con dos valores.
- `value` es el valor almacenado o el zero value del tipo del value.
- `ok` es `true` exactamente cuando la key está presente.
- Un zero value almacenado produce `zero, true`; una key ausente produce `zero, false`.
- Usa esta forma solo cuando presencia y cero tengan significados diferentes.

## 2. Conceptos clave

| Resultado de la lectura | Significado |
|---|---|
| `5, true` | La key existe con valor `5` |
| `0, true` | La key existe con valor almacenado `0` |
| `0, false` | La key no existe; `0` es el zero value |
| `_, ok := m[key]` | Comprueba presencia ignorando el valor |

## 3. Cómo funciona en Go

```go
scores := map[string]int{
    "ready": 5,
    "waiting": 0,
}

value, ok := scores["waiting"]
fmt.Println(value, ok) // 0 true

value, ok = scores["missing"]
fmt.Println(value, ok) // 0 false
```

Para comprobar solo la presencia:

```go
if _, ok := scores["ready"]; ok {
    fmt.Println("present")
}
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** `value == 0` no implica que la key esté ausente; solo `ok` responde a la pregunta de presencia.

Una lectura de un solo valor sigue siendo preferible cuando el zero value ya modela correctamente la ausencia, como en un contador sencillo. Usa `comma ok` cuando la ausencia necesite una rama separada o cuando cero sea un dato almacenado válido.

El idioma también aparece al recibir de channels y en type assertions, pero cada contexto da a `ok` un significado específico.

## 5. Chuleta rápida

- Forma → `value, ok := m[key]`
- `value` → valor almacenado o zero value
- `ok == true` → la key está presente
- `ok == false` → la key está ausente
- `0, true` → cero almacenado
- `0, false` → key ausente
- Solo presencia → `_, ok := m[key]`
- Una lectura simple basta cuando cero representa bien la ausencia
- `comma ok` también aparece en otras operaciones de Go

### Comprueba que realmente lo sabes

1. ¿Qué información se pierde al usar únicamente `value := m[key]`?
2. ¿Cuándo es más apropiada una lectura de un solo valor que `comma ok`?
3. ¿Cómo representan estados diferentes `0, true` y `0, false`?
