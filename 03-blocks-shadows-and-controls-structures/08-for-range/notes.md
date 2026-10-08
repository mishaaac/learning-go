# `for-range` in Go

`for-range` iterates directly over the contents of arrays, slices, maps, and strings. The values it produces depend on the operand, and the iteration value is a copy rather than an assignable reference to the original element.

## 1. What I Need to Understand

- Arrays and slices produce an index and a copied element value.
- Maps produce a key and copied value in an unspecified order.
- Strings produce a byte offset and a decoded `rune`.
- `_` discards an unwanted first value; omitting the second variable keeps only the first.
- Since Go 1.22 semantics, variables declared by a range clause with `:=` are new for each iteration.

## 2. Key Concepts

| Operand | First value | Second value |
|---|---|---|
| Array/slice | Index | Element copy |
| Map | Key | Value copy |
| String | Starting byte offset | Decoded `rune` |

Common forms are `for i, value := range values`, `for _, value := range values`, and `for key := range values`.

## 3. How It Works in Go

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

For a string:

```go
for offset, r := range "aπ!" {
    fmt.Println(offset, r, string(r))
}
// offsets: 0, 1, 3
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** the first value for a string is a byte offset, not a character number.

Changing the range value variable does not change a slice element. Use the index when mutation is intended:

```go
for index := range values {
    values[index] *= 2
}
```

Never rely on map iteration order. Invalid UTF-8 in a string is decoded as `utf8.RuneError`, advancing one byte for that invalid sequence.

> ⚠️ **Important correction:** per-iteration variables are new when the range clause declares them with `:=` under Go 1.22+ language semantics. Variables supplied with `=` were declared earlier and are assigned again on each iteration.

## 5. Quick Cheat Sheet

- Array/slice → index, element copy
- Map → key, value copy
- String → byte offset, `rune`
- Map order is unspecified
- Ignore first value → `_`
- Need only first value → omit the second
- Range value does not mutate the original element
- Mutate a slice through its index
- Multibyte runes make string offsets jump
- Invalid UTF-8 → `utf8.RuneError`
- Go 1.22+ `:=` variables are new per iteration

### Test Yourself

1. Why does assigning to a range value variable not modify a slice element?
2. Why might string offsets be `0`, `1`, and `3` rather than consecutive character numbers?
3. How do `:=` and `=` differ for range variables under Go 1.22+ semantics?
