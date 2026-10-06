# Using `make` for Slices

`make` creates an initialized slice with a chosen length and optional capacity. The crucial distinction is that length creates usable elements, while extra capacity only reserves room for future growth.

## 1. What I Need to Understand

- `make([]T, length)` creates `length` elements initialized to the zero value.
- Without a third argument, capacity equals length.
- `make([]T, length, capacity)` separates existing elements from reserved space.
- Indices are valid only below `len`, not below `cap`.
- For building with `append`, `make([]T, 0, capacity)` is usually the intended form.

## 2. Key Concepts

| Expression | `len` | `cap` | Initial elements |
|---|---:|---:|---|
| `make([]int, 5)` | 5 | 5 | Five zero values |
| `make([]int, 5, 10)` | 5 | 10 | Five zero values |
| `make([]int, 0, 10)` | 0 | 10 | None |

The required relationship is `length <= capacity`.

## 3. How It Works in Go

```go
indexed := make([]int, 3)
indexed[0] = 10

collected := make([]int, 0, 3)
collected = append(collected, 10, 20)

fmt.Println(indexed)   // [10 0 0]
fmt.Println(collected) // [10 20]
```

```text
make([]T, len, cap)
          │    └── reserved extent
          └─────── existing elements
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `make([]int, 5)` is not an empty slice with room for five values; it already contains five values.

This common mistake adds a sixth element:

```go
values := make([]int, 5)
values = append(values, 10) // [0 0 0 0 0 10]
```

Use direct indexing when the final length is known and all positions will be filled. Use length `0` plus capacity when values will arrive progressively through `append`.

If constant arguments specify `length > capacity`, compilation fails. If runtime values produce that relationship, the call panics.

## 5. Quick Cheat Sheet

- `make([]T, n)` → `len == cap == n`
- Created elements contain zero values
- `make([]T, n, c)` requires `n <= c`
- `len` → existing, indexable elements
- `cap` → reserved growth space
- `make([]T, 0, c)` → empty, non-`nil`, preallocated slice
- Capacity alone does not make an index valid
- `append` adds after the current length
- Known final length → allocate length and index
- Progressive collection → length `0`, capacity estimate, then `append`

### Test Yourself

1. Why does appending to `make([]int, 5)` produce a sixth element?
2. When should the second argument to `make` be zero?
3. What determines whether an index is currently valid: length or capacity?
