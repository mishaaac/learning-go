# Slices in Go

A slice is a flexible view over a contiguous portion of an underlying array. It is Go's usual choice for variable-length sequences because the length is not part of the slice type.

## 1. What I Need to Understand

- `[]T` is a slice type; its current length is not part of that type.
- A slice describes underlying storage with a length and capacity.
- The zero value of a slice is `nil`, with length and capacity `0`.
- Indexing is limited by `len`, even when more capacity exists.
- Slices cannot be compared with each other using `==`; they may only be compared directly with `nil`.

## 2. Key Concepts

| Concept | Meaning | Why it matters |
|---|---|---|
| `[]int{1, 2, 3}` | Slice literal | Creates a ready-to-use slice with three elements |
| `var values []int` | `nil` slice | Safe with `len`, `cap`, `append`, and `range` |
| `[][]int` | Slice of slices | Inner slices may have different lengths |
| `slices.Equal` | Element-by-element equality | Use it instead of `==` for comparable elements |
| `slices.EqualFunc` | Equality using a supplied function | Supports custom comparison logic |

## 3. How It Works in Go

```go
values := []int{10, 20, 30}
values[1] = 25

var empty []int
fmt.Println(values, len(values)) // [10 25 30] 3
fmt.Println(empty == nil)        // true
```

Indexed elements in a literal may leave gaps at the zero value:

```go
sparse := []int{0: 1, 3: 8} // [1 0 0 8]
```

For content comparison:

```go
fmt.Println(slices.Equal([]int{1, 2}, []int{1, 2})) // true
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** a slice can grow through `append`, but writing directly to `values[len(values)]` is still out of range.

`nil` and `[]int{}` both have length `0`, but only the first compares equal to `nil`. Most operations treat them similarly, so distinguish them only when the representation matters.

> ⚠️ **Important correction:** slices are not directly comparable to other slices. `slices.Equal` requires comparable element values; use `slices.EqualFunc` when equality needs custom logic. Comparing slices of unrelated element types is also a compile-time error unless an appropriate `EqualFunc` call supports both types.

`reflect.DeepEqual` has broader semantics, including treating `nil` and non-`nil` empty slices as different. Prefer the specific `slices` functions when their semantics match the task.

## 5. Quick Cheat Sheet

- `[]T` → slice of `T`
- Length is not part of the slice type
- `var s []T` → `nil`, `len == 0`, `cap == 0`
- `[]T{}` → empty but non-`nil`
- Valid indices depend on `len`, not `cap`
- `[][]T` → slice of slices
- `slice == slice` → compile-time error
- `slice == nil` → valid
- `slices.Equal` → normal element equality
- `slices.EqualFunc` → custom equality

### Test Yourself

1. Why can slices of different lengths still have the same type?
2. Why does extra capacity not make `s[len(s)]` a valid index?
3. When would `slices.EqualFunc` be more appropriate than `slices.Equal`?
