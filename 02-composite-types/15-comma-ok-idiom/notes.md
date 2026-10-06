# The `comma ok` Idiom in Go

A one-value map lookup cannot distinguish a missing key from a stored zero value. The `comma ok` form returns both the value and an explicit presence result.

## 1. What I Need to Understand

- `value, ok := m[key]` performs a two-value map lookup.
- `value` is the stored value or the value type's zero value.
- `ok` is `true` exactly when the key is present.
- A stored zero value produces `zero, true`; an absent key produces `zero, false`.
- Use this form only when presence and zero have different meanings.

## 2. Key Concepts

| Lookup result | Meaning |
|---|---|
| `5, true` | Key exists with value `5` |
| `0, true` | Key exists with stored value `0` |
| `0, false` | Key does not exist; `0` is the zero value |
| `_, ok := m[key]` | Check presence while ignoring the value |

## 3. How It Works in Go

```go
scores := map[string]int{
    "ready": 5,
    "waiting": 0,
}

value, ok := scores["waiting"]
fmt.Println(value, ok) // 0 true

value, ok = scores["missing"]
fmt.Println(value, ok) // 0 false
```

For presence alone:

```go
if _, ok := scores["ready"]; ok {
    fmt.Println("present")
}
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** `value == 0` does not imply that the key is absent; only `ok` answers the presence question.

A single-value lookup remains preferable when the zero value already models absence correctly, as with a simple counter. Use `comma ok` when absence requires a separate branch or when zero is valid stored data.

The idiom also appears with channel receives and type assertions, but each context gives `ok` a context-specific meaning.

## 5. Quick Cheat Sheet

- Form → `value, ok := m[key]`
- `value` → stored value or zero value
- `ok == true` → key is present
- `ok == false` → key is absent
- `0, true` → stored zero
- `0, false` → missing key
- Presence only → `_, ok := m[key]`
- Simple lookup is enough when zero safely represents absence
- `comma ok` also appears in other Go operations

### Test Yourself

1. What information is lost when you use only `value := m[key]`?
2. When is a single-value lookup more appropriate than `comma ok`?
3. How do `0, true` and `0, false` represent different states?
