# Converting Between Arrays and Slices

Arrays and slices can expose or reproduce the same sequence in different ways. The important question is whether an operation shares the original storage or copies its elements.

## 1. What I Need to Understand

- `array[:]` creates a slice view and shares the array's storage.
- `[N]T(slice)` creates an independent array value by copying elements.
- `(*[N]T)(slice)` creates an array pointer that shares storage with the slice.
- For either slice-to-array form, `N` cannot exceed `len(slice)`.
- The target array length must be explicit; `[...]T(slice)` is not a valid conversion.

## 2. Key Concepts

| Operation | Copies? | Shares storage? |
|---|---:|---:|
| `array[:]` | No | Yes |
| `array[1:3]` | No | Yes |
| `[N]T(slice)` | Yes | No |
| `(*[N]T)(slice)` | No | Yes |

## 3. How It Works in Go

```go
array := [4]int{1, 2, 3, 4}
view := array[:]
view[0] = 10
fmt.Println(array) // [10 2 3 4]

slice := []int{5, 6, 7, 8}
copyArray := [4]int(slice)
sharedArray := (*[4]int)(slice)

slice[0] = 50
fmt.Println(copyArray)    // [5 6 7 8]
fmt.Println(sharedArray)  // &[50 6 7 8]
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `[4]int(slice)` copies, while `(*[4]int)(slice)` shares the slice's underlying storage.

A shorter array conversion copies only the required prefix:

```go
firstTwo := [2]int(slice)
```

If `len(slice) < N`, converting to `[N]T` or `*[N]T` panics at runtime. Extra capacity does not help; the rule uses length.

If an API repeatedly converts arrays of different lengths just to accept them, a slice parameter is often the clearer design because array length is part of the type.

## 5. Quick Cheat Sheet

- Array to slice → `array[:]`
- Array slicing shares storage
- Slice to array value → `[N]T(slice)`
- Array value conversion copies
- Slice to array pointer → `(*[N]T)(slice)`
- Array pointer conversion shares storage
- Required rule → `N <= len(slice)`
- `cap(slice)` does not make a too-large conversion valid
- Smaller target arrays copy the prefix
- `[...]T(slice)` is not a valid conversion

### Test Yourself

1. Which conversion should you use when the result must be independent of the slice?
2. Why can a slice with `len == 2` and `cap == 8` not convert to `[4]int`?
3. How can modifying `array[:]` affect the original array?
