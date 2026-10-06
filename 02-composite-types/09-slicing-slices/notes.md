# Slicing Slices in Go

A slice expression creates a new view over part of existing storage. Its compact syntax is useful, but shared elements and capacity make later mutations—especially `append`—important to reason about.

## 1. What I Need to Understand

- `s[low:high]` includes `low` and excludes `high`.
- Omitting `low` means `0`; omitting `high` means `len(s)`.
- Slicing does not copy elements; the result normally shares the backing array.
- A subslice may have capacity beyond the elements currently visible in it.
- `s[low:high:max]` limits the resulting capacity to `max-low`.

## 2. Key Concepts

| Expression | Length | Capacity behavior |
|---|---:|---|
| `s[:2]` | `2` | May extend beyond index `2` |
| `s[1:]` | `len(s)-1` | Extends through remaining capacity |
| `s[:]` | `len(s)` | Shares the same visible range |
| `s[low:high:max]` | `high-low` | Limited to `max-low` |

## 3. How It Works in Go

```go
letters := []string{"a", "b", "c", "d"}
left := letters[:2]
right := letters[1:3]

left[1] = "B"
fmt.Println(letters) // [a B c d]
fmt.Println(right)   // [B c]
```

```text
backing array: [a][B][c][d]
                └ left ┘
                   └ right ┘
```

A capacity-limited view is written as:

```go
left = letters[:2:2] // len 2, cap 2
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `sub := original[:2]` creates a view; it does not copy the first two elements.

With `sub := original[:2]`, `append(sub, value)` may overwrite `original[2]` when spare shared capacity exists. Restricting capacity with `sub := original[:2:2]` forces an append beyond length two to allocate separate storage.

The full expression controls future growth, not element sharing inside the existing range. Even after `s[:2:2]`, assigning `sub[0]` still modifies the shared element.

Invalid slice bounds are rejected at compile time when provably constant or otherwise panic at runtime.

## 5. Quick Cheat Sheet

- `s[low:high]` → interval `[low, high)`
- `s[:high]` → starts at `0`
- `s[low:]` → ends at `len(s)`
- `s[:]` → full visible slice
- Slicing normally shares elements
- Mutating shared elements is visible through other views
- Subslice capacity can exceed its length
- `append` may overwrite shared later elements
- `s[low:high:max]` → capacity `max-low`
- Limiting capacity affects append, not current element sharing

### Test Yourself

1. Why can changing a subslice change the original slice?
2. Under what condition can appending to a subslice overwrite a later original element?
3. What does the third index in `s[low:high:max]` control?
