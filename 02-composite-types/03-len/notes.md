# `len` in Go

`len` is a built-in function that reports the length of several Go types. Its exact meaning depends on the operand, so understanding what it counts prevents indexing and text-processing mistakes.

## 1. What I Need to Understand

- For arrays and slices, `len` is the number of elements.
- For strings, `len` counts bytes, not Unicode code points.
- For maps, `len` counts entries.
- For channels, `len` counts elements currently queued in the buffer.
- A `nil` slice or map has length `0`.

## 2. Key Concepts

| Operand | Meaning of `len(value)` |
|---|---|
| Array | Number of array elements |
| Slice | Current number of slice elements |
| String | Number of bytes |
| Map | Number of key-value entries |
| Channel | Number of buffered elements currently queued |

`len` has language-defined typing rules. A normal function cannot reproduce exactly the same single operation over all these unrelated categories of types.

## 3. How It Works in Go

```go
array := [3]int{10, 20, 30}
slice := []int{10, 20}
var nilSlice []int
counts := map[string]int{"go": 2}

fmt.Println(len(array))    // 3
fmt.Println(len(slice))    // 2
fmt.Println(len(nilSlice)) // 0
fmt.Println(len(counts))   // 1
fmt.Println(len("🌞"))      // 4 UTF-8 bytes
```

For a non-empty slice, valid indices run from `0` through `len(slice)-1`.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `len(slice) == 0` does not prove that the slice is `nil`; a non-`nil` empty slice also has length `0`.

Calling `len` on a supported `nil` value is safe:

```go
var values []int
if len(values) == 0 {
    fmt.Println("no values")
}
```

Do not use `len(text)` as a character count for arbitrary UTF-8 text. Also, channel length is only a momentary buffer observation and is usually not a reliable basis for coordinating concurrent work.

## 5. Quick Cheat Sheet

- `len(array)` → number of elements
- `len(slice)` → current number of elements
- `len(string)` → bytes
- `len(map)` → entries
- `len(channel)` → currently buffered elements
- `len(nilSlice)` → `0`
- `len(nilMap)` → `0`
- Last valid slice index → `len(s)-1`, only when non-empty
- `len(s) == 0` does not imply `s == nil`
- Unsupported operand type → compile-time error

### Test Yourself

1. Why can `len("🌞")` differ from the number of visible symbols?
2. What can and cannot be concluded from `len(s) == 0`?
3. Why is `len(channel)` a poor synchronization mechanism?
