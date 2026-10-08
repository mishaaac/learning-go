# Labels in Go `for` Loops

Labels let `break` or `continue` target an enclosing control structure instead of the nearest one. They are most useful in nested loops when an inner decision determines what should happen to the outer loop.

## 1. What I Need to Understand

- A label is an identifier followed by `:` before a statement.
- Unlabeled `break` and `continue` target the nearest applicable structure.
- `continue label` begins the next iteration of the labeled `for`.
- `break label` terminates the labeled `for`, `switch`, or `select` as allowed.
- A `continue` label must identify an enclosing `for` loop.

## 2. Key Concepts

| Statement | Effect inside nested loops |
|---|---|
| `continue` | Next iteration of inner loop |
| `continue outer` | Next iteration of labeled outer loop |
| `break` | Exit nearest applicable structure |
| `break outer` | Exit labeled outer structure |

## 3. How It Works in Go

```go
samples := []string{"hello", "apple"}

outer:
for _, sample := range samples {
    for _, r := range sample {
        if r == 'l' {
            continue outer
        }
        fmt.Printf("%c", r)
    }
    fmt.Println(" accepted")
}
```

When an inner rune is `'l'`, execution skips the rest of both the inner work and the current outer iteration.

## 4. Examples, Differences, and Common Mistakes

> **Do not confuse:** plain `continue` affects the inner loop; `continue outer` affects the explicitly labeled loop.

A common pattern rejects an outer item as soon as one inner value is invalid:

```go
outer:
for _, item := range items {
    for _, part := range item.Parts {
        if !valid(part) {
            continue outer
        }
    }
    accept(item)
}
```

Labels should clarify which loop is targeted. If nested flow remains difficult to follow, extracting logic into a function may be clearer than adding more labels.

## 5. Quick Cheat Sheet

- Label syntax → `outer:`
- Place label immediately before the target statement
- Plain `continue` → nearest `for`
- `continue outer` → next outer iteration
- Plain `break` → nearest applicable structure
- `break outer` → exit labeled structure
- `continue` label must name an enclosing `for`
- Labeled control is mainly for nested loops
- It also skips remaining outer-body code
- Prefer descriptive labels

### Test Yourself

1. What code is skipped when an inner loop executes `continue outer`?
2. Why can `continue` target only a labeled `for`?
3. When might extracting a function be clearer than using a label?
