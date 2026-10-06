# Slice Capacity in Go

A slice records both its current length and the capacity available from its starting position in the backing array. Capacity explains when `append` can reuse storage and when it must obtain new storage.

## 1. What I Need to Understand

- `len(s)` counts current elements; `cap(s)` measures how far the slice can extend in its backing array.
- The invariant is `0 <= len(s) <= cap(s)`.
- `append` reuses the backing array when enough capacity exists.
- If capacity is insufficient, `append` allocates new storage and copies existing elements.
- Reserving a reasonable capacity can avoid repeated allocations and copies.

## 2. Key Concepts

| Concept | Meaning |
|---|---|
| Length | Elements currently belonging to the slice |
| Capacity | Elements available from the slice start to the end of its backing storage |
| Backing array | Storage containing the slice's elements |
| Reallocation | New backing storage obtained when growth does not fit |
| Preallocation | Reserving capacity before repeated appends |

## 3. How It Works in Go

```go
values := make([]int, 0, 4)
fmt.Println(len(values), cap(values)) // 0 4

values = append(values, 10, 20, 30)
fmt.Println(len(values), cap(values)) // 3 4

values = append(values, 40)
fmt.Println(len(values), cap(values)) // 4 4
```

```text
backing array
[10][20][30][  ]
 ↑---- len ----↑
 ↑------ cap -------↑
```

The next append requires more capacity, so it may move the elements to a new backing array.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** capacity is not the number of currently valid indices. Only indices below `len(s)` are directly accessible.

For an array, `cap(array) == len(array)`. For a `nil` slice, both values are `0`.

> ⚠️ **Important correction:** exact capacity growth factors are implementation details, not promises of the Go language. Current runtime behavior may change across Go versions, element sizes, and allocation constraints; code must not depend on capacities such as `1, 2, 4, 8` after successive appends.

Preallocate when an expected size is known, but avoid treating capacity as a fixed maximum: `append` can still grow beyond it.

## 5. Quick Cheat Sheet

- `len(s)` → current elements
- `cap(s)` → available extent in backing storage
- Always `0 <= len(s) <= cap(s)`
- Valid direct indices depend on `len`
- Spare capacity lets `append` reuse storage
- Exhausted capacity may trigger allocation and copying
- Reallocation can break sharing with older slices
- `make([]T, 0, n)` reserves capacity for appends
- `cap(nilSlice)` → `0`
- Exact growth factors are not API guarantees

### Test Yourself

1. Why can a slice have capacity greater than its length without allowing access to all that capacity by index?
2. What work may occur when `append` exceeds the current capacity?
3. Why should a program not rely on an observed sequence of capacity values?
