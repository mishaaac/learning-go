# Clearing a Slice with `clear`

For a slice, `clear` replaces every current element with the element type's zero value. It resets contents without changing the slice's length or capacity.

## 1. What I Need to Understand

- `clear(slice)` acts on all elements from index `0` through `len(slice)-1`.
- Each element becomes the zero value of its type.
- The slice keeps the same length and capacity.
- `clear` does not remove elements or make the slice `nil`.
- Clearing a `nil` slice is a safe no-op.

## 2. Key Concepts

| Before | Operation | After |
|---|---|---|
| `[]string{"first", "second"}` | `clear(values)` | `[]string{"", ""}` |
| `[]int{4, 5}` | `clear(values)` | `[]int{0, 0}` |
| `len == 2` | `clear(values)` | `len == 2` |

`clear` was added as a built-in in Go 1.21.

## 3. How It Works in Go

```go
labels := []string{"first", "second", "third"}

clear(labels)

fmt.Printf("%q\n", labels) // ["" "" ""]
fmt.Println(len(labels))    // 3
fmt.Println(cap(labels))    // unchanged
```

The operation is conceptually equivalent to assigning the zero value to each current element.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** clearing values is not the same as changing the slice to length zero.

To reuse storage while making the slice logically empty, reslice it:

```go
labels = labels[:0]
```

That changes the length but does not zero the old elements. By contrast, `clear(labels)` zeroes current elements but preserves length.

> ⚠️ **Important correction:** “emptying a slice” is ambiguous. `clear(s)`, `s = s[:0]`, and `s = nil` produce different length, `nil` status, and data-retention behavior.

## 5. Quick Cheat Sheet

- `clear(s)` → set current elements to zero values
- `[]string` elements become `""`
- `[]int` elements become `0`
- `len(s)` does not change
- `cap(s)` does not change
- The slice does not become `nil`
- `clear(nilSlice)` → safe no-op
- `s = s[:0]` → length becomes `0`, values are not cleared first
- `s = nil` → slice becomes `nil`

### Test Yourself

1. What remains unchanged after `clear` is applied to a slice?
2. How does `clear(s)` differ from `s = s[:0]`?
3. What values result from clearing a slice of structs?
