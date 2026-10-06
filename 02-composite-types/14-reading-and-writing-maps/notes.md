# Reading and Writing Maps in Go

The expression `m[key]` accesses a map entry, while `m[key] = value` creates or replaces one. A missing key reads as the value type's zero value, which is convenient but sometimes ambiguous.

## 1. What I Need to Understand

- Use `m[key]` to read a value.
- Use `m[key] = value` to insert or replace an entry.
- A missing key returns the value type's zero value.
- `:=` cannot assign directly to a map index because the index expression is not a new variable.
- Numeric map values can use operations such as `m[key]++`.

## 2. Key Concepts

| Operation | Meaning |
|---|---|
| `value := m[key]` | Read an entry |
| `m[key] = value` | Insert or replace an entry |
| `m[key]++` | Increment a numeric value, starting from zero if absent |
| `len(m)` | Count current entries |
| Missing key | Return the value type's zero value |

## 3. How It Works in Go

```go
totalWins := map[string]int{}

totalWins["Orcas"] = 1
totalWins["Lions"] = 2

fmt.Println(totalWins["Orcas"])   // 1
fmt.Println(totalWins["Kittens"]) // 0

totalWins["Kittens"]++
totalWins["Lions"] = 3

fmt.Println(totalWins["Kittens"]) // 1
fmt.Println(totalWins["Lions"])   // 3
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** reading a missing key does not insert that key. It only produces the zero value.

This is invalid:

```go
totalWins["Orcas"] := 1 // compile-time error
```

Use `=` because a map index is an assignment target, not an identifier being declared.

The zero-value behavior makes maps concise counters: an absent `int` value behaves as `0`, so `counts[word]++` works without a separate initialization branch. If a stored zero and an absent key mean different things, use the `comma ok` form instead of a single-value lookup.

## 5. Quick Cheat Sheet

- Read → `m[key]`
- Insert → `m[key] = value`
- Replace → same assignment syntax
- Missing key → zero value
- Missing lookup does not add an entry
- `m[key] := value` → invalid
- Numeric counter → `m[key]++`
- Absent numeric key starts conceptually from `0`
- Need presence information → use `value, ok := m[key]`

### Test Yourself

1. Why does `counts[word]++` work when `word` is not yet a key?
2. Does reading a missing key change `len(m)`? Why?
3. Why is `:=` invalid on the left side of a map-index assignment?
