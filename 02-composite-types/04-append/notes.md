# `append` in Go

`append` adds elements to the end of a slice and returns the resulting slice. Keeping that return value is essential because growth may change the slice's length, capacity, and backing array.

## 1. What I Need to Understand

- `append` accepts a destination slice followed by one or more elements.
- It works directly with a `nil` slice.
- It returns the updated slice; the usual pattern is `s = append(s, value)`.
- `source...` expands another slice so its elements can be appended.
- Go passes the slice value to `append` by value, so the resulting slice descriptor must be returned.

## 2. Key Concepts

| Form | Effect |
|---|---|
| `append(s, value)` | Adds one element |
| `append(s, a, b)` | Adds several elements |
| `append(s, other...)` | Adds every element from `other` |
| `append(bytes, text...)` | Appends a string's bytes to `[]byte` |
| `s = append(...)` | Preserves the returned slice |

## 3. How It Works in Go

```go
var numbers []int

numbers = append(numbers, 10)
numbers = append(numbers, 20, 30)

more := []int{40, 50}
numbers = append(numbers, more...)

fmt.Println(numbers) // [10 20 30 40 50]
```

Conceptually:

```text
slice value → append → updated slice value
                   ├── greater len
                   └── possibly new backing array
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `append(s, other)` tries to append `other` as one element; `append(s, other...)` appends its elements.

Ignoring the return value is invalid as a standalone call and, more importantly, would lose changes to the slice descriptor:

```go
numbers = append(numbers, 60) // correct
```

If there is enough capacity, `append` may reuse the existing backing array. Otherwise, it allocates larger storage and copies existing elements, which is why other slices that shared the old storage may not observe later changes.

> ⚠️ **Important correction:** call-by-value is part of the explanation, but it does not mean slice elements are copied when a slice is passed. The copied slice descriptor still refers to the same backing array until an operation such as growth replaces it.

## 5. Quick Cheat Sheet

- `append` adds elements at the end
- It works on a `nil` slice
- Always keep the result
- One value → `s = append(s, value)`
- Several values → `s = append(s, a, b)`
- Another slice → `s = append(s, other...)`
- `...` expands the source slice's elements
- `append` increases `len`
- Existing capacity may be reused
- Insufficient capacity may produce a new backing array

### Test Yourself

1. Why must the result of `append` be assigned or otherwise used?
2. What does `...` change when appending another slice?
3. How can an allocation during `append` change the relationship between two slices?
