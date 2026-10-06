# Declaring a Slice in Go

Go offers several valid slice declarations, and each represents a different initial state. Choosing among `var`, a literal, and `make` depends on whether values already exist and how the slice will be populated.

## 1. What I Need to Understand

- Use `var s []T` when a slice may remain unused; it starts as `nil`.
- Use a slice literal when initial values are already known.
- Use `make([]T, n)` when `n` elements must exist immediately.
- Use `make([]T, 0, n)` when starting empty and collecting values with `append`.
- A capacity estimate reduces growth but does not limit the final size.

## 2. Key Concepts

| Situation | Declaration | Initial state |
|---|---|---|
| May stay empty | `var data []int` | `nil`, length `0` |
| Must be empty but non-`nil` | `data := []int{}` | non-`nil`, length `0` |
| Known values | `data := []int{2, 4}` | length `2` |
| Exact indexable length | `make([]int, n)` | `n` zero-valued elements |
| Append with estimated size | `make([]int, 0, n)` | empty with reserved capacity |

## 3. How It Works in Go

```go
var optional []int
fixedValues := []int{2, 4, 6}

transformed := make([]int, len(fixedValues))
for i, value := range fixedValues {
    transformed[i] = value * 2
}

collected := make([]int, 0, 10)
collected = append(collected, 7)
```

The first `make` supports indexed writes; the second supports progressive appends without creating placeholder elements.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `[]T{}` and `var s []T` both have length zero, but only the latter is `nil`.

This declaration already contains five elements:

```go
data := make([]int, 5)
```

Appending does not fill those elements; it adds after them. If the final length is only an estimate, prefer `make([]T, 0, estimate)` and `append` so unused zero values do not remain in the result.

The `nil` versus empty distinction can matter at serialization or API boundaries, but ordinary slice operations often work with both.

## 5. Quick Cheat Sheet

- Possibly unused → `var s []T`
- Known values → `s := []T{...}`
- Empty non-`nil` required → `s := []T{}`
- Exact length, indexed writes → `make([]T, n)`
- Progressive appends → `make([]T, 0, n)`
- `make([]T, n)` creates `n` real elements
- Capacity is an estimate, not a maximum
- `append` always grows from the current length
- Prefer the declaration that matches how values will be written

### Test Yourself

1. Why is `make([]T, n)` suitable for indexed transformation output?
2. Why is `make([]T, 0, n)` safer when the final count is only estimated?
3. In what situation might the difference between `nil` and empty non-`nil` matter?
