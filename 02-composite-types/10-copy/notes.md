# `copy` in Go

The built-in `copy` transfers elements into positions that already exist in a destination slice. It is useful for independent copies, selected ranges, and safe copying between overlapping regions.

## 1. What I Need to Understand

- The form is `copy(destination, source)`.
- It copies `min(len(destination), len(source))` elements.
- Capacity does not increase the amount copied; only current lengths matter.
- `copy` does not grow the destination.
- Overlapping source and destination regions are allowed.

## 2. Key Concepts

| Operation | Result |
|---|---|
| `n := copy(dst, src)` | Copies elements and returns their count |
| `copy(dst, src[2:])` | Copies a selected source range |
| `copy(s[:3], s[1:])` | Safely copies overlapping regions |
| `copy(dst, array[:])` | Uses an array through a slice view |

## 3. How It Works in Go

```go
source := []int{1, 2, 3, 4}
destination := make([]int, len(source))

count := copy(destination, source)
destination[0] = 99

fmt.Println(count)       // 4
fmt.Println(source)      // [1 2 3 4]
fmt.Println(destination) // [99 2 3 4]
```

For an overlapping shift:

```go
values := []int{1, 2, 3, 4}
copy(values[:3], values[1:])
fmt.Println(values) // [2 3 4 4]
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `copy` replaces existing destination elements; `append` adds elements and increases length.

This destination has capacity but no length, so nothing is copied:

```go
destination := make([]int, 0, 4)
count := copy(destination, []int{1, 2}) // count == 0
```

Create sufficient destination length when making an independent copy. The returned count may be ignored when the lengths already make the outcome certain.

Slicing an array with `array[:]` lets it serve as the source or destination because `copy` operates on slices.

## 5. Quick Cheat Sheet

- Syntax → `copy(dst, src)`
- First argument → destination
- Second argument → source
- Return value → elements copied
- Count → `min(len(dst), len(src))`
- Capacity does not determine the count
- `copy` does not grow `dst`
- Slicing selects the copied region
- Overlap is supported
- `array[:]` exposes an array as a slice

### Test Yourself

1. Why does copying into `make([]int, 0, 10)` copy zero elements?
2. How would you create a fully independent copy of a slice?
3. Why is overlapping `copy` useful for shifting elements?
