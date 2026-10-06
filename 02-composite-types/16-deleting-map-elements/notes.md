# Deleting Map Elements in Go

The built-in `delete` removes the entry associated with one map key. It is intentionally safe when the key is absent or the map is `nil`.

## 1. What I Need to Understand

- The syntax is `delete(m, key)`.
- If the key exists, its entire key-value entry is removed.
- If the key is absent, the operation does nothing.
- If the map is `nil`, the operation also does nothing.
- `delete` returns no value.

## 2. Key Concepts

| Situation | Result of `delete(m, key)` |
|---|---|
| Key exists | Entry is removed |
| Key is absent | No-op |
| Map is `nil` | No-op |
| Caller needs old value | Read it before deleting |

## 3. How It Works in Go

```go
scores := map[string]int{
    "hello": 5,
    "world": 10,
}

delete(scores, "hello")
delete(scores, "missing")

fmt.Println(len(scores)) // 1

var nilScores map[string]int
delete(nilScores, "hello") // safe
```

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** assigning the zero value keeps the key present; `delete` removes the key itself.

```go
scores["world"] = 0       // key still exists
delete(scores, "world")  // key no longer exists
```

There is no need to perform a separate presence check before deleting. If the removed value or its previous presence matters, read it first with `comma ok`, then call `delete`.

Attempting to assign the result of `delete` is invalid because the built-in has no return value.

## 5. Quick Cheat Sheet

- Delete one entry → `delete(m, key)`
- Existing key → removed
- Missing key → no-op
- `nil` map → no-op
- No preliminary existence check required
- `delete` returns nothing
- Assigning zero is not deletion
- Need old value → read before deleting
- `len(m)` decreases only when an existing key is removed

### Test Yourself

1. Why is setting `m[key] = 0` different from deleting the key?
2. What happens when `delete` receives a `nil` map?
3. How would you preserve the value and presence information before deletion?
