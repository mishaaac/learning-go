# Clearing a Map with `clear`

For a map, `clear` removes every entry in one operation. The map remains usable if it was initialized, but its length becomes zero.

## 1. What I Need to Understand

- `clear(m)` removes all key-value entries.
- Afterward, `len(m) == 0`.
- An initialized map remains initialized and writable.
- Calling `clear` on a `nil` map is a safe no-op.
- `clear` behaves differently for maps and slices.

## 2. Key Concepts

| Need | Operation | Result |
|---|---|---|
| Remove one map entry | `delete(m, key)` | Only that key is removed |
| Remove every map entry | `clear(m)` | `len(m) == 0` |
| Reset slice elements | `clear(s)` | Elements become zero values; length stays |

`clear` became a built-in in Go 1.21.

## 3. How It Works in Go

```go
scores := map[string]int{
    "hello": 5,
    "world": 10,
}

clear(scores)
fmt.Println(len(scores)) // 0

scores["new"] = 1       // still writable
fmt.Println(scores)      // map[new:1]
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `clear(map)` removes entries, while `clear(slice)` preserves elements and replaces their values with zero values.

For a `nil` map:

```go
var scores map[string]int
clear(scores) // safe; scores is still nil
```

Clearing an initialized map does not make it `nil`. If the distinction matters, `m = nil` is a different operation; future writes would then require initialization again.

Use `delete` when selecting one key and `clear` when every current entry should be removed.

## 5. Quick Cheat Sheet

- Remove all entries → `clear(m)`
- Result → `len(m) == 0`
- Initialized map remains writable
- Initialized map does not become `nil`
- `clear(nilMap)` → no-op
- One key → `delete(m, key)`
- All keys → `clear(m)`
- `clear(map)` removes entries
- `clear(slice)` zeroes elements without changing length

### Test Yourself

1. Can an initialized map be written immediately after `clear`? Why?
2. How does `clear` differ between a map and a slice?
3. How is `clear(m)` different from assigning `m = nil`?
