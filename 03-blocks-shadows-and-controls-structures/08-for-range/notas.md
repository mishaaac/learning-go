# `for-range` en Go

`for-range` itera directamente sobre el contenido de arrays, slices, maps y strings. Los valores que produce dependen del operando y el valor de iteración es una copia, no una referencia asignable al elemento original.

## 1. Qué debo entender

- Los arrays y slices producen un índice y una copia del valor del elemento.
- Los maps producen una key y una copia del value en un orden no especificado.
- Los strings producen un byte offset y un `rune` decodificado.
- `_` descarta un primer valor no deseado; omitir la segunda variable conserva solo la primera.
- Con la semántica de Go 1.22, las variables declaradas por una cláusula range con `:=` son nuevas en cada iteración.

## 2. Conceptos clave

| Operando | Primer valor | Segundo valor |
|---|---|---|
| Array/slice | Índice | Copia del elemento |
| Map | Key | Copia del value |
| String | Byte offset inicial | `rune` decodificado |

Las formas habituales son `for i, value := range values`, `for _, value := range values` y `for key := range values`.

## 3. Cómo funciona en Go

```go
values := []int{2, 4, 6}
for index, value := range values {
    fmt.Println(index, value)
}

for _, value := range values {
    value *= 2
}
fmt.Println(values) // [2 4 6]
```

Para un string:

```go
for offset, r := range "aπ!" {
    fmt.Println(offset, r, string(r))
}
// desplazamientos: 0, 1, 3
```

## 4. Ejemplos, diferencias y errores comunes

> **No confundir:** el primer valor para un string es un byte offset, no un número de carácter.

Cambiar la variable de valor de range no cambia un elemento del slice. Usa el índice cuando quieras mutarlo:

```go
for index := range values {
    values[index] *= 2
}
```

Nunca dependas del orden de iteración de un map. El UTF-8 inválido de un string se decodifica como `utf8.RuneError` y avanza un byte para esa secuencia inválida.

> ⚠️ **Corrección importante:** las variables son nuevas por iteración cuando la cláusula range las declara con `:=` bajo la semántica del lenguaje Go 1.22+. Las variables proporcionadas con `=` fueron declaradas antes y vuelven a recibir una asignación en cada iteración.

## 5. Chuleta rápida

- Array/slice → índice, copia del elemento
- Map → key, copia del value
- String → byte offset, `rune`
- El orden de un map no está especificado
- Ignorar primer valor → `_`
- Solo necesitas el primero → omite el segundo
- El value de range no muta el elemento original
- Muta un slice mediante su índice
- Los runes multibyte hacen saltar los offsets
- UTF-8 inválido → `utf8.RuneError`
- En Go 1.22+, las variables con `:=` son nuevas por iteración

### Comprueba que realmente lo sabes

1. ¿Por qué asignar a la variable value de range no modifica un elemento del slice?
2. ¿Por qué los offsets de un string podrían ser `0`, `1` y `3` en vez de números de carácter consecutivos?
3. ¿En qué se diferencian `:=` y `=` para variables de range con la semántica Go 1.22+?
